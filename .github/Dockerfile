
# Copyright 2022 Kirill Scherba <kirill@scherba.ru>.  All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.
#
# Teonet teofortune server application docker file
#
# Docker build (included private repositories):
# --------------------------------------------
#
#   docker build --build-arg github_user="${USERNAME}" --build-arg github_personal_token="${TOKEN_FOR_REPOSITORIES}" -t teofortune -f ./.github/Dockerfile .
#
# Docker test run:
# ---------------
#
#   docker run --network host -it --rm teofortune teofortune
#
# Publish to github:
# -----------------
#
#   docker login ghcr.io -u USERNAME -p TOKEN_FOR_PACKAGES
#   #
#   docker tag teocfg ghcr.io/teonet-go/teofortune:0.0.0
#   docker push ghcr.io/teonet-go/teofortune:0.0.0
#   #
#   docker tag teofortune ghcr.io/teonet-go/teofortune:0.0.0
#   docker push ghcr.io/teonet-go/teofortune:0.0.0
#
# Run from gihub repo:
# -------------------
#
#   # Copy static and tmpl folders to the host $HOME/.config/teonet/teofortune
#   # cp -r static $HOME/.config/teonet/teofortune
#   # cp -r tmpl $HOME/.config/teonet/teofortune
#   docker run --network=host -it --restart=always --name teofortune -v $HOME/.config/teonet/teofortune:/root/.config/teonet/teofortune ghcr.io/teonet-go/teofortune:0.0.1 teofortune -loglevel=debug
#

# Build stage
#
FROM golang:alpine AS builder
# RUN apk add --no-cache git build-base
RUN apk add --no-cache git npm
WORKDIR /go/src/app
COPY . .

# Add the keys from arguments
# ARG github_user
# ENV github_user=$github_user
# ARG github_personal_token
# ENV github_personal_token=$github_personal_token
ARG version
ENV version=$version

# Change github url
# RUN git config \
#     --global \
#     url."https://${github_user}:${github_personal_token}@github.com".insteadOf \
#     "https://github.com"

# RUN npm i @vue/cli-service @vue/cli-plugin-babel
# RUN cd frontend && npm i && npm run build
RUN go get
RUN go install ./

# Final stage
#
FROM alpine:latest
RUN apk --no-cache add ca-certificates fortune
COPY --from=builder /go/bin/* /usr/local/bin/
CMD "teofortune"
LABEL Name=teofortune Version=$(version)
