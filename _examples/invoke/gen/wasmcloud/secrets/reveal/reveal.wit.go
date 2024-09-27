// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package reveal represents the imported interface "wasmcloud:secrets/reveal@0.1.0-draft".
package reveal

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/wasmCloud/component-sdk-go/_examples/invoke/gen/wasmcloud/secrets/store"
)

// Reveal represents the imported function "reveal".
//
// Reveals the value of a secret to the caller.
// This lets you easily audit your code to discover where secrets are being used.
//
//	reveal: func(s: borrow<secret>) -> secret-value
//
//go:nosplit
func Reveal(s store.Secret) (result store.SecretValue) {
	s0 := cm.Reinterpret[uint32](s)
	wasmimport_Reveal((uint32)(s0), &result)
	return
}

//go:wasmimport wasmcloud:secrets/reveal@0.1.0-draft reveal
//go:noescape
func wasmimport_Reveal(s0 uint32, result *store.SecretValue)
