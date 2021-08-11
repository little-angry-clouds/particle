FROM ubuntu:20.04

ARG KUBEVAL_VERSION=v0.16.1
ARG YAMLLINT_VERSION=1.26.2
ARG BATS_VERSION=v1.4.1
ARG KIND_VERSION=v0.11.1
ARG HELM_VERSION=v3.6.3

RUN apt update && apt install --no-install-recommends -y python3-pip curl git apt-transport-https gnupg lsb-release ca-certificates
# Install docker
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
RUN echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | tee /etc/apt/sources.list.d/docker.list > /dev/null
RUN apt update && apt install --no-install-recommends -y docker-ce docker-ce-cli containerd.io
# Clean apt caché
RUN apt-get clean autoclean && apt-get autoremove --yes && rm -rf /var/lib/apt/lists/*
# Install kind
RUN curl -fsSL -o kind https://kind.sigs.k8s.io/dl/$KIND_VERSION/kind-linux-amd64 && chmod +x ./kind && mv kind /usr/local/bin
# Install kubeval
RUN curl -fsSL -o kubeval-linux-amd64.tar.gz https://github.com/instrumenta/kubeval/releases/download/$KUBEVAL_VERSION/kubeval-linux-amd64.tar.gz && tar xf kubeval-linux-amd64.tar.gz && mv kubeval /usr/local/bin && rm kubeval-linux-amd64.tar.gz
# Install yamllint
RUN pip install yamllint==$YAMLLINT_VERSION
# Install bats
RUN git clone https://github.com/bats-core/bats-core.git && cd bats-core && ./install.sh /usr/local && cd .. && rm -r bats-core
# Install helm
RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 && chmod 700 get_helm.sh && bash get_helm.sh --version $HELM_VERSION && rm get_helm.sh
# Install kubectl
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.22.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.20.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.18.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.18.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.16.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.16.0

COPY particle /usr/local/bin/
