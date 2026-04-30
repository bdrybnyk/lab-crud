package main

type Guitar struct {
	ID           string `json:"id" validate:"required,uuid4"`
	Brand        string `json:"brand" validate:"required"`
	Model        string `json:"model" validate:"required"`
	StringsCount int    `json:"strings_count" validate:"required,gte=4,lte=12"`
	IsElectric   bool   `json:"is_electric"`
}

type GuitarRepository interface {
	Create(g Guitar) error
	GetAll() ([]Guitar, error)
	GetByID(id string) (Guitar, error)
	Update(id string, g Guitar) error
	PatchElectric(id string, isElectric bool) error
	Delete(id string) error
}
