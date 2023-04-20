package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/ultimathul3/notes-backend/docs"
	"github.com/ultimathul3/notes-backend/internal/auth"
	"github.com/ultimathul3/notes-backend/internal/config"
	"github.com/ultimathul3/notes-backend/internal/middleware"
	"github.com/ultimathul3/notes-backend/internal/note"
	"github.com/ultimathul3/notes-backend/internal/notebook"
	"github.com/ultimathul3/notes-backend/internal/session"
	sharedNote "github.com/ultimathul3/notes-backend/internal/sharednote"
	"github.com/ultimathul3/notes-backend/internal/user"
	"github.com/ultimathul3/notes-backend/pkg/hash"
	"github.com/ultimathul3/notes-backend/pkg/jwtauth"
	"github.com/ultimathul3/notes-backend/pkg/postgresql"
)

func Run(cfg *config.Config) {
	gin.SetMode(cfg.HTTP.GinMode)
	router := gin.New()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = cfg.CORS.AllowOrigins
	router.Use(cors.New(corsConfig))

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
	noteRepo := note.NewRepositoryPostgres(pgConn)
	sharedNoteRepo := sharedNote.NewRepositoryPostgres(pgConn)

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
	noteUsecase := note.NewUsecase(noteRepo)
	sharedNoteUsecase := sharedNote.NewUsecase(sharedNoteRepo)

	// handlers
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth.NewHandlerHTTP(router, userUsecase, sessionUsecase)
	notebook.NewHandlerHTTP(router, notebookUsecase, tokenChecker.Handle())
	note.NewHandlerHTTP(router, noteUsecase, tokenChecker.Handle())
	sharedNote.NewHandlerHTTP(router, sharedNoteUsecase, userUsecase, tokenChecker.Handle())

	addr := fmt.Sprintf("%s:%d", cfg.HTTP.IP, cfg.HTTP.Port)

	server := &http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    cfg.HTTP.ReadTimeout * time.Second,
		WriteTimeout:   cfg.HTTP.WriteTimeout * time.Second,
		IdleTimeout:    cfg.HTTP.IdleTimeout * time.Second,
		MaxHeaderBytes: cfg.HTTP.MaxHeaderMebibytes << 20,
	}

	log.Info("server is listening on ", addr)

	shutdownGracefully(server, cfg.HTTP.ShutdownTimeout)
}
