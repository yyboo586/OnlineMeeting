package file

import (
	"OnlineMeeting/api/v1/common"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadReq struct {
	g.Meta `path:"/" method:"post" mime:"multipart/form-data" tags:"文件管理" summary:"上传文件" `
	common.Author
	File         *ghttp.UploadFile `type:"file" v:"required" dc:"File to upload"`
	RoomNumber   string            `v:"required#会议ID必须" dc:"会议ID"`
	UploaderID   string            `v:"required#上传者ID必须" dc:"上传者ID"`
	UploaderName string            `v:"required#上传者名称必须" dc:"上传者名称"`
}

type UploadRes struct {
	ID string
}

type DownloadReq struct {
	g.Meta `path:"/:id" method:"get" tags:"文件管理" summary:"下载文件" `
	common.Author
	ID int64 `p:"id" v:"required#文件ID必须" dc:"文件ID"`
}

type DownloadRes struct {
	FilePath string
}

type DeleteReq struct {
	g.Meta `path:"/:id" method:"delete" tags:"文件管理" summary:"删除文件"`
	common.Author
	ID          int64  `p:"id" v:"required#文件ID必须" dc:"文件ID"`
	DeletorID   string `v:"required#删除者ID必须" dc:"删除者ID"`
	DeletorName string `v:"required#删除者名称必须" dc:"删除者名称"`
}

type DeleteRes struct {
	common.EmptyRes
}

type ListReq struct {
	g.Meta `path:"/list" tags:"文件管理" method:"get" summary:"获取单场会议的所有文件"`
	common.Author
	common.PageReq
	RoomNumber string `v:"required#会议ID必须" dc:"会议ID"`
}

type FileEntity struct {
	ID       string
	FileName string
}

type ListRes struct {
	List []*FileEntity
	common.PageRes
}
