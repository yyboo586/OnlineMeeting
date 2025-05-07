package entity

import (
	"github.com/gogf/gf/os/gtime"
)

type FileStatus int

const (
	_ FileStatus = iota
	FileStatusNormal
	FileStatusDeleted
)

type FileStatusText string

const (
	FileStatusNormalText  FileStatusText = "正常"
	FileStatusDeletedText FileStatusText = "已删除"
)

type FileType int

const (
	_             FileType = iota
	FileTypeImage          // 图片
	FileNormal             // 普通文件
	FileTypeVideo          // 视频
	FileTypeAudio          // 音频
)

type File struct {
	ID                string
	MeetingRoomNumber string
	FileName          string
	SaveName          string
	FileSize          int
	FileType          FileType
	StoragePath       string
	Status            FileStatus
	UploaderID        string
	UploaderName      string
	UploadTime        *gtime.Time
	DeletorID         string
	DeletorName       string
	DeleteTime        *gtime.Time
}

type FileDB struct {
	ID                string
	MeetingRoomNumber string
	FileName          string
	SaveName          string
	FileSize          int
	FileType          int
	StoragePath       string
	Status            int
	UploaderID        string
	UploaderName      string
	UploadTime        *gtime.Time
	DeletorID         string
	DeletorName       string
	DeleteTime        *gtime.Time
}

func ConvertToEntityModel(dbFile *FileDB) *File {
	return &File{
		ID:                dbFile.ID,
		MeetingRoomNumber: dbFile.MeetingRoomNumber,
		FileName:          dbFile.FileName,
		SaveName:          dbFile.SaveName,
		FileSize:          dbFile.FileSize,
		FileType:          FileType(dbFile.FileType),
		StoragePath:       dbFile.StoragePath,
	}
}
