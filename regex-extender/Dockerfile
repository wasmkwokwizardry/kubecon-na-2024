FROM python:3.13-alpine

RUN pip install --no-cache uv

WORKDIR /app
COPY pyproject.toml uv.lock ./
RUN uv sync --locked

COPY . ./

CMD ["uv", "run", "fastapi", "run", "main.py"]
