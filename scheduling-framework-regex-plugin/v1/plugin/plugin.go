package plugin

import (
	"context"
	"fmt"
	"regexp"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

// RegexScheduling is a plugin that schedules pods based on a regex annotation.
type RegexScheduling struct{}

var _ framework.FilterPlugin = &RegexScheduling{}

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "RegexScheduling"
	// regexAnnotationKey is the key for the pod annotation that defines the regex.
	regexAnnotationKey = "scheduler.wasmkwokwizardry.io/regex"
)

// Name returns name of the plugin. It is used in logs, etc.
func (r *RegexScheduling) Name() string {
	return Name
}

// Filter filters out nodes that do not match the regex in the pod annotation, if it is defined, otherwise it returns success.
func (r *RegexScheduling) Filter(_ context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.InfoS("execute Filter on RegexScheduling plugin", "pod", klog.KObj(pod), "node", klog.KObj(nodeInfo.Node()))

	// If the regex annotation is not found, return success.
	pattern, ok := pod.Annotations[regexAnnotationKey]
	if !ok {
		return framework.NewStatus(framework.Success, "")
	}

	// If the regex is not valid, return an error status.
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("failed to compile regex %q: %s", pattern, err))
	}

	// If the regex matches the node name, return success.
	if regex.MatchString(nodeInfo.Node().GetName()) {
		return framework.NewStatus(framework.Success, "")
	}

	// Otherwise, return an unschedulable status.
	return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("node %q does not match regex %q", nodeInfo.Node().GetName(), regex))
}

// New initializes a new RegexScheduling plugin and returns it.
func New(_ context.Context, _ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &RegexScheduling{}, nil
}
