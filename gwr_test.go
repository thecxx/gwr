package gwr_test

import (
	"net/http"
	"testing"

	"github.com/thecxx/gwr"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Params
	var params struct {
		Age     int    `wr-param:"{url.age: $0<10||20<$0, default: 20, #: \"参数校验失败\"}"`
		FromUri string `wr-param:"{header.from_uri: /^[0-9]+$/.match($0), @default: \"https://www.baidu.com\", @error: \"参数校验失败\"}"`
	}
	wr := gwr.Attach(w, r)
	// Auto flush
	defer wr.Flush()

	err := wr.Extract(&params, nil)
	if err != nil {

	}

}

func TestQuery(t *testing.T) {

}
