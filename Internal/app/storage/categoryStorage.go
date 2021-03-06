package storage

import (
	"restapisrever/Internal/app/models"
)

type CategoryArray struct {
	Arr    []models.Category
	nextId int
}

func (c *CategoryArray) CategoryStorageInit() *CategoryArray {
	categories := []string{"cat", "dog", "fish"}
	c.nextId = 1
	c.Arr = make([]models.Category, 3)
	for i := 0; i < 3; i++ {
		c.Arr[i] = models.NewCategory(categories[i], false, categories[i], c.nextId)
		c.nextId++
	}
	return c
}
func (c *CategoryArray) GetCategoryById(id int) (*models.Category, int) {
	for i := 0; i < len(c.Arr); i++ {
		category := c.Arr[i]
		if id == category.IdCategory {
			return &category, 0
		}
	}
	return nil, -1
}

func (c *CategoryArray) GetAllCategories() CategoryArray {
	return *c
}

func (c *CategoryArray) CreateCategory(category models.Category) {
	category.IdCategory = c.nextId
	c.nextId++
	c.Arr = append(c.Arr, category)
}

func (c *CategoryArray) EditCategory(category models.Category, id int) {
	for i := 0; i < len(c.Arr); i++ {
		tmp := c.Arr[i]
		if id == tmp.IdCategory {
			c.Arr[i] = category
			c.Arr[i].IdCategory = id
		}
	}
	return
}

func (c *CategoryArray) DeleteCategory(id int) {
	for i := 0; i < len(c.Arr); i++ {
		tmp := c.Arr[i]
		if id == tmp.IdCategory {
			c.Arr = removeCategory(c.Arr, i)
		}
	}
}

func removeCategory(s []models.Category, i int) []models.Category {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
