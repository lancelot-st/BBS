package model

type VoteData struct {
	PostID    string `json:"post_id" `
	Direction int8   `json:"direction" validate:"required,oneof=1 0 -1"` //赞成票是+1，反对票是-1,validate字段之间不能有空格
}
