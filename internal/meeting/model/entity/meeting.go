package entity

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

type MeetingMode int

const (
	MeetingModeUnknown      MeetingMode = iota // 未知模式
	MeetingModeRoming                          // 漫游模式
	MeetingModeMeetingRoom                     // 会议室模式
	MeetingModeVirtualHuman                    // 虚拟人模式
)

const (
	MeetingModeUnknownText      string = "未知模式"
	MeetingModeRomingText       string = "漫游模式"
	MeetingModeMeetingRoomText  string = "会议室模式"
	MeetingModeVirtualHumanText string = "虚拟人模式"
)

type MeetingType int

const (
	MeetingTypeUnknown   MeetingType = iota
	MeetingTypeInstant               // 即时会议
	MeetingTypeScheduled             // 预约会议
)

const (
	MeetingTypeUnknownText   string = "未知类型"
	MeetingTypeInstantText   string = "即时会议"
	MeetingTypeScheduledText string = "预约会议"
)

type MeetingStatus int

const (
	MeetingStatusUnknown  MeetingStatus = iota
	MeetingStatusCreated                // 已创建
	MeetingStatusStarted                // 进行中
	MeetingStatusEnded                  // 已结束
	MeetingStatusCanceled               // 已取消
)

const (
	MeetingStatusUnknownText  string = "未知状态"
	MeetingStatusCreatedText  string = "待开始"
	MeetingStatusStartedText  string = "进行中"
	MeetingStatusEndedText    string = "已结束"
	MeetingStatusCanceledText string = "已取消"
)

type Meeting struct {
	ID            int64
	RoomNumber    string
	Topic         string
	Mode          MeetingMode
	Distance      int
	Type          MeetingType
	Status        MeetingStatus
	Location      string
	Description   string
	CreatorInfo   *UserInfo
	ModeratorInfo *UserInfo
	CreateTime    *gtime.Time
	StartTime     *gtime.Time
	EndTime       *gtime.Time
	Members       []*UserInfo
}

type MeetingDB struct {
	ID          int64
	RoomNumber  string
	Topic       string
	Mode        int
	Distance    int
	Type        int
	Status      int
	Location    string
	CreatorId   string
	Description string
	CreateTime  *gtime.Time
	StartTime   *gtime.Time
	EndTime     *gtime.Time
	Members     []*MeetingParticipantDB
}

func ConvertToEntityModel(sqDB *MeetingDB) *Meeting {
	res := &Meeting{
		ID:          sqDB.ID,
		RoomNumber:  sqDB.RoomNumber,
		Topic:       sqDB.Topic,
		Mode:        MeetingMode(sqDB.Mode),
		Distance:    sqDB.Distance,
		Type:        MeetingType(sqDB.Type),
		Status:      MeetingStatus(sqDB.Status),
		Location:    sqDB.Location,
		Description: sqDB.Description,
		CreateTime:  sqDB.CreateTime,
		StartTime:   sqDB.StartTime,
		EndTime:     sqDB.EndTime,
	}
	for _, v := range sqDB.Members {
		userInfo := &UserInfo{
			UserID:     v.UserID,
			UserName:   v.UserName,
			Role:       MeetingParticipantRole(v.Role),
			Status:     MeetingParticipantStatus(v.Status),
			UpdateTime: v.UpdateTime,
			JoinTime:   v.JoinTime,
			ExitTime:   v.ExitTime,
		}
		if IsAdmin(MeetingParticipantRole(v.Role)) {
			res.CreatorInfo = userInfo
		}
		if IsModerator(MeetingParticipantRole(v.Role)) {
			res.ModeratorInfo = userInfo
		}
		if !IsAdmin(MeetingParticipantRole(v.Role)) && !IsModerator(MeetingParticipantRole(v.Role)) && IsMember(MeetingParticipantRole(v.Role)) {
			res.Members = append(res.Members, userInfo)
		}
	}

	return res
}

func GetParticipantStatusText(status int) string {
	switch MeetingParticipantStatus(status) {
	case ParticipantStatusUnhandled:
		return ParticipantStatusUnhandledText
	case ParticipantStatusAccepted:
		return ParticipantStatusAcceptedText
	case ParticipantStatusRejected:
		return ParticipantStatusRejectedText
	default:
		return ParticipantStatusUnknownText
	}
}
func GetParticipantRole(roleText string) MeetingParticipantRole {
	switch roleText {
	case ParticipantRoleModeratorText:
		return ParticipantRoleModerator
	case ParticipantRoleAdminText:
		return ParticipantRoleAdmin
	case ParticipantRoleMemberText:
		return ParticipantRoleMember
	default:
		return MeetingParticipantRole(0)
	}
}

