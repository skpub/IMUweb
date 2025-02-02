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

	infrastructure "IMUbackend/internal/infrastructure"
	repository "IMUbackend/internal/repository"
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
		port = os.Getenv("BACKEND_PORT")
		address = os.Getenv("BACKEND_ADDR")
	)
	u, err := url.Parse(address + ":" + port)
	if err != nil {
		panic(fmt.Errorf("specify BACKEND_ADDR correctly."))
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
	bucket := os.Getenv("MDBUCKET")

	client, err := infrastructure.NewS3Client()
	if err != nil {
		log.Fatal(ctx, err)
	}

	repo := repository.NewMarkdownRepository(client, bucket)

	svc := service.NewMarkdownService(repo)
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