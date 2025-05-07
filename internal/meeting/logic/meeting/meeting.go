package meeting

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"OnlineMeeting/api/v1/meeting"
	"OnlineMeeting/internal/consts"
	"OnlineMeeting/internal/meeting/dao"
	"OnlineMeeting/internal/meeting/model/do"
	"OnlineMeeting/internal/meeting/model/entity"
	"OnlineMeeting/internal/meeting/service"

	// "OnlineMeeting/internal/meeting/dao/do"

	"OnlineMeeting/library/liberr"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func init() {
	service.RegisterMeeting(New())
}

func New() service.IMeeting {
	return &Meeting{}
}

type Meeting struct {
}

func (s *Meeting) Create(ctx context.Context, req *meeting.CreateReq) (res *meeting.CreateRes, err error) {
	// 数据格式转换
	modelData := &entity.Meeting{
		RoomNumber:  s.GetRoomID(ctx),
		Topic:       req.Topic,
		Mode:        entity.GetMode(req.Mode),
		Distance:    req.Distance,
		Type:        entity.GetType(req.Type),
		Status:      entity.MeetingStatusCreated,
		Location:    req.Location,
		Description: req.Description,
		CreatorInfo: &entity.UserInfo{
			UserID:   req.CreatorID,
			UserName: req.CreatorName,
		},
		ModeratorInfo: &entity.UserInfo{
			UserID:   req.ModeratorID,
			UserName: req.ModeratorName,
		},
		CreateTime: gtime.Now(),
		StartTime:  req.StartTime,
	}
	if err = entity.ValidateMeetingInfo(modelData); err != nil {
		return
	}
	// 封装会议信息
	meetingData := do.TMeeting{
		RoomNumber:  modelData.RoomNumber,
		Topic:       modelData.Topic,
		Mode:        modelData.Mode,
		Distance:    modelData.Distance,
		Type:        modelData.Type,
		Status:      modelData.Status,
		Location:    modelData.Location,
		CreatorId:   modelData.CreatorInfo.UserID,
		Description: modelData.Description,
		CreateTime:  modelData.CreateTime,
		StartTime:   modelData.StartTime,
		EndTime:     modelData.EndTime,
	}
	// 封装管理员、主持人信息
	invitesData := make([]*do.TMeetingParticipant, 0)
	if modelData.CreatorInfo.UserID == modelData.ModeratorInfo.UserID { // 创建者和主持人是同一人
		invitesData = append(invitesData, &do.TMeetingParticipant{
			MRoomNumber: modelData.RoomNumber,
			UserId:      modelData.CreatorInfo.UserID,
			UserName:    modelData.CreatorInfo.UserName,
			Role:        entity.ParticipantRoleAdmin | entity.ParticipantRoleModerator | entity.ParticipantRoleMember, // 主持人 + 管理员 + 成员
			Status:      entity.ParticipantStatusUnhandled,
		})
	} else { // 创建者和主持人不是同一人, 创建者默认为管理员
		invitesData = append(invitesData, &do.TMeetingParticipant{
			MRoomNumber: modelData.RoomNumber,
			UserId:      modelData.CreatorInfo.UserID,
			UserName:    modelData.CreatorInfo.UserName,
			Role:        entity.ParticipantRoleAdmin | entity.ParticipantRoleMember, // 管理员 + 成员
			Status:      entity.ParticipantStatusUnhandled,
		})
		invitesData = append(invitesData, &do.TMeetingParticipant{
			MRoomNumber: modelData.RoomNumber,
			UserId:      modelData.ModeratorInfo.UserID,
			UserName:    modelData.ModeratorInfo.UserName,
			Role:        entity.ParticipantRoleModerator | entity.ParticipantRoleMember, // 主持人 + 成员
			Status:      entity.ParticipantStatusUnhandled,
		})
	}
	// 封装成员信息
	for _, v := range req.MemberInfos {
		invitesData = append(invitesData, &do.TMeetingParticipant{
			MRoomNumber: modelData.RoomNumber,
			UserId:      v.UserID,
			UserName:    v.UserName,
			Role:        entity.ParticipantRoleMember,
			Status:      entity.ParticipantStatusUnhandled,
		})
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.TMeeting.Ctx(ctx).TX(tx).Insert(meetingData)
			liberr.ErrIsNil(ctx, err, "会议信息添加失败")
			_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).Insert(invitesData)
			liberr.ErrIsNil(ctx, err, "参会人员信息添加失败")
		})
		return err
	})

	res = &meeting.CreateRes{
		RoomNumber:    modelData.RoomNumber,
		Topic:         modelData.Topic,
		CreatorName:   modelData.CreatorInfo.UserName,
		Type:          entity.GetTypeText(modelData.Type),
		Mode:          entity.GetModeText(modelData.Mode),
		Distance:      modelData.Distance,
		StartTime:     modelData.StartTime,
		ModeratorName: modelData.ModeratorInfo.UserName,
		Description:   modelData.Description,
	}

	return
}

