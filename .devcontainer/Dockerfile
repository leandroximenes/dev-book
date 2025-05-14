# Use your existing base image
FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install necessary packages
RUN apt-get update && apt-get install -y \
    curl \
    jq \
    git \
    sudo \
    && rm -rf /var/lib/apt/lists/*

# Install GitHub CLI
RUN mkdir -p -m 755 /etc/apt/keyrings && \
    wget -qO- https://cli.github.com/packages/githubcli-archive-keyring.gpg | tee /etc/apt/keyrings/githubcli-archive-keyring.gpg > /dev/null && \
    chmod go+r /etc/apt/keyrings/githubcli-archive-keyring.gpg && \
    echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | tee /etc/apt/sources.list.d/github-cli.list > /dev/null && \
    apt-get update && \
    apt-get install gh -y

# Install Delve for debugging (if needed)
RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Add /root/bin to PATH (if needed)
ENV PATH="${PATH}:/root/bin"

# Set the workspace directory
ENV EXERCISM_WORKSPACE=/workspaces/exercism

# Add /workspaces as a safe directory in Git configuration
RUN git config --global --add safe.directory /workspaces

# Set the working directory
WORKDIR /workspace

# Set devcontainer 
RUN chown -R vscode:vscode /go

# Expose the debugging port
EXPOSE 2345 3000