package httpserver

import (
	"context"
	"errors"
	"github.com/zibilal/logwrapper"
	"github.com/zibilal/simpleapi/api"
	"github.com/zibilal/simpleapi/api/v1"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type HttpServer struct {
	apiEngine api.ApiEngine
	commandName string
	address string
}

func NewHttpServer(engine api.ApiEngine, commandName string) *HttpServer {
	httpServer := new(HttpServer)
	httpServer.apiEngine = engine
	httpServer.commandName = commandName
	return httpServer
}

func (s *HttpServer) registerVersions() error {
	err := s.apiEngine.RegisterVersion(v1.VersionOneApi())
	if err != nil {
		return err
	}

	return nil
}

func (s *HttpServer) Run () error {

	// Start the server engine
	if s.commandName == "serve" {

		err := s.registerVersions()
		if err != nil {
			return err
		}

		go func() {
			if err:=s.apiEngine.Execute(); err != nil && err != http.ErrServerClosed {
				logwrapper.Fatal("unable to initiate server due to ", err.Error())
			}
		}()

		// wait for interrupt signal to gracefully shutdown the server
		// with timeout of 5 seconds
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<- quit

		// waits for 5 seconds
		parentContext := context.Background()
		ctx, cancel := context.WithTimeout(parentContext, 5 * time.Second)
		defer cancel()

		if err := s.apiEngine.Shutdown(ctx); err != nil {
			logwrapper.Fatal("Server shutdown: ", err)
		}

		logwrapper.Info("Server exiting")

		return nil
	} else {
		return errors.New("please check your implementation, need address value and command name value")
	}
}