import re
from fastapi import FastAPI, Request
from fastapi.responses import JSONResponse

app = FastAPI()


@app.post("/filter")
async def filter(request: Request):
    # Decode the JSON object from the request body
    data = await request.json()

    # Check if the regex pattern is provided in the Pod annotations
    try:
        pattern = data["Pod"]["metadata"]["annotations"][
            "scheduler.wasmkwokwizardry.io/regex"
        ]
    except KeyError:
        return JSONResponse(
            content={
                "Nodes": data["Nodes"],
            }
        )

    # Compile the regex pattern
    try:
        regex = re.compile(pattern)
    except re.error:
        return JSONResponse(
            content={
                "Error": f"Invalid regex pattern: {pattern}",
            }
        )

    # Filter the nodes based on the regex pattern
    nodes = []
    failed_nodes = {}

    for node in data["Nodes"]["items"]:
        name = node["metadata"]["name"]
        if regex.search(name) is None:
            failed_nodes[name] = (
                f"Node {name} does not match the regex pattern {pattern}"
            )
        else:
            nodes.append(node)

    # Return the filtered nodes
    return JSONResponse(
        content={
            "Nodes": {
                "items": nodes,
            },
            "FailedAndUnresolvableNodes": failed_nodes,
        }
    )
