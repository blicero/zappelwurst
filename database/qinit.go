// /home/krylon/go/src/github.com/blicero/zappelwurst/database/qinit.go
// -*- mode: go; coding: utf-8; -*-
// Created on 31. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:57:00 krylon>

package database

// These queries initialize a freshly opened database.
// Go doesn't allow const arrays or slices, but consider these constants.

var initQueries = []string{
	`
CREATE TABLE host (
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    os TEXT NOT NULL DEFAULT '',
    last_contact INTEGER NOT NULL DEFAULT 0
) STRICT
`,
	`
CREATE TABLE report (
    id INTEGER PRIMARY KEY,
    host_id INTEGER NOT NULL,
    timestamp INTEGER NOT NULL,
    load1 REAL NOT NULL,
    load5 REAL NOT NULL,
    load15 REAL NOT NULL,
    uptime INTEGER,
    sensors TEXT NOT NULL DEFAULT '',
    UNIQUE (host_id, timestamp),
    FOREIGN KEY (host_id) REFERENCES host (id)
	ON UPDATE RESTRICT
        ON DELETE CASCADE,
    CHECK (load1 >= 0 AND load5 >= 0 AND load15 >= 0),
    CHECK (json_valid(sensors))
) STRICT
`,
	"CREATE INDEX report_host_idx ON report (host_id)",
	"CREATE INDEX report_time_idx ON report (timestamp)",
}
