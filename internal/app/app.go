package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ultimathul3/notes-backend/internal/config"
	"github.com/ultimathul3/notes-backend/internal/user"
	"github.com/ultimathul3/notes-backend/pkg/postgresql"
)

func init() {
	file, err := os.OpenFile("log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(io.MultiWriter(os.Stdout, file))
}

func Run() {
	cfg, err := config.ReadEnvFile()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(cfg.HTTP.GinMode)
	router := gin.New()

	pgConn, err := postgresql.NewConnection(
		context.Background(),
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		fmt.Sprint(cfg.PostgreSQL.Port),
		cfg.PostgreSQL.DB,
	)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := user.NewRepositoryPostgres(pgConn)
	userUsecase := user.NewUsecase(userRepo)
	user.NewHandlerHTTP(router, userUsecase)

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.HTTP.Port),
		Handler:        router,
		ReadTimeout:    cfg.HTTP.ReadTimeout * time.Second,
		WriteTimeout:   cfg.HTTP.WriteTimeout * time.Second,
		IdleTimeout:    cfg.HTTP.IdleTimeout * time.Second,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderMebibytes << 20,
	}

	log.Info("server is starting...")

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
