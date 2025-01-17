// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"doushen_by_liujun/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddFollowsReq             = pb.AddFollowsReq
	AddFollowsResp            = pb.AddFollowsResp
	CheckIsFollowReq          = pb.CheckIsFollowReq
	CheckIsFollowResp         = pb.CheckIsFollowResp
	CheckUserReq              = pb.CheckUserReq
	CheckUserResp             = pb.CheckUserResp
	DelFollowsReq             = pb.DelFollowsReq
	DelFollowsResp            = pb.DelFollowsResp
	Follows                   = pb.Follows
	GetFollowersByIdReq       = pb.GetFollowersByIdReq
	GetFollowersByIdResp      = pb.GetFollowersByIdResp
	GetFollowersCountByIdReq  = pb.GetFollowersCountByIdReq
	GetFollowersCountByIdResp = pb.GetFollowersCountByIdResp
	GetFollowsByIdReq         = pb.GetFollowsByIdReq
	GetFollowsByIdResp        = pb.GetFollowsByIdResp
	GetFriendsByIdReq         = pb.GetFriendsByIdReq
	GetFriendsByIdResp        = pb.GetFriendsByIdResp
	GetUserByIdReq            = pb.GetUserByIdReq
	GetUserByIdResp           = pb.GetUserByIdResp
	GetUserListByIdListReq    = pb.GetUserListByIdListReq
	GetUserListByIdListResp   = pb.GetUserListByIdListResp
	GetUserinfoByIdReq        = pb.GetUserinfoByIdReq
	GetUserinfoByIdResp       = pb.GetUserinfoByIdResp
	GetUsersByIdsReq          = pb.GetUsersByIdsReq
	GetUsersByIdsResp         = pb.GetUsersByIdsResp
	SaveUserReq               = pb.SaveUserReq
	SaveUserResp              = pb.SaveUserResp
	Userinfo                  = pb.Userinfo
	Usersinfo                 = pb.Usersinfo

	User interface {
		AddFollows(ctx context.Context, in *AddFollowsReq, opts ...grpc.CallOption) (*AddFollowsResp, error)
		DelFollows(ctx context.Context, in *DelFollowsReq, opts ...grpc.CallOption) (*DelFollowsResp, error)
		GetFollowsById(ctx context.Context, in *GetFollowsByIdReq, opts ...grpc.CallOption) (*GetFollowsByIdResp, error)
		SaveUser(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserResp, error)
		CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserResp, error)
		GetUserinfoById(ctx context.Context, in *GetUserinfoByIdReq, opts ...grpc.CallOption) (*GetUserinfoByIdResp, error)
		GetUsersByIds(ctx context.Context, in *GetUsersByIdsReq, opts ...grpc.CallOption) (*GetUsersByIdsResp, error)
		GetFollowersById(ctx context.Context, in *GetFollowersByIdReq, opts ...grpc.CallOption) (*GetFollowersByIdResp, error)
		GetFollowersCountById(ctx context.Context, in *GetFollowersCountByIdReq, opts ...grpc.CallOption) (*GetFollowersCountByIdResp, error)
		CheckIsFollow(ctx context.Context, in *CheckIsFollowReq, opts ...grpc.CallOption) (*CheckIsFollowResp, error)
		GetFriendsById(ctx context.Context, in *GetFriendsByIdReq, opts ...grpc.CallOption) (*GetFriendsByIdResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error)
		GetUserListByIdList(ctx context.Context, in *GetUserListByIdListReq, opts ...grpc.CallOption) (*GetUserListByIdListResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) AddFollows(ctx context.Context, in *AddFollowsReq, opts ...grpc.CallOption) (*AddFollowsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.AddFollows(ctx, in, opts...)
}

func (m *defaultUser) DelFollows(ctx context.Context, in *DelFollowsReq, opts ...grpc.CallOption) (*DelFollowsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.DelFollows(ctx, in, opts...)
}

func (m *defaultUser) GetFollowsById(ctx context.Context, in *GetFollowsByIdReq, opts ...grpc.CallOption) (*GetFollowsByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetFollowsById(ctx, in, opts...)
}

func (m *defaultUser) SaveUser(ctx context.Context, in *SaveUserReq, opts ...grpc.CallOption) (*SaveUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.SaveUser(ctx, in, opts...)
}

func (m *defaultUser) CheckUser(ctx context.Context, in *CheckUserReq, opts ...grpc.CallOption) (*CheckUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.CheckUser(ctx, in, opts...)
}

func (m *defaultUser) GetUserinfoById(ctx context.Context, in *GetUserinfoByIdReq, opts ...grpc.CallOption) (*GetUserinfoByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserinfoById(ctx, in, opts...)
}

func (m *defaultUser) GetUsersByIds(ctx context.Context, in *GetUsersByIdsReq, opts ...grpc.CallOption) (*GetUsersByIdsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUsersByIds(ctx, in, opts...)
}

func (m *defaultUser) GetFollowersById(ctx context.Context, in *GetFollowersByIdReq, opts ...grpc.CallOption) (*GetFollowersByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetFollowersById(ctx, in, opts...)
}

func (m *defaultUser) GetFollowersCountById(ctx context.Context, in *GetFollowersCountByIdReq, opts ...grpc.CallOption) (*GetFollowersCountByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetFollowersCountById(ctx, in, opts...)
}

func (m *defaultUser) CheckIsFollow(ctx context.Context, in *CheckIsFollowReq, opts ...grpc.CallOption) (*CheckIsFollowResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.CheckIsFollow(ctx, in, opts...)
}

func (m *defaultUser) GetFriendsById(ctx context.Context, in *GetFriendsByIdReq, opts ...grpc.CallOption) (*GetFriendsByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetFriendsById(ctx, in, opts...)
}

func (m *defaultUser) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserByIdResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUser) GetUserListByIdList(ctx context.Context, in *GetUserListByIdListReq, opts ...grpc.CallOption) (*GetUserListByIdListResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserListByIdList(ctx, in, opts...)
}
