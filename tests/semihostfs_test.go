// Copyright 2024 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package simple

import (
	"io"
	"os"
	"testing"
)

func fatalErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func TestFiles(t *testing.T) {
	const s = "1234567890\n"
	for _, name := range []string{"./test.txt", "./dir/test.txt"} {
		f, err := os.Create(name[1:]) // path without leading "."
		fatalErr(t, err)
		_, err = io.WriteString(f, s)
		fatalErr(t, err)
		fatalErr(t, f.Close())

		f, err = os.Open(name[2:]) // path without leading "./"
		fatalErr(t, err)
		data, err := io.ReadAll(f)
		fatalErr(t, err)
		if string(data) != s {
			t.Fatalf("%s: %s != %s", name, data, s)
		}
		fatalErr(t, f.Close())

		fatalErr(t, os.Remove(name))
	}
}
