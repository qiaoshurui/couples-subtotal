package dto

type AddPhoto struct {
	AlbumId int64 `json:"albumId"` //相册id
}
type PhotoList struct {
	Page     int   `json:"page"`     // 页码
	PageSize int   `json:"pageSize"` // 每页大小
	AlbumId  int64 `json:"albumId"`  //相册id
}
type PhotoListRes struct {
	Id      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	AlbumId int64  `json:"albumId"`
	ImgUrl  string `json:"imgUrl"`
}
