package main

import (
	"context"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/gorilla/mux"
	sal "github.com/salrashid123/oauth2/google"
)

const (
	address  = "demosvr-chzpysq24q-ew.a.run.app:443"
	audience = "https://demosvr-chzpysq24q-ew.a.run.app"
)

func getRpcCredsFromFile(ctx context.Context) (credentials.PerRPCCredentials, error) {
	scopes := "https://www.googleapis.com/auth/userinfo.email"
	data, err := ioutil.ReadFile("service-account.json")
	if err != nil {
		return nil, err
	}
	creds, err := google.CredentialsFromJSON(ctx, data, scopes)
	if err != nil {
		log.Fatal(err)
	}

	idTokenSource, err := sal.IdTokenSource(
		&sal.IdTokenConfig{
			Credentials: creds,
			Audiences:   []string{audience},
		},
	)
	rpcCreds, err := sal.NewIDTokenRPCCredential(ctx, idTokenSource)
	if err != nil {
		return nil, err
	}

	return rpcCreds, nil
}

func getRpcCreds(ctx context.Context) (credentials.PerRPCCredentials, error) {
	scopes := "https://www.googleapis.com/auth/userinfo.email"

	creds, err := google.FindDefaultCredentials(ctx, scopes)
	if err != nil {
		return nil, err
	}

	idTokenSource, err := sal.IdTokenSource(
		&sal.IdTokenConfig{
			Credentials: creds,
			Audiences:   []string{audience},
		},
	)
	rpcCreds, err := sal.NewIDTokenRPCCredential(ctx, idTokenSource)
	if err != nil {
		return nil, err
	}

	return rpcCreds, nil
}

func multiply(x int32, y int32) (int32, error) {
	ctx := context.Background()
	pool, _ := x509.SystemCertPool()
	creds := credentials.NewClientTLSFromCert(pool, "")
	perRPC, err := getRpcCreds(ctx)
	if err != nil {
		return 0, err
	}

	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(perRPC),
		grpc.WithBlock(),
	)
	if err != nil {
		return 0, err
	}

	defer conn.Close()

	c := NewTestServerClient(conn)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	r, err := c.Multiply(ctx, &MultiplyRequest{
		X: x,
		Y: y,
	})

	return r.GetResult(), err
}

func handler(w http.ResponseWriter, r *http.Request) {
	result, err := multiply(7, 8)

	if err != nil {
		log.Fatalf(fmt.Sprintf("Multiply failed: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Result: %v", result)))
	w.WriteHeader(http.StatusOK)
}

func main() {
	log.Print("Demo Client Services starting...")

	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
