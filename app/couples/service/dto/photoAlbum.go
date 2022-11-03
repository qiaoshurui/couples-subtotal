package dto

type AddPhotoAlbum struct {
	Name string `json:"name"`
	Type int8   `json:"type"` //(0 情侣相册；1 个人相册)
}
type AlbumListRes struct {
	Page     int  `json:"page"`     // 页码
	PageSize int  `json:"pageSize"` // 每页大小
	Type     int8 `json:"type"`     //(0 情侣相册；1 个人相册)
}
type AlbumListReq struct {
	Id       int64  `json:"id"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	OwnerId  int64  `json:"owner_id"`
	Type     string `json:"type"`
	AlbumUrl string `json:"album_url"`
}
