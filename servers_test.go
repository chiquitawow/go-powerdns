package powerdns

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

var expectedServerStruct = Server{
	ConfigUrl:  "/api/v1/servers/localhost/config{/config_setting}",
	DaemonType: "authoritative",
	ID:         "localhost",
	ServerType: "Server",
	URL:        "/api/v1/servers/localhost",
	Version:    "4.2.0-alpha1",
	ZonesURL:   "/api/v1/servers/localhost/zones{/zone}",
}

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
		w.Write(wantServers)
	})

	got, _, err := client.Servers.Get(context.Background())
	if err != nil {
		t.Errorf("Servers.Get returned error: %v", err)
	}

	want := []Server{expectedServerStruct}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Servers.Get returned %+v,\n want %+v", got, want)
	}
}
