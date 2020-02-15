FROM golang:1.12 as builder
ENV TZ=Asia/Tokyo
ARG approot=/app

# basic settings
RUN apt-get update -yq && apt-get install -yq --no-install-recommends \
        postgresql-client \
    && rm -rf /var/lib/apt/lists/*

# go setup
RUN go get -u -t github.com/volatiletech/sqlboiler && \
    go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql && \
    go get github.com/derekparker/delve/cmd/dlv && \
    go get github.com/oxequa/realize

# make woking directory
WORKDIR $approot
ADD . $approot

# build 
RUN make build