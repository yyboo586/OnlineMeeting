// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeeting is the golang structure of table t_meeting for DAO operations like Where/Data.
type TMeeting struct {
	g.Meta      `orm:"table:t_meeting, do:true"`
	Id          interface{} // ID
	RoomNumber  interface{} // ID
	Topic       interface{} //
	Mode        interface{} //
	Distance    interface{} //
	Type        interface{} //
	Status      interface{} //
	Location    interface{} //
	CreatorId   interface{} // ID
	Description interface{} //
	CreateTime  *gtime.Time //
	StartTime   *gtime.Time //
	EndTime     *gtime.Time //
}