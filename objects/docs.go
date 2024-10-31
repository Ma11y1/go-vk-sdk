package objects

import (
	"fmt"
)

type Document struct {
	AccessKey  string          `json:"access_key"`
	Date       int             `json:"date"`
	Ext        string          `json:"ext"`
	ID         int             `json:"id"`
	IsLicensed BoolInt         `json:"is_licensed"`
	OwnerID    int             `json:"owner_id"`
	Preview    DocumentPreview `json:"preview"`
	Size       int             `json:"size"`
	Title      string          `json:"title"`
	Type       int             `json:"type"`
	URL        string          `json:"url"`
	DocumentPreviewAudioMessage
	DocumentPreviewGraffiti
}

func (d *Document) ToAttachment() string {
	return fmt.Sprintf("d%d_%d", d.OwnerID, d.ID)
}

type DocumentPreview struct {
	Photo        DocumentPreviewPhoto        `json:"photo"`
	Graffiti     DocumentPreviewGraffiti     `json:"graffiti"`
	Video        DocumentPreviewVideo        `json:"video"`
	AudioMessage DocumentPreviewAudioMessage `json:"audio_message"`
}

type DocumentPreviewPhoto struct {
	Sizes []DocumentPreviewPhotoSizes `json:"sizes"`
}

func (p *DocumentPreviewPhoto) MaxSize() (maxPhotoSize DocumentPreviewPhotoSizes) {
	var maxSize float64

	for _, photoSize := range p.Sizes {
		size := photoSize.Height * photoSize.Width
		if size > maxSize {
			maxSize = size
			maxPhotoSize = photoSize
		}
	}

	return
}

func (p *DocumentPreviewPhoto) MinSize() (minPhotoSize DocumentPreviewPhotoSizes) {
	var minSize float64

	for _, photoSize := range p.Sizes {
		size := photoSize.Height * photoSize.Width
		if size < minSize || minSize == 0 {
			minSize = size
			minPhotoSize = photoSize
		}
	}

	return
}

type DocumentPreviewPhotoSizes struct {
	Height float64 `json:"height"`
	Src    string  `json:"src"`
	Type   string  `json:"type"`
	Width  float64 `json:"width"`
}

type DocumentPreviewGraffiti struct {
	Src    string `json:"src"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DocumentPreviewVideo struct {
	FileSize int    `json:"file_size"`
	Height   int    `json:"height"`
	Src      string `json:"src"`
	Width    int    `json:"width"`
}

type DocumentPreviewAudioMessage struct {
	Duration        int    `json:"duration"`
	Waveform        []int  `json:"waveform"`
	LinkOgg         string `json:"link_ogg"`
	LinkMp3         string `json:"link_mp3"`
	Transcript      string `json:"transcript"`
	TranscriptState string `json:"transcript_state"`
}

type DocumentTypes struct {
	Count int    `json:"count"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
}

type DocumentUploadResponse struct {
	File string `json:"file"`
}
