package testutil

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//DBテスト用のhelper
//環境変数からでなく、ハードコーディングする。
//ローカルマシン環境は固定かれている為
//github actionsも固定化されている

func OpenDBForTest(t *testing.T) *sqlx.DB {
	//開発環境
	port := 33306

	//github actionsのみ
	if _, defined := os.LookupEnv("CI"); defined {
		port = 3306
	}
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("todo:todo@tcp(127.0.0.1:%d)/todo?parseTime=true", port),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(
		func() { _ = db.Close() },
	)
	return sqlx.NewDb(db, "mysql")
}
