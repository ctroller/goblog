FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Set environment variables for Go
ENV CGO_ENABLED=0 \
    GOPATH=/home/vscode/go \
    GOBIN=/home/vscode/go/bin

# Update and install dependencies
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git curl && \
    # Clean up
    apt-get autoremove -y && apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install act for local workflow testing
RUN curl --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

# Tailwind Standalone
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.6/tailwindcss-linux-x64 && \
    chmod +x tailwindcss-linux-x64 && \
    mv tailwindcss-linux-x64 /usr/local/bin/tailwindcss

# Install Go Tools
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/golang/mock/mockgen@latest && \
    go install golang.org/x/tools/cmd/godoc@latest && \
    go install golang.org/x/perf/cmd/benchstat@latest

RUN chown -R vscode:vscode /home/vscode/go

# Switch to a non-root user - this needs to be down here so the above commands will be run as root
USER vscode

# Set the working directory
WORKDIR /workspace

# Keep the container running
CMD ["sleep", "infinity"]