FROM ubuntu:18.04 as build
RUN apt-get update
RUN apt-get -y upgrade

ENV GOLANG_VERSION 1.16.10

# Install wget
RUN apt update && apt install -y build-essential wget
# Install Go
RUN wget https://golang.org/dl/go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN rm -f go${GOLANG_VERSION}.linux-amd64.tar.gz

ENV PATH "$PATH:/usr/local/go/bin"

RUN go version
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
RUN go mod tidy
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o ./etlTrigger

FROM clickhouse/clickhouse-server:22.3.2

RUN apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3E4AD4719DDE9A38

RUN apt-get update && apt-get install -y \
    curl \
    groff \
    less \
    unzip \
    wget \
    libffi-dev \
    uuid-runtime \
    python3-pip;
RUN curl https://dl.google.com/dl/cloudsdk/release/google-cloud-sdk.tar.gz > /tmp/google-cloud-sdk.tar.gz
RUN pip3 install "pip>=20"
RUN pip3 install pyopenssl
RUN mkdir -p /usr/local/gcloud \
  && tar -C /usr/local/gcloud -xvf /tmp/google-cloud-sdk.tar.gz \
  && /usr/local/gcloud/google-cloud-sdk/install.sh
# Adding the package path to local
ENV PATH $PATH:/usr/local/gcloud/google-cloud-sdk/bin

RUN mkdir /app
COPY --from=build /app/etlTrigger /app/
COPY --from=build /app/src /app/src

ENTRYPOINT [ ]
CMD ["/app/etlTrigger"]