package dto

import "github.com/google/uuid"

type Vote struct {
	VoteId    uuid.UUID `json:"vote_id"`
	ProductID uuid.UUID `json:"product_id"`
	Score     int32     `json:"score"`
}

type CreateVoteReq struct {
	SessionID uuid.UUID `json:"session_id"`
	Vote      Vote      `json:"vote"`
}

type CreateVotesReq struct {
	SessionID uuid.UUID `db:"session_id"`
	Votes     []Vote    `json:"votes"`
}

type UpdateVoteReq struct {
	SessionID uuid.UUID `json:"session_id"`
	Vote      Vote      `json:"vote"`
}

type UpdateVotesReq struct {
	SessionID uuid.UUID `json:"session_id"`
	Votes     []Vote    `json:"votes"`
}
