package main

import (
	stdlog "log"
	"net/http"
	"time"

	"github.com/pachyderm/s3server"
	"github.com/pachyderm/s3server/example/controllers"
	"github.com/pachyderm/s3server/example/models"
	"github.com/sirupsen/logrus"
)

func main() {
	db := models.NewStorage()

	logger := logrus.WithFields(logrus.Fields{
		"source": "s3server-example",
	})

	s3 := s3server.NewS3()
	s3.Root = controllers.RootController{DB: db}
	s3.Bucket = controllers.BucketController{DB: db}
	s3.Object = controllers.ObjectController{DB: db}

	router := s3.Router(logger)

	server := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Infof("http request: %s %s", r.Method, r.RequestURI)
			router.ServeHTTP(w, r)
		}),
		ErrorLog:     stdlog.New(logger.Writer(), "", 0),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	server.ListenAndServe()
}