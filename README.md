# kubecon-na-2024

Materials and examples for the talk presented at KubeCon 2024 @ Salt Lake City - KWOK + WASM Wizardry: Writing and Testing Scheduler Plugins at Scale

## Scheduler Plugins

This repository provides examples based on an experimental plugin called Regex Scheduling.

The plugin implements the `Filter` interface and checks if a Pod has the `scheduler.wasmkwokwizardry.io/regex` annotation.
If this annotation is present, the plugin filters Nodes to match the specified regex pattern.

To explore these examples interactively, we recommend using the [kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator) project.
This simulator offers a rich UI that provides valuable insights and observability into custom scheduling plugins.

### Regex Scheduling as a Scheduling Framework Plugin

Try out the Scheduling Framework in Go implementation by following these [instructions](./scheduling-framework-regex-plugin/README.md).

### Regex Scheduling as a Scheduler Extender

Try out the Python Scheduler Extender in Python implementation by following these [instructions](./scheduler-extender-regex/README.md).

### Regex Scheduling as WASM Extension

Try out the WASM Extension implementation by following these [instructions](./wasm-extension-regex-plugin/README.md).
