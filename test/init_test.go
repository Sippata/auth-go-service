package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/Sippata/medods-test/src/mongo"
)

func TestMain(m *testing.M) {
	var db mongo.MongoDB
	if err := db.Open(); err != nil {
		panic(fmt.Sprintf("Db connection fault: %v", err))
	}
	defer db.Close()
	session, err := db.Client.StartSession()
	if err != nil {
		panic(err)
	}
	callback := func(sessionCtx) (interface{}, error) {
		return nil, err
	}
	_, err = session.WithTransaction(context.Background(), callback)

	os.Exit(result.(int))
}
