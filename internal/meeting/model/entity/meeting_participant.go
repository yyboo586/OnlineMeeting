package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type UserInfo struct {
	UserID     string
	UserName   string
	Role       MeetingParticipantRole
	Status     MeetingParticipantStatus
	AcceptTime *gtime.Time
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
	AcceptTime *gtime.Time `orm:"accept_time"`
	JoinTime   *gtime.Time `orm:"join_time"`
	ExitTime   *gtime.Time `orm:"exit_time"`
}
