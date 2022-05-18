package api

import "github.com/gogf/gf/v2/frame/g"

type DetailUserReq struct {
	g.Meta `path:"/user/:id" method:"get" tags:"user" summary:"get users"`
	ID     int64 `json:"id" desc:"主键" v:"required"`
}

type DetailUserRes struct {
}