package main

import (
	"context"
	"database/sql"

	"github.com/backslash-security-tests/golang-supermarket/inventory"
	_ "github.com/lib/pq"
)

type inventoryServer struct {
	postgresqlDB sql.DB
}

func (srv *inventoryServer) FindItem(ctx context.Context, input *inventory.FindItemInput) (output *inventory.FindItemOutput, err error) {
	name := input.Name

	var sku interface{}
	srv.postgresqlDB.QueryRow("SELECT sku FROM inventory WHERE name = $1", name).Scan(&sku)
	output.Sku = sku

	return output, nil
}
