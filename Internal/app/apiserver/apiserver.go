package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"restapisrever/Internal/app/storage"
	"restapisrever/Internal/app/store"
	"strconv"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
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
	s.router.HandleFunc("/root/api/{id}", s.getBannerById())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}
	s.store = st
	return nil
}

func (s *APIServer) getBannerById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		id, err := strconv.Atoi(getNparamFromUrl(3, request.URL.String()))
		if err != nil {
			return
		}
		var stor storage.BannerArray
		stor.BannerStorageInit()
		banner, success := stor.GetBannerById(id)
		if success == -1 {
			return
		}
		bannerJson, err := json.Marshal(banner)
		if err != nil {
			return
		}
		_, _ = io.WriteString(writer, string(bannerJson))
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
