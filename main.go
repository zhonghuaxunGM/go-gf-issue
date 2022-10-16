package main

import (
	"fmt"
	"go-gf-issue/internal/dao"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
)

// CREATE TABLE `test_1665845241938231900` (
// 	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
// 	`passport` varchar(45) DEFAULT NULL,
// 	`password` char(32) DEFAULT NULL,
// 	`nickname` varchar(45) DEFAULT NULL,
// 	`create_time` timestamp NULL DEFAULT NULL,
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
// INSERT INTO `gf_issue`.`test_1665845241938231900` (`id`, `passport`, `password`, `nickname`, `create_time`) VALUES ('11', 'user_11', 'pass_11', 'name_11', '2018-10-24 10:00:00');

func main() {
	issue2181()
	// issue2105()
	//cmd.Main.Run(gctx.New())
}
func issue2181() {
	startTime := gtime.New(1664640000)
	fmt.Println(startTime) // 2022-10-02 00:00:00

	endTime := gtime.New(1665308819)
	fmt.Println(endTime) //2022-10-09 17:46:59

	cnt, err := dao.Test1665845241938231900.Ctx(gctx.New()).Where("create_time >= ? and create_time <= ?", startTime, endTime).Count()
	fmt.Println(cnt, err)
}

func issue2105() {
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
}
