package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type MeetingParticipantRole uint

const (
	ParticipantRoleAdmin     MeetingParticipantRole = 1 << iota // 管理员
	ParticipantRoleModerator                                    // 主持人
	ParticipantRoleMember                                       // 普通成员
)

const (
	ParticipantRoleUnknownText   string = "未知角色"
	ParticipantRoleAdminText     string = "管理员"
	ParticipantRoleModeratorText string = "主持人"
	ParticipantRoleMemberText    string = "普通成员"
)

type MeetingParticipantStatus int

const (
	ParticipantStatusUnknown   MeetingParticipantStatus = iota // 未知
	ParticipantStatusUnhandled                                 // 未处理
	ParticipantStatusAccepted                                  // 已接受
	ParticipantStatusRejected                                  // 已拒绝
)

const (
	ParticipantStatusUnknownText   string = "未知"
	ParticipantStatusUnhandledText string = "未处理"
	ParticipantStatusAcceptedText  string = "已接受"
	ParticipantStatusRejectedText  string = "已拒绝"
)

type UserInfo struct {
	UserID     string
	UserName   string
	Role       MeetingParticipantRole
	Status     MeetingParticipantStatus
	UpdateTime *gtime.Time
	JoinTime   *gtime.Time
	ExitTime   *gtime.Time
}

type MeetingParticipantDB struct {
	ID         int64       `orm:"id,primary"`
	RoomNumber string      `orm:"m_room_number"`
	UserID     string      `orm:"user_id"`
	UserName   string      `orm:"user_name"`
	Role       int         `orm:"role"`
	Status     int         `orm:"status"`
	UpdateTime *gtime.Time `orm:"update_time"`
	JoinTime   *gtime.Time `orm:"join_time"`
	ExitTime   *gtime.Time `orm:"exit_time"`
}
