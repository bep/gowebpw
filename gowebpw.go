// Copyright 2025 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

package gowebpw

import (
	"bytes"
	"compress/gzip"
	"errors"
	"image"
	"io"
	"math/bits"

	"github.com/bep/textandbinaryreader"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type EncodeRequest struct {
	Src image.Image
	Dst io.Writer
}

var foo textandbinaryreader.Reader

type EncodeResponse struct{}

// Response:
// JSON
// body: some magic keyword + id + length + binary data.
// Need to create a special reader for this that can be passed to json.Decoder which collects body and skips any binary data.

type ServerConfig struct {
	Binary []byte // Gzipped WebAssembly binary.
}

func NewServer(conf ServerConfig) *Server {
	return &Server{
		conf: conf,
	}
}

type Server struct {
	// Config.
	conf ServerConfig

	// The WASI instance.
	mod api.Module

	// State.
	started bool
	closed  bool
}

func (s *Server) checkStartedAndNotClosed() {
	if !s.started {
		panic("server not started")
	}
	if s.closed {
		panic("server is closed")
	}
}

func (s *Server) Encode(req EncodeRequest) error {
	s.checkStartedAndNotClosed()
	return errors.New("todo")
}

func (s *Server) Start() error {
	if s.closed {
		return errors.New("server is closed")
	}
	s.started = true

	cfg := wazero.NewRuntimeConfig()
	cfg = cfg.WithCoreFeatures(api.CoreFeaturesV2)
	if bits.UintSize < 64 {
		cfg = cfg.WithMemoryLimitPages(512) // 32MB
	} else {
		cfg = cfg.WithMemoryLimitPages(4096) // 256MB
	}
	// ctx := context.Background()
	// rt := wazero.NewRuntimeWithConfig(ctx, cfg)

	r, err := gzip.NewReader(bytes.NewReader(s.conf.Binary))
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var data bytes.Buffer
	_, err = data.ReadFrom(r)
	if err != nil {
		return err
	}

	/*cm, err := rt.CompileModule(experimental.WithCompilationWorkers(ctx, runtime.GOMAXPROCS(0)/4), data.Bytes())
	if err != nil {
		return err
	}

	wasi_snapshot_preview1.MustInstantiate(ctx, rt)

	mc := wazero.NewModuleConfig().WithStderr(os.Stderr).WithStdout(io.Discard) // .WithStdin(io.Discard)

	mc = mc.WithStartFunctions() // Delay _start until later.*/

	return nil
}

func (s *Server) Close() error {
	s.closed = true
	return nil
}