func GetParticipantRoles(role int) (result []string) {
	r := MeetingParticipantRole(role)
	if r&ParticipantRoleModerator != 0 {
		result = append(result, string(ParticipantRoleModeratorText))
	}
	if r&ParticipantRoleAdmin != 0 {
		result = append(result, string(ParticipantRoleAdminText))
	}
	if r&ParticipantRoleMember != 0 {
		result = append(result, string(ParticipantRoleMemberText))
	}
	return result
}

func GetParticipantRoleText(role int) string {
	switch MeetingParticipantRole(role) {
	case ParticipantRoleModerator:
		return ParticipantRoleModeratorText
	case ParticipantRoleAdmin:
		return ParticipantRoleAdminText
	case ParticipantRoleMember:
		return ParticipantRoleMemberText
	default:
		return ParticipantRoleUnknownText
	}
}

func GetStatus(statusText string) MeetingStatus {
	switch statusText {
	case MeetingStatusCreatedText:
		return MeetingStatusCreated
	case MeetingStatusStartedText:
		return MeetingStatusStarted
	case MeetingStatusEndedText:
		return MeetingStatusEnded
	case MeetingStatusCanceledText:
		return MeetingStatusCanceled
	default:
		return MeetingStatusUnknown
	}
}

func GetStatusText(status MeetingStatus) string {
	switch status {
	case MeetingStatusCreated:
		return MeetingStatusCreatedText
	case MeetingStatusStarted:
		return MeetingStatusStartedText
	case MeetingStatusEnded:
		return MeetingStatusEndedText
	case MeetingStatusCanceled:
		return MeetingStatusCanceledText
	default:
		return MeetingStatusUnknownText
	}
}

func GetType(typeText string) MeetingType {
	switch typeText {
	case MeetingTypeInstantText:
		return MeetingTypeInstant
	case MeetingTypeScheduledText:
		return MeetingTypeScheduled
	default:
		return MeetingTypeUnknown
	}
}

func GetTypeText(meetingType MeetingType) string {
	switch MeetingType(meetingType) {
	case MeetingTypeInstant:
		return MeetingTypeInstantText
	case MeetingTypeScheduled:
		return MeetingTypeScheduledText
	default:
		return MeetingTypeUnknownText
	}
}

func GetMode(modeText string) MeetingMode {
	switch modeText {
	case MeetingModeRomingText:
		return MeetingModeRoming
	case MeetingModeMeetingRoomText:
		return MeetingModeMeetingRoom
	case MeetingModeVirtualHumanText:
		return MeetingModeVirtualHuman
	default:
		return MeetingModeUnknown
	}
}

func GetModeText(mode MeetingMode) string {
	switch mode {
	case MeetingModeRoming:
		return MeetingModeRomingText
	case MeetingModeMeetingRoom:
		return MeetingModeMeetingRoomText
	case MeetingModeVirtualHuman:
		return MeetingModeVirtualHumanText
	default:
		return MeetingModeUnknownText
	}
}

// ValidateMeetingInfo 校验会议信息
func ValidateMeetingInfo(data *Meeting) (err error) {
	// 会议模式，参数校验
	if data.Mode == MeetingModeUnknown {
		return gerror.New(string(GetModeText(data.Mode)))
	}
	if data.Mode == MeetingModeRoming {
		if data.Distance <= 0 {
			return gerror.New("电子围栏半径不能小于等于0")
		}
	} else {
		data.Distance = 0
	}

	// 会议时间，参数校验
	if data.StartTime.Before(gtime.Now()) {
		return gerror.New("会议开始时间必须大于当前时间")
	}
	if data.StartTime.After(gtime.Now().AddDate(0, 3, 0)) {
		return gerror.New("会议预约时间不能超过三个月")
	}

	return nil
}

// 判断当前用户是否为管理员
func IsAdmin(role MeetingParticipantRole) bool {
	return role&ParticipantRoleAdmin != 0
}

// 判断当前用户是否为主持人
func IsModerator(role MeetingParticipantRole) bool {
	return role&ParticipantRoleModerator != 0
}

func IsMember(role MeetingParticipantRole) bool {
	return role&ParticipantRoleMember != 0
}
