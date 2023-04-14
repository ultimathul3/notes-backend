package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
	"github.com/ultimathul3/notes-backend/internal/auth"
	"github.com/ultimathul3/notes-backend/internal/config"
	"github.com/ultimathul3/notes-backend/internal/middleware"
	"github.com/ultimathul3/notes-backend/internal/notebook"
	"github.com/ultimathul3/notes-backend/internal/session"
	"github.com/ultimathul3/notes-backend/internal/user"
	"github.com/ultimathul3/notes-backend/pkg/hash"
	"github.com/ultimathul3/notes-backend/pkg/jwtauth"
	"github.com/ultimathul3/notes-backend/pkg/postgresql"
)

func Run(cfg *config.Config) {
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
	notebookRepo := notebook.NewRepositoryPostgres(pgConn)

	// providers
	sha256Hasher := hash.NewSHA256Hasher([]byte(cfg.PasswordSalt))
	jwt := jwtauth.NewJWT(cfg.Auth.AccessTokenTTL, cfg.Auth.JwtSecretKey)

	// middlewares
	tokenChecker := middleware.NewTokenChecker(jwt)

	// usecases
	userUsecase := user.NewUsecase(userRepo, sha256Hasher)
	sessionUsecase := session.NewUsecase(
		sessionRepo,
		jwt,
		cfg.Auth.RefreshTokenTTL,
		cfg.Auth.MaxUserSessionsCount,
	)
	notebookUsecase := notebook.NewUsecase(notebookRepo)

	// handlers
	auth.NewHandlerHTTP(router, userUsecase, sessionUsecase)
	notebook.NewHandlerHTTP(router, notebookUsecase, tokenChecker.Handle())

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
