// Package main is the entrypoint of the %.wasm file, compiled with
// '-target=wasi'. See /guest/RATIONALE.md for details.
package main

import (
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/config"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/klog"
	pluginapi "sigs.k8s.io/kube-scheduler-wasm-extension/guest/plugin"

	"github.com/wasmkwokwizardry/kubecon-na-2024/wasm-extension-regex-plugin/v1/plugin"
)

// main is compiled to a WebAssembly function named "_start", called by the
// wasm scheduler plugin during initialization.
func main() {
	p, err := plugin.New(klog.Get(), config.Get())
	if err != nil {
		panic(err)
	}

	pluginapi.Set(p)
}
