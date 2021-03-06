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
	"context"

	"github.com/ethereum/go-ethereum/statediff/indexer/database/sql/postgres"

	"github.com/jmoiron/sqlx"
)

// NewDB returns a new sqlx.DB instance using the provided config
func NewDB(ctx context.Context, conf postgres.Config) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", conf.DbConnectionString())
	if err != nil {
		return nil, err
	}
	if conf.MaxConns > 0 {
		db.SetMaxOpenConns(conf.MaxConns)
	}
	if conf.MaxIdle > 0 {
		db.SetMaxIdleConns(conf.MaxIdle)
	}
	if conf.MaxConnLifetime > 0 {
		lifetime := conf.MaxConnLifetime
		db.SetConnMaxLifetime(lifetime)
	}
	return db, nil
}
