package module

import (
	"context"
	"log"
	"testing"

	"github.com/AfsanehHabibi/neveshtedan/graph/model"
	"github.com/AfsanehHabibi/neveshtedan/internal/auth"
	"github.com/AfsanehHabibi/neveshtedan/pkg/logic"
	"github.com/AfsanehHabibi/neveshtedan/pkg/repository/postgres/util/test"
)

var (
	ctx    = context.Background()
	module logic.Neveshtedan
)

func TestMain(m *testing.M) {
	con := test.PreparePostgresForTest()
	module = NewNeveshtedanModule(con)
	m.Run()
}

func logsBasicUserInAndFillContext() context.Context {
	_, err := module.CreateUser(ctx, model.NewUser{Username: "Ahmad", Password: "03jf9efk"})
	if err != nil {
		log.Fatalln("failed to setup test user")
	}
	return auth.AddUserToContext(ctx, 0)
}

func clear() {
	test.EmptyTables()
}
