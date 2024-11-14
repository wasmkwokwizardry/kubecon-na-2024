// Package main is the entrypoint of the %.wasm file, compiled with
// '-target=wasi'. See /guest/RATIONALE.md for details.
package main

import (
	_ "github.com/wasilibs/nottinygc"

	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/config"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/filter"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/klog"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/prefilter"

	"github.com/wasmkwokwizardry/kubecon-na-2024/wasm-extension-regex-plugin/v2/plugin"
)

// main is compiled to a WebAssembly function named "_start", called by the
// wasm scheduler plugin during initialization.
func main() {
	p, err := plugin.New(klog.Get(), config.Get())
	if err != nil {
		panic(err)
	}

	prefilter.SetPlugin(p)
	filter.SetPlugin(p)
}
