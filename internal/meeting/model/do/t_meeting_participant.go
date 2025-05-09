// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeetingParticipant is the golang structure of table t_meeting_participant for DAO operations like Where/Data.
type TMeetingParticipant struct {
	g.Meta      `orm:"table:t_meeting_participant, do:true"`
	Id          interface{} // ID
	MRoomNumber interface{} // ID, t_meetingroom_number
	UserId      interface{} // ID
	UserName    interface{} //
	Role        interface{} // , 1:, 2:, 3:
	Status      interface{} // , 1:, 2:, 3:
	UpdateTime  *gtime.Time //
	JoinTime    *gtime.Time //
	ExitTime    *gtime.Time //
}
