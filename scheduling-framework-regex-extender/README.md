# Regex Scheduling Plugin implemented as a Kubernetes Scheduler Extender in Python

This is a simple extender sample in Python.

It filters out nodes that do not match a regular expression found in the `scheduler.wasmkwokwizardry.io/regex` pod annotation.

Add the following lines to the scheduler config to enable the extender:

```yaml
extenders:
  - urlPrefix: http://extender:8000/
    filterVerb: filter
    weight: 10
```