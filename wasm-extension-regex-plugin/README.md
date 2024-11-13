# Regex Scheduling Plugin as a WebAssembly Extension

This is an implementation of the RegexScheduling plugin which gets compiled to WebAssembly
and executed through the [kube-scheduler-wasm-extension](https://github.com/kubernetes-sigs/kube-scheduler-wasm-extension) project
and visualized via the [kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator) project.

## How to run

1. Run `make start-simulator-wasm` from the project root.
2. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
3. Click the cog icon in the upper-left corner to open the Scheduler configuration.
   1. Edit the `plugins` field so it contains the following:
      ```yaml
      plugins:
        multiPoint:
          enabled:
            - name: WasmRegexScheduling
      ```
   2. Edit the `pluginConfig` field so it contains the following:
      ```yaml
        pluginConfig:
          - name: WasmRegexScheduling
            args:
              guestURL: http://static-webserver/regex.wasm
      ```
4. Create a Node with a default name and additional two Nodes named `kubecon-1` and `kubecon-2`
5. Create a Pod with an annotation `scheduler.wasmkwokwizardry.io/regex: 'kubecon-.*'`
    ```yaml
    annotations:
      scheduler.wasmkwokwizardry.io/regex: 'kubecon-.*'
    ```
6. Add a regex annotation for a non-existing Node Group and show result:
    ```yaml
    annotations:
      scheduler.wasmkwokwizardry.io/regex: 'cncf-.*'
    ```