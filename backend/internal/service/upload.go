package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joey17520/ailiaili/internal/cache"
	"github.com/joey17520/ailiaili/internal/domain/dto"
	"github.com/joey17520/ailiaili/internal/domain/model"
	"github.com/joey17520/ailiaili/internal/domain/vo"
	"github.com/joey17520/ailiaili/internal/global"
	"github.com/joey17520/ailiaili/utils"
)

func UploadImg(ctx *gin.Context, file *multipart.FileHeader) (string, error) {
	suffix := path.Ext(file.Filename)
	fileName := generateImgFilename(suffix)

	objectKey := "image/" + fileName
	filePath := "./upload/image/" + fileName
	// 参数校验
	if !utils.IsImgType(suffix) { // 文件后缀
		return "", errors.New("文件类型错误")
	}

	//文件大小限制
	if !utils.FileSize(file.Size, 1, global.Config.File.MaxImgSize) {
		return "", errors.New("文件大小超出限制")
	}

	//保存文件
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return "", errors.New("文件上传失败")
	}

	url := generateFileUrl(objectKey)
	if global.Config.Storage.OssType != "local" {
		// 上传到OSS
		global.Storage.PutObjectFromFile(objectKey, filePath)
	}

	// 缓存url
	userId := ctx.GetUint("userId")
	cache.SetUploadImage(url, userId)

	return url, nil
}

func UploadVideoCreate(ctx *gin.Context, videoFileReq dto.VideoFileReq) (vo.ResourceResp, error) {
	userId := ctx.GetUint("userId")
	var fileInfo model.VideoFile
	if err := global.Mysql.Where("hash = ? and uid = ?", videoFileReq.Hash, userId).Find(&fileInfo).Error; err != nil {
		utils.ErrorLog("视频文件信息不存在", "upload", videoFileReq.Hash)
		return vo.ResourceResp{}, errors.New("视频文件不存在")
	}

	// 先创建视频记录
	uploadVideoPath := "./upload/video/" + fileInfo.DirName + "/upload.mp4"
	vid, _ := initVideo(userId, uploadVideoPath, fileInfo.OriginalName)
	if vid == 0 {
		return vo.ResourceResp{}, errors.New("创建失败")
	}

	resource, err := CompleteUploadVideo(vid, userId, fileInfo.DirName, fileInfo.OriginalName)
	if err != nil {
		return vo.ResourceResp{}, err
	}

	return resource, nil
}

func UploadVideoAdd(ctx *gin.Context, vid uint, videoFileReq dto.VideoFileReq) (vo.ResourceResp, error) {
	userId := ctx.GetUint("userId")
	var fileInfo model.VideoFile
	if err := global.Mysql.Where("hash = ? and uid = ?", videoFileReq.Hash, userId).Find(&fileInfo).Error; err != nil {
		utils.ErrorLog("视频文件信息不存在", "upload", videoFileReq.Hash)
		return vo.ResourceResp{}, errors.New("视频文件不存在")
	}

	resource, err := CompleteUploadVideo(vid, userId, fileInfo.DirName, fileInfo.OriginalName)
	if err != nil {
		return vo.ResourceResp{}, err
	}

	return resource, nil
}

func UploadVideoCheck(ctx *gin.Context, videoFileReq dto.VideoFileReq) ([]int, error) {
	userId := ctx.GetUint("userId")
	var fileInfo model.VideoFile
	if err := global.Mysql.Where("hash = ? and uid = ?", videoFileReq.Hash, userId).Find(&fileInfo).Error; err != nil {
		utils.ErrorLog("视频文件信息不存在", "upload", videoFileReq.Hash)
		return nil, errors.New("视频文件不存在")
	}

	var checks []int
	fileDir := "./upload/video/" + fileInfo.DirName
	for i := 0; i < fileInfo.ChunksCount; i++ {
		if utils.IsFileExists(fmt.Sprintf("%s/chunks/%d.part", fileDir, i)) {
			checks = append(checks, i)
		}
	}

	return checks, nil
}

func UploadVideoChunk(ctx *gin.Context, file *multipart.FileHeader) error {
	userId := ctx.GetUint("userId")

	// 获取分片信息
	fileHash := ctx.PostForm("hash")
	fileName := ctx.PostForm("name")
	chunkIndex, _ := strconv.Atoi(ctx.PostForm("chunkIndex"))
	totalChunks, _ := strconv.Atoi(ctx.PostForm("totalChunks"))

	suffix := path.Ext(fileName)
	if !utils.IsVideoType(suffix) { // 文件后缀
		return errors.New("视频上传失败")
	}

	if !utils.FileSize(file.Size, int64(totalChunks), global.Config.File.MaxVideoSize) {
		return errors.New("文件大小超出限制")
	}

	// 查询文件表如果哈希存在则
	var dirName string
	var videoFileInfo model.VideoFile
	global.Mysql.Where("uid = ? and hash = ?", userId, fileHash).First(&videoFileInfo)
	if videoFileInfo.ID == 0 {
		dirName = generateVideoFilename()
		global.Mysql.Create(&model.VideoFile{Uid: userId, Hash: fileHash, DirName: dirName, OriginalName: fileName, ChunksCount: totalChunks})
	} else {
		dirName = videoFileInfo.DirName
	}

	fileDir := "./upload/video/" + dirName
	chunksPath := fileDir + "/chunks/" + strconv.Itoa(chunkIndex) + ".part"
	if err := ctx.SaveUploadedFile(file, chunksPath); err != nil {
		return errors.New("文件上传失败")
	}

	return nil
}

