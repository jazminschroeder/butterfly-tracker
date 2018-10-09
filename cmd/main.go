package main

import(
	"github.com/jazminschroeder/butterflytracker"
	"golang.org/x/net/context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"github.com/go-kit/kit/log"
)

func main() {
	ctx := context.Background()
	var service butterflytracker.Service
	service = butterflytracker.ButterflyTracker{}
	endpoints := butterflytracker.Endpoints{
		TrackButterflyEndpoint: butterflytracker.MakeTrackButterflyEndpoint(service),
	}
	var (
		httpAddr = flag.String("http.addr", ":8081", "HTTP listen address")
	)
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	

	var h http.Handler
	{
		h = butterflytracker.MakeHTTPHandler(ctx, endpoints, log.With(logger, "component", "HTTP"))
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	logger.Log("exit", <-errs)
}