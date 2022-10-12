package dto

type CouplesInfo struct {
	UserId        int64  `json:"user_d"`
	UserNickName  string `json:"user_nick_name"`
	UserHeaderImg string `json:"user_header_img"`

	LoverId        int64  `json:"lover_id"`
	LoverNickName  string `json:"lover_nick_name"`
	LoverHeaderImg string `json:"lover_header_img"`

	MemorialDay int `json:"memorial_day"` //纪念日已过去天数
}
