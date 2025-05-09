// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TMeeting is the golang structure for table t_meeting.
type TMeeting struct {
	Id          int64       `json:"Id"          orm:"id"          ` // ID
	RoomNumber  string      `json:"RoomNumber"  orm:"room_number" ` // ID
	Topic       string      `json:"Topic"       orm:"topic"       ` //
	Mode        int         `json:"Mode"        orm:"mode"        ` //
	Distance    int         `json:"Distance"    orm:"distance"    ` //
	Type        int         `json:"Type"        orm:"type"        ` //
	Status      int         `json:"Status"      orm:"status"      ` //
	Location    string      `json:"Location"    orm:"location"    ` //
	CreatorId   string      `json:"CreatorId"   orm:"creator_id"  ` // ID
	Description string      `json:"Description" orm:"description" ` //
	CreateTime  *gtime.Time `json:"CreateTime"  orm:"create_time" ` //
	StartTime   *gtime.Time `json:"StartTime"   orm:"start_time"  ` //
	EndTime     *gtime.Time `json:"EndTime"     orm:"end_time"    ` //
}
