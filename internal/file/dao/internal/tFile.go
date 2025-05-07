// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TFileDao is the data access object for the table t_file.
type TFileDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TFileColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TFileColumns defines and stores column names for the table t_meeting.
type TFileColumns struct {
	Id         		string // ID
	MRoomNumber 	string // ID
	FileName    	string //
	SaveName    	string //
	FileSize    	string //
	FileType    	string //
	StoragePath 	string //
	UploaderId  	string //
	UploaderName   	string //
	UploadTime 		string //
	DeletorId  		string //
	DeletorName   	string //
	DeleteTime     	string //
	Status 			string
}

// tFileColumns holds the columns for the table t_meeting.
var tFileColumns = TFileColumns{
	Id:      		"id",
	MRoomNumber: 	"m_room_number",
	FileName:    	"file_name",
	SaveName:       "save_name",
	FileSize:    	"file_size",
	FileType:       "file_type",
	StoragePath:    "storage_path",
	UploaderId:   	"uploader_id",
	UploaderName: 	"upload_name",
	DeletorId:  	"deletor_id",
	DeletorName: 	"deletor_name",
	DeleteTime:  	"create_time",
	Status:			"status",
}

// NewTFileDao creates and returns a new DAO object for table data access.
func NewTFileDao(handlers ...gdb.ModelHandler) *TFileDao {
	return &TFileDao{
		group:    "default",
		table:    "t_file",
		columns:  tFileColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TFileDao) Columns() TFileColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TFileDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *TFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}