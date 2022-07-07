package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"restapisrever/Internal/app/models"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
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
	s.router.HandleFunc("/root/api/{id}", s.firstBanner())
}

func (s *APIServer) firstBanner() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		firstBanner := models.NewBanner("Banner1", 100, "Test", false, 1)
		b, err := json.Marshal(firstBanner)
		_, err = io.WriteString(writer, string(b))
		if err != nil {
			return
		}
	}
}
