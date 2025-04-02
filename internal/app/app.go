package app

import (
	"app1/internal/config"
	httpServer "app1/internal/delivery/http"
	logger "app1/internal/log"
	"app1/internal/repository/memory"
	"app1/internal/usecase"
	"context"
	"fmt"
	fatalLog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	// Считать одну переменную окружения. Если её нет или пустая - вернёт local
	appConfigFile := config.GetAppConfigFile()

	// config
	cfg, err := config.NewConfig(appConfigFile)
	if err != nil {
		fatalLog.Fatalf("#app_e1. Error create new config: %v", err)
	}

	// log
	log, err := logger.NewLogger(cfg)
	if err != nil {
		fatalLog.Fatalf("#app_e2. Error create new logger: %v", err)
	}

	// repository
	contactRepo := memory.NewContactRepo()

	// usecase
	contactUsecase := usecase.NewContactUsecase(contactRepo)

	// contacts handler
	contactHandler := httpServer.NewContactHandler(contactUsecase)

	// http handlers
	httpHandlers, err := httpServer.NewHandler(cfg, contactHandler)
	if err != nil {
		errTxt := fmt.Sprintf("#app_e3. Error create new httpHandlers: %v", err)
		log.Error(errTxt)
		fatalLog.Fatal(errTxt)
	}

	// http server
	httpSrv, err := httpServer.NewServer(cfg, httpHandlers)
	if err != nil {
		errTxt := fmt.Sprintf("#app_e4. Error create new httpServer: %v", err)
		log.Error(errTxt)
		fatalLog.Fatal(errTxt)
	}
	go func() {
		if err := httpSrv.Run(); err != nil && err != http.ErrServerClosed {
			errTxt := fmt.Sprintf("#app_e5. Error run httpServer: %v", err)
			log.Error(errTxt)
			fatalLog.Fatal(errTxt)
		}
	}()

	// wait quit
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// http server - stop
	if err := httpSrv.Stop(context.Background()); err != nil {
		errTxt := fmt.Sprintf("#app_e6. Error stop httpServer: %v", err)
		log.Error(errTxt)
		fatalLog.Fatal(errTxt)
	}
}
