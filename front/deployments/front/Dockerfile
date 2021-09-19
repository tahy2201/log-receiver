FROM google/cloud-sdk:347.0.0-alpine

# front階層想定
WORKDIR /usr/local/front

RUN apk update && apk add make musl-dev go python2

# Configure Go
RUN export GOPATH=/root/go
RUN export PATH=${GOPATH}/bin:/usr/local/go/bin:$PATH
RUN export GOBIN=$GOROOT/bin
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
RUN export GO111MODULE=on
RUN go version

RUN gcloud components install app-engine-python app-engine-go cloud-datastore-emulator
RUN gcloud config set project dev-gcp-project

ENTRYPOINT [ "dev_appserver.py" , "--skip_sdk_update_check=yes", "--host=0.0.0.0", "--application=lcl-log-receiver",\
    "--env_var", "PUBSUB_EMULATOR_HOST=pubsub:8085", "app.yaml"]
