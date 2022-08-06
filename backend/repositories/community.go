package repositories

import (
	"BBS/model"
)

//-------社区相关的函数------

var CommunityRepository = newCommunityRepository()

func newCommunityRepository() *communityRepository {
	return &communityRepository{}
}

type communityRepository struct {
}

func (c *communityRepository) GetCommunityDetailById(id int64) (Community *model.CommunityDetail, err error) {
	db := model.DB()
	db.Where("community_id = ?", id).Table("community").Find(&Community)
	return
}

func (c *communityRepository) GetCommunityNameById(idStr string) (Community *model.Community, err error) {
	db := model.DB()
	db.Where("community_id = ?", idStr).Table("community").Find(&Community)
	return Community, nil
}

func (c *communityRepository) GetCommunityList() (Community []*model.Community, err error) {
	Db := model.DB()
	Db.Table("community").Find(&Community)
	return
}
