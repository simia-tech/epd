package main

import (
	"syscall/js"

	"github.com/simia-tech/epd"
)

var beforeUnloadChan = make(chan struct{})

func main() {
	generateProfileCb := js.NewCallback(returnToPromise(generateProfile))
	defer generateProfileCb.Release()
	js.Global().Set("_generateProfile", generateProfileCb)

	beforeUnloadCb := js.NewEventCallback(0, beforeUnload)
	defer beforeUnloadCb.Release()
	addEventListener := js.Global().Get("addEventListener")
	addEventListener.Invoke("beforeunload", beforeUnloadCb)

	<-beforeUnloadChan
}

func generateProfile(arguments []js.Value) ([]interface{}, error) {
	name := arguments[0].String()

	document, privateKey, err := epd.Generate(name)
	if err != nil {
		return nil, err
	}
	privateKeyArray := js.TypedArrayOf([]uint8(privateKey))
	defer privateKeyArray.Release()

	documentBytes, err := document.MarshalBinary()
	if err != nil {
		return nil, err
	}
	documentArray := js.TypedArrayOf([]uint8(documentBytes))
	defer documentArray.Release()

	return []interface{}{documentArray, privateKeyArray}, nil
}

func beforeUnload(event js.Value) {
	beforeUnloadChan <- struct{}{}
}

func returnToPromise(fn func([]js.Value) ([]interface{}, error)) func([]js.Value) {
	return func(arguments []js.Value) {
		if len(arguments) < 2 {
			panic("not enought arguments")
		}
		resolveFn := arguments[len(arguments)-2]
		rejectFn := arguments[len(arguments)-1]

		result, err := fn(arguments)
		if err != nil {
			rejectFn.Invoke(err.Error())
			return
		}

		resolveFn.Invoke(result...)
	}
}
