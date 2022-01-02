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

package public_blocks

import (
	"github.com/jmoiron/sqlx"
	"github.com/vulcanize/migration-tools/pkg/interfaces"
)

// Writer struct for writing v3 DB public.blocks models
type Writer struct {
	db *sqlx.DB
}

// NewWriter satisfies interfaces.WriterConstructor for public.blocks
func NewWriter(db *sqlx.DB) interfaces.Writer {
	return &Writer{db: db}
}

// Write satisfies interfaces.Writer for public.blocks
func (w *Writer) Write(models [][]interface{}) error {

}
