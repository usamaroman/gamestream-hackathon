FROM python:3.12.1-slim

ENV PYTHONPATH "/app"

RUN apt-get update && apt-get install -y \
    ffmpeg libsm6 libxext6 \
    libpq-dev \
    gcc \
    && apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY poetry.lock pyproject.toml /app/

RUN pip install poetry==1.7.1 && poetry install

COPY . .

RUN chmod +x /app/entrypoint.sh

CMD ["/app/entrypoint.sh"]