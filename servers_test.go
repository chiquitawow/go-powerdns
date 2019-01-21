package powerdns

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

var wantServers = []byte(`[{
	"config_url": "/api/v1/servers/localhost/config{/config_setting}",
	"daemon_type": "authoritative",
	"id": "localhost",
	"type": "Server",
	"url": "/api/v1/servers/localhost",
	"version": "4.2.0-alpha1",
	"zones_url": "/api/v1/servers/localhost/zones{/zone}"
	}]`)

func TestServerService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/servers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		//fmt.Fprint(w, `[{
		//	"config_url": "/api/v1/servers/localhost/config{/config_setting}",
		//	"daemon_type": "authoritative",
		//	"id": "localhost",
		//	"type": "Server",
		//	"url": "/api/v1/servers/localhost",
		//"version": "4.2.0-alpha1",
		//"zones_url": "/api/v1/servers/localhost/zones{/zone}"
		//	}]`)
		w.Write(wantServers)
	})

	got, _, err := client.Servers.Get(context.Background())
	if err != nil {
		t.Errorf("Servers.Get returned error: %v", err)
	}

	ptrSrvs := Server{}
	json.Unmarshal(wantServers, ptrSrvs)
	fmt.Printf("> %v\n", ptrSrvs)

	if !reflect.DeepEqual(got, wantServers) {
		t.Errorf("Servers.Get returned %+v,\n want %+v", got, wantServers)
	}
}
