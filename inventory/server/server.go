package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/inventory"
)

type inventoryServer struct {
	inventory.UnimplementedInventoryServer
	postgresqlDB *sql.DB
	restClient   http.Client
}

func (srv *inventoryServer) FindItem(ctx context.Context, input *inventory.FindItemInput) (output *inventory.FindItemOutput, err error) {
	name := input.Name

	var sku interface{}
	srv.postgresqlDB.QueryRow(fmt.Sprintf("SELECT sku FROM inventory WHERE name = '%s'", name)).Scan(&sku)
	output.Sku = sku

	return output, nil
}
