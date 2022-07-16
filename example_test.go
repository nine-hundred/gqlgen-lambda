package main

import (
	"context"
	"entgo.io/ent/dialect"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"gqlgen-lambda/ent"
	"log"
	"testing"
)

func Example_Todo() {
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening")
	}

	defer client.Close()
	ctx := context.Background()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resourece: %v", err)
	}

	task1, err := client.Todo.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Println(task1.ID)
}

func TestExample_Todo(t *testing.T) {
	t.Run("test Example_Todo()", func(t *testing.T) {
		Example_Todo()
	})
}
