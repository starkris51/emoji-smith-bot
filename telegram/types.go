package telegram

type Update struct {
	UpdateID int      `json:"update_id"`
	Message  *Message `json:"message,omitempty"`
}

type Message struct {
	MessageID int         `json:"message_id"`
	Chat      Chat        `json:"chat"`
	Text      string      `json:"text,omitempty"`
	Photo     []PhotoSize `json:"photo,omitempty"`
	Video     *Video      `json:"video,omitempty"`
	Animation *Animation  `json:"animation,omitempty"`
	Sticker   *Sticker    `json:"sticker,omitempty"`
}

type Chat struct {
	ID int64 `json:"id"`
}

type Sticker struct {
	FileID       string `json:"file_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	IsAnimated   bool   `json:"is_animated"`
	IsVideo      bool   `json:"is_video"`
	Emoji        string `json:"emoji,omitempty"`
	SetName      string `json:"set_name,omitempty"`
	FileSize     int    `json:"file_size"`
	MimeType     string `json:"mime_type,omitempty"`
	FileUniqueID string `json:"file_unique_id,omitempty"`
}

type Animation struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
	FileName string `json:"file_name"`
	FileSize int    `json:"file_size"`
}

type Video struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
	FileName string `json:"file_name"`
	FileSize int    `json:"file_size"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}
