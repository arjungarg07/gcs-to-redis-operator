FROM golang:1.19-buster as build

WORKDIR /app

RUN go env -w GOPRIVATE=github.com/ShareChat

ARG GITHUB_TOKEN

RUN git config \
 --global \
 url."https://${GITHUB_TOKEN}@github.com".insteadOf \
 "https://github.com"

RUN apt-get -y update
RUN apt-get install -y curl
RUN curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > /tmp/google-cloud-sdk.tar.gz
RUN echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
RUN apt-get install -y apt-transport-https ca-certificates gnupg
RUN curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
RUN apt-get update && apt-get install -y google-cloud-sdk

# Copy go mod and sum files
COPY go.mod /app
COPY go.sum /app

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

ENV ACTIVE_ENV PRODUCTION

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN chmod +x /app/run.sh

RUN go build -o main main.go

ENTRYPOINT ["./run.sh"]
