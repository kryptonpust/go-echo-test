package sqlTools

import (
	"github.com/kryptonpust/go-echo-test/infrastructure/datastore"
	"github.com/jmoiron/sqlx"
)

func NewSqlxTransaction(mdbi *datastore.MasterDbInstance) *SqlxTransaction {
	return &SqlxTransaction{
		db: mdbi.DBX(),
	}
}

type SqlxTransaction struct {
	db *sqlx.DB
}

func (s SqlxTransaction) Init() (tx *sqlx.Tx, err error) {
	tx, err = s.db.Beginx()
	return
}
