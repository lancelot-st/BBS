package model

import "time"

type Community struct {
	CommunityId int64  `gorm:"column:community_id;unique;" json:"CommunityId" form:"CommunityId"`
	Name        string `gorm:"column:community_name;size:32;unique;"json:"CommunityName" form:"CommunityName" `
}

type CommunityDetail struct {
	CommunityID   int64     `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"` // omitempty 当Introduction为空时不展示
	CreateTime    time.Time `json:"create_time" db:"create_time"`
}
