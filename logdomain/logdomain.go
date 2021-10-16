// /home/krylon/go/src/github.com/blicero/zappelwurst/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 15. 10. 2021 by Benjamin Walkenhorst
// (c) 2021 Benjamin Walkenhorst
// Time-stamp: <2021-10-15 23:46:18 krylon>

// Package logdomain provides constants for log sources.
package logdomain

//go:generate stringer -type=ID

// ID represents a log source
type ID uint8

// These constants signify the various parts of the application.
const (
	Common ID = iota
	Client
	DBPool
	Database
	Web
)

// AllDomains returns a slice of all the known log sources.
func AllDomains() []ID {
	return []ID{
		Common,
		Client,
		DBPool,
		Database,
		Web,
	}
} // func AllDomains() []ID
