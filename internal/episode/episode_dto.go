package episode

type createDTO struct {
	Name          string `json:"name"  validate:"required,min=5"`
	EpisodeNumber string `json:"episode_number" validate:"required,min=1"`
}
