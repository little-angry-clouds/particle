FROM ubuntu:20.04 as main

ARG YAMLLINT_VERSION=1.26.3
ARG KIND_VERSION=v0.11.1
ARG HELM_VERSION=v3.7.0
ARG BATS_VERSION=v1.4.1
ARG KUBEVAL_VERSION=v0.16.1
ARG KUBE_SCORE_VERSION=1.12.0
ARG KUBE_LINTER_VERSION=0.2.5
ARG HELMFILE_VERSION=v0.141.0

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
# Install yamllint
RUN pip install yamllint==$YAMLLINT_VERSION
# Install helm
RUN curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 && chmod 700 get_helm.sh && bash get_helm.sh --version $HELM_VERSION && rm get_helm.sh
# Install kubectl
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.22.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.20.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.20.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.18.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.18.0
RUN curl -fsSL -o kubectl https://dl.k8s.io/release/v1.16.0/bin/linux/amd64/kubectl && chmod 700 kubectl && mv kubectl /usr/local/bin/kubectl-1.16.0
# Install bats
RUN git clone https://github.com/bats-core/bats-core.git && cd bats-core && ./install.sh /usr/local && cd .. && rm -r bats-core
# Install kubeval
RUN curl -fsSL -o kubeval-linux-amd64.tar.gz https://github.com/instrumenta/kubeval/releases/download/$KUBEVAL_VERSION/kubeval-linux-amd64.tar.gz && tar xf kubeval-linux-amd64.tar.gz && mv kubeval /usr/local/bin && rm kubeval-linux-amd64.tar.gz
# Install kube-score
RUN curl -fsSL -o kube-score https://github.com/zegl/kube-score/releases/download/v$KUBE_SCORE_VERSION/kube-score_${KUBE_SCORE_VERSION}_linux_amd64 && chmod +x kube-score && mv kube-score /usr/local/bin
# Install kube-linter
RUN curl -fsSL -o kube-linter.tar.gz https://github.com/stackrox/kube-linter/releases/download/$KUBE_LINTER_VERSION/kube-linter-linux.tar.gz && tar xf kube-linter.tar.gz && mv kube-linter /usr/local/bin && rm kube-linter.tar.gz
# Install helmfile
RUN curl -fsSL -o helmfile https://github.com/roboll/helmfile/releases/download/$HELMFILE_VERSION/helmfile_linux_amd64 && chmod +x helmfile && mv helmfile /usr/local/bin

COPY particle /usr/local/bin/particle
