package main

import (
	"os"

	"github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/pkg/plugin/regex"

	"k8s.io/component-base/cli"
	_ "k8s.io/component-base/metrics/prometheus/clientgo" // for rest client metric registration
	_ "k8s.io/component-base/metrics/prometheus/version"  // for version metric registration
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	// Register custom plugins to the scheduler framework.
	// Later they can consist of scheduler profile(s) and hence
	// used by various kinds of workloads.
	command := app.NewSchedulerCommand(
		app.WithPlugin(regex.Name, regex.New),
	)

	code := cli.Run(command)
	os.Exit(code)
}
