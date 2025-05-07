package model

type UserAction int

const (
	_            UserAction = iota
	ActionInvite            // 会议邀请处理（忽略/接受/拒绝）
	ActionJoin              // 加入会议
	ActionExit              // 退出会议
)

type HandleUserAction struct {
	RoomNumber string
	UserID     string
	Action     UserAction
	Status     int
}