func UploadVideoMerge(ctx *gin.Context, videoFileReq dto.VideoFileReq) error {
	userId := ctx.GetUint("userId")
	var fileInfo model.VideoFile
	if err := global.Mysql.Where("hash = ? and uid = ?", videoFileReq.Hash, userId).Find(&fileInfo).Error; err != nil {
		utils.ErrorLog("视频文件信息不存在", "upload", videoFileReq.Hash)
		return errors.New("视频文件不存在")
	}

	fileDir := "./upload/video/" + fileInfo.DirName
	if err := mergeChunks(fileDir, fileInfo.ChunksCount); err != nil {
		utils.ErrorLog("合并分片失败", "upload", err.Error())
		return errors.New("合并分片失败")
	}

	if err := os.RemoveAll(fileDir + "/chunks/"); err != nil {
		utils.ErrorLog("删除临时文件夹失败", "upload", err.Error())
	}

	return nil
}

func mergeChunks(fileDir string, totalChunks int) error {
	outputPath := fileDir + "/upload.mp4"
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := 0; i < totalChunks; i++ {
		chunkPath := fmt.Sprintf("%s/chunks/%d.part", fileDir, i)
		chunk, err := os.ReadFile(chunkPath)
		if err != nil {
			return err
		}
		if _, err := outFile.Write(chunk); err != nil {
			return err
		}
	}

	return nil
}

func CompleteUploadVideo(vid, userId uint, videoName, title string) (vo.ResourceResp, error) {
	uploadVideoPath := "./upload/video/" + videoName + "/upload.mp4"
	transcodingInfo, err := ProcessVideoInfo(uploadVideoPath)
	if err != nil {
		return vo.ResourceResp{}, errors.New("读取视频信息失败")
	}

	// 去掉后缀名
	titleWithoutExt := title[:len(title)-len(path.Ext(title))]

	// 存入数据库
	resource := model.Resource{
		Vid:       vid,
		Uid:       userId,
		Title:     titleWithoutExt,
		CodecName: transcodingInfo.CodecName,
		Status:    global.VIDEO_PROCESSING,
		Duration:  transcodingInfo.Duration,
	}
	if err := global.Mysql.Create(&resource).Error; err != nil {
		return vo.ResourceResp{}, errors.New("保存视频失败")
	}

	// 启动转码服务
	transcodingInfo.VideoID = vid
	transcodingInfo.DirName = videoName
	transcodingInfo.ResourceID = resource.ID
	transcodingInfo.OutputDir = "./upload/video/" + videoName + "/"
	transcodingInfo.InputFile = transcodingInfo.OutputDir + "upload.mp4"
	go VideoTransCoding(transcodingInfo)

	return vo.ResourceToResourceResp(resource), nil
}

// 生成文件url
func generateFileUrl(objectKey string) string {
	if global.Config.Storage.OssType != "local" {
		global.Storage.GetObjectUrl(objectKey)
	}

	return "/api/" + objectKey
}

// 初始化视频
func initVideo(userId uint, videoPath, title string) (uint, error) {
	// 生成封面
	coverName := generateImgFilename(".jpg")
	objectKey := "image/" + coverName
	filePath := "./upload/image/" + coverName

	GenerateCover(videoPath, filePath)
	if global.Config.Storage.OssType != "local" {
		// 上传到OSS
		global.Storage.PutObjectFromFile(objectKey, filePath)
	}
	// 去掉后缀名
	titleWithoutExt := title[:len(title)-len(path.Ext(title))]

	videoId, err := CreateVideo(&model.Video{
		Uid:       userId,
		Cover:     generateFileUrl(objectKey),
		Title:     titleWithoutExt,
		Copyright: true,
		Status:    global.CREATED_VIDEO,
	})
	if err != nil {
		return 0, err
	}

	return videoId, nil
}

// 随机生成图片文件名
func generateImgFilename(suffix string) string {
	id := global.SnowflakeNode.Generate()
	return id.String() + suffix
}

// 随机视频文件名
func generateVideoFilename() string {
	id := global.SnowflakeNode.Generate()
	return id.String()
}
