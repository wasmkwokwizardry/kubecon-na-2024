// Package main is the entrypoint of the %.wasm file, compiled with
// '-target=wasi'. See /guest/RATIONALE.md for details.
package main

import (
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/config"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/klog"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/plugin"

	"github.com/wasmkwokwizardry/wasm-extension-regex-plugin/pkg/plugin/regex"
)

// main is compiled to a WebAssembly function named "_start", called by the
// wasm scheduler plugin during initialization.
func main() {
	p, err := regex.New(klog.Get(), config.Get())
	if err != nil {
		panic(err)
	}
	plugin.Set(p)
}
