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

package sql

import (
	"github.com/jmoiron/sqlx"
)

// Writer struct for writing v3 DB public.nodes models
type Writer struct {
	db *sqlx.DB
}

// NewWriter satisfies interfaces.WriterConstructor for public.nodes
func NewWriter(db *sqlx.DB) *Writer {
	return &Writer{db: db}
}

// Write satisfies interfaces.Writer for v3 database
func (w *Writer) Write(pgStr WritePgStr, models interface{}) error {
	rows, err := w.db.NamedQuery(string(pgStr), models)
	if err != nil {
		return err
	}
	return rows.Close()
}

// Close satisfies io.Closer
func (w *Writer) Close() error {
	return w.db.Close()
}
