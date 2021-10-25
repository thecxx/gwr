package gwr_test

import (
	"net/http"
	"testing"

	"github.com/thecxx/gwr"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	// Params
	var params struct {
		FromUri string `wr-p:"(v)=>{no-empty||参数校验失败}(Q.from_uri)"`
		Age     int    `wr-p:"range(1,20)||参数校验失败"`
	}
	wr := gwr.Attach(w, r)
	// Auto flush
	defer wr.Flush()

	err := wr.Import(&params, nil)
	if err != nil {

	}

	wr.Remote().IsMobile()
	wr.Form("name")
	wr.Query("name")
	wr.Param("name")

}

func TestQuery(t *testing.T) {

}
