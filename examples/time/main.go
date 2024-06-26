// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"time"

	_ "github.com/embeddedgo/noostest/init"
)

func main() {
	for {
		fmt.Println(time.Now())
		time.Sleep(time.Second)
	}
}
