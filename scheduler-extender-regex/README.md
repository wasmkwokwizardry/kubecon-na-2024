# Regex Scheduling Plugin implemented as a Scheduler Extender in Python

This is a simple Scheduler Extender implementation in Python.

## v1

This version is the basic version of the extender. It receives the full node list in the request and filters out nodes that do not match the regex. It then returns the filtered list via 2 fields:
- `Nodes` contains the list of node that match the regex, with the full node info
- `FailedAndUnresolvableNodes` is a map of node name to error message for nodes that do not match.

The code can be found in [`regex_v1.py`](./regex_v1.py).

### How to run

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

## v2

This version adds support for "caching" the node list. This would normally be done by having the extender maintain a cache of the node list and updating it periodically.

When the `nodeCacheCapable` field is set to `true`, the extender will receive a list of node _names_ in the `NodeNames` field of the request instead of the full node info.
In our case, we don't even need the node info, but just the name, so we can actually use this field to filter the nodes. We can then return the list of node names that match the regex in the `Nodes` field instead of the full node info list.

The code can be found in [`regex_v2.py`](./regex_v2.py).

### How to run

1. Run `make start-simulator-extender-v2` from the project root.
2. Open the Simulator at [http://localhost:3000](http://localhost:3000) in your browser.
3. Click the cog icon in the upper-left corner to open the Scheduler configuration and edit the `extenders` section to include the `nodeCacheCapable` field:
    ```yaml
    extenders:
    - urlPrefix: http://extender:8000/
      filterVerb: filter
      weight: 10
      nodeCacheCapable: true
    ```
4. Profit!