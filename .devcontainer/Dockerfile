FROM golang:1.26

# Instala dependências
RUN apt-get update && apt-get install -y \
    zsh \
    git \
    curl \
    wget \
    unzip \
    sudo \
    && rm -rf /var/lib/apt/lists/*

# Define usuário vscode (melhor prática no devcontainer)
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=$USER_UID

RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME -s /bin/zsh \
    && echo "$USERNAME ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers

USER $USERNAME
WORKDIR /home/$USERNAME

# Instala Oh My Zsh
RUN sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended

# Instala plugins extras
RUN git clone https://github.com/zsh-users/zsh-autosuggestions \
    ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions

RUN git clone https://github.com/zsh-users/zsh-syntax-highlighting \
    ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting

RUN git clone https://github.com/marlonrichert/zsh-autocomplete \
    ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autocomplete

# Ativa plugins
RUN sed -i 's/plugins=(git)/plugins=(git zsh-autosuggestions zsh-autocomplete zsh-syntax-highlighting)/' ~/.zshrc

# Define zsh como padrão
CMD ["zsh"]
