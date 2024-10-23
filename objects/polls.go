package objects

import (
	"fmt"
)

type PollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

type PollsPoll struct {
	AnswerID      int             `json:"answer_id"`
	Answers       []PollsAnswer   `json:"answers"`
	Created       int             `json:"created"`
	ID            int             `json:"id"`
	OwnerID       int             `json:"owner_id"`
	Question      string          `json:"question"`
	Votes         int             `json:"votes"`
	AnswerIDs     []int           `json:"answer_ids"`
	EndDate       int             `json:"end_date"`
	Anonymous     NumberFlagBool  `json:"anonymous"`
	Closed        NumberFlagBool  `json:"closed"`
	IsBoard       NumberFlagBool  `json:"is_board"`
	CanEdit       NumberFlagBool  `json:"can_edit"`
	CanVote       NumberFlagBool  `json:"can_vote"`
	CanReport     NumberFlagBool  `json:"can_report"`
	CanShare      NumberFlagBool  `json:"can_share"`
	Multiple      NumberFlagBool  `json:"multiple"`
	DisableUnvote NumberFlagBool  `json:"disable_unvote"`
	Photo         Photo           `json:"photo"`
	AuthorID      int             `json:"author_id"`
	Background    PollsBackground `json:"background"`
	Friends       []PollsFriend   `json:"friends"`
	Profiles      []UserFull      `json:"profiles"`
	Groups        []GroupFull     `json:"groups"`
	EmbedHash     string          `json:"embed_hash"`
}

func (poll PollsPoll) ToAttachment() string {
	return fmt.Sprintf("poll%d_%d", poll.OwnerID, poll.ID)
}

type PollsFriend struct {
	ID int `json:"id"`
}

type PollsVoters struct {
	AnswerID int              `json:"answer_id"`
	Users    PollsVotersUsers `json:"users"`
}

type PollsVotersUsers struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}

type PollsVotersFields struct {
	AnswerID int                    `json:"answer_id"`
	Users    PollsVotersUsersFields `json:"users"`
}

type PollsVotersUsersFields struct {
	Count int        `json:"count"`
	Items []UserFull `json:"items"`
}

type PollsBackground struct {
	Type   string `json:"type"`
	Angle  int    `json:"angle"`
	Color  string `json:"color"`
	Points []struct {
		Position float64 `json:"position"`
		Color    string  `json:"color"`
	} `json:"points"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PollsPhoto struct {
	ID     int           `json:"id"`
	Color  string        `json:"color"`
	Images []PhotosImage `json:"images"`
}
