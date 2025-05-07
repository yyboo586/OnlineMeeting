package service

import (
	"OnlineMeeting/api/v1/meeting"
	"OnlineMeeting/internal/meeting/model"
	"context"
)

type IMeeting interface {
	Create(ctx context.Context, req *meeting.CreateReq) (res *meeting.CreateRes, err error)
	GetByRoomNumber(ctx context.Context, roomNumber string) (res *meeting.GetDetailsRes, err error)
	GetByScope(ctx context.Context, req *meeting.ListReq, scope string) (res *meeting.ListRes, err error)
	ListAll(ctx context.Context, req *meeting.ListAllReq) (res *meeting.ListRes, err error)
	// Edit(ctx context.Context, req *meeting.EditReq) (err error)

	// 更新会议状态
	UpdateStatusByRoomNumber(ctx context.Context, roomNumber string, status int) (err error)
	// 校验status是否在合法的枚举范围内
	CheckMeetingStatusValid(status int) (valid bool)
	// 检查会议是否存在
	CheckMeetingExists(ctx context.Context, roomNumber string) (exists bool, err error)

	// 邀请人员加入会议
	InviteParticipants(ctx context.Context, roomNumber string, userInfos []*model.UserInfo) (err error)
	// 移除参会人员
	RemoveParticipant(ctx context.Context, roomNumber string, userID string) (err error)
	// 校验status是否在合法的枚举范围内
	CheckParticipantStatusValid(status int) (valid bool)
	// 仅处理用户行为：接受会议邀请/加入会议/退出会议
	HandleUserAction(ctx context.Context, actionInfo *model.HandleUserAction) (err error)
}

var localMeeting IMeeting

func Meeting() IMeeting {
	if localMeeting == nil {
		panic("implement not found for interface IMeeting, forgot register?")
	}
	return localMeeting
}

func RegisterMeeting(i IMeeting) {
	localMeeting = i
}
