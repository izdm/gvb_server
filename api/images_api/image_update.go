package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// 在这段代码中，binding:"required"通常用于表单数据绑定或请求体数据绑定的场景。具体来说：
// 当使用一个处理 HTTP 请求的框架（如gin）时，binding:"required"指示该框架在绑定请求数据（如 JSON 数据、表单数据等）到结构体时，对应的字段是必需的。
// 如果请求数据中缺少ID或Name字段，框架会返回一个错误，表示这些字段是必需的但未提供。这有助于确保数据的完整性和合法性。
// 具体来说，msg标签用于提供自定义的错误消息。例如：
// 在ID字段的结构体标签中，msg:"请选择文件id"表示当在绑定（例如，从 HTTP 请求中绑定数据到结构体）过程中ID字段出现问题（比如缺少该字段，因为binding:"required"要求该字段必须存在）时，将返回的错误消息是 “请选择文件 id”。
// 同样，Name字段的msg:"请输入文件名称"表示当Name字段在绑定过程中出现问题时，将返回 “请输入文件名称” 这个错误消息。
// 这种方式常用于在处理用户输入（例如，通过 API 接收的数据）时，为用户提供更友好、更具针对性的错误提示信息。
type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请选择文件id"`
	Name string `json:"name" binding:"required" msg:"请输入文件名称"`
}

func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.BannerModel
	err := global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}

	res.OkWithMessage("图片名称修改成功", c)
	return
}
