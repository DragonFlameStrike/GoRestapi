package storage

import (
	"restapisrever/Internal/app/models"
	"strconv"
)

type BannerArray struct {
	Arr []models.Banner
}

func (b *BannerArray) BannerStorageInit() *BannerArray {
	b.Arr = make([]models.Banner, 10)
	for i := 0; i < 10; i++ {
		b.Arr[i] = models.NewBanner("Banner"+strconv.Itoa(i), i, "text", false, i)
	}
	return b
}
func (b *BannerArray) GetBannerById(id int) (*models.Banner, int) {
	for i := 0; i < len(b.Arr); i++ {
		banner := b.Arr[i]
		if id == banner.IdBanner {
			return &banner, 0
		}
	}
	return nil, -1
}

func (b *BannerArray) GetAllBanners() BannerArray {
	return *b
}
