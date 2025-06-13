// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	_ "github.com/embeddedgo/noostest/init"

	"github.com/embeddedgo/noostest/examples/panic/testpanic"
)

func main() {
	testpanic.Run(3)
}
