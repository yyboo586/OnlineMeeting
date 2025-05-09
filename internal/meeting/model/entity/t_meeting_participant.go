// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeetingParticipant is the golang structure for table t_meeting_participant.
type TMeetingParticipant struct {
	Id          int64       `json:"Id"          orm:"id"            ` // ID
	MRoomNumber string      `json:"MRoomNumber" orm:"m_room_number" ` // ID, t_meetingroom_number
	UserId      string      `json:"UserId"      orm:"user_id"       ` // ID
	UserName    string      `json:"UserName"    orm:"user_name"     ` //
	Role        uint        `json:"Role"        orm:"role"          ` // , 1:, 2:, 3:
	Status      int         `json:"Status"      orm:"status"        ` // , 1:, 2:, 3:
	UpdateTime  *gtime.Time `json:"UpdateTime"  orm:"update_time"   ` //
	JoinTime    *gtime.Time `json:"JoinTime"    orm:"join_time"     ` //
	ExitTime    *gtime.Time `json:"ExitTime"    orm:"exit_time"     ` //
}
