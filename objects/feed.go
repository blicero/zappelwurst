// /home/krylon/go/src/github.com/blicero/zappelwurst/objects/feed.go
// -*- mode: go; coding: utf-8; -*-
// Created on 16. 10. 2021 by Benjamin Walkenhorst
// (c) 2021 Benjamin Walkenhorst
// Time-stamp: <2022-04-17 13:34:39 krylon>

package objects

import "time"

// Feed represents a Podcast. Duh.
type Feed struct {
	ID          int64
	Name        string
	Description string
	URL         string
	Homepage    string
	Interval    time.Duration
	LastUpdate  time.Time
	Active      bool
}
