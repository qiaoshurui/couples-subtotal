package dto

import "time"

type AddDynamic struct {
	Content string `json:"content"`
	UserId  int64  `json:"userId"`
	Status  int8   `json:"status"`
}

type DeleteDynamic struct {
	Id        int64 `json:"id"`
	IsDeleted int8  `json:"isDeleted"` //是否删除

}

type UpdateDynamic struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Status  int8   `json:"status"`
}

// todo 考虑通用的排序方式
// 考虑 定义通用的分页、排序结构体

type GetDynamicList struct {
	Page     int    `json:"page"`     // 页码
	PageSize int    `json:"pageSize"` // 每页大小
	Content  string `json:"content"`
}

type DynamicListInfo struct {
	Id           int64  `json:"id"`
	Content      string `json:"content"`
	UserId       int64  `json:"userId"`       //动态发布者id
	UserNickName string `json:"userNickName"` //动态发布者的昵称
}

type SimpleDynamicDetail struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	UserId    int64     `json:"userId"`    //动态发布者id
	CreatedAt time.Time `json:"createdAt"` //动态发布时间
}
