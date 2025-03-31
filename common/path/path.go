// /home/krylon/go/src/github.com/blicero/scrollmaster/common/path/path.go
// -*- mode: go; coding: utf-8; -*-
// Created on 21. 08. 2024 by Benjamin Walkenhorst
// (c) 2024 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:08:29 krylon>

package path

//go:generate stringer -type=Path

type Path uint8

const (
	Base Path = iota
	Log
	Database
	AgentConfig
	SessionStore
	Cookiejar
)
