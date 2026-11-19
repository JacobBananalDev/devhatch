ARG GO_VERSION=1.27.4
FROM golang:${GO_VERSION}

ARG DEV_UID=1000
ARG DEV_GID=1000

RUN set -eux; \
    curl --fail --location --silent --show-error \
        https://oils.pub/download/oils-for-unix-0.38.0.tar.gz \
        --output /tmp/oils-for-unix.tar.gz; \
    printf '%s  %s\n' \
        a33453722819b55ee552bfd7f3c2bab8f1940def55d5c8b46af16ce95bdf8803 \
        /tmp/oils-for-unix.tar.gz \
        | sha256sum --check --status; \
    tar -xzf /tmp/oils-for-unix.tar.gz -C /tmp; \
    cd /tmp/oils-for-unix-0.38.0; \
    ./configure --without-readline; \
    ./_build/oils.sh; \
    ./install; \
    rm -rf /tmp/oils-for-unix-0.38.0 /tmp/oils-for-unix.tar.gz

RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends gnupg nano; \
    install -d -m 0755 /etc/apt/keyrings; \
    curl --fail --location --silent --show-error \
        https://packages.microsoft.com/keys/microsoft.asc \
        | gpg --dearmor --yes -o /etc/apt/keyrings/microsoft.gpg; \
    chmod 0644 /etc/apt/keyrings/microsoft.gpg; \
    printf 'deb [arch=%s signed-by=/etc/apt/keyrings/microsoft.gpg] https://packages.microsoft.com/repos/code stable main\n' \
        "$(dpkg --print-architecture)" \
        > /etc/apt/sources.list.d/vscode.list; \
    apt-get update; \
    apt-get install -y --no-install-recommends code; \
    rm -rf /var/lib/apt/lists/*

WORKDIR /workspace

RUN set -eux; \
    if ! getent group "$DEV_GID" >/dev/null; then \
        groupadd --gid "$DEV_GID" devhatch; \
    fi; \
    if ! getent passwd "$DEV_UID" >/dev/null; then \
        useradd --uid "$DEV_UID" --gid "$DEV_GID" --no-create-home --shell /bin/bash devhatch; \
    fi; \
    mkdir -p /home/devhatch; \
    chown "$DEV_UID:$DEV_GID" /home/devhatch; \
    git config --system --add safe.directory /workspace
