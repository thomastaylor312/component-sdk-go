// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package preopens represents the imported interface "wasi:filesystem/preopens@0.2.0".
package preopens

import (
	"github.com/wasmCloud/component-sdk-go/_examples/http-server/gen/wasi/filesystem/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// GetDirectories represents the imported function "get-directories".
//
// Return the set of preopened directories, and their path.
//
//	get-directories: func() -> list<tuple<descriptor, string>>
//
//go:nosplit
func GetDirectories() (result cm.List[cm.Tuple[types.Descriptor, string]]) {
	wasmimport_GetDirectories(&result)
	return
}

//go:wasmimport wasi:filesystem/preopens@0.2.0 get-directories
//go:noescape
func wasmimport_GetDirectories(result *cm.List[cm.Tuple[types.Descriptor, string]])
