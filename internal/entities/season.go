package entities

type Season struct {
	ID       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name" validate:"required,min=3"`
	Number   uint8  `json:"number" db:"number" validate:"required,min=1"`
	Slug     string `json:"slug" db:"slug"`
	ImageUrl string `json:"image_url" db:"image_url" validate:"required,min=6"`
}
