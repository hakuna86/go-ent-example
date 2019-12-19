package main

import (
	"context"
	"fmt"
	"github.com/hakuna86/go-ent-example/ent"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	//client, err := ent.Open("mysql","root:11111@tcp(localhost:3306)/test")
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "11111", "test")
	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	ctx := context.Background()
	// Run migration.
	// Dump migration changes to stdout.
	//f, err := os.Create("migrate.sql")
	//if err != nil {
	//	log.Fatalf("create migrate file: %v", err)
	//}
	//defer f.Close()
	//if err := client.Schema.WriteTo(ctx, f); err != nil {
	//	log.Fatalf("failed printing schema changes: %v", err)
	//}

	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}

	//if err := client.Schema.Create(ctx); err != nil {
	//	log.Fatalf("failed printing schema changes: %v", err)
	//}
}