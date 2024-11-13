package regex

import (
	"context"
	"fmt"
	"regexp"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	// Name is the name of the plugin used in Registry and configurations.
	Name = "RegexScheduling"
	// RegexAnnotationKey is the key for the pod annotation that defines the regex.
	RegexAnnotationKey = "scheduler.wasmkwokwizardry.io/regex"
)

type RegexScheduling struct{}

// New initializes a new RegexScheduling plugin and returns it.
func New(_ context.Context, _ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &RegexScheduling{}, nil
}

// Name returns name of the plugin. It is used in logs, etc.
func (r RegexScheduling) Name() string {
	return Name
}

// Filter filters out nodes that do not match the regex in the pod annotation, if it is defined, otherwise it returns success.
func (r RegexScheduling) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	regex, ok := pod.Annotations[RegexAnnotationKey]
	if !ok {
		return framework.NewStatus(framework.Success, "")
	}
	match, err := regexp.MatchString(regex, nodeInfo.Node().Name)
	if err != nil {
		return framework.NewStatus(framework.Error, fmt.Sprintf("Failed to compile regex %q: %s", regex, err))
	}
	if !match {
		return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("Node %q does not match regex %q", nodeInfo.Node().Name, regex))
	}
	return framework.NewStatus(framework.Success, "")
}

var _ framework.FilterPlugin = &RegexScheduling{}
