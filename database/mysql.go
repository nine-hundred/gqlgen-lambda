package database

import (
	"context"
	"entgo.io/ent/dialect"
	"gqlgen-lambda/ent"
	"gqlgen-lambda/ent/migrate"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlClient *ent.Client

func init() {
	client, err := ent.Open(dialect.MySQL, "root:password@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalf("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true));
		err != nil {
		log.Fatalf("opening ent client", err)
	}
	MysqlClient = client
}
