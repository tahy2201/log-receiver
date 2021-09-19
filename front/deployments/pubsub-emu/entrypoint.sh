#!/bin/bash

set -eu

gcloud beta emulators pubsub start --project=lcl-log-receiver --host-port=0.0.0.0:8085 &

# create topic&subscription
cd /root/python-pubsub/samples/snippets/

array=( \
 "tp-trace-log-a"\
 "tp-trace-log-a"
)

for tp_name in "${array[@]}"
do
    python3 publisher.py lcl-log-receiver create ${tp_name}
    python3 subscriber.py lcl-log-receiver create ${tp_name} "sub-${tp_name}"
done

# shell終了回避
while : ; do sleep 1; done
