package entities

type Episode struct {
	ID       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name" validate:"required,min=3"`
	Number   uint8  `json:"number" db:"number" validate:"required,min=1"`
	Duration string `json:"duration" db:"duration" validate:"required,min=3"`
	Url      string `json:"url" db:"url" validate:"required,min=10"`
	Slug     string `json:"slug" db:"slug"`
	SeasonId uint64 `json:"season_id" db:"season_id" validate:"required,min=1"`
}

type EpisodeWithSeasonSlug struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Number   uint8  `json:"number"`
	Duration string `json:"duration"`
	Url      string `json:"url" `
	Slug     string `json:"slug"`
	Season   struct {
		Slug string `json:"slug"`
	} `json:"season"`
}

type EpisodeWithImage struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Season struct {
		ImageUrl string `json:"image_url"`
	} `json:"season"`
}
