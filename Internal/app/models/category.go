package models

type Category struct {
	Name string `json:"name"`
	//Banners    []string `json:"banners"`
	Deleted    bool   `json:"deleted"`
	IdRequest  string `json:"idRequest"`
	IdCategory int    `json:"idCategory"`
}

func NewCategory(name string, deleted bool, idRequest string, idCategory int) Category {
	return Category{
		Name: name,
		//Banners:    make([]string, 1),
		Deleted:    deleted,
		IdCategory: idCategory,
		IdRequest:  idRequest,
	}
}
