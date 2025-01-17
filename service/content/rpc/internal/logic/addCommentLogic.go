package logic

import (
	"context"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/content/rpc/internal/model"
	"errors"

	"fmt"

	"log"

	"time"

	"doushen_by_liujun/service/content/rpc/internal/svc"
	"doushen_by_liujun/service/content/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCommentLogic) AddComment(in *pb.AddCommentReq) (*pb.AddCommentResp, error) {

	//1. 雪花算法生成 id
	snowflake, err1 := util.NewSnowflake(3)
	if err1 != nil {
		return nil, errors.New("rpc-AddComment-新增评论，snowflake生成id失败")
	}
	if err := l.svcCtx.KqPusherClient.Push("content_rpc_addCommentLogic_AddComment_NewSnowflake_false"); err != nil {
		log.Fatal(err)
	}
	snowId := snowflake.Generate()
	//2. 新增评论信息到 comment 表项
	_, err := l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
		Id:         snowId,
		UserId:     in.UserId,
		VideoId:    in.VideoId,
		Content:    in.Content,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		IsDelete:   0,
	})
	if err != nil {
		return nil, errors.New("rpc-AddComment-新增评论数据失败")
	}

	fmt.Println("【rpc-AddComment-新增评论数据成功】")

	if err := l.svcCtx.KqPusherClient.Push("content_rpc_addCommentLogic_AddComment_Insert_false"); err != nil {
		log.Fatal(err)
	}
	logx.Error("rpc-AddComment-新增评论数据成功")
	if err := l.svcCtx.KqPusherClient.Push("content_rpc_addCommentLogic_AddComment_success"); err != nil {
		log.Fatal(err)
	}

	return &pb.AddCommentResp{}, nil
}
