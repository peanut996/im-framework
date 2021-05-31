// Package api
// @Title  api_request.go
// @Description record some defined request struct.
// @Author  peanut996
// @Update  peanut996  2021/5/22 21:57
package api

import "framework/api/model"

type PushChatRequest struct {
	Message *model.ChatMessage `json:"message"`
	Target  string             `json:"target"`
}
type ChatRequest struct {
	//From sender user id
	From    string      `json:"from"`
	To      string      `json:"to,omitempty"`
	Time    int64       `json:"time,omitempty"`
	Type    string      `json:"type"`
	Content interface{} `json:"content"`
}

type UserRequest struct {
	UID string `json:"uid"`
}

type AuthRequest struct {
	Token string `json:"token"`
}

type FriendRequest struct {
	FriendA string `json:"friendA"`
	FriendB string `json:"friendB"`
}

type GroupRequest struct {
	UID        string `json:"uid,omitempty"`
	GroupID    string `json:"groupID,omitempty"`
	GroupName  string `json:"groupName,omitempty"`
	GroupAdmin string `json:"groupAdmin,omitempty"`
}

//LoadRequest 用户初始化请求
type LoadRequest struct {
	UID string `json:"uid"`
}

//FindRequest 模糊查找请求
type FindRequest struct {
	Account   string `json:"account,omitempty"`
	GroupName string `json:"groupName,omitempty"`
}

//InviteRequest 用户邀请进群请求
type InviteRequest struct {
	Friends []string `json:"friends,omitempty"`
	GroupID string   `json:"groupID,omitempty"`
}

//PullRequest 分页拉取消息请求
type PullRequest struct {
	UID      string `json:"uid,omitempty"`
	GroupID  string `json:"groupID,omitempty"`
	FriendID string `json:"friendID,omitempty"`
	Current  int64  `json:"current"`
	PageSize int64  `json:"pageSize"`
}

//UpdateRequest 更新用户信息
type UpdateRequest struct {
	UID      string `json:"uid"`
	Password string `json:"password,omitempty"`
	Account  string `json:"account,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
