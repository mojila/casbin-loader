package main

import (
	"context"
	"fmt"
	"os"

	casbin "github.com/casbin/casbin-go-client/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()

	casbinHost := os.Getenv("CASBIN_HOST")

	client, err := casbin.NewClient(
		ctx,
		casbinHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	connectionString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	enforcer, err := client.NewEnforcer(ctx, casbin.Config{
		DriverName:    "postgres",
		ConnectString: connectionString,
		DbSpecified:   true,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	// Load the policy from DB.
	err = enforcer.LoadPolicy(ctx)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(3)
	}

	fmt.Println("Policy loaded from DB.")
	os.Exit(0)
}
