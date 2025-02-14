package main

import (
	"IMUbackend/gen/imubackend"
	"context"
	"flag"
	"net/url"
	"sync"
	"syscall"

	"goa.design/clue/debug"
	"goa.design/clue/log"

	dbb "IMUbackend/db"

	infrastructure "IMUbackend/internal/infrastructure"
	"IMUbackend/internal/repository"
	service "IMUbackend/internal/service"
	"fmt"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
)

func main() {
	var (
		dbgF = flag.Bool("debug", false, "Log request and response bodies")
	)
	err := godotenv.Load("../.env")
	if err != nil {
		panic(fmt.Errorf("error loading .env file"))
	}
	var (
		port    = os.Getenv("PRIVATE_BACKEND_PORT")
		address = os.Getenv("PRIVATE_BACKEND_ADDR")
	)
	u, err := url.Parse(address + ":" + port)
	if err != nil {
		panic(fmt.Errorf("specify BACKEND_ADDR correctly"))
	}

	format := log.FormatJSON
	if log.IsTerminal() {
		format = log.FormatTerminal
	}
	ctx := log.Context(context.Background(), log.WithFormat(format))
	if *dbgF {
		ctx = log.Context(ctx, log.WithDebug())
		log.Debugf(ctx, "debug logs enabled.")
	}
	log.Print(ctx, log.KV{K: "http-port", V: port})

	var (
		imubackendSvc imubackend.Service
	)
	//
	// DI START
	//

	// minio
	endpoint := os.Getenv("MINIO_SERVER_URL")
	accessKeyID := os.Getenv("MINIO_ROOT_USER")
	secret := os.Getenv("MINIO_ROOT_PASSWORD")
	client, err := infrastructure.NewObjectStorageConnection(endpoint, accessKeyID, secret)
	if err != nil {
		panic(err)
	}
	// bucket := os.Getenv("MDBUCKET")
	// end minio

	// postgres
	pg_host := os.Getenv("PG_HOST")
	pg_user := os.Getenv("PG_USER")
	pg_password := os.Getenv("PG_PASSWORD")
	pg_dbname := os.Getenv("PG_DBNAME")
	salt := os.Getenv("SALT")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", pg_user, pg_password, pg_host, "5432", pg_dbname)

	db, err := infrastructure.NewDBConnection(dsn)
	if err != nil {
		panic(err)
	}
	dbWrapper := dbb.New(db.GetDBTX())
	userRepo := repository.NewStudentRepository(dbWrapper)
	articleRepo := repository.NewArticleRepository(dbWrapper, client, "mds")
	// end postgres

	// jwt
	jwtSecret := os.Getenv("JWT_SECRET")
	// end jwt

	svc := service.NewIMUSrv(articleRepo, userRepo, jwtSecret, salt, db)
	//
	// DI END
	//
	{
		imubackendSvc = svc
	}

	var (
		imubackendEndpoints *imubackend.Endpoints
	)
	{
		imubackendEndpoints = imubackend.NewEndpoints(imubackendSvc)
		imubackendEndpoints.Use(debug.LogPayloads())
		imubackendEndpoints.Use(log.Endpoint)
	}

	errc := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(ctx)

	handleHTTPServer(ctx, u, imubackendEndpoints, &wg, errc, *dbgF)

	log.Printf(ctx, "exiting (%v)", <-errc)

	cancel()

	wg.Wait()
	log.Printf(ctx, "exited")

	// //
	// //
	// //

	// endpoints := gen.NewEndpoints(svc)
	// mux := goahttp.NewMuxer()

	// server.New(endpoints, mux, decoder, encoder, errorHandler, errorStatus)

	// http.Handle("/api/", mux)
	// log.Fatal(http.ListenAndServe(":8080", nil))

}
