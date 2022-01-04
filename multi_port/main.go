package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	var eg errgroup.Group
	r := gin.Default()

	insecureServer := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	secureServer := &http.Server{
		Addr: ":8443",
		Handler: r,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	eg.Go(func() error {
		err := insecureServer.ListenAndServe()
		if err != nil {
			logrus.Error(err)
		}
		return err
	})

	eg.Go(func() error {
		err := secureServer.ListenAndServeTLS("server.pem", "server.key")
		if err != nil {
			logrus.Error(err)
		}
		return err
	})

	if err := eg.Wait(); err != nil {
		logrus.Fatal(err)
	}

}
