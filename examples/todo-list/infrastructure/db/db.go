package db

import (
	"context"
	"fmt"
	"log"
	"todolist/ent"
)

func Provide() *ent.Client {
	fmt.Printf("providing db client\n")
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed schema resources: %v", err)
	}
	return client
}
