# log-receiver

いろんなところからログ(trace-log-a)を受け取ってBigQueryに格納していくサンプル
Client -> GAE[このApp] -> Pub/Sub -> data-workflow -> BigQuery

### logalテスト
```
docker-compose up
```

### dev環境デプロイ
```
gcloud dataflow jobs run importing-tracelog-a-from-gcloud \
    --gcs-location gs://dataflow-templates/latest/PubSub_Subscription_to_BigQuery \
    --region asia-northeast1 \
    --staging-location gs://stream-import/tmp-tracelog-a \
    --parameters \
inputSubscription=projects/dev-log-receiver/subscriptions/trace-log-a-sub,\
outputTableSpec=dev-log-receiver:dev_tracelogs.trace_log_a,\
outputDeadletterTable=dev-log-receiver:dev_tracelogs.trace_log_a_error_records

```