func (s *Meeting) GetByRoomNumber(ctx context.Context, roomNumber string) (res *meeting.GetDetailsRes, err error) {
	meetingInfo, err := dao.TMeeting.GetByRoomID(ctx, roomNumber, nil)
	if err != nil {
		return
	}
	modelData := entity.ConvertToEntityModel(meetingInfo)

	res = &meeting.GetDetailsRes{}
	res.RoomNumber = modelData.RoomNumber
	res.Topic = modelData.Topic
	res.Mode = entity.GetModeText(modelData.Mode)
	res.Distance = modelData.Distance
	res.Type = entity.GetTypeText(modelData.Type)
	res.Status = entity.GetStatusText(modelData.Status)
	res.Location = modelData.Location
	res.Description = modelData.Description
	res.CreateTime = modelData.CreateTime
	res.StartTime = modelData.StartTime
	res.EndTime = modelData.EndTime
	for _, v := range modelData.Members {
		userInfo := &meeting.UserInfo{
			UserID:     v.UserID,
			UserName:   v.UserName,
			Roles:      entity.GetParticipantRoles(int(v.Role)),
			Status:     entity.GetParticipantStatusText(int(v.Status)),
			AcceptTime: v.AcceptTime,
			JoinTime:   v.JoinTime,
			ExitTime:   v.ExitTime,
		}
		res.MemberInfos = append(res.MemberInfos, userInfo)
	}
	res.CreatorInfo = &meeting.UserInfo{
		UserID:     modelData.CreatorInfo.UserID,
		UserName:   modelData.CreatorInfo.UserName,
		Roles:      entity.GetParticipantRoles(int(modelData.CreatorInfo.Role)),
		Status:     entity.GetParticipantStatusText(int(modelData.CreatorInfo.Status)),
		AcceptTime: modelData.CreatorInfo.AcceptTime,
		JoinTime:   modelData.CreatorInfo.JoinTime,
		ExitTime:   modelData.CreatorInfo.ExitTime,
	}
	res.ModeratorInfo = &meeting.UserInfo{
		UserID:     modelData.ModeratorInfo.UserID,
		UserName:   modelData.ModeratorInfo.UserName,
		Roles:      entity.GetParticipantRoles(int(modelData.ModeratorInfo.Role)),
		Status:     entity.GetParticipantStatusText(int(modelData.ModeratorInfo.Status)),
		AcceptTime: modelData.ModeratorInfo.AcceptTime,
		JoinTime:   modelData.ModeratorInfo.JoinTime,
		ExitTime:   modelData.ModeratorInfo.ExitTime,
	}
	if res.CreatorInfo.UserID == res.ModeratorInfo.UserID {
		res.Count = len(res.MemberInfos) + 1
	} else {
		res.Count = len(res.MemberInfos) + 2
	}

	return
}

