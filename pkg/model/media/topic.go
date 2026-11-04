package media

type Topic struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	FileName     string `json:"file_name"`
	MainImageUrl string `json:"main_image_url"`
	MainImageKey string `json:"main_image_key"`
}
