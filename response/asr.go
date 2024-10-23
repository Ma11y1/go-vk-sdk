package response

import "go-vk-sdk/constants"

// Doc: https://dev.vk.com/ru/method/asr

type AsrCheckStatusResponse struct {
	BaseResponse
	Response struct {
		ID string `json:"id"` // Identifier of the created task for processing audio recordings in UUID format.
		// Status of the audio recording processing task. Possible values:
		// processing—the audio recording is being processed.
		// finished—processing of the audio recording is finished.
		// internal_error — internal errors of the VK speech recognition service.
		// transcoding_error — error in transcoding the audio recording into the internal format. Try downloading the audio in a different supported format.
		// recognition_error—speech recognition error, difficulty in recognition. Try speaking more clearly or reducing background noise.
		Status constants.AsrStatus `json:"status"`
		Text   string              `json:"text"` // Decoding the text. It matters if the status parameter is finished.
	} `json:"response"`
}

type AsrGetUploadURLResponse struct {
	BaseResponse
	Response struct {
		UploadURL string `json:"upload_url"` // link to the server address for downloading audio recordings
	} `json:"response"`
}

type AsrProcessResponse struct {
	BaseResponse
	Response struct {
		TaskID string `json:"task_id"` // identifier of the created task for processing audio recordings in UUID format
	} `json:"response"`
}
