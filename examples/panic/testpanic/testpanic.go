// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testpanic

import "fmt"

func panicZero(n int) {
	if n == 0 {
		panic("intentional panic")
	}
	Run(n-1)
}

func Run(n int) {
	fmt.Println(n)
	panicZero(n)
}
