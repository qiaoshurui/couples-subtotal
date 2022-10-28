package dto

type AddPhotoAlbum struct {
	Name string `json:"name"`
	Type string `json:"type"` //(0 情侣相册；1 个人相册)
}
