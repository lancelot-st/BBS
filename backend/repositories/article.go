package repositories

import (
	"BBS/model"
	"BBS/model/redis"
	"log"
)

var ArticleRepository = newArticleRepository()

func newArticleRepository() *articleRepository {
	return &articleRepository{}
}

type articleRepository struct {
}

func (a *articleRepository) CreateArticle(p *model.Article) error {
	db := model.DB()

	db.Table("post").Create(&p)

	err := redis.CreatArticle(p.PostID, p.CommunityID)
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func (a *articleRepository) GetArticleByID(id int64) (Data *model.Article, err error) {
	db := model.DB()
	db.Where("id = ?", id).Table("post").Find(&Data)
	return Data, nil
}

//获取帖子列表

func (a *articleRepository) GetArticleList(page, size int64) (articles []*model.Article, err error) {
	db := model.DB()
	var total int64
	db.Count(&total).Table("post")
	offset := (page - 1) * size
	err = db.Order("create_time desc").Offset(int(offset)).Limit(int(size)).Table("post").Find(&articles).Error
	if err != nil {
		return
	}
	return articles, nil
}

func (a *articleRepository) GetArticleList2(p *model.ArticleList) (articles []*model.Article, err error) {
	//去redis查询id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		log.Println("ID列表获取失败")
		return nil, err
	}
	//返回的数据按照指定顺序排序
	if len(ids) == 0 {
		return nil, nil
	}
	return a.GetArticleListByIDs(ids)

}

func (a *articleRepository) GetArticleListByIDs(ids []string) (articleList []*model.Article, err error) {
	db := model.DB()
	log.Println(ids)
	for _, id := range ids {
		p := new(model.Article)
		err = db.Where("post_id = ?", id).Table("post").Find(&p).Error
		articleList = append(articleList, p)
	}

	return articleList, err
}
