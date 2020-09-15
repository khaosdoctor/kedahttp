package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http/httputil"
	"os"
	"sync"
	"time"

	"github.com/arschles/containerscaler/externalscaler"
	"github.com/arschles/containerscaler/pkg/k8s"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"google.golang.org/grpc"
)

const (
	clusterID = "test-cluster"
	clientID  = "cscaler-client"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	reqCounter := &reqCounter{i: 0, mut: new(sync.RWMutex)}

	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler

	e.Use(middleware.Logger())
	e.Use(userAgentHandler())
	countM := countMiddleware(reqCounter)

	// don't put this inside the middleware so we don't print out an incorrect
	// counter
	e.GET("/counter", func(c echo.Context) error {
		fmt.Fprintf(c.Response(), "%d", reqCounter.get())
		return nil
	})

	// Azure front door health check
	e.GET("/pong", pongHandler)

	e.GET("/azfdhealthz", newHealthCheckHandler())
	e.Any("/", newForwardingHandler(), countM)

	clientset, dynCl, err := k8s.NewClientset()
	if err != nil {
		log.Fatal(err)
	}

	e.POST("/admin/deploy", newAdminCreateDeploymentHandler(clientset, dynCl))
	e.DELETE("/admin/deploy", newAdminDeleteDeploymentHandler(clientset, dynCl))

	port := "8080"
	listenPortEnv := os.Getenv("LISTEN_PORT")
	portEnv := os.Getenv("PORT")

	if listenPortEnv != "" {
		port = listenPortEnv
	} else if portEnv != "" {
		port = portEnv
	}

	go func() {
		log.Printf("HTTP listening on port %s", port)
		portStr := fmt.Sprintf(":%s", port)
		log.Fatal(e.Start(portStr))
	}()
	go func() {
		log.Printf("GRPC listening on port 9090")
		log.Fatal(startGrpcServer(reqCounter))
	}()
	select {}
}

func startGrpcServer(ctr *reqCounter) error {
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	externalscaler.RegisterExternalScalerServer(grpcServer, newImpl(ctr))
	return grpcServer.Serve(lis)
}

func pongHandler(c echo.Context) error {
	reqBytes, err := httputil.DumpRequest(c.Request(), true)
	if err != nil {
		return c.String(500, err.Error())
	}
	return c.String(200, string(reqBytes))
}
