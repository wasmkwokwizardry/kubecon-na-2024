package regex

import (
	"fmt"
	"regexp"

	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/api"
	"sigs.k8s.io/kube-scheduler-wasm-extension/guest/api/proto"
	klogapi "sigs.k8s.io/kube-scheduler-wasm-extension/guest/klog/api"
)

func New(klog klogapi.Klog, jsonConfig []byte) (api.Plugin, error) {
	return &RegexScheduling{log: klog}, nil
}

// RegexScheduling is a plugin that schedules pods based on a regex annotation.
type RegexScheduling struct {
	log klogapi.Klog
}

const (
	// Name is the name of the plugin used in the plugin registry and configurations.
	Name = "RegexScheduling"
	// RegexAnnotationKey is the key for the pod annotation that defines the regex.
	RegexAnnotationKey = "scheduler.example.com/regex"
)

func (r *RegexScheduling) Filter(state api.CycleState, pod proto.Pod, nodeInfo api.NodeInfo) *api.Status {
	regex, ok := pod.GetAnnotations()[RegexAnnotationKey]
	if !ok {
		return &api.Status{Code: api.StatusCodeSuccess}
	}
	match, err := regexp.MatchString(regex, nodeInfo.Node().GetName())
	if err != nil {
		return &api.Status{Code: api.StatusCodeError, Reason: fmt.Sprintf("Failed to compile regex %q: %s", regex, err)}
	}
	if !match {
		return &api.Status{Code: api.StatusCodeUnschedulable, Reason: fmt.Sprintf("Node %q does not match regex %q", nodeInfo.Node().GetName(), regex)}
	}
	return &api.Status{Code: api.StatusCodeSuccess}
}