// TODO: 这个方法应该加上分页参数
func (s *Meeting) GetByScope(ctx context.Context, req *meeting.ListReq, scope string) (res *meeting.ListRes, err error) {
	roomIDs := make([]string, 0)
	result, err := dao.TMeetingParticipant.Ctx(ctx).
		Fields(dao.TMeetingParticipant.Columns().MRoomNumber).
		Where(dao.TMeetingParticipant.Columns().UserId, req.UserID).
		Array()
	if err != nil {
		return nil, gerror.New(fmt.Sprintf("Logic.ListAll: %v", err.Error()))
	}
	for _, v := range result {
		roomIDs = append(roomIDs, v.Val().(string))
	}
	if len(roomIDs) == 0 {
		return
	}

	m := dao.TMeeting.Ctx(ctx).Fields(dao.TMeeting.Columns().RoomNumber).WhereIn(dao.TMeeting.Columns().RoomNumber, roomIDs)
	if scope == "history" {
		m = m.Where(dao.TMeeting.Columns().Status, "3").
			WhereOr(dao.TMeeting.Columns().Status, "4").
			WhereGT(dao.TMeeting.Columns().CreateTime, gtime.Now().AddDate(0, -3, 0))
	} else if scope == "future" {
		m = m.Where(dao.TMeeting.Columns().Status, "1").
			WhereOr(dao.TMeeting.Columns().Status, "2")
	} else {
		return
	}
	total, err := m.Count()
	if err != nil {
		return
	}

	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	result, err = m.Page(req.PageNum, req.PageSize).OrderDesc(dao.TMeeting.Columns().CreateTime).Array()
	if err != nil {
		return
	}
	ids := make([]string, 0)
	for _, v := range result {
		ids = append(ids, v.Val().(string))
	}
	if len(ids) == 0 {
		return
	}

	meetingInfos := []*meeting.GetDetailsRes{}
	for _, id := range ids {
		info, err := s.GetByRoomNumber(ctx, id)
		if err != nil {
			return nil, err
		}
		meetingInfos = append(meetingInfos, info)
	}

	res = &meeting.ListRes{
		Meetings: meetingInfos,
		PageRes: &meeting.PageRes{
			CurrentPage: req.PageNum,
			Total:       total,
		},
	}

	return
}

