package file

import (
	"OnlineMeeting/api/v1/common"
	"OnlineMeeting/api/v1/file"
	"OnlineMeeting/internal/consts"
	"OnlineMeeting/internal/file/dao"
	"OnlineMeeting/internal/file/model/do"
	"OnlineMeeting/internal/file/service"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

var (
	err error
)

func init() {
	service.RegisterFile(New())
}

func New() service.IFile {
	f := &File{}
	f.node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	return f
}

type File struct {
	node *snowflake.Node
}

func (f *File) Upload(ctx context.Context, req *file.UploadReq) (res *file.UploadRes, err error) {
	if req.File == nil {
		return nil, gerror.New("文件不能为空")
	}
	storagePath := fmt.Sprintf("uploads/%s/", req.RoomNumber)
	saveName, err := req.File.Save(storagePath, true)
	if err != nil {
		deleteFile(storagePath + saveName)
		return nil, err
	}
	id := f.node.Generate()
	fileInfo := do.TFile{
		Id:           id.Int64(),
		MRoomNumber:  req.RoomNumber,
		FileName:     req.File.Filename,
		SaveName:     saveName,
		FileSize:     req.File.Size, // 字节
		FileType:     1,
		StoragePath:  "uploads/" + req.RoomNumber,
		Status:       1,
		UploaderId:   req.UploaderID,
		UploaderName: req.UploaderName,
		UploadTime:   gtime.Now(),
	}

	_, err = dao.TFile.Ctx(ctx).Insert(fileInfo)
	if err != nil {
		if dErr := deleteFile(storagePath + saveName); dErr != nil {
			log.Println("删除文件失败", dErr)
		}
	}

	res = &file.UploadRes{
		ID: id.String(),
	}

	return
}

func (f *File) Download(ctx context.Context, fileId int64) (res *file.DownloadRes, err error) {
	// 1. 查询文件信息
	fileInfo, err := dao.TFile.Ctx(ctx).Where("id", fileId).One()
	if err != nil {
		return nil, gerror.Wrap(err, "查询文件失败")
	}

	if fileInfo.IsEmpty() {
		return nil, gerror.New("文件不存在")
	}

	if fileInfo["status"].Int() != 1 { // 状态为1表示正常
		return nil, gerror.New("文件已被删除")
	}

	// 2. 构造文件访问路径（本地示例）
	storagePath := fileInfo["storage_path"].String()
	saveName := fileInfo["save_name"].String()
	filePath := fmt.Sprintf("/%s/%s", storagePath, saveName) // 假设前端可通过此路径访问

	res = &file.DownloadRes{
		FilePath: filePath,
	}

	return res, nil
}
func (f *File) Delete(ctx context.Context, fileId int64, deletorId, deletorName string) (err error) {
	// 1. 查询文件信息
	fileInfo, err := dao.TFile.Ctx(ctx).Where("id", fileId).One()
	if err != nil {
		return gerror.Wrap(err, "查询文件失败")
	}

	if fileInfo.IsEmpty() {
		return gerror.New("文件不存在")
	}

	// 2. 更新数据库记录（软删除）
	_, err = dao.TFile.Ctx(ctx).Where("id", fileId).Update(do.TFile{
		Status:      2, // 已删除
		DeletorId:   deletorId,
		DeletorName: deletorName,
		DeleteTime:  gtime.Now(),
	})

	if err != nil {
		return gerror.Wrap(err, "更新文件状态失败")
	}

	// 3. 构造文件路径
	storagePath := fileInfo["storage_path"].String()
	saveName := fileInfo["save_name"].String()
	filePath := fmt.Sprintf("%s/%s", storagePath, saveName)

	// 4. 删除磁盘文件
	if err = deleteFile(filePath); err != nil {
		log.Printf("删除文件失败: %v", err)
		// 可选择继续执行数据库操作或直接返回错误
	}

	return nil
}

func (f *File) ListByRoom(ctx context.Context, req *file.ListReq) (res *file.ListRes, err error) {
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	// 1. 查询文件列表
	sqlModel := dao.TFile.Ctx(ctx).Where("m_room_number", req.RoomNumber).Where("status", 1)
	total, err := sqlModel.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询文件列表失败")
	}

	fileList, err := sqlModel.Page(req.PageNum, req.PageSize).All()
	if err != nil {
		return nil, gerror.Wrap(err, "查询文件列表失败")
	}
	list := make([]*file.FileEntity, 0)
	for _, v := range fileList.List() {
		list = append(list, &file.FileEntity{
			ID:       fmt.Sprintf("%d", v["id"].(int64)),
			FileName: v["file_name"].(string),
		})
	}

	res = &file.ListRes{
		List: list,
		PageRes: common.PageRes{
			CurrentPage: req.PageNum,
			Total:       total,
		},
	}
	return
}

func deleteFile(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	return os.Remove(path)
}
