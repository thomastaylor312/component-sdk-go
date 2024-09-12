// Code generated by wit-bindgen-go. DO NOT EDIT.

package handler

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"go.wasmcloud.dev/component/gen/wasmcloud/messaging/types"
)

func lift_OptionString(f0 uint32, f1 *uint8, f2 uint32) (v cm.Option[string]) {
	if f0 == 0 {
		return
	}
	return cm.Some[string](cm.LiftString[string]((*uint8)(f1), (uint32)(f2)))
}

func lift_BrokerMessage(f0 *uint8, f1 uint32, f2 *uint8, f3 uint32, f4 uint32, f5 *uint8, f6 uint32) (v types.BrokerMessage) {
	v.Subject = cm.LiftString[string](f0, f1)
	v.Body = cm.LiftList[cm.List[uint8]](f2, f3)
	v.ReplyTo = lift_OptionString(f4, f5, f6)
	return
}
