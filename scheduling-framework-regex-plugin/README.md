# Regex Scheduling Plugin implemented as a Scheduler Framework Plugin in Golang

This is an implementation which uses the [Scheduling Framework](https://github.com/kubernetes-sigs/scheduler-plugins) implementation in Go.

## How to run

1. Go to the Simulator project by running `cd kube-scheduler-simulator`
2. Add the `regex.go` file to the `simulator/pkg/plugin/regex` directory
3. Add the plugin in the debuggable scheduler by editing the `cmd/scheduler/scheduler.go` and changing the following line:
    ```go
        import (
            "github.com/kubernetes-sigs/scheduler-plugins/pkg/plugin/regex"
        )        

        command, cancelFn, err := debuggablescheduler.NewSchedulerCommand(
            debuggablescheduler.WithPlugin(regex.Name, regex.New),
        )
    ```
4. Build the new Scheduler and run it locally by running the command `make docker_build docker_up_local`
5. Enable the new Scheduler by editing the `KubeSchedulerConfiguration` from the Simulator UI by adding the following line to the `default-scheduler` profile:
    ```yaml
    multiPoint:
      enabled:  
      - name: RegexScheduling
    ```
6. Create a Node with a default name and additional two Nodes named `kubecon-1` and `kubecon-2`
7. Create a Pod with an annotation `"scheduler.example.com/regex": "kubecon-.*"`
    ```yaml
    annotations:
      scheduler.example.com/regex: "kubecon-.*"
    ```
8. Add a regex annotation for a non-existing Node Group and show result:
    ```yaml
    annotations:
      scheduler.example.com/regex: "cncf-.*"
    ```
9. Add an invalid regex annotation and show result:
    ```yaml
    annotations:
      scheduler.example.com/regex: 'kubecon-.*\'
    ```