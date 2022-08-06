package services

import (
	"BBS/backend/repositories"
	"BBS/model"
)

var CommunityService = newCommunityService()

func newCommunityService() *communityService {
	return &communityService{}
}

type communityService struct {
}

func (c *communityService) GetCommunityDetailById(id int64) (Community *model.CommunityDetail, err error) {
	return repositories.CommunityRepository.GetCommunityDetailById(id)
}

func (c *communityService) GetCommunityNameById(idStr string) (Community *model.Community, err error) {
	return repositories.CommunityRepository.GetCommunityNameById(idStr)
}

func (c *communityService) GetCommunityList() (Community []*model.Community, err error) {
	return repositories.CommunityRepository.GetCommunityList()
}
