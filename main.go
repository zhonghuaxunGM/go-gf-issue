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
}

func issue2181() {
	startTime := gtime.New(1664640000)
	fmt.Println(startTime) // 2022-10-02 00:00:00

	endTime := gtime.New(1665308819)
	fmt.Println(endTime) //2022-10-09 17:46:59
	type User struct {
		Id         int
		Passport   string
		Password   string
		NickName   string
		CreateTime gtime.Time
	}
	user := make([]User, 0)
	err := dao.Test1665845241938231900.Ctx(gctx.New()).Where("create_time >= ? and create_time <= ?", startTime, endTime).Scan(&user)
	fmt.Println(err)
	g.Dump(user)
}
