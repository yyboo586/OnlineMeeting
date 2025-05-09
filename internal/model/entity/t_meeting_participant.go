// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeetingParticipant is the golang structure for table t_meeting_participant.
type TMeetingParticipant struct {
	Id          int64       `json:"id"          orm:"id"            ` // ID
	MRoomNumber string      `json:"mRoomNumber" orm:"m_room_number" ` // ID, t_meetingroom_number
	UserId      string      `json:"userId"      orm:"user_id"       ` // ID
	UserName    string      `json:"userName"    orm:"user_name"     ` //
	Role        uint        `json:"role"        orm:"role"          ` // , 1:, 2:, 3:
	Status      int         `json:"status"      orm:"status"        ` // , 1:, 2:, 3:
	UpdateTime  *gtime.Time `json:"updateTime"  orm:"update_time"   ` //
	JoinTime    *gtime.Time `json:"joinTime"    orm:"join_time"     ` //
	ExitTime    *gtime.Time `json:"exitTime"    orm:"exit_time"     ` //
}
