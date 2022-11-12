package dto

type AddPhotoAlbum struct {
	Name string `json:"name"`
	Type int8   `json:"type"` //(0 情侣相册；1 个人相册)
}
type AlbumListReq struct {
	Page     int  `json:"page"`     // 页码
	PageSize int  `json:"pageSize"` // 每页大小
	Type     int8 `json:"type"`     //(0 情侣相册；1 个人相册)
}
type AlbumList struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Type     int8   `json:"type"`
	AlbumUrl string `json:"album_url"`
}
type AlbumListRes struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Type       int8   `json:"type"`
	AlbumUrl   string `json:"album_url"`
	PhotoCount int64  `json:"photo_count"` //照片个数
}
