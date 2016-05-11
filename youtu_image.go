package youtu

import (
    	"encoding/base64"
)

//ImageRsp  请求 结构体
type ImageRsp struct {
	AppID string     `json:"app_id"`         //App的 API ID
	Image string     `json:"image"`          //base64编码的二进制图片数据
}

//FuzzyDetectRsp 判断一个图像的模糊程度 返回 结构体
type FuzzyDetectRsp struct {
	Fuzzy  bool   `json:"fuzzy"` //是否模糊
    FuzzyConfidence float64 `json:"fuzzy_confidence"` //模糊参考值,范围 0-1的浮点数,越大置信度越高
	ErrorCode int32  `json:"errorcode"` //返回状态码
	ErrorMsg  string `json:"errormsg"`  //返回错误消息
}

//FuzzyDetect 判断一个图像的模糊程度
func (y *Youtu) FuzzyDetect(image []byte) (rsp FuzzyDetectRsp, err error) {
    b64Image := base64.StdEncoding.EncodeToString(image)
	req := ImageRsp{
		AppID:  y.appID(),
		Image:  b64Image,
	}
	err = y.interfaceRequest("youtu/imageapi/fuzzydetect", req, &rsp)
	return
}

//ImageTag 类型说明
type ImageTag struct {
	TagName  string   `json:"tag_name"` //返回图像标签的名字
    TagConfidence int32 `json:"tag_confidence"` //图像标签的置信度,整形范围 0-100,越大置信度越高
	TagConfideceF float64  `json:"tag_confidence_f"` //图像标签的置信度,整形范围 0-100,越大置信度越高
}

//ImageTags ImageTag的切片
type ImageTags []ImageTag

//ImageTagRsp 识别一个图像的标签信息,对图像分类 返回 结构体
type ImageTagRsp struct {
	Tags  ImageTags   `json:"tags"` //图像的分类标签
	ErrorCode int32  `json:"errorcode"` //返回状态码
	ErrorMsg  string `json:"errormsg"`  //返回错误消息
}

//ImageTag 识别一个图像的标签信息,对图像分类
func (y *Youtu) ImageTag(image []byte) (rsp ImageTagRsp, err error) {
    b64Image := base64.StdEncoding.EncodeToString(image)
	req := ImageRsp{
		AppID:  y.appID(),
		Image:  b64Image,
	}
	err = y.interfaceRequest("youtu/imageapi/imagetag", req, &rsp)
	return
}
