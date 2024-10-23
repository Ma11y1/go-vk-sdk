package objects

import (
	"fmt"
)

type NotesNote struct {
	CanComment     NumberFlagBool `json:"can_comment"` // Information whether current user can comment the note
	Comments       int            `json:"comments"`    // Comments number
	Date           int            `json:"date"`        // Date when the note has been created in Unixtime
	ID             int            `json:"id"`          // Note ID
	OwnerID        int            `json:"owner_id"`    // Note owner's ID
	Text           string         `json:"text"`        // Note text
	TextWiki       string         `json:"text_wiki"`   // Note text in wiki format
	Title          string         `json:"title"`       // Note title
	ViewURL        string         `json:"view_url"`    // url of the page with note preview
	ReadComments   int            `json:"read_comments"`
	PrivacyView    []interface{}  `json:"privacy_view"`
	PrivacyComment []interface{}  `json:"privacy_comment"`
}

func (note NotesNote) ToAttachment() string {
	return fmt.Sprintf("note%d_%d", note.OwnerID, note.ID)
}

type NotesNoteComment struct {
	Date    int    `json:"date"`
	ID      int    `json:"id"`
	Message string `json:"message"`
	NID     int    `json:"nid"`
	OID     int    `json:"oid"`
	ReplyTo int    `json:"reply_to"`
	UID     int    `json:"uid"`
}
