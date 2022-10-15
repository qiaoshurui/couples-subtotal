package dto

import "time"

// todo 全部json统一使用 小驼峰
type AddDynamic struct {
	Content string `json:"content"`
	UserId  int64  `json:"user_id"`
	Status  int8   `json:"status"`
}
type DeleteDynamic struct {
	Id        int64 `json:"id"`
	IsDeleted int8  `json:"is_deleted"` //是否删除

}
type UpdateDynamic struct {
	Id        int64     `json:"id"`
	Content   string    `json:"content"`
	Status    int8      `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
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
	UserId       int64  `json:"user_id"`        //动态发布者id
	UserNickName string `json:"user_nick_name"` //动态发布者的昵称
}
