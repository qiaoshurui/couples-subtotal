package dto

type AddPhotoAlbum struct {
	Name string `json:"name"`
	Type int8   `json:"type"` //(0 情侣相册；1 个人相册)
}
