package services

import (
	"BBS/backend/repositories"
	"BBS/model"
	"strconv"
)

var VoteService = newVoteService()

func newVoteService() *voteService {
	return &voteService{}
}

type voteService struct {
}

//投票限制

func (v *voteService) VoteForArticle(userID int64, p *model.VoteData) error {
	postID := p.PostID
	userid := strconv.FormatInt(userID, 10)
	errVote := repositories.VoteRepository.VoteForArticle(userid, postID, float64(p.Direction))
	return errVote
}
