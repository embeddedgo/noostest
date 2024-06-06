// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple

import "testing"

func TestMul(t *testing.T) {
	x := 2 * 3 * 4
	if x != 24 {
		t.Errorf("2*3*4 = %d; want 24", x)
	}
}
