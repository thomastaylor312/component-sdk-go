// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package lattice represents the imported interface "wasmcloud:bus/lattice@1.0.0".
package lattice

import (
	"github.com/ydnar/wasm-tools-go/cm"
)

// CallTargetInterface represents the imported resource "wasmcloud:bus/lattice@1.0.0#call-target-interface".
//
// Interface target. This represents an interface, which can be selected by `set-link-name`.
//
//	resource call-target-interface
type CallTargetInterface cm.Resource

// ResourceDrop represents the imported resource-drop for resource "call-target-interface".
//
// Drops a resource handle.
//
//go:nosplit
func (self CallTargetInterface) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CallTargetInterfaceResourceDrop((uint32)(self0))
	return
}

//go:wasmimport wasmcloud:bus/lattice@1.0.0 [resource-drop]call-target-interface
//go:noescape
func wasmimport_CallTargetInterfaceResourceDrop(self0 uint32)

// NewCallTargetInterface represents the imported constructor for resource "call-target-interface".
//
//	constructor(namespace: string, %package: string, %interface: string)
//
//go:nosplit
func NewCallTargetInterface(namespace string, package_ string, interface_ string) (result CallTargetInterface) {
	namespace0, namespace1 := cm.LowerString(namespace)
	package0, package1 := cm.LowerString(package_)
	interface0, interface1 := cm.LowerString(interface_)
	result0 := wasmimport_NewCallTargetInterface((*uint8)(namespace0), (uint32)(namespace1), (*uint8)(package0), (uint32)(package1), (*uint8)(interface0), (uint32)(interface1))
	result = cm.Reinterpret[CallTargetInterface]((uint32)(result0))
	return
}

//go:wasmimport wasmcloud:bus/lattice@1.0.0 [constructor]call-target-interface
//go:noescape
func wasmimport_NewCallTargetInterface(namespace0 *uint8, namespace1 uint32, package0 *uint8, package1 uint32, interface0 *uint8, interface1 uint32) (result0 uint32)

// SetLinkName represents the imported function "set-link-name".
//
// Set an optional link name to use for all interfaces specified. This is advanced
// functionality only available within wasmcloud and, as such, is exposed here as
// part of the
// wasmcloud:bus package. This is used when you are linking multiple of the same interfaces
// (i.e. a keyvalue implementation for caching and another one for secrets) to a component
//
//	set-link-name: func(name: string, interfaces: list<call-target-interface>)
//
//go:nosplit
func SetLinkName(name string, interfaces cm.List[CallTargetInterface]) {
	name0, name1 := cm.LowerString(name)
	interfaces0, interfaces1 := cm.LowerList(interfaces)
	wasmimport_SetLinkName((*uint8)(name0), (uint32)(name1), (*CallTargetInterface)(interfaces0), (uint32)(interfaces1))
	return
}

//go:wasmimport wasmcloud:bus/lattice@1.0.0 set-link-name
//go:noescape
func wasmimport_SetLinkName(name0 *uint8, name1 uint32, interfaces0 *CallTargetInterface, interfaces1 uint32)
