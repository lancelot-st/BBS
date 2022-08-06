package redis

//redis尽量使用命名空间方便业务去拆分
const (
	KeyPostInfoHashPrefix  = "BBS:post:"
	KeyPostTimeZSet        = "BBS:post:time"
	KeyPostScoreZSet       = "BBS:post:score"
	KeyPostVotedZSetPrefix = "BBS:post:voted:"

	KeyCommunityPostSetPrefix = "BBS:community:" // set保存每个分区下帖子的id
)

const (
	OneWeekInSeconds         = 7 * 24 * 3600
	VoteScore        float64 = 432 // 每一票的值432分
	PostPerAge               = 20
)
