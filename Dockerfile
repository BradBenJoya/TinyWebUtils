FROM debian:trixie-slim
WORKDIR /usr/local/app

RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        make ninja-build clang clang-tools golang ca-certificates git python3-pip \
    && pip install --no-cache-dir --break-system-packages cmake \
    && rm -rf /var/lib/apt/lists/*

ENV CC=clang CXX=clang++

COPY . .
RUN make install build

EXPOSE 8000
CMD ["./bin/tinywebutils/server"]