FROM gobuffalo/buffalo:v0.16.14 as builder

ENV GO111MODULE on
ENV GOPROXY http://proxy.golang.org

RUN mkdir -p /src/github.com/pasiasty/archer
WORKDIR /src/github.com/pasiasty/archer

# this will cache the npm install step, unless package.json changes
ADD package.json .
ADD yarn.lock .
RUN yarn install --no-progress
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

ADD . .
RUN buffalo build --static -o /bin/app

FROM alpine

LABEL "com.centurylinklabs.watchtower.enable"="true"

RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

COPY --from=builder /bin/app .
COPY --from=builder /src/github.com/pasiasty/archer/public public

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

CMD /bin/app pop migrate; /bin/app
