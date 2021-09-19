package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	pubsub "cloud.google.com/go/pubsub"
)

func TracelogB(w http.ResponseWriter, r *http.Request) {
	fmt.Println("call TracelogB")

	// Body の内容をそのまま[]byteに
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	bsJson := make(map[string]string)
	json.Unmarshal(bs, &bsJson)
	fmt.Printf("Published message with custom attributes; msg ID: %s\n", bsJson)

	// PtojectIDを取得してクライアント作成
	// ctx := appengine.NewContext(r)
	// client, err := pubsub.NewClient(ctx, appengine.AppID(ctx))
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "dev-log-receiver")
	if err != nil {
		panic(err)
	}

	// トピックに配信
	tp := client.Topic("tp-tracelog-b")
	pubResult := tp.Publish(ctx, &pubsub.Message{Attributes: bsJson})
	id, err := pubResult.Get(ctx)
	if err != nil {
		fmt.Printf("Get: %v", err)
	}
	fmt.Printf("Published message with custom attributes; msg ID: %v\n", id)

	// if _, err := tp.Publish(ctx, &pubsub.Message{Data: bs}); err != nil {
	// 	panic(err)
	// }
	w.WriteHeader(http.StatusNotModified)
}
