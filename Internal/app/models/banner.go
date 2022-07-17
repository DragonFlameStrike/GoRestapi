package models

type Banner struct {
	Name       string     `json:"name"`
	Price      int        `json:"price"`
	Text       string     `json:"text"`
	Categories []Category `json:"categories"`
	Deleted    bool       `json:"deleted"`
	IdBanner   int        `json:"idBanner"`
}

func NewBanner(name string, price int, text string, deleted bool, idBanner int) Banner {
	return Banner{
		Name:       name,
		Price:      price,
		Text:       text,
		Categories: make([]Category, 0),
		Deleted:    deleted,
		IdBanner:   idBanner,
	}
}
