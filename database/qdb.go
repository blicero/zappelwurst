// /home/krylon/go/src/github.com/blicero/zappelwurst/database/qdb.go
// -*- mode: go; coding: utf-8; -*-
// Created on 31. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 19:15:52 krylon>

package database

import "github.com/blicero/zappelwurst/database/query"

var dbQueries = map[query.ID]string{
	query.HostAdd:           "INSERT INTO host (name, os) VALUES (?, ?) RETURNING id",
	query.HostUpdateContact: "UPDATE host SET last_contact = ? WHERE id = ?",
}
