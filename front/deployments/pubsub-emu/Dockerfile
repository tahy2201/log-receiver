FROM getalma/gcp-pubsub-emulator:latest

WORKDIR /root
RUN pip3 install virtualenv
RUN virtualenv env
RUN apk update && apk add g++ python3-dev linux-headers

# とても時間がかかる。pythonイメージ + gcloudの方がいい気がする。
RUN git clone https://github.com/googleapis/python-pubsub.git
RUN pip3 install -r python-pubsub/samples/snippets/requirements.txt

ENTRYPOINT ["gcloud", "beta", "emulators", "pubsub", "start", "--project=lcl-log-receiver", "--host-port=0.0.0.0:8085"]
