package controller

import (
	"context"
	"fmt"
	"go-gf-issue/api"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/container/gvar"
)

var UserInfo = User{}

type User struct {
}

func getURI(ctx context.Context, key string) *gvar.Var {
	return g.RequestFromCtx(ctx).Get(key)
}

func (e *User) Detail(ctx context.Context, req *api.DetailUserReq) (res *api.DetailUserRes, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%v", e)
			g.Log().Errorf(ctx, "User Detail: %s", err.Error())
		}
	}()
	fmt.Println(getURI(ctx, "id"))
	fmt.Println(getURI(ctx, "id").Int64())
	return nil, err
}
