// Copyright (c) 2020, Hugues GUILLEUS <ghugues@netc.fr>. All rights reserved.
// Use of this source code is governed by a BSD
// license that can be found in the LICENSE file.

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetName(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		in := "../front/app.js"
		n, f := getName(in)
		assert.Equal(t, "FrontApp", n, in, "name")
		assert.Equal(t, in, f, "files")
	})
	t.Run("equal", func(t *testing.T) {
		in := "xxx=../front/app.js"
		n, f := getName(in)
		assert.Equal(t, "xxx", n, in, "name")
		assert.Equal(t, "../front/app.js", f, "files")
	})
}

func TestFname(t *testing.T) {
	assert.Equal(t, "FrontApp", fname("../front/app.js"))
}
