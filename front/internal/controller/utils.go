package controller

import (
	"encoding/json"
	"net/http"

	pubsub "cloud.google.com/go/pubsub"
	"google.golang.org/appengine"
)

type StdResponse struct {
	Result  int
	Message string
	MsgId   string
}

func writeErrorResponse(w http.ResponseWriter, msg string) http.ResponseWriter {
	res, _ := json.Marshal(StdResponse{1, msg, ""})
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return w
}

/**
指定されたPub/sub topicにメッセージを送る。
tn:topic name
d:message for publishing to topic
*/
func publishMessage(r *http.Request, tn string, val map[string]string) (string, string, error) {

	d, err := json.Marshal(val)
	if err != nil {
		return "", "error convert json.", err
	}

	// ProjectIDを取得してクライアント作成
	ctx := appengine.NewContext(r)
	client, err := pubsub.NewClient(ctx, appengine.AppID(ctx))
	if err != nil {
		return "", "error creating pubsub client.", err
	}

	// トピックに配信
	tp := client.Topic(tn)
	pubResult := tp.Publish(ctx, &pubsub.Message{Data: d})
	id, err := pubResult.Get(ctx)
	if err != nil {
		return "", "error from publishing " + tn, err
	}
	return id, "", nil
}
