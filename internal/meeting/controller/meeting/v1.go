package controller

import (
	"OnlineMeeting/api/v1/meeting"
	"context"

	meetingService "OnlineMeeting/internal/meeting/service"
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
func (c *meetingsController) Edit(ctx context.Context, req *meeting.EditReq) (res *meeting.EditRes, err error) {
	err = meetingService.Meeting().Edit(ctx, req)
	return
}

// UpdateMeetingStatus 更改会议状态
func (c *meetingsController) UpdateMeetingStatus(ctx context.Context, req *meeting.UpdateStatusReq) (res *meeting.UpdateStatusRes, err error) {
	err = meetingService.Meeting().UpdateStatusByRoomNumber(ctx, req)
	return
}

func (c *meetingsController) RemoveParticipants(ctx context.Context, req *meeting.RemoveParticipantsReq) (res *meeting.RemoveParticipantsRes, err error) {
	err = meetingService.Meeting().RemoveParticipants(ctx, req)
	return
}

func (c *meetingsController) UpdateParticipantStatus(ctx context.Context, req *meeting.UpdateParticipantStatusReq) (res *meeting.UpdateParticipantStatusRes, err error) {
	res, err = meetingService.Meeting().UpdateParticipantStatus(ctx, req)
	return
}
func (c *meetingsController) JoinMeeting(ctx context.Context, req *meeting.JoinMeetingReq) (res *meeting.JoinMeetingRes, err error) {
	res, err = meetingService.Meeting().Join(ctx, req)
	return
}
func (c *meetingsController) ExitMeeting(ctx context.Context, req *meeting.ExitMeetingReq) (res *meeting.ExitMeetingRes, err error) {
	res, err = meetingService.Meeting().Exit(ctx, req)
	return
}
