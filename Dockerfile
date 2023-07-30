FROM ubuntu

WORKDIR /tune-bot/

COPY . ./core

RUN apt-get update && apt-get install -y dos2unix
RUN dos2unix core/infrastructure/install.sh
RUN dos2unix core/infrastructure/create.sql

RUN core/infrastructure/install.sh

RUN rm -rf core && apt-get clean && rm -rf /var/lib/apt/lists/*

RUN bin/database
EXPOSE 3306

ENTRYPOINT ["tail", "-f", "/dev/null"]
