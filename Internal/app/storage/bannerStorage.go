package storage

import (
	"restapisrever/Internal/app/models"
	"strconv"
)

type BannerArray struct {
	arr []models.Banner
}

func (b *BannerArray) BannerStorageInit() *BannerArray {
	b.arr = make([]models.Banner, 10)
	for i := 0; i < 10; i++ {
		b.arr[i] = models.NewBanner("Banner"+strconv.Itoa(i), i, "text", false, i)
	}
	return b
}
func (b *BannerArray) GetBannerById(id int) (*models.Banner, int) {
	for i := 0; i < len(b.arr); i++ {
		banner := b.arr[i]
		if id == banner.IdBanner {
			return &banner, 0
		}
	}
	return nil, -1
}
