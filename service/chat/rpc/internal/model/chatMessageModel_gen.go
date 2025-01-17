// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	chatMessageFieldNames          = builder.RawFieldNames(&ChatMessage{})
	chatMessageRows                = strings.Join(chatMessageFieldNames, ",")
	chatMessageRowsExpectAutoSet   = strings.Join(stringx.Remove(chatMessageFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	chatMessageRowsWithPlaceHolder = strings.Join(stringx.Remove(chatMessageFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLiujunChatChatMessageIdPrefix = "cache:liujunChat:chatMessage:id:"
)

type (
	chatMessageModel interface {
		Insert(ctx context.Context, data *ChatMessage) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ChatMessage, error)
		Update(ctx context.Context, data *ChatMessage) error
		Delete(ctx context.Context, id int64) error
		GetChatMsgByIds(ctx context.Context, userId, toUserId int64)(*[]*ChatMessage, error)
	}

	defaultChatMessageModel struct {
		sqlc.CachedConn
		table string
	}

	ChatMessage struct {
		Id         int64     `db:"id"`          // 主键
		UserId     int64     `db:"user_id"`     // 发送人id
		ToUserId   int64     `db:"to_user_id"`  // 接收人id
		Message    string    `db:"message"`     // 消息内容
		CreateTime time.Time `db:"create_time"` // 该条记录创建时间
		UpdateTime time.Time `db:"update_time"` // 该条最后一次更新时间
		IsDelete   int64     `db:"is_delete"`   // 逻辑删除
	}
)

func newChatMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultChatMessageModel {
	return &defaultChatMessageModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`chat_message`",
	}
}

func (m *defaultChatMessageModel) withSession(session sqlx.Session) *defaultChatMessageModel {
	return &defaultChatMessageModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`chat_message`",
	}
}

func (m *defaultChatMessageModel) GetChatMsgByIds(ctx context.Context, userId, toUserId int64) (*[]*ChatMessage, error) {
	var result []*ChatMessage
	querySql := fmt.Sprintf("SELECT * FROM %s WHERE `user_id` = ? AND `to_user_id` = ? AND `is_delete` = 0 ORDER BY `create_time` DESC", m.table)
	if err := m.QueryRowsNoCacheCtx(ctx, &result, querySql, userId, toUserId); err != nil {
		return nil, fmt.Errorf("fail to getChatMessage by ids, error = %s", err)
	}
	fmt.Println(result)
	return &result, nil
}

func (m *defaultChatMessageModel) Delete(ctx context.Context, id int64) error {
	liujunChatChatMessageIdKey := fmt.Sprintf("%s%v", cacheLiujunChatChatMessageIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, liujunChatChatMessageIdKey)
	return err
}

func (m *defaultChatMessageModel) FindOne(ctx context.Context, id int64) (*ChatMessage, error) {
	liujunChatChatMessageIdKey := fmt.Sprintf("%s%v", cacheLiujunChatChatMessageIdPrefix, id)
	var resp ChatMessage
	err := m.QueryRowCtx(ctx, &resp, liujunChatChatMessageIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatMessageRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultChatMessageModel) Insert(ctx context.Context, data *ChatMessage) (sql.Result, error) {
	liujunChatChatMessageIdKey := fmt.Sprintf("%s%v", cacheLiujunChatChatMessageIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, chatMessageRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.UserId, data.ToUserId, data.Message, data.IsDelete)
	}, liujunChatChatMessageIdKey)
	return ret, err
}

func (m *defaultChatMessageModel) Update(ctx context.Context, data *ChatMessage) error {
	liujunChatChatMessageIdKey := fmt.Sprintf("%s%v", cacheLiujunChatChatMessageIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, chatMessageRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.Message, data.IsDelete, data.Id)
	}, liujunChatChatMessageIdKey)
	return err
}

func (m *defaultChatMessageModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLiujunChatChatMessageIdPrefix, primary)
}

func (m *defaultChatMessageModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", chatMessageRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultChatMessageModel) tableName() string {
	return m.table
}
