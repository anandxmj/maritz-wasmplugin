package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	types.DefaultVMContext
}

type pluginContext struct {
	types.DefaultPluginContext
	contextID uint32
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	proxywasm.LogInfo("Creating VM Context")
	return &pluginContext{}
}

type httpHeaders struct {
	types.DefaultHttpContext
	contextID uint32
}

func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpHeaders{contextID: contextID}
}

func (ctx *httpHeaders) OnHttpRequestHeaders(numHeaders int, eos bool) types.Action {
	proxywasm.LogInfo("Printing Request header")
	hs, err := proxywasm.GetHttpRequestHeaders()
	if err != nil {
		proxywasm.LogCriticalf("AJX Failed to get request  headers: %v", err)
	}
	for _, h := range hs {
		proxywasm.LogInfof("AJX Req Header %s:%s", h[0], h[1])
	}
	return types.ActionContinue
}

func (ctx *httpHeaders) OnHttpResponseHeaders(numHeaders int, eos bool) types.Action {
	proxywasm.LogInfo("Injecting Response Header")
	proxywasm.AddHttpResponseHeader("Cartoon", "HEMAN")
	return types.ActionContinue
}
