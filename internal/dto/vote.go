package dto

type Vote struct {
	VoteId    string `json:"vote_id"`
	ProductID string `json:"product_id"`
	Score     int32  `json:"score"`
}

type CreateVoteReq struct {
	SessionID string `json:"session_id"`
	Vote      Vote   `json:"vote"`
}

type CreateVotesReq struct {
	SessionID string `db:"session_id"`
	Votes     []Vote `json:"votes"`
}

type UpdateVoteReq struct {
	SessionID string `json:"session_id"`
	Vote      Vote   `json:"vote"`
}

type UpdateVotesReq struct {
	SessionID string `json:"session_id"`
	Votes     []Vote `json:"votes"`
}
