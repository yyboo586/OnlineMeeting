// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TMeetingParticipantDao is the data access object for the table t_meeting_participant.
type TMeetingParticipantDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  TMeetingParticipantColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// TMeetingParticipantColumns defines and stores column names for the table t_meeting_participant.
type TMeetingParticipantColumns struct {
	Id          string // ID
	MRoomNumber string // ID, t_meetingroom_number
	UserId      string // ID
	UserName    string //
	Role        string // , 1:, 2:, 3:
	Status      string // , 1:, 2:, 3:
	UpdateTime  string //
	JoinTime    string //
	ExitTime    string //
}

// tMeetingParticipantColumns holds the columns for the table t_meeting_participant.
var tMeetingParticipantColumns = TMeetingParticipantColumns{
	Id:          "id",
	MRoomNumber: "m_room_number",
	UserId:      "user_id",
	UserName:    "user_name",
	Role:        "role",
	Status:      "status",
	UpdateTime:  "update_time",
	JoinTime:    "join_time",
	ExitTime:    "exit_time",
}

// NewTMeetingParticipantDao creates and returns a new DAO object for table data access.
func NewTMeetingParticipantDao(handlers ...gdb.ModelHandler) *TMeetingParticipantDao {
	return &TMeetingParticipantDao{
		group:    "default",
		table:    "t_meeting_participant",
		columns:  tMeetingParticipantColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *TMeetingParticipantDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *TMeetingParticipantDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *TMeetingParticipantDao) Columns() TMeetingParticipantColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *TMeetingParticipantDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *TMeetingParticipantDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *TMeetingParticipantDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
