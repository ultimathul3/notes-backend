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
	"github.com/ultimathul3/notes-backend/internal/auth"
	"github.com/ultimathul3/notes-backend/internal/config"
	"github.com/ultimathul3/notes-backend/internal/session"
	"github.com/ultimathul3/notes-backend/internal/user"
	"github.com/ultimathul3/notes-backend/pkg/hash"
	"github.com/ultimathul3/notes-backend/pkg/jwtauth"
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

	// repositories
	userRepo := user.NewRepositoryPostgres(pgConn)
	sessionRepo := session.NewRepositoryPostgres(pgConn)

	// user
	sha256Hasher := hash.NewSHA256Hasher([]byte(cfg.PasswordSalt))
	userUsecase := user.NewUsecase(userRepo, sha256Hasher)
	user.NewHandlerHTTP(router, userUsecase)

	// auth
	jwt := jwtauth.NewJWT(cfg.Auth.AccessTokenTTL, cfg.Auth.JwtSecretKey)
	sessionUsecase := session.NewUsecase(sessionRepo)
	auth.NewHandlerHTTP(
		router,
		userUsecase,
		sessionUsecase,
		jwt,
		cfg.Auth.RefreshTokenTTL,
		cfg.Auth.MaxUserSessionsCount,
	)

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
