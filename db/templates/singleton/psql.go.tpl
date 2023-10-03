import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

const (
	bulkInsertCount        = 10000
	accountBulkInsertCount = 3000
)

func MustGetDB(host, user, pass, dbName string, port int) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", user, pass, host, port, dbName))
	if err != nil {
		panic(err)
	}

	return db
}

func CloseTx(tx *sql.Tx) {
    if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
        log.Error(err)
    }
}
