package models

type Banner struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Text     string `json:"text"`
	Deleted  bool   `json:"deleted"`
	IdBanner int    `json:"idBanner"`
}

func NewBanner(name string, price int, text string, deleted bool, idBanner int) Banner {
	return Banner{
		Name:     name,
		Price:    price,
		Text:     text,
		Deleted:  deleted,
		IdBanner: idBanner,
	}
}
