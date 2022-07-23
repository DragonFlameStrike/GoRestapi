package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"restapisrever/Internal/app/models"
	"restapisrever/Internal/app/storage"
	"strconv"
)

type APIServer struct {
	config          *Config
	logger          *logrus.Logger
	router          *mux.Router
	bannerStorage   *storage.BannerArray
	categoryStorage *storage.CategoryArray
}

func New(config *Config) *APIServer {
	var bannerStorage storage.BannerArray
	var categoryStorage storage.CategoryArray
	return &APIServer{
		config:          config,
		logger:          logrus.New(),
		router:          mux.NewRouter(),
		categoryStorage: categoryStorage.CategoryStorageInit(),
		bannerStorage:   bannerStorage.BannerStorageInit(categoryStorage),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.configureRouter()

	//if err := s.configureStore(); err != nil {
	//	return err
	//}

	s.logger.Info("starting api server")
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/root/api/create", s.createBanner())
	s.router.HandleFunc("/root/api/search-random", s.getRandomBannerBySearchValue())
	s.router.HandleFunc("/root/api/{id}", s.getBannerById())
	s.router.HandleFunc("/root/api/{id}/delete", s.deleteBannerById())
	s.router.HandleFunc("/root/api/search/", s.getAllBanners())
	s.router.HandleFunc("/root/api/search/{value}", s.getBannerBySearchValue())
	s.router.HandleFunc("/root/api/categories/create", s.createCategory())
	s.router.HandleFunc("/root/api/categories/{id}", s.getCategoryById())
	s.router.HandleFunc("/root/api/categories/{id}/delete", s.deleteCategoryById())
	s.router.HandleFunc("/root/api/categories/search/", s.getAllCategories())
	s.router.HandleFunc("/root/api/categories/search/{value}", s.getCategoryBySearchValue())

}

func (s *APIServer) getBannerById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		id, err := strconv.Atoi(getNparamFromUrl(3, request.URL.String()))
		if err != nil {
			return
		}
		if request.Method == "GET" {
			banner, success := s.bannerStorage.GetBannerById(id)
			if success == -1 {
				return
			}
			bannerJson, err := json.Marshal(banner)
			if err != nil {
				return
			}
			_, _ = io.WriteString(writer, string(bannerJson))
		}
		if request.Method == "POST" {
			banner := models.Banner{}
			decoder := json.NewDecoder(request.Body)
			err := decoder.Decode(&banner)
			if err != nil {
				return
			}
			s.bannerStorage.EditBanner(banner, id)
			defer request.Body.Close()
		}
	}
}

func (s *APIServer) deleteBannerById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		id, err := strconv.Atoi(getNparamFromUrl(3, request.URL.String()))
		if err != nil {
			return
		}
		s.bannerStorage.DeleteBanner(id)
		defer request.Body.Close()
	}
}

func (s *APIServer) getAllBanners() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		banners := s.bannerStorage.GetAllBanners()
		bannersJson, err := json.Marshal(banners)
		bannersJson = bannersJson[7 : len(bannersJson)-1]
		if err != nil {
			return
		}
		_, _ = io.WriteString(writer, string(bannersJson))
	}
}

func (s *APIServer) getBannerBySearchValue() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}
}

func (s *APIServer) createBanner() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if request.Method == "POST" || request.Method == "OPTIONS" {
			banner := models.Banner{}
			decoder := json.NewDecoder(request.Body)
			err := decoder.Decode(&banner)
			if err != nil {
				return
			}
			s.bannerStorage.CreateBanner(banner)
			defer request.Body.Close()
		}
		return
	}
}

func (s *APIServer) createCategory() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if request.Method == "POST" || request.Method == "OPTIONS" {
			category := models.Category{}
			decoder := json.NewDecoder(request.Body)
			err := decoder.Decode(&category)
			if err != nil {
				return
			}
			s.categoryStorage.CreateCategory(category)
			defer request.Body.Close()
		}
		return
	}
}

func (s *APIServer) getCategoryById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		id, err := strconv.Atoi(getNparamFromUrl(4, request.URL.String()))
		if err != nil {
			return
		}
		if request.Method == "GET" {
			category, success := s.categoryStorage.GetCategoryById(id)
			if success == -1 {
				return
			}
			categoryJson, err := json.Marshal(category)
			if err != nil {
				return
			}
			_, _ = io.WriteString(writer, string(categoryJson))
		}
		if request.Method == "POST" {
			category := models.Category{}
			decoder := json.NewDecoder(request.Body)
			err := decoder.Decode(&category)
			if err != nil {
				return
			}
			s.categoryStorage.EditCategory(category, id)
			defer request.Body.Close()
		}
	}
}

func (s *APIServer) deleteCategoryById() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		id, err := strconv.Atoi(getNparamFromUrl(4, request.URL.String()))
		if err != nil {
			return
		}
		s.categoryStorage.DeleteCategory(id)
		defer request.Body.Close()
	}
}

func (s *APIServer) getAllCategories() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		categories := s.categoryStorage.GetAllCategories()
		categoriesJson, err := json.Marshal(categories)
		categoriesJson = categoriesJson[7 : len(categoriesJson)-1]
		if err != nil {
			return
		}
		_, _ = io.WriteString(writer, string(categoriesJson))
	}
}

func (s *APIServer) getCategoryBySearchValue() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}
}

func (s *APIServer) getRandomBannerBySearchValue() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		banners := s.bannerStorage.GetAllBannersBySearchValue()
		banner := banners.GetRandom()
		s.bannerStorage.IncreaseRandSeed()
		bannerJson, err := json.Marshal(banner)
		if err != nil {
			return
		}
		_, _ = io.WriteString(writer, string(bannerJson))

		return
	}
}

func getNparamFromUrl(n int, url string) string {
	var count int = 0
	var id string = ""
	for i, c := range url {
		if c == '/' {
			count++
			continue
		}
		if count == n {
			fmt.Printf("%d %c\n\n", i, c)
			id += string(c)
		}
	}
	return id
}
