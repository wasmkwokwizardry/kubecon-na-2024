FROM python:3.13-alpine

RUN pip install --no-cache uv

WORKDIR /app
COPY pyproject.toml uv.lock .
RUN uv sync --locked

# workaround the fact that uv currently eats our signals
# once https://github.com/astral-sh/uv/issues/6724 ships, we can remove this
STOPSIGNAL SIGKILL

ENV PYTHONUNBUFFERED=1

ARG main=regex.py
COPY ${main} .

ENV MAIN=${main}
CMD uv run fastapi run ${MAIN}