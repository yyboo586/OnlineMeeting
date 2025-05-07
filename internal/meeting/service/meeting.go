package service

import (
	"OnlineMeeting/api/v1/meeting"
	"context"
)

type IMeeting interface {
	Create(ctx context.Context, req *meeting.CreateReq) (res *meeting.CreateRes, err error)
	GetByRoomNumber(ctx context.Context, roomNumber string) (res *meeting.GetDetailsRes, err error)
	GetByScope(ctx context.Context, req *meeting.ListReq, scope string) (res *meeting.ListRes, err error)
	ListAll(ctx context.Context, req *meeting.ListAllReq) (res *meeting.ListRes, err error)
	Edit(ctx context.Context, req *meeting.EditReq) (err error)
	UpdateStatusByRoomNumber(ctx context.Context, req *meeting.UpdateStatusReq) (err error)
	RemoveParticipants(ctx context.Context, req *meeting.RemoveParticipantsReq) (err error)
	UpdateParticipantStatus(ctx context.Context, req *meeting.UpdateParticipantStatusReq) (res *meeting.UpdateParticipantStatusRes, err error)
	Join(ctx context.Context, req *meeting.JoinMeetingReq) (res *meeting.JoinMeetingRes, err error)
	Exit(ctx context.Context, req *meeting.ExitMeetingReq) (res *meeting.ExitMeetingRes, err error)
}

var localMeeting IMeeting

func Meeting() IMeeting {
	if localMeeting == nil {
		panic("implement not found for interface ISqmeeting, forgot register?")
	}
	return localMeeting
}

func RegisterMeeting(i IMeeting) {
	localMeeting = i
}
