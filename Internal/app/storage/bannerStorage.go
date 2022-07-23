package storage

import (
	"math/rand"
	"restapisrever/Internal/app/models"

	"strconv"
)

type BannerArray struct {
	Arr      []models.Banner
	nextId   int
	randSeed int64
}

func (b *BannerArray) BannerStorageInit(categoryStorage CategoryArray) *BannerArray {
	b.nextId = 1
	b.Arr = make([]models.Banner, 10)
	for i := 0; i < 10; i++ {
		categories := make([]models.Category, 1)
		categories[0] = categoryStorage.Arr[i%len(categoryStorage.Arr)]
		b.Arr[i] = models.NewBanner("Banner"+strconv.Itoa(b.nextId), i, "text", false, b.nextId, categories)
		b.nextId++
	}
	b.randSeed = 1
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
			b.Arr[i].IdBanner = id
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

func (b *BannerArray) GetAllBannersBySearchValue() BannerArray {
	return *b
}

func (b *BannerArray) GetRandom() models.Banner {
	rand.Seed(b.randSeed)
	number := rand.Intn(len(b.Arr))
	banner := b.Arr[number]
	return banner
}

func (b *BannerArray) IncreaseRandSeed() {
	b.randSeed++
	return
}

func removeBanner(s []models.Banner, i int) []models.Banner {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
