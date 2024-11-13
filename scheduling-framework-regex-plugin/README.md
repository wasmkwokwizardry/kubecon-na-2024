# Regex Scheduling Plugin implemented as a Scheduler Framework Plugin in Golang

This is an implementation which uses the [Scheduling Framework](https://github.com/kubernetes-sigs/scheduler-plugins) implementation in Go.

## How to run

1. Run `make start-simulator-extender` from the project root.
2. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
3. Click the cog icon in the upper-left corner to open the Scheduler configuration and add the following top-level config:
    ```yaml
    extenders:
    - urlPrefix: http://extender:8000/
      filterVerb: filter
      weight: 10
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
7. Add an invalid regex annotation and show result:
    ```yaml
    annotations:
      scheduler.wasmkwokwizardry.io/regex: 'kubecon-.*\'
    ```
