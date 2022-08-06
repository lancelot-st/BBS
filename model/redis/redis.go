package redis

import (
	"BBS/model"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

var (
	client *redis.Client
	Nil    = redis.Nil
)

type SliceCmd = redis.SliceCmd
type StringStringMapCmd = redis.StringStringMapCmd

// Init 初始化连接
func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("Redis.host"), viper.GetInt("Redis.port")),
		Password: viper.GetString("Redis.password"), // no password set
		DB:       viper.GetInt("Redis.db"),          // use default DB
		PoolSize: viper.GetInt("Redis.pool_size"),
	})
	_, err = client.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func CreatArticle(postID int64, communityID int64) error {

	// 帖子时间
	pipeline := client.TxPipeline()
	pipeline.ZAdd(KeyPostTimeZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	communityKey := KeyCommunityPostSetPrefix + strconv.Itoa(int(communityID))
	pipeline.SAdd(communityKey, postID)
	// 帖子分数
	pipeline.ZAdd(KeyPostScoreZSet, redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err

}

func CreatePostTime(PostID int64) error {
	_, err := client.SAdd(KeyPostTimeZSet, time.Now().Unix()).Result()
	if err != nil {
		return err
	}

	return nil
}

func GetPostVoteData(ids []string) ([]int64, error) {

	//利用事务减少请求
	pipeline := client.TxPipeline()
	for _, id := range ids {
		key := KeyPostVotedZSetPrefix + id
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	data := make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, err
}

//获取redis储存的分数最高的size个帖子

func GetPostIDsInOrder(p *model.ArticleList) ([]string, error) {
	key := KeyPostTimeZSet
	if p.Order == "Score" {
		key = KeyPostScoreZSet
	}

	start := p.Page * p.Size //页码
	end := start + p.Size - 1
	return client.ZRevRange(key, start, end).Result() //排序获取分数最高的的从start开始到end结束的key_value
}

func GetRedis() *redis.Client {
	return client
}

func Close() {
	_ = client.Close()
}
