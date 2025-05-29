// Copyright 2025 Kevin Fisher
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

// Package db handles database related operations.
package db

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool = nil

// InitPool initializes the connection pool.
func InitPool() error {
	pool, err := pgxpool.New(context.Background(), "dbname=artie")
	if err != nil {
		return err
	} else {
		Pool = pool
		return nil
	}
}

// Close
func Close() {
	Pool.Close()
}

// InitOpticalDriveInfo checks the database to see if there is an entry for `sn`
// and adds if not.
func InitOpticalDriveInfo(ctx context.Context, sn string) error {
	stmt := "SELECT id FROM optical_drive WHERE serial_number=@sn"
	args := pgx.NamedArgs{"sn": sn}
	var id int
	err := Pool.QueryRow(ctx, stmt, args).Scan(&id)
	if err == nil {
		slog.Debug("Optical drive information already in database.", "id", id)
		return nil
	} else if err != pgx.ErrNoRows {
		return fmt.Errorf("query row failed: %w", err)
	}

	slog.Debug("Adding optical drive info to database.")

	stmt = "INSERT INTO optical_drive (serial_number) VALUES (@sn)"
	tag, err := Pool.Exec(ctx, stmt, args)
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}

	if tag.RowsAffected() != 1 {
		return errors.New("insert affected more than 1 row")
	}

	return nil
}
