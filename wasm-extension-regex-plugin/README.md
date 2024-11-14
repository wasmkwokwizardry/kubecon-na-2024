# Regex Scheduling Plugin as a WebAssembly Extension

This is an implementation of the RegexScheduling plugin which gets compiled to WebAssembly
and executed through the [kube-scheduler-wasm-extension](https://github.com/kubernetes-sigs/kube-scheduler-wasm-extension) project
and visualized via the [kube-scheduler-simulator](https://github.com/kubernetes-sigs/kube-scheduler-simulator) project.

## v1

This version is the basic version of the plugin. It gets called via the `Filter` method and receives the pod and node. It checks for the `scheduler.wasmkwokwizardry.io/regex` annotation on the pod and uses its value as a regex to match the node name. If the node name does not match the regex, it returns an `Unschedulable` status to let the scheduler know that the node is not a valid candidate.

The code can be found under the [`v1`](./v1) directory.

### How to run

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
              guestURL: http://static-webserver/regex_v1.wasm
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

## v2

This version uses multiple optimizations:
- It uses the more efficient `nottinygc` garbage collector.
- It manually configures lifecycle hooks to avoid no-op overhead.
- It implements the `PreFilter` method to compile the regex once and reuse it for all nodes via the state.

The code can be found under the [`v2`](./v2) directory.

### How to run

1. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
2. Click the cog icon in the upper-left corner to open the Scheduler configuration and edit the `guestURL` field to point to the `regex_v2.wasm` file.
   ```yaml
    pluginConfig:
      - name: WasmRegexScheduling
        args:
          guestURL: http://static-webserver/regex_v2.wasm
    ```
3. Profit!