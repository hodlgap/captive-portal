import (
	"database/sql"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	_ "github.com/newrelic/go-agent/v3/integrations/nrpq"
)

const (
	bulkInsertCount        = 10000
	accountBulkInsertCount = 3000
)

func NewDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
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
