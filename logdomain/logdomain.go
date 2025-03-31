// /home/krylon/go/src/github.com/blicero/randomland/logdomain/logdomain.go
// -*- mode: go; coding: utf-8; -*-
// Created on 21. 12. 2024 by Benjamin Walkenhorst
// (c) 2024 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:09:31 krylon>

package logdomain

//go:generate stringer -type=ID

type ID uint8

const (
	Database ID = iota
	Server
	Client
)

func AllDomains() []ID {
	return []ID{
		Database,
		Server,
		Client,
	}
} // func AllDomains() []ID
