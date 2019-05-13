FROM node:12.2.0

WORKDIR /app

RUN apt-get update && \
    apt-get install -y ssh git build-essential libssl-dev curl && \
    curl -o- -L https://yarnpkg.com/install.sh | sh

ENV PATH $HOME/.yarn/bin:$HOME/.config/yarn/global/node_modules/.bin:$PATH

ENV CHOKIDAR_USEPOLLING=true
RUN yarn global add @vue/cli
RUN yarn global add @vue/cli-init
