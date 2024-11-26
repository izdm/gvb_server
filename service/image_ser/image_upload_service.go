package image_ser

import (
	"errors"
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

var (
	//图片上传白名单
	WhiteImageList = []string{
		"jpg", "jpeg", "png", "ico", "gif", "svg", "webp",
	}
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功？
}

func (ImageService) ImageUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	fileName := file.Filename
	res.FileName = fileName

	// 校验文件合法性
	if err := validateFile(file); err != nil {
		res.IsSuccess = false
		res.Msg = err.Error()
		return
	}

	// 读取文件内容
	fileObj, err := file.Open()
	if err != nil {
		global.Log.Error(err.Error())
		res.IsSuccess = false
		res.Msg = "文件打开失败"
		return
	}
	defer fileObj.Close()

	byteData, err := io.ReadAll(fileObj)
	if err != nil {
		global.Log.Error(err.Error())
		res.IsSuccess = false
		res.Msg = "文件读取失败"
		return
	}

	// 检查文件是否已存在
	imageHash := utils.Md5(byteData)
	if existingFilePath, found := checkFileInDB(imageHash); found {
		res.FileName = existingFilePath
		res.IsSuccess = true
		res.Msg = "图片已存在"
		return
	}

	// 上传文件（本地或七牛云）
	filePath, fileType, uploadErr := uploadFile(fileName, byteData)
	if uploadErr != nil {
		res.IsSuccess = false
		res.Msg = uploadErr.Error()
		return
	}

	// 图片入库
	global.DB.Create(&models.BannerModel{
		Path:      filePath,
		Hash:      imageHash,
		Name:      fileName,
		ImageType: fileType,
	})

	res.FileName = filePath
	res.IsSuccess = true
	res.Msg = "上传成功"
	return

}
func validateFile(file *multipart.FileHeader) error {
	fileName := file.Filename

	// 校验后缀名
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !utils.InList(suffix, WhiteImageList) {
		return errors.New("非法文件类型")
	}

	// 校验大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		return fmt.Errorf("图片大小超过限制，当前大小为%.2fMB，最大为:%dMB", size, global.Config.Upload.Size)
	}

	return nil
}
func checkFileInDB(imageHash string) (string, bool) {
	var bannerModel models.BannerModel
	err := global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
	if err == nil {
		// 文件已存在
		return bannerModel.Path, true
	}
	return "", false
}
func uploadFile(fileName string, byteData []byte) (filePath string, fileType ctype.ImageType, err error) {
	if global.Config.QiNiu.Enable {
		// 七牛云上传
		filePath, err = qiniu.UploadImage(byteData, fileName, global.Config.QiNiu.Prefix)
		if err != nil {
			global.Log.Error("七牛云上传失败：" + err.Error())
			return "", ctype.Local, err
		}
		return filePath, ctype.QiNiu, nil
	}

	// 本地上传
	basePath := global.Config.Upload.Path
	filePath = path.Join(basePath, fileName)

	err = os.WriteFile(filePath, byteData, fs.ModePerm)
	if err != nil {
		global.Log.Error("本地上传失败：" + err.Error())
		return "", ctype.Local, err
	}
	return filePath, ctype.Local, nil
}
