package objects

import "encoding/json"

type NotificationsFeedback struct {
	Attachments []WallWallpostAttachment `json:"attachments"`
	FromID      int                      `json:"from_id"`
	Geo         Geo                      `json:"geo"`
	ID          int                      `json:"id"`
	Likes       LikesInfo                `json:"likes"`
	Text        string                   `json:"text"`
	ToID        int                      `json:"to_id"`
}

type NotificationsNotification struct {
	Date     int                `json:"date"`
	Feedback json.RawMessage    `json:"feedback"`
	Parent   json.RawMessage    `json:"parent"`
	Reply    NotificationsReply `json:"reply"`
	Type     string             `json:"type"`
}

type NotificationsNotificationsComment struct {
	Date    int          `json:"date"`
	ID      int          `json:"id"`
	OwnerID int          `json:"owner_id"`
	Photo   Photo        `json:"photo"`
	Post    WallWallpost `json:"post"`
	Text    string       `json:"text"`
	Topic   BoardTopic   `json:"topic"`
	Video   Video        `json:"video"`
}

type NotificationsReply struct {
	Date string `json:"date"`
	ID   int    `json:"id"`
	Text string `json:"text"`
}
