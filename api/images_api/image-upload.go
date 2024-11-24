package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"io/fs"
	"os"
	"path"
)

type FileUploadResponse struct {
	FileName  string `json:"file_name"`
	Msg       string `json:"msg"`        //消息
	IsSuccess bool   `json:"is_success"` //是否上传成功？
}

// 上传单个图片 返回图片url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	//对post请求做判断
	if c.ContentType() != "multipart/form-data" {
		res.FailWithMessage("请求格式不正确，请使用multipart/form-data格式上传文件", c)
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fileList, ok := form.File["images"]

	if !ok {
		res.FailWithMessage("不存在文件", c)
	}

	//判断路劲是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err.Error())
		}
	}

	//不存在就创建
	var resList []FileUploadResponse

	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过，当前大小为%.2fMB，设定最大为:%dMB", size, global.Config.Upload.Size),
			})
			continue
		}

		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				FileName:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})
			continue
		}

		resList = append(resList, FileUploadResponse{
			FileName:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})
	}
	res.OkWithData(resList, c)
	//fileHeader, err := c.FormFile("image")
	//if err != nil {
	//	res.FailWithMessage(err.Error(), c)
	//	return
	//}
	//fmt.Println(fileHeader.Header)
	//fmt.Println(fileHeader.Size)
	//fmt.Println(fileHeader.Filename)
}
