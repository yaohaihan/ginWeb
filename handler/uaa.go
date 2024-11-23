package handler

import (
	"golang.org/x/net/context"
)

type UaaHandler interface {
	Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error
}

type uaa struct {
}

func NewUaaHandler() UaaHandler {
	return &uaa{}
}

type LoginReq struct {
}

type LoginRsp struct {
	Id   string `json:"id"`
	Name string `json:"name"` //结构体标签 它表示这个字段在 JSON 格式中对应的键名为 "name"
}

func (u *uaa) Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) error {
	rsp.Id, rsp.Name = "1", "小明"
	return nil
}
