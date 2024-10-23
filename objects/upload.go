package objects

import "io"

type UploadFile struct {
	Name      string    `json:"name"`
	FieldName string    `json:"field_name,omitempty"`
	Data      io.Reader `json:"data"`
}

func NewUploadFile(name string, data io.Reader) *UploadFile {
	return &UploadFile{
		Name:      name,
		FieldName: "file",
		Data:      data,
	}
}

func NewUploadFilePhoto(name string, data io.Reader) *UploadFile {
	return &UploadFile{
		Name:      name,
		FieldName: "photo",
		Data:      data,
	}
}

func NewUploadFileVideo(name string, data io.Reader) *UploadFile {
	return &UploadFile{
		Name:      name,
		FieldName: "video",
		Data:      data,
	}
}

func NewUploadFileImage(name string, data io.Reader) *UploadFile {
	return &UploadFile{
		Name:      name,
		FieldName: "image",
		Data:      data,
	}
}
