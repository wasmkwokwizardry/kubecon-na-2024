# Regex Scheduling Plugin implemented as a Scheduler Framework Plugin in Golang

This is an implementation which uses the [Scheduling Framework](https://github.com/kubernetes-sigs/scheduler-plugins) implementation in Go.


## v1

This is the basic version of the plugin. It is triggered by the `Filter` method,
receiving both the pod and node as parameters.

The plugin checks for the `scheduler.wasmkwokwizardry.io/regex` annotation on the pod,
using its value as a regex pattern to match against the node name.
If the node name does not match the specified regex,
the plugin returns an `Unschedulable` status to indicate that the node is not a suitable candidate for scheduling the pod.

The code can be found under the [`v1`](./v1) directory.

### How to run

1. Go to the Simulator project by running `cd kube-scheduler-simulator/simulator`
2. Add the new plugin as a dependency by running `go get github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin`
3. Open [this file](../kube-scheduler-simulator/simulator/scheduler/config/plugin.go) and import the following module: `regex "github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/v1/plugin"`
4. Add the Regex Scheduling plugin to the out-of-tree registry:
    ```go
    var outOfTreeRegistries = runtime.Registry{
        regex.Name: regex.New,
    }
    ```
5. Go back to the project root and run the Simulator by running `make start-simulator`.
6. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
7. Click the cog icon in the upper-left corner to open the Scheduler configuration and edit the `plugins` field so it contains the following:
    ```yaml
    multiPoint:
      enabled:  
      - name: RegexScheduling
    ```
8. Create a Node with a default name and additional two Nodes named `kubecon-1` and `kubecon-2`
9. Create a Pod with an annotation `scheduler.wasmkwokwizardry.io/regex: 'kubecon-.*'`
    ```yaml
    annotations:
      scheduler.wasmkwokwizardry.io/regex: 'kubecon-.*'
    ```
10. Create a Pod with an annotation for a non-existing Node Group and observe result:
     ```yaml
     annotations:
       scheduler.wasmkwokwizardry.io/regex: 'cncf-.*'
     ```

## v2

This version adds a `PreFilter` step where it compiles the regex pattern
and caches it for later use in the `Filter` step.
This avoids recompiling the regex pattern in each invocation of the `Filter` method for a given pod and available nodes.

The code can be found under the [`v2`](./v2) directory.

### How to run

1. Open the file `simulator/scheduler/config/plugin.go` and edit import from the previous version: `regex "github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/v2/plugin"`
2. Restart the Simulator by running `make start-simulator` from the project root.
3. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
4. Click the cog icon in the upper-left corner to open the Scheduler configuration and edit the `plugins` field so it contains the following:
    ```yaml
    multiPoint:
      enabled:  
      - name: RegexScheduling
    ```
5. Profit!