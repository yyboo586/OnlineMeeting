package service

import (
	"OnlineMeeting/api/v1/file"
	"context"
)

type IFile interface {
	Upload(ctx context.Context, req *file.UploadReq) (res *file.UploadRes, err error)
	ListByRoom(ctx context.Context, req *file.ListReq) (res *file.ListRes, err error)
	Delete(ctx context.Context, id int64, deletorID, deletorName string) (err error)
	Download(ctx context.Context, id int64) (res *file.DownloadRes, err error)
}

var localFile IFile

func File() IFile {
	if localFile == nil {
		panic("implement not found for interface ISqmeeting, forgot register?")
	}
	return localFile
}

func RegisterFile(i IFile) {
	localFile = i
}
