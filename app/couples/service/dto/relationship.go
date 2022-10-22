package dto

type CouplesInfo struct {
	UserId        int64  `json:"userId"`
	UserNickName  string `json:"userNickName"`
	UserHeaderImg string `json:"userHeaderImg"`

	LoverId        int64  `json:"loverId"`
	LoverNickName  string `json:"loverNickName"`
	LoverHeaderImg string `json:"loverHeaderImg"`

	MemorialDay int `json:"memorialDay"` //纪念日已过去天数
}
type RelationshipBinding struct {
	RegistrationCode string `json:"registrationCode"`
	UserId           int64  `json:"userId"`
}
