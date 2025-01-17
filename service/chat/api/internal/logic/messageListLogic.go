package logic

import (
	"context"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/chat/api/internal/svc"
	"doushen_by_liujun/service/chat/api/internal/types"
	"doushen_by_liujun/service/chat/rpc/pb"
	"fmt"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageListLogic) MessageList(req *types.MessageChatReq) (*types.MessageChatReqResp, error) {
	var resp *types.MessageChatReqResp
	// parse token
	res, err := util.ParseToken(req.Token)
	if err != nil {
		if err = l.svcCtx.KqPusherClient.Push("chat_api_messageListLogic_MessageList_ParseToken_false"); err != nil {
			log.Fatal(err)
		}
		resp = &types.MessageChatReqResp{
			StatusCode:  1,
			StatusMsg:   "fail to parse token",
			MessageList: nil,
		}
		return resp, fmt.Errorf("fail to parse token, error = %s", err)
	}

	// get params
	userId := res.UserID
	toUserId := req.ToUserId

	request := pb.GetChatMessageByIdReq{
		UserId:     userId,
		ToUserId:   toUserId,
		PreMsgTime: req.PreMsgTime,
	}
	// get chat messages
	message, err := l.svcCtx.ChatRpcClient.GetChatMessageById(l.ctx, &request)
	if err != nil {
		if err = l.svcCtx.KqPusherClient.Push("chat_api_messageListLogic_MessageList_GetChatMessageById_false"); err != nil {
			log.Fatal(err)
		}
		resp = &types.MessageChatReqResp{
			StatusCode:  1,
			StatusMsg:   "fail to get chat message",
			MessageList: nil,
		}
		return resp, fmt.Errorf("fail to get chat message, error = %s", err)
	}

	var messages []types.Message
	for _, item := range message.MessageList {
		msg := types.Message{
			Id:         item.Id,
			ToUserId:   item.ToUserId,
			FromUserId: item.FromUserId,
			Content:    item.Content,
			CreateTime: item.CreateTime,
		}
		messages = append(messages, msg)
	}

	resp = &types.MessageChatReqResp{
		StatusCode:  0,
		StatusMsg:   "get chat messages successfully",
		MessageList: messages,
	}
	if err = l.svcCtx.KqPusherClient.Push("chat_api_messageListLogic_MessageList_success"); err != nil {
		log.Fatal(err)
	}
	return resp, nil
}
