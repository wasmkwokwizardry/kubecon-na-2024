package v2

import (
	"context"
	"fmt"
	"regexp"

	"github.com/wasmkwokwizardry/kubecon-na-2024/scheduling-framework-regex-plugin/pkg/plugin/regex"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

var (
	// Name is the name of the plugin used in Registry and configurations.
	Name = regex.Name
)

type RegexScheduling struct{}

// New initializes a new RegexScheduling plugin and returns it.
func New(ctx context.Context, obj runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &RegexScheduling{}, nil
}

// Name returns name of the plugin. It is used in logs, etc.
func (r RegexScheduling) Name() string {
	return regex.Name
}

func (r RegexScheduling) PreFilter(ctx context.Context, state *framework.CycleState, pod *v1.Pod) (*framework.PreFilterResult, *framework.Status) {
	logger := klog.FromContext(ctx)

	regexString := pod.Annotations[regex.RegexAnnotationKey]
	if regexString == "" {
		return nil, framework.NewStatus(framework.Success, "")
	}

	regex, err := regexp.Compile(regexString)
	if err != nil {
		logger.Error(err, "Failed to compile regex", "regex", regexString)
		return nil, framework.NewStatus(framework.Success, "")
	}

	logger.Info("Caching the compiled regex", "regex", regexString)
	state.Write(framework.StateKey(regexString), &PreFilterState{regex: regex})

	return nil, framework.NewStatus(framework.Success, "")
}

// PreFilterState computed at PreFilter and used at Filter.
type PreFilterState struct {
	regex *regexp.Regexp
}

// Clone the preFilter state.
func (s *PreFilterState) Clone() framework.StateData {
	return s
}

func (r RegexScheduling) PreFilterExtensions() framework.PreFilterExtensions {
	return r
}

func (r RegexScheduling) AddPod(
	_ context.Context,
	_ *framework.CycleState,
	_ *v1.Pod,
	_ *framework.PodInfo,
	_ *framework.NodeInfo,
) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

func (r RegexScheduling) RemovePod(
	_ context.Context,
	_ *framework.CycleState,
	_ *v1.Pod,
	_ *framework.PodInfo,
	_ *framework.NodeInfo,
) *framework.Status {
	return framework.NewStatus(framework.Success, "")
}

// Filter filters out nodes that do not match the regex in the pod annotation, if it is defined, otherwise it returns success.
func (r RegexScheduling) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	logger := klog.FromContext(ctx)

	regexString := pod.Annotations[regex.RegexAnnotationKey]
	if regexString == "" {
		return framework.NewStatus(framework.Success, "")
	}
	stateData, err := state.Read(framework.StateKey(regexString))
	if err != nil {
		return framework.NewStatus(framework.Error, err.Error())
	}
	if stateData != nil {
		logger.Info("Found existing state", "key", regexString)
		preFilterState, ok := stateData.(*PreFilterState)
		if !ok {
			return framework.NewStatus(framework.Error, fmt.Sprintf("invalid state: expected &preFilterState, got %T", stateData))
		}
		if preFilterState.regex == nil {
			return framework.NewStatus(framework.Error, "invalid state: compiled regex is nil")
		}
		if match := preFilterState.regex.MatchString(nodeInfo.Node().Name); !match {
			return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("Node %q does not match regex %q", nodeInfo.Node().Name, regexString))
		}
	} else {
		logger.Info("State not found", "key", regexString)
		regex, err := regexp.Compile(regexString)
		if err != nil {
			return framework.NewStatus(framework.Error, fmt.Sprintf("Failed to compile regex %q: %s", regexString, err))
		}
		if match := regex.MatchString(nodeInfo.Node().Name); !match {
			return framework.NewStatus(framework.Unschedulable, fmt.Sprintf("Node %q does not match regex %q", nodeInfo.Node().Name, regexString))
		}
	}
	return framework.NewStatus(framework.Success, "")
}

var _ framework.FilterPlugin = &RegexScheduling{}
var _ framework.PreFilterPlugin = &RegexScheduling{}
