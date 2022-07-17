package storage

import (
	"restapisrever/Internal/app/models"
	"strconv"
)

type BannerArray struct {
	Arr    []models.Banner
	nextId int
}

func (b *BannerArray) BannerStorageInit() *BannerArray {
	b.nextId = 1
	b.Arr = make([]models.Banner, 10)
	for i := 0; i < 10; i++ {
		b.Arr[i] = models.NewBanner("Banner"+strconv.Itoa(b.nextId), i, "text", false, b.nextId)
		b.nextId++
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

func (b *BannerArray) CreateBanner(banner models.Banner) {
	banner.IdBanner = b.nextId
	b.nextId++
	b.Arr = append(b.Arr, banner)
}

func (b *BannerArray) EditBanner(banner models.Banner, id int) {
	for i := 0; i < len(b.Arr); i++ {
		tmp := b.Arr[i]
		if id == tmp.IdBanner {
			b.Arr[i] = banner
		}
	}
	return
}

func (b *BannerArray) DeleteBanner(id int) {
	for i := 0; i < len(b.Arr); i++ {
		tmp := b.Arr[i]
		if id == tmp.IdBanner {
			b.Arr = removeBanner(b.Arr, i)
		}
	}
}

func removeBanner(s []models.Banner, i int) []models.Banner {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
