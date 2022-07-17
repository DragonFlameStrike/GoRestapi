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
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.BannerArray
}

func New(config *Config) *APIServer {
	var bannerStorage storage.BannerArray
	return &APIServer{
		config:  config,
		logger:  logrus.New(),
		router:  mux.NewRouter(),
		storage: bannerStorage.BannerStorageInit(),
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
	s.router.HandleFunc("/root/api/{id}", s.getBannerById())
	s.router.HandleFunc("/root/api/{id}/delete", s.deleteBannerById())
	s.router.HandleFunc("/root/api/search/", s.getAllBanners())
	s.router.HandleFunc("/root/api/search/{value}", s.getBannerBySearchValue())

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
			banner, success := s.storage.GetBannerById(id)
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
			s.storage.EditBanner(banner, id)
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
		s.storage.DeleteBanner(id)
		defer request.Body.Close()
	}
}

func (s *APIServer) getAllBanners() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		banners := s.storage.GetAllBanners()
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
			s.storage.CreateBanner(banner)
			defer request.Body.Close()
		}
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
