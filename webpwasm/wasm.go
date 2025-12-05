// Copyright 2025 Bjørn Erik Pedersen
// SPDX-License-Identifier: MIT

// Package webpwasm embeds the WebP WebAssembly binary.
// The reason it's embedded in its own package is so that it can be optional.
package webpwasm

import _ "embed"

//go:embed webp.wasm.gz
var Binary []byte
