// VulcanizeDB
// Copyright © 2022 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package migration_tools

import (
	"github.com/jmoiron/sqlx"
)

const subChannelBufferSize = 1024

func NewSubChannelSet(tables []TableName) map[TableName]chan [2]uint64 {
	subChans := make(map[TableName]chan [2]uint64, len(tables))
	for _, tableName := range tables {
		subChans[tableName] = make(chan [2]uint64, subChannelBufferSize)
	}
	return subChans
}

// Reader struct for reading v2 DB eth.log_cids models
type Reader struct {
	db *sqlx.DB
}

// NewReader satisfies interfaces.ReaderConstructor for eth.log_cids
func NewReader(db *sqlx.DB) *Reader {
	return &Reader{db: db}
}

// Read satisfies interfaces.Reader for eth.log_cids
// Read is safe for concurrent use, as the only shared state is the concurrent safe *sqlx.DB
func (r *Reader) Read(blockRange [2]uint64, pgStr ReadPgStr, models interface{}) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			rollback(tx)
			panic(p)
		} else if err != nil {
			rollback(tx)
		} else {
			err = tx.Commit()
		}
	}()
	err = tx.Select(models, string(pgStr))
	return err
}