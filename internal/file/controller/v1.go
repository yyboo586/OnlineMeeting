package controller

import (
	"OnlineMeeting/api/v1/file"
	"OnlineMeeting/internal/file/service"
	meetingService "OnlineMeeting/internal/meeting/service"

	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

type fileController struct{}

var FileController = new(fileController)

// Upload 文件上传接口
func (c *fileController) Upload(ctx context.Context, req *file.UploadReq) (res *file.UploadRes, err error) {
	exists, err := meetingService.Meeting().CheckMeetingExists(ctx, req.RoomNumber)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, gerror.New("会议不存在，无法上传文件")
	}

	res, err = service.File().Upload(ctx, req)

	return
}

func (c *fileController) Delete(ctx context.Context, req *file.DeleteReq) (res *file.DeleteRes, err error) {
	res = new(file.DeleteRes)
	err = service.File().Delete(ctx, req.ID, req.DeletorID, req.DeletorName)
	return
}

func (c *fileController) Download(ctx context.Context, req *file.DownloadReq) (res *file.DownloadRes, err error) {
	res, err = service.File().Download(ctx, req.ID)
	return
}

func (c *fileController) List(ctx context.Context, req *file.ListReq) (res *file.ListRes, err error) {
	res, err = service.File().ListByRoom(ctx, req)
	return
}
