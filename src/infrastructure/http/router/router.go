package router

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/Nagoya-Caravan-Hackathon-PDD/Github-api/src/infrastructure/http/middleware"
	"github.com/diegosz/go-graphql-client"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	slogecho "github.com/samber/slog-echo"
)

type router struct {
	db            *sql.DB
	graphqlClient *graphql.Client
	echo          *echo.Echo
	middleware    middleware.Middleware
}

func NewRouter(db *sql.DB, graphqlClient *graphql.Client) *echo.Echo {
	router := &router{
		db:            db,
		graphqlClient: graphqlClient,
		echo:          echo.New(),
		middleware:    middleware.NewMiddleware(),
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logConfig := slogecho.Config{
		DefaultLevel:     slog.LevelDebug,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		WithRequestID:      true,
		WithRequestBody:    true,
		WithResponseBody:   true,
		WithRequestHeader:  true,
		WithResponseHeader: true,
		WithSpanID:         true,
		WithTraceID:        true,
	}

	// router.echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
	// 	AllowOrigins: []string{"http://localhost:3000"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	// 	AllowHeaders: []string{middleware.AuthorizationHeaderKey},
	// }), slogecho.NewWithConfig(logger, logConfig), router.middleware.FirebaseAuth, echoMiddleware.Recover())

	router.echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{middleware.AuthorizationHeaderKey, "Content-Type"},
	}), slogecho.NewWithConfig(logger, logConfig), echoMiddleware.Recover())

	router.GithubAPI()

	return router.echo
}
