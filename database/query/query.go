// /home/krylon/go/src/github.com/blicero/zappelwurst/database/query/query.go
// -*- mode: go; coding: utf-8; -*-
// Created on 31. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:42:30 krylon>

// Package query provides symbolic constants to refer to database queries.
package query

//go:generate stringer -type=ID

// ID identifies a database query
type ID uint8

const (
	HostAdd ID = iota
	HostUpdateContact
	HostDelete
	HostGetByID
	HostGetAll
	ReportAdd
	ReportDelete
	ReportDeleteByHost
	ReportGetByHost
	ReportGetByPeriod
	ReportGetByID
	ReportGetMulti
)
