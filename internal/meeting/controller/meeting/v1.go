package controller

import (
	"OnlineMeeting/api/v1/meeting"
	"context"

	"OnlineMeeting/internal/meeting/model"
	meetingService "OnlineMeeting/internal/meeting/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type meetingsController struct {
}

var MeetingController = new(meetingsController)

// Create 添加会议信息表
func (c *meetingsController) Create(ctx context.Context, req *meeting.CreateReq) (res *meeting.CreateRes, err error) {
	res, err = meetingService.Meeting().Create(ctx, req)
	return
}

// Get 获取会议详细信息
func (c *meetingsController) Get(ctx context.Context, req *meeting.GetDetailsReq) (res *meeting.GetDetailsRes, err error) {
	res = new(meeting.GetDetailsRes)
	res, err = meetingService.Meeting().GetByRoomNumber(ctx, req.RoomNumber)
	return
}

func (c *meetingsController) GetHistory(ctx context.Context, req *meeting.ListHistoryReq) (res *meeting.ListRes, err error) {
	res = new(meeting.ListRes)
	res, err = meetingService.Meeting().GetByScope(ctx, &req.ListReq, "history")
	return
}

// Get 获取历史会议信息
func (c *meetingsController) GetFuture(ctx context.Context, req *meeting.ListFutureReq) (res *meeting.ListRes, err error) {
	res = new(meeting.ListRes)
	res, err = meetingService.Meeting().GetByScope(ctx, &req.ListReq, "future")
	return
}

// ListAll 列表
func (c *meetingsController) ListAll(ctx context.Context, req *meeting.ListAllReq) (res *meeting.ListRes, err error) {
	res = new(meeting.ListRes)
	res, err = meetingService.Meeting().ListAll(ctx, req)
	return
}

// Edit 修改会议信息表
// func (c *meetingsController) Edit(ctx context.Context, req *meeting.EditReq) (res *meeting.EditRes, err error) {
// 	err = meetingService.Meeting().Edit(ctx, req)
// 	return
// }

// UpdateMeetingStatus 更改会议状态
func (c *meetingsController) UpdateMeetingStatus(ctx context.Context, req *meeting.UpdateStatusReq) (res *meeting.UpdateStatusRes, err error) {
	res = new(meeting.UpdateStatusRes)
	valid := meetingService.Meeting().CheckMeetingStatusValid(req.Status)
	if !valid {
		err = gerror.New("请求参数错误: Status参数值非法")
		return
	}

	err = meetingService.Meeting().UpdateStatusByRoomNumber(ctx, req.RoomNumber, req.Status)

	return
}

func (c *meetingsController) InviteParticipants(ctx context.Context, req *meeting.InviteParticipantsReq) (res *meeting.InviteParticipantsRes, err error) {
	res = new(meeting.InviteParticipantsRes)

	userInfos := make([]*model.UserInfo, 0)
	for _, v := range req.UserInfos {
		userInfos = append(userInfos, &model.UserInfo{
			ID:   v.UserID,
			Name: v.UserName,
		})
	}
	err = meetingService.Meeting().InviteParticipants(ctx, req.RoomNumber, userInfos)

	return
}

func (c *meetingsController) RemoveParticipants(ctx context.Context, req *meeting.RemoveParticipantsReq) (res *meeting.RemoveParticipantsRes, err error) {
	res = new(meeting.RemoveParticipantsRes)

	err = meetingService.Meeting().RemoveParticipant(ctx, req.RoomNumber, req.UserID)

	return
}

func (c *meetingsController) UpdateParticipantStatus(ctx context.Context, req *meeting.UpdateParticipantStatusReq) (res *meeting.UpdateParticipantStatusRes, err error) {
	res = new(meeting.UpdateParticipantStatusRes)
	valid := meetingService.Meeting().CheckParticipantStatusValid(req.Status)
	if !valid {
		err = gerror.New("请求参数错误: Status参数值非法")
		return
	}

	actionInfo := &model.HandleUserAction{
		RoomNumber: req.RoomNumber,
		UserID:     req.UserID,
		Action:     model.ActionInvite,
		Status:     req.Status,
	}
	err = meetingService.Meeting().HandleUserAction(ctx, actionInfo)

	return
}
func (c *meetingsController) JoinMeeting(ctx context.Context, req *meeting.JoinReq) (res *meeting.JoinRes, err error) {
	res = new(meeting.JoinRes)

	actionInfo := &model.HandleUserAction{
		RoomNumber: req.RoomNumber,
		UserID:     req.UserID,
		Action:     model.ActionJoin,
	}
	err = meetingService.Meeting().HandleUserAction(ctx, actionInfo)

	return
}
func (c *meetingsController) ExitMeeting(ctx context.Context, req *meeting.ExitReq) (res *meeting.ExitRes, err error) {
	res = new(meeting.ExitRes)

	actionInfo := &model.HandleUserAction{
		RoomNumber: req.RoomNumber,
		UserID:     req.UserID,
		Action:     model.ActionExit,
	}
	err = meetingService.Meeting().HandleUserAction(ctx, actionInfo)

	return
}
