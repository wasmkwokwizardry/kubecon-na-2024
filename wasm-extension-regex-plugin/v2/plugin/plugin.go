package plugin

import (
	"fmt"
	"regexp"

	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/api"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/api/proto"
	klog "sigs.k8s.io/kube-scheduler-wasm-extension/guest/klog/api"
)

// RegexScheduling is a plugin that schedules pods based on a regex annotation.
type RegexScheduling struct {
	klog klog.Klog
}

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "RegexScheduling"
	// regexAnnotationKey is the key for the pod annotation that defines the regex.
	regexAnnotationKey = "scheduler.wasmkwokwizardry.io/regex"
	// preFilterStateKey is the key for the state that stores the regex.
	preFilterStateKey = "PreFilterRegex" + Name
)

// preFilterState computed at PreFilter and used at Filter.
type preFilterState struct {
	regex *regexp.Regexp
}

// PreFilter looks for a regex annotation in the pod and stores the compiled version in the state.
func (r *RegexScheduling) PreFilter(state api.CycleState, pod proto.Pod) (nodeNames []string, status *api.Status) {
	r.klog.InfoS("execute PreFilter on RegexScheduling plugin", "pod", klog.KObj(pod))

	// If the regex annotation is not found, return success.
	pattern, ok := pod.GetAnnotations()[regexAnnotationKey]
	if !ok {
		return nil, &api.Status{Code: api.StatusCodeSuccess}
	}

	// If the regex is not valid, return an error status.
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, &api.Status{Code: api.StatusCodeError, Reason: fmt.Sprintf("failed to compile regex %q: %s", pattern, err)}
	}

	// Store the regex in the state.
	state.Write(preFilterStateKey, &preFilterState{regex: regex})

	return nil, &api.Status{Code: api.StatusCodeSuccess}
}

// Filter filters out nodes that do not match the regex in the state, if it is defined, otherwise it returns success.
func (r *RegexScheduling) Filter(state api.CycleState, pod proto.Pod, nodeInfo api.NodeInfo) *api.Status {
	r.klog.InfoS("execute Filter on RegexScheduling plugin", "pod", klog.KObj(pod), "node", klog.KObj(nodeInfo.Node()))

	// Load the regex from the state.
	data, ok := state.Read(preFilterStateKey)

	// If the regex is not found in the state, return success.
	if !ok {
		return &api.Status{Code: api.StatusCodeSuccess}
	}

	// If the regex matches the node name, return success.
	regex := data.(*preFilterState).regex
	if regex.MatchString(nodeInfo.Node().GetName()) {
		return &api.Status{Code: api.StatusCodeSuccess}
	}

	// Otherwise, return an unschedulable status.
	return &api.Status{Code: api.StatusCodeUnschedulable, Reason: fmt.Sprintf("node %q does not match regex %q", nodeInfo.Node().GetName(), regex)}
}

// New initializes a new RegexScheduling plugin and returns it.
func New(klog klog.Klog, jsonConfig []byte) (*RegexScheduling, error) {
	return &RegexScheduling{klog: klog}, nil
}
