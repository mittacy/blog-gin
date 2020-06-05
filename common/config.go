package common

import "github.com/crazychat/blog-gin/model"

var InitAdmin = model.Admin {
	Name:      "admin",
	Password:  Encryption("admin"),
	Views:		 7156,
	Cname:     "陈铭涛",
	Introduce: "就读佛山大学 - 大三 - 计算机系",
	Github:    "https://github.com/crazychat",
	Mail:      "mail@mittacy.com",
	Bilibili:  "https://space.bilibili.com/384942135",
}
