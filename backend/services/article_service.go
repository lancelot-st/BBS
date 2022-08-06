package services

import (
	"BBS/backend/repositories"
	"BBS/model"
	"BBS/model/redis"
	"log"
)

var ArticleService = newArticleService()

func newArticleService() *articleService {
	return &articleService{}
}

type articleService struct {
}

func (a *articleService) CreateArticleServ(p *model.Article) error {

	err := repositories.ArticleRepository.CreateArticle(p)
	if err != nil {
		return err
	}
	return nil
}

func (a *articleService) GetArticleByIDServ(id int64) (Data *model.ArticleDetail, err error) {
	data, errData := repositories.ArticleRepository.GetArticleByID(id)

	if errData != nil {
		log.Fatalf("获取帖子失败")
	}
	user, errUser := repositories.UserRepository.GetUserByUid(data.AuthorID)
	if errUser != nil {
		log.Fatal("用户获取失败")
	}
	community, errCom := repositories.CommunityRepository.GetCommunityDetailById(data.AuthorID)
	if errCom != nil {
		log.Fatal("社区详细信息获取失败")
	}

	var ArticleDetail = &model.ArticleDetail{
		Article:         data,
		CommunityDetail: community,
		AuthorName:      user.Username,
		VoteNum:         2,
	}
	return ArticleDetail, nil
}

func (a *articleService) GetArticleListServ(page, size int64) (List []*model.ArticleDetail, err error) {
	Posts, errP := repositories.ArticleRepository.GetArticleList(page, size)
	if errP != nil {
		log.Fatal("获取帖子列表失败")
		return nil, errP
	}

	data := make([]*model.ArticleDetail, 0, len(Posts))

	for _, post := range Posts {
		user, errUser := repositories.UserRepository.GetUserByUid(post.AuthorID)
		if errUser != nil {
			log.Println("用户获取失败")
			continue
		}
		community, errCom := repositories.CommunityRepository.GetCommunityDetailById(post.CommunityID)
		if errCom != nil {
			log.Println("社区详细信息获取失败")
			continue
		}
		ArticleDetail := &model.ArticleDetail{
			Article:         post,
			AuthorName:      user.Username,
			CommunityDetail: community,
			VoteNum:         2,
		}
		data = append(data, ArticleDetail)
	}
	return data, nil
}

func (a *articleService) GetArticleListServ2(p *model.ArticleList) (List []*model.ArticleDetail, err error) {
	Posts, errP := repositories.ArticleRepository.GetArticleList2(p)
	if errP != nil {
		log.Fatal("获取帖子列表失败")
		return nil, errP
	}
	ids, err := redis.GetPostIDsInOrder(p)

	VoteData, err := redis.GetPostVoteData(ids)

	data := make([]*model.ArticleDetail, 0, len(Posts))

	for idx, post := range Posts {
		user, errUser := repositories.UserRepository.GetUserByUid(post.AuthorID)
		if errUser != nil {
			log.Println("用户获取失败")
			continue
		}
		community, errCom := repositories.CommunityRepository.GetCommunityDetailById(post.CommunityID)
		if errCom != nil {
			log.Println("社区获取失败")
			continue
		}
		ArticleDetail := &model.ArticleDetail{
			Article:         post,
			AuthorName:      user.Username,
			CommunityDetail: community,
			VoteNum:         VoteData[idx],
		}
		data = append(data, ArticleDetail)
	}
	return data, nil
}
