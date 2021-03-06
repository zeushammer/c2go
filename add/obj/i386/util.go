// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package i386

const (
	fmtLong = 1 << iota
)

func bool2int(b bool) int {
	if b {
		return 1
	}
	return 0
}
