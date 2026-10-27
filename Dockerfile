FROM debian:trixie-slim
WORKDIR /usr/local/app

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        make ninja-build clang golang ca-certificates git python3-pip clang-tools \
    && pip install --no-cache-dir --break-system-packages cmake \
    && rm -rf /var/lib/apt/lists/*

COPY . .
RUN mkdir -p bin/corruptor bin/tinywebutils
EXPOSE 8000

CMD ["make"]