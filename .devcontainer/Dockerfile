FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

# Set environment variables for Go
ENV CGO_ENABLED=0

# Update and install dependencies
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y git curl && \
    # Clean up
    apt-get autoremove -y && apt-get clean -y && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install Go Tools
RUN go install golang.org/x/tools/gopls@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install golang.org/x/lint/golint@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install golang.org/x/tools/cmd/goimports@latest && \
    go install github.com/golang/mock/mockgen@latest && \
    go install golang.org/x/tools/cmd/godoc@latest && \
    go install golang.org/x/perf/cmd/benchstat@latest

# Switch to a non-root user
USER vscode

# Set the working directory
WORKDIR /workspace

# Keep the container running
CMD ["sleep", "infinity"]