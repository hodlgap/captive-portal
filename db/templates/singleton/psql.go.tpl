import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

    _ "github.com/lib/pq"
)

const (
	bulkInsertCount        = 10000
	accountBulkInsertCount = 3000
)

func NewDB(host, user, pass, dbName string, port int) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, pass, host, port, dbName,
	))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, errors.WithStack(db.Ping())
}

func CloseTx(tx *sql.Tx) {
    if err := tx.Rollback(); err != nil && err != sql.ErrTxDone {
        log.Errorf("%+v", errors.WithStack(err))
    }
}
