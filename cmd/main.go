package main

import (
	"context"
	"log"
	"net/http"
	"os"
	l "service-hf-client-p5/external/logger"
	clientrpc "service-hf-client-p5/internal/adapters/rpc"
	"service-hf-client-p5/internal/core/application"
	httpH "service-hf-client-p5/internal/handler/http"

	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		l.Errorf("", "error set envs %v", " | ", err)
	}
}

func main() {

	router := http.NewServeMux()

	ctx := context.Background()

	clientRPC := clientrpc.NewClientRPC(ctx, os.Getenv("HOST_CLIENT"), os.Getenv("PORT_CLIENT"))

	clientWorkerRPC := clientrpc.NewClientWorkerRPC(ctx, os.Getenv("HOST_CLIENT"), os.Getenv("PORT_CLIENT"))
	
	app := application.NewApplication(ctx, clientRPC, clientWorkerRPC)

	h := httpH.NewHandler(app)

	router.Handle("/hermes_foods/client/", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))
	router.Handle("/hermes_foods/client", http.StripPrefix("/", httpH.Middleware(h.HandlerClient)))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_HTTP_PORT"), router))
}
