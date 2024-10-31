package objects

type BoardTopic struct {
	Comments  int     `json:"comments"`
	Created   int     `json:"created"`
	CreatedBy int     `json:"created_by"`
	ID        int     `json:"id"`
	IsClosed  BoolInt `json:"is_closed"`
	IsFixed   BoolInt `json:"is_fixed"`
	Title     string  `json:"title"`
	Updated   int     `json:"updated"`
	UpdatedBy int     `json:"updated_by"`
}

type BoardTopicComment struct {
	Attachments []WallCommentAttachment `json:"attachments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Text        string                  `json:"text"`
	Likes       LikesInfo               `json:"likes"`
	CanEdit     BoolInt                 `json:"can_edit"`
}

type BoardTopicPoll struct {
	AnswerID int           `json:"answer_id"`
	Answers  []PollsAnswer `json:"answers"`
	Created  int           `json:"created"`
	IsClosed BoolInt       `json:"is_closed"`
	OwnerID  int           `json:"owner_id"`
	PollID   int           `json:"poll_id"`
	Question string        `json:"question"`
	Votes    int           `json:"votes"`
}
