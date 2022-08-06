package repositories

import (
	"BBS/model/redis"
	"errors"
	redis2 "github.com/go-redis/redis"
	"math"
	"time"
)

//投票功能的开发

type IVoteRepository interface {
	VoteForArticle(userID string, postID string, value float64) error
}

const (
	OneWeekSecond = 3600 * 24 * 7
	ScorePerVote  = 432 // 每一票占的值
)

var (
	ErrorVoteTimeout = errors.New("超过投票时间")
	ErrorVotedTwice  = errors.New("不可重复投相同的票")
)

/* PostVote 为帖子投票
投票分为四种情况：1.投赞成票 2.投反对票 3.取消投票 4.反转投票

记录文章参与投票的人
更新文章分数：赞成票要加分；反对票减分

v=1时，有两种情况
	1.之前没投过票，现在要投赞成票		--> 更新分数和投票记录		差值的绝对值：1  +432
	2.之前投过反对票，现在要改为赞成票	--> 更新分数和投票记录		差值的绝对值：2  +432*2
v=0时，有两种情况
	1.之前投过反对票，现在要取消			--> 更新分数和投票记录		差值的绝对值：1  +432
	2.之前投过赞成票，现在要取消			--> 更新分数和投票记录		差值的绝对值：1  -432
v=-1时，有两种情况
	1.之前没投过票，现在要投反对票		--> 更新分数和投票记录		差值的绝对值：1  -432
	2.之前投过赞成票，现在要改为反对票	--> 更新分数和投票记录		差值的绝对值：2  -432*2
*/

//投票限制
//每个帖子发表后的一星期内允许投票超过一个星期不再投票
//超过一个星期的投票数就不在储存在redis 内就将他储存在mysql里

var VoteRepository = newVoteRepository()

func newVoteRepository() *voteRepository {
	return &voteRepository{}
}

type voteRepository struct {
}

func (v *voteRepository) VoteForArticle(userID string, postID string, value float64) error {
	client := redis.GetRedis()
	postTime := client.ZScore(redis.KeyPostTimeZSet, postID).Val()

	if float64(time.Now().Unix())-postTime > OneWeekSecond {
		return ErrorVoteTimeout
	}
	ov := client.ZScore(redis.KeyPostVotedZSetPrefix+postID, userID).Val()

	var dir float64
	if value == ov {
		return ErrorVotedTwice
	}
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) //计算两次投票的差值

	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(redis.KeyPostScoreZSet, dir*diff*ScorePerVote, postID)
	if value == 0 {
		pipeline.ZRem(redis.KeyPostVotedZSetPrefix+postID, userID)

	} else {
		pipeline.ZAdd(redis.KeyPostVotedZSetPrefix+postID, redis2.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
