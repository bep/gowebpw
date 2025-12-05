// Copyright 2025 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

package golibtemplate

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/bep/gowebpw/webpwasm"
	qt "github.com/frankban/quicktest"
)

func TestEncode(t *testing.T) {
	c := qt.New(t)

	const sunset = "sunset.jpg"

	r, err := os.Open(filepath.Join("testdata/images", sunset))
	c.Assert(err, qt.IsNil)

	defer r.Close()

	img, _, err := image.Decode(r)
	c.Assert(err, qt.IsNil)

	s := NewServer(
		ServerConfig{
			Binary: webpwasm.Binary,
		},
	)
	c.Assert(s.Start(), qt.IsNil)

	defer func() {
		c.Assert(s.Close(), qt.IsNil)
	}()

	req := EncodeRequest{
		Src: img,
		Dst: io.Discard,
	}

	c.Assert(s.Encode(req), qt.IsNil)
}
