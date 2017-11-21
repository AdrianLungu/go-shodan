package shodan

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestClient_GetHostsForQuery_DifferentVersionFormats(t *testing.T) {
	setUpTestServe()
	defer tearDownTestServe()

	mux.HandleFunc(hostSearchPath, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.Write(getStub(t, "host/version"))
	})

	options := &HostQueryOptions{Query: "argentina"}

	response, err := client.GetHostsForQuery(options)
	if err != nil {
		t.Error(err)
	}

	_, err = json.Marshal(response)

	assert.Nil(t, err)
}
