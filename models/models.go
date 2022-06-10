package models

type Media string

const (
	Img Media = "image"
	Vdo Media = "video"
)

type FileType string

const (
	JPG  FileType = "jpg"
	PNG  FileType = "png"
	SVG  FileType = "svg"
	GIF  FileType = "gif"
	WEBP FileType = "webp"
)

// ImageUploadResponse defines the response schema of the upload image response by imgur.com
type ImageUploadResponse struct {
	Data struct {
		Id          string   `json:"id"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Datetime    int      `json:"datetime"`
		Type        string   `json:"type"`
		Animated    bool     `json:"animated"`
		Width       int      `json:"width"`
		Height      int      `json:"height"`
		Size        int      `json:"size"`
		Views       int      `json:"views"`
		Bandwidth   int      `json:"bandwidth"`
		Vote        string   `json:"vote"`
		Favorite    bool     `json:"favorite"`
		Nsfw        string   `json:"nsfw"`
		Section     string   `json:"section"`
		AccountUrl  string   `json:"account_url"`
		AccountId   int      `json:"account_id"`
		IsAd        bool     `json:"is_ad"`
		InMostViral bool     `json:"in_most_viral"`
		Tags        []string `json:"tags"`
		AdType      int      `json:"ad_type"`
		AdUrl       string   `json:"ad_url"`
		InGallery   bool     `json:"in_gallery"`
		Deletehash  string   `json:"deletehash"`
		Name        string   `json:"name"`
		Link        string   `json:"link"`
		HasSound    bool     `json:"has_sound"`
		Mp4         string   `json:"mp4"`
		Hls         string   `json:"hls"`
	} `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
