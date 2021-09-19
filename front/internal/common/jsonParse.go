package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

const (
	TRACE_LOG_A = "trace-log-a"
	TRACE_LOG_B = "trace-log-b"
)

// json定義を読み込む
// k:name, v:型
var logSchema = make(map[string]map[string]BqSchemaElement)

func init() {
	logSchema[TRACE_LOG_A] = loadJsonSchema("./schema/trace-log-a.json")
}

type BqSchemaElement struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Mode string `json:"mode"`
}

func loadJsonSchema(filePath string) map[string]BqSchemaElement {
	// ファイル読み込み
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	bqSchemaList := []BqSchemaElement{}
	err2 := json.Unmarshal(bytes, &bqSchemaList)
	if err2 != nil {
		panic(err2)
	}

	bqSchemaMap := make(map[string]BqSchemaElement)
	for _, e := range bqSchemaList {
		bqSchemaMap[e.Name] = e
	}

	return bqSchemaMap
}

// 該当のjsonを受け取り、必要なものだけ返す
func GetParamMap(logType string, jb []byte) (map[string]string, error) {
	pm := make(map[string]string)
	err := json.Unmarshal(jb, &pm)
	if err != nil {
		return nil, err
	}
	retMap := make(map[string]string)

	for k, v := range pm {
		// BqSchemaで定義されていないカラムは無視する。
		if val, ok := logSchema[logType][k]; ok {
			if (val.Type == "STRING") ||
				(val.Type == "INTEGER" && regexp.MustCompile(`\d+`).MatchString(v)) ||
				(val.Type == "TIMESTAMP" && regexp.MustCompile(`\d{4}-\d{2}-\d{2}\s\d{2}:\d{2}:\d{2}\s\+\d{2}:\d{2}`).MatchString(v)) ||
				(val.Type == "DOUBLE" && regexp.MustCompile(`\d+\.\d+`).MatchString(v)) {
				retMap[k] = v
			} else {
				return nil, fmt.Errorf("validation error. column :%v, value :%v", k, v)
			}
		}
	}
	return retMap, nil
}
