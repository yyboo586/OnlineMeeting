// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMeetingDao is the data access object for the table t_meeting.
type TMeetingDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  TMeetingColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// TMeetingColumns defines and stores column names for the table t_meeting.
type TMeetingColumns struct {
	Id          string // ID
	RoomNumber  string // ID
	Topic       string //
	Mode        string //
	Distance    string //
	Type        string //
	Status      string //
	Location    string //
	CreatorId   string // ID
	Description string //
	CreateTime  string //
	StartTime   string //
	EndTime     string //
}

// tMeetingColumns holds the columns for the table t_meeting.
var tMeetingColumns = TMeetingColumns{
	Id:          "id",
	RoomNumber:  "room_number",
	Topic:       "topic",
	Mode:        "mode",
	Distance:    "distance",
	Type:        "type",
	Status:      "status",
	Location:    "location",
	CreatorId:   "creator_id",
	Description: "description",
	CreateTime:  "create_time",
	StartTime:   "start_time",
	EndTime:     "end_time",
}

// NewTMeetingDao creates and returns a new DAO object for table data access.
func NewTMeetingDao(handlers ...gdb.ModelHandler) *TMeetingDao {
	return &TMeetingDao{
		group:    "default",
		table:    "t_meeting",
		columns:  tMeetingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMeetingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMeetingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMeetingDao) Columns() TMeetingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMeetingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMeetingDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TMeetingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}