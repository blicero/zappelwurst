// /home/krylon/go/src/github.com/blicero/zappelwurst/main.go
// -*- mode: go; coding: utf-8; -*-
// Created on 31. 03. 2025 by Benjamin Walkenhorst
// (c) 2025 Benjamin Walkenhorst
// Time-stamp: <2025-03-31 18:15:08 krylon>

package main

import (
	"fmt"

	"github.com/blicero/zappelwurst/common"
)

func main() {
	fmt.Printf("%s %s built on %s\n",
		common.AppName,
		common.Version,
		common.BuildStamp.Format(common.TimestampFormat))

	fmt.Println("Nothing to see here, yet, keep moving.")
}
