package response

import "go-vk-sdk/objects"

// Doc: https://dev.vk.com/ru/method/notes

type NotesAddResponse struct {
	BaseResponse
	Response struct {
		Items []objects.NewsfeedNewsfeedItem `json:"items"`
		objects.UsersAndGroups
		NextFrom string `json:"next_from"`
	} `json:"response"`
}

type NotesCreateCommentResponse struct {
	BaseResponse
	Response struct {
		Members []int `json:"members"`
		Groups  []int `json:"groups"`
	} `json:"response"`
}

type NotesDeleteResponse struct {
	BaseResponse
	Response struct {
		objects.UsersAndGroups
	} `json:"response"`
}

type NotesDeleteCommentResponse struct {
	BaseResponse
	Response struct {
		Items []objects.NewsfeedNewsfeedItem `json:"items"`
		objects.UsersAndGroups
		NextFrom string `json:"next_from"`
	} `json:"response"`
}

type NotesEditResponse struct {
	BaseResponse
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID        int    `json:"id"`
			Title     string `json:"title"`
			NoReposts int    `json:"no_reposts"`
			SourceIDs []int  `json:"source_ids"`
		} `json:"items"`
	} `json:"response"`
}

type NotesEditCommentResponse struct {
	BaseResponse
	Response struct {
		Count int                        `json:"count"`
		Items []objects.WallWallpostToID `json:"items"`
	} `json:"response"`
}

type NotesGetResponse struct {
	BaseResponse
	Response struct {
		Count int                 `json:"count"`
		Items []objects.NotesNote `json:"items"`
	} `json:"response"`
}

type NotesGetByIDResponse struct {
	BaseResponse
	Response objects.NotesNote `json:"response"`
}

type NotesGetCommentsResponse struct {
	BaseResponse
	Response struct {
		Count int                        `json:"count"`
		Items []objects.NotesNoteComment `json:"items"`
	} `json:"response"`
}

type NotesRestoreCommentResponse struct {
	BaseResponse
	Response struct {
		Items      []objects.NewsfeedNewsfeedItem `json:"items"`
		Profiles   []objects.UserFull             `json:"profiles"`
		Groups     []objects.GroupFull            `json:"groups"`
		NextOffset string                         `json:"next_offset"`
		NextFrom   string                         `json:"next_from"`
	} `json:"response"`
}