func (s *Meeting) Edit(ctx context.Context, req *meeting.EditReq) (err error) {
	// 数据格式转换
	newData := &entity.Meeting{
		RoomNumber: req.RoomNumber,
		Topic:      req.Topic,
		StartTime:  req.StartTime,
		Mode:       entity.GetMode(req.Mode),
		Distance:   req.Distance,
		Location:   req.Location,
		Type:       entity.GetType(req.Type),
		ModeratorInfo: &entity.UserInfo{
			UserID:   req.ModeratorID,
			UserName: req.ModeratorName,
		},
		Description: req.Description,
	}
	if err = entity.ValidateMeetingInfo(newData); err != nil {
		return
	}

	result, err := dao.TMeeting.GetByRoomID(ctx, req.RoomNumber, nil)
	if err != nil {
		return
	}
	oldData := entity.ConvertToEntityModel(result)
	if oldData.Status == entity.MeetingStatusEnded || oldData.Status == entity.MeetingStatusCanceled {
		return gerror.New("会议已结束/取消，无法修改")
	}

	// 封装会议信息
	updateData := do.TMeeting{}
	if newData.Topic != oldData.Topic {
		updateData.Topic = newData.Topic
	}
	if newData.StartTime != oldData.StartTime {
		updateData.StartTime = newData.StartTime
	}
	if newData.Mode != oldData.Mode {
		updateData.Mode = newData.Mode
	}
	if newData.Distance != oldData.Distance {
		updateData.Distance = newData.Distance
	}
	if newData.Type != oldData.Type {
		updateData.Type = newData.Type
	}
	if newData.Location != oldData.Location {
		updateData.Location = newData.Location
	}
	if newData.Description != oldData.Description {
		updateData.Description = newData.Description
	}
	// 封装参会人员信息
	insertData := make([]*do.TMeetingParticipant, 0)
	deleteData := make([]string, 0)
	var updateCreatorInfo map[string]interface{}
	var updateModeratorInfo map[string]interface{}
	var updateDataInfo map[string]interface{}
	// 主持人变更处理
	if newData.ModeratorInfo.UserID != oldData.ModeratorInfo.UserID {
		if newData.ModeratorInfo.UserID == oldData.CreatorInfo.UserID { // 判断新主持人是否为创建人
			// 创建人成为主持人
			// 1、原创建人 添加 主持人角色
			updateCreatorInfo = map[string]interface{}{
				dao.TMeetingParticipant.Columns().Role: oldData.CreatorInfo.Role | entity.ParticipantRoleModerator,
			}
			// 2、原主持人 移除 主持人角色
			updateModeratorInfo = map[string]interface{}{
				dao.TMeetingParticipant.Columns().Role: oldData.ModeratorInfo.Role &^ entity.ParticipantRoleModerator,
			}
		} else { // 新主持人不是创建人
			var userInfo *entity.UserInfo
			// 判断新主持人是否在成员列表已存在
			for index, v := range oldData.Members {
				if v.UserID == newData.ModeratorInfo.UserID {
					oldData.Members = append(oldData.Members[:index], oldData.Members[index+1:]...)
					userInfo = v
					break
				}
			}
			log.Printf("%+v\n", oldData.Members)
			if userInfo != nil {
				if entity.IsAdmin(oldData.CreatorInfo.Role) &&
					entity.IsModerator(oldData.CreatorInfo.Role) { // 如果创建者既为管理员、也为主持人，移除创建者的 主持人角色
					updateCreatorInfo = map[string]interface{}{
						dao.TMeetingParticipant.Columns().Role: oldData.CreatorInfo.Role &^ entity.ParticipantRoleModerator,
					}
				} else { // 原主持人 移除 主持人角色
					updateModeratorInfo = map[string]interface{}{
						dao.TMeetingParticipant.Columns().Role: oldData.ModeratorInfo.Role &^ entity.ParticipantRoleModerator,
					}
				}
				// 更新
				updateDataInfo = map[string]interface{}{
					dao.TMeetingParticipant.Columns().Role: userInfo.Role | entity.ParticipantRoleModerator,
				}
			} else {
				// 1、原主持人 移除 主持人角色
				updateModeratorInfo = map[string]interface{}{
					dao.TMeetingParticipant.Columns().Role: entity.ParticipantRoleMember,
				}
				// 2、插入新主持人
				insertData = append(insertData, &do.TMeetingParticipant{
					MRoomNumber: newData.RoomNumber,
					UserId:      newData.ModeratorInfo.UserID,
					UserName:    newData.ModeratorInfo.UserName,
					Role:        entity.ParticipantRoleModerator | entity.ParticipantRoleMember,
					Status:      entity.ParticipantStatusUnhandled,
				})
			}
		}
	}
	// 成员信息变更处理
	m := make(map[string]string)
	for _, v := range req.MemberInfos {
		m[v.UserID] = v.UserName
	}
	for _, v := range oldData.Members {
		if _, ok := m[v.UserID]; !ok {
			deleteData = append(deleteData, v.UserID)
		} else {
			delete(m, v.UserID)
		}
	}
	for k, v := range m {
		insertData = append(insertData, &do.TMeetingParticipant{
			MRoomNumber: req.RoomNumber,
			UserId:      k,
			UserName:    v,
			Role:        entity.ParticipantRoleMember,
			Status:      entity.ParticipantStatusUnhandled,
		})
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.TMeeting.Ctx(ctx).TX(tx).
				Where(dao.TMeeting.Columns().RoomNumber, req.RoomNumber).
				Update(updateData)
			if err != nil {
				return
			}
			if len(deleteData) > 0 {
				_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).
					Where(dao.TMeetingParticipant.Columns().MRoomNumber, req.RoomNumber).
					WhereIn(dao.TMeetingParticipant.Columns().UserId, deleteData).
					Delete()
				if err != nil {
					return
				}
			}
			if len(insertData) > 0 {
				_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).
					Insert(insertData)
				if err != nil {
					return
				}
			}
			if updateCreatorInfo != nil {
				_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).
					Where(dao.TMeetingParticipant.Columns().UserId, oldData.CreatorInfo.UserID).
					Update(updateCreatorInfo)
				if err != nil {
					return
				}
			}
			if updateModeratorInfo != nil {
				_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).
					Where(dao.TMeetingParticipant.Columns().UserId, oldData.ModeratorInfo.UserID).
					Update(updateModeratorInfo)
				if err != nil {
					return
				}
			}
			if updateDataInfo != nil {
				_, err = dao.TMeetingParticipant.Ctx(ctx).TX(tx).
					Where(dao.TMeetingParticipant.Columns().UserId, req.ModeratorID).
					Update(updateDataInfo)
				if err != nil {
					return
				}
			}
		})
		return err
	})

	return
}

func (s *Meeting) ListAll(ctx context.Context, req *meeting.ListAllReq) (res *meeting.ListRes, err error) {
	log.Println(req.PageReq)
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res = &meeting.ListRes{}
	m := dao.TMeeting.Ctx(ctx).
		Fields(dao.TMeeting.Columns().RoomNumber).
		WhereGT(dao.TMeeting.Columns().CreateTime, gtime.Now().AddDate(0, -3, 0))
	total, err := m.Count()
	if err != nil {
		return
	}
	roomIDs := make([]string, 0)
	result, err := m.Page(req.PageNum, req.PageSize).
		OrderDesc(dao.TMeeting.Columns().CreateTime).
		Array()
	if err != nil {
		return nil, gerror.New(fmt.Sprintf("Logic.ListAll: %v", err.Error()))
	}
	for _, v := range result {
		roomIDs = append(roomIDs, v.Val().(string))
	}
	if len(roomIDs) == 0 {
		return
	}

	for _, roomID := range roomIDs {
		mInfo, err := s.GetByRoomNumber(ctx, roomID)
		if err != nil {
			return nil, gerror.New(fmt.Sprintf("Logic.ListAll: %v", err.Error()))
		}
		res.Meetings = append(res.Meetings, mInfo)
	}

	res.PageRes = &meeting.PageRes{
		Total:       total,
		CurrentPage: req.PageNum,
	}

	return
}

