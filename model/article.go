package model

import "time"

type Article struct {

	//排序讲究内存对齐
	PostID      int64     `json:"post_id" sql:"post_id" gorm:"unique" form:"post_id"`
	AuthorID    int64     `json:"author_id" sql:"author_id" gorm:"unique" form:"author_id"`
	CommunityID int64     `json:"community_id" sql:"community_id"`
	Status      int32     `json:"status" sql:"status"`
	Title       string    `json:"title" sql:"title" form:"title"`
	Content     string    `json:"content" sql:"content"`
	CreateTime  time.Time `json:"create_time" sql:"create_time"`
}

type ArticleDetail struct {
	*Article                            // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
	AuthorName       string             `json:"author_name"`
	VoteNum          int64              `json:"vote_num"`
}

type ArticleList struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
