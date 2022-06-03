package main

import (
	"context"
	"fmt"
	"gobase/api/internal/controller/relationship"
	"gobase/api/internal/controller/user"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/friendsofgo/errors"
	"gobase/api/cmd/banner"
	"gobase/api/cmd/serverd/router"
	"gobase/api/internal/repository"
	"gobase/api/pkg/app"
	"gobase/api/pkg/db/pg"
	"gobase/api/pkg/env"
	"gobase/api/pkg/httpserv"
)

func main() {
	ctx := context.Background()

	banner.Print()

	appCfg := app.Config{
		ProjectName:      env.GetAndValidateF("PROJECT_NAME"),
		AppName:          env.GetAndValidateF("APP_NAME"),
		SubComponentName: env.GetAndValidateF("PROJECT_COMPONENT"),
		Env:              app.Env(env.GetAndValidateF("APP_ENV")),
		Version:          env.GetAndValidateF("APP_VERSION"),
		Server:           env.GetAndValidateF("SERVER_NAME"),
		ProjectTeam:      os.Getenv("PROJECT_TEAM"),
	}
	if err := appCfg.IsValid(); err != nil {
		log.Fatal(err)
	}

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Exiting...")
}

func run(ctx context.Context) error {
	log.Println("Starting app initialization")
	dbOpenConns, err := strconv.Atoi(env.GetAndValidateF("DB_POOL_MAX_OPEN_CONNS"))
	if err != nil {
		return errors.WithStack(fmt.Errorf("invalid db pool max open conns: %w", err))
	}
	dbIdleConns, err := strconv.Atoi(env.GetAndValidateF("DB_POOL_MAX_IDLE_CONNS"))
	if err != nil {
		return errors.WithStack(fmt.Errorf("invalid db pool max idle conns: %w", err))
	}

	conn, err := pg.NewPool(env.GetAndValidateF("DB_URL"), dbOpenConns, dbIdleConns)
	if err != nil {
		return err
	}

	defer conn.Close()

	rtr, err := initRouter(ctx, conn)

	log.Println("App initialization completed")

	httpserv.NewServer(rtr.Handler()).Start(ctx)

	return nil
}

func initRouter(
	ctx context.Context,
	dbConn pg.BeginnerExecutor) (router.Router, error) {
	return router.New(
		ctx,
		strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
		os.Getenv("GQL_INTROSPECTION_ENABLED") == "true",
		user.New(repository.New(dbConn)),
		relationship.New(repository.New(dbConn)),
	), nil
}
