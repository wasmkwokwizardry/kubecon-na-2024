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

var (
	_ framework.FilterPlugin    = &RegexScheduling{}
	_ framework.PreFilterPlugin = &RegexScheduling{}
)

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "RegexScheduling"
	// regexAnnotationKey is the key for the pod annotation that defines the regex.
	regexAnnotationKey = "scheduler.wasmkwokwizardry.io/regex"
	// preFilterStateKey is the key for the state that stores the regex.
	preFilterStateKey = "PreFilterRegex" + Name
)

// Name returns name of the plugin. It is used in logs, etc.
func (r *RegexScheduling) Name() string {
	return Name
}

// preFilterState computed at PreFilter and used at Filter.
type preFilterState struct {
	regex *regexp.Regexp
}

// Clone implements the mandatory Clone interface. We don't really copy the data since
// there is no need for that.
func (s *preFilterState) Clone() framework.StateData {
	return s
}

// PreFilter looks for a regex annotation in the pod and stores the compiled version in the state.
func (r *RegexScheduling) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	klog.InfoS("execute PreFilter on RegexScheduling plugin", "pod", klog.KObj(pod))

	// If the regex annotation is not found, return success.
	pattern, ok := pod.Annotations[regexAnnotationKey]
	if !ok {
		return nil, framework.NewStatus(framework.Success, "")
	}

	// If the regex is not valid, return an error status.
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, framework.NewStatus(framework.Error, fmt.Sprintf("failed to compile regex %q: %s", pattern, err))
	}

	// Store the regex in the state.
	state.Write(preFilterStateKey, &preFilterState{regex: regex})

	return nil, framework.NewStatus(framework.Success, "")
}

// Filter filters out nodes that do not match the regex in the state, if it is defined, otherwise it returns success.
func (r *RegexScheduling) Filter(_ context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.InfoS("execute Filter on RegexScheduling plugin", "pod", klog.KObj(pod), "node", klog.KObj(nodeInfo.Node()))

	// Load the regex from the state.
	data, err := state.Read(preFilterStateKey)
	if err != nil {
		// If the regex is not found in the state, return success.
		if err == framework.ErrNotFound {
			return framework.NewStatus(framework.Success, "")
		}

		// This should never happen.
		return framework.NewStatus(framework.Error, err.Error())
	}

	// If the regex matches the node name, return success.
	regex := data.(*preFilterState).regex
	if regex.MatchString(nodeInfo.Node().GetName()) {
		return framework.NewStatus(framework.Success, "")
	}

	// Otherwise, return an unschedulable status.
	return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("node %q does not match regex %q", nodeInfo.Node().GetName(), regex))
}

// PreFilterExtensions implements the mandatory PreFilterExtensions interface.
func (r *RegexScheduling) PreFilterExtensions() framework.PreFilterExtensions {
	return r
}

// AddPod implements the mandatory AddPod method of the framework.PreFilterPlugin interface.
// It's a no-op for this plugin.
func (r *RegexScheduling) AddPod(
	_ context.Context,
	_ *framework.CycleState,
	_ *v1.Pod,
	_ *framework.PodInfo,
	_ *framework.NodeInfo,
) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// RemovePod implements the mandatory RemovePod method of the framework.PreFilterPlugin interface.
// It's a no-op for this plugin.
func (r *RegexScheduling) RemovePod(
	_ context.Context,
	_ *framework.CycleState,
	_ *v1.Pod,
	_ *framework.PodInfo,
	_ *framework.NodeInfo,
) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// New initializes a new RegexScheduling plugin and returns it.
func New(_ context.Context, _ runtime.Object, _ framework.Handle) (framework.Plugin, error) {
	return &RegexScheduling{}, nil
}
