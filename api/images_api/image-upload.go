package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/service/image_ser"
	"io/fs"
	"os"
)

var (
	//图片上传白名单
	WhiteImageList = []string{
		"jpg", "jpeg", "png", "ico", "gif", "svg", "webp",
	}
)

// 上传多个文件
func (ImagesApi) ImageUploadView(c *gin.Context) {
	// 检查请求格式
	if c.ContentType() != "multipart/form-data" {
		res.FailWithMessage("请求格式不正确，请使用multipart/form-data格式上传文件", c)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}

	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在文件", c)
		return
	}

	// 确保本地存储路径存在
	basePath := global.Config.Upload.Path
	if _, err := os.ReadDir(basePath); err != nil {
		if err = os.MkdirAll(basePath, fs.ModePerm); err != nil {
			global.Log.Error(err.Error())
			res.FailWithMessage("创建本地上传目录失败", c)
			return
		}
	}

	var resList []image_ser.FileUploadResponse

	for _, file := range fileList {
		// 调用服务层上传逻辑
		serviceRes := service.ServiceApp.ImageService.ImageUploadService(file)
		resList = append(resList, serviceRes)
	}

	// 返回上传结果
	res.OkWithData(resList, c)
}

//fileHeader, err := c.FormFile("image")
//if err != nil {
//	res.FailWithMessage(err.Error(), c)
//	return
//}
//fmt.Println(fileHeader.Header)
//fmt.Println(fileHeader.Size)
//fmt.Println(fileHeader.Filename)
