package main

import (
	"fmt"
	"go-gf-issue/internal/dao"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

// https://github.com/gogf/gf/pull/2213
func main() {
	type JsonItem struct {
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	}
	type Test struct {
		Id   string      `json:"id,omitempty"`
		Json []*JsonItem `json:"json,omitempty"`
	}

	ctx := gctx.New()
	var list []*Test
	err := dao.Test.Ctx(ctx).Scan(&list)
	if err != nil {
		return
	}
	g.Dump(list[0])
	fmt.Println(list[0].Json)
	fmt.Println(len(list[0].Json))
	g.Dump(list[1])
	fmt.Println(list[1].Json)
	fmt.Println(len(list[1].Json))
	//cmd.Main.Run(gctx.New())
}
