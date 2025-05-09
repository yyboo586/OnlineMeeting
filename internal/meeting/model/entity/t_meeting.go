// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeeting is the golang structure for table t_meeting.
type TMeeting struct {
	Id          int64       `json:"id"          orm:"id"          ` // ID
	RoomNumber  string      `json:"roomNumber"  orm:"room_number" ` // ID
	Topic       string      `json:"topic"       orm:"topic"       ` //
	Mode        int         `json:"mode"        orm:"mode"        ` //
	Distance    int         `json:"distance"    orm:"distance"    ` //
	Type        int         `json:"type"        orm:"type"        ` //
	Status      int         `json:"status"      orm:"status"      ` //
	Location    string      `json:"location"    orm:"location"    ` //
	CreatorId   string      `json:"creatorId"   orm:"creator_id"  ` // ID
	Description string      `json:"description" orm:"description" ` //
	CreateTime  *gtime.Time `json:"createTime"  orm:"create_time" ` //
	StartTime   *gtime.Time `json:"startTime"   orm:"start_time"  ` //
	EndTime     *gtime.Time `json:"endTime"     orm:"end_time"    ` //
}