func (s *Meeting) UpdateStatusByRoomNumber(ctx context.Context, req *meeting.UpdateStatusReq) (err error) {
	if entity.MeetingStatus(req.Status) == entity.MeetingStatusUnknown {
		return gerror.New("状态值错误")
	}
	if entity.MeetingStatus(req.Status) == entity.MeetingStatusCreated {
		return gerror.New("不能将会议修改为已创建")
	}

	var data g.Map = g.Map{}
	data[dao.TMeeting.Columns().Status] = req.Status
	if entity.MeetingStatus(req.Status) == entity.MeetingStatusEnded {
		data[dao.TMeeting.Columns().EndTime] = gtime.Now()
	}

	_, err = dao.TMeeting.Ctx(ctx).Where(dao.TMeeting.Columns().RoomNumber, req.RoomNumber).Update(data)
	if err != nil {
		return
	}

	return
}

func (s *Meeting) RemoveParticipants(ctx context.Context, req *meeting.RemoveParticipantsReq) (err error) {
	if len(req.UserIDs) == 0 {
		return
	}

	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.TMeetingParticipant.Ctx(ctx).
			Where(dao.TMeetingParticipant.Columns().MRoomNumber, req.RoomNumber).
			WhereIn(dao.TMeetingParticipant.Columns().UserId, req.UserIDs).
			Delete()
	})

	return
}

func (s *Meeting) GetRoomID(ctx context.Context) (roomID string) {
	for {
		roomID = s.generateRoomID()
		num, err := dao.TMeeting.Ctx(ctx).Count(fmt.Sprintf("room_number = %s", roomID))
		if err == nil && num == 0 {
			break
		}
		log.Println("生成会议室ID失败，正在重试...")
		time.Sleep(time.Millisecond * 500)
	}

	return roomID
}
func (s *Meeting) generateRoomID() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%09d", rand.Intn(1000000000))
}

func (s *Meeting) UpdateParticipantStatus(ctx context.Context, req *meeting.UpdateParticipantStatusReq) (res *meeting.UpdateParticipantStatusRes, err error) {
	if req.Status != int(entity.ParticipantStatusAccepted) &&
		req.Status != int(entity.ParticipantStatusRejected) &&
		req.Status != int(entity.ParticipantStatusUnhandled) {
		return nil, gerror.New("未知的状态")
	}
	data := make(map[string]interface{})
	data[dao.TMeetingParticipant.Columns().Status] = req.Status
	data[dao.TMeetingParticipant.Columns().AcceptTime] = gtime.Now()
	_, err = dao.TMeetingParticipant.Ctx(ctx).
		Where(dao.TMeetingParticipant.Columns().MRoomNumber, req.RoomNumber).
		Where(dao.TMeetingParticipant.Columns().UserId, req.UserID).
		Update(data)

	return
}

func (s *Meeting) Join(ctx context.Context, req *meeting.JoinMeetingReq) (res *meeting.JoinMeetingRes, err error) {
	data := make(map[string]interface{})
	data[dao.TMeetingParticipant.Columns().JoinTime] = gtime.Now()
	_, err = dao.TMeetingParticipant.Ctx(ctx).
		Where(dao.TMeetingParticipant.Columns().MRoomNumber, req.RoomNumber).
		Where(dao.TMeetingParticipant.Columns().UserId, req.UserID).
		Update(data)

	return
}
func (s *Meeting) Exit(ctx context.Context, req *meeting.ExitMeetingReq) (res *meeting.ExitMeetingRes, err error) {
	data := make(map[string]interface{})
	data[dao.TMeetingParticipant.Columns().ExitTime] = gtime.Now()
	_, err = dao.TMeetingParticipant.Ctx(ctx).
		Where(dao.TMeetingParticipant.Columns().MRoomNumber, req.RoomNumber).
		Where(dao.TMeetingParticipant.Columns().UserId, req.UserID).
		Update(data)

	return
}
