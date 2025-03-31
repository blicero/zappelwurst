// /home/krylon/go/src/github.com/blicero/zappelwurst/model/model.go
// -*- mode: go; coding: utf-8; -*-
// Created on 31. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:20:20 krylon>

// Package model provides data types used in the rest of the application.
package model

import "time"

// Host is a machine the Agent runs on.
type Host struct {
	ID          int64
	Name        string
	OS          string
	LastContact time.Time
}

// Report is the data set a Host transmits to the Server.
type Report struct {
	ID        int64
	HostID    int64
	Timestamp time.Time
	Load      []float64
	Uptime    int64
	Sensors   map[string]string
}
