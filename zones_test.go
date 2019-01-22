package powerdns

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var expectedZoneStruct = Zone{
	Account:        "",
	DNSSec:         false,
	ID:             "example.com.",
	Kind:           "Native",
	LastCheck:      0,
	Masters:        []string{},
	Name:           "example.com.",
	NotifiedSerial: 0,
	Serial:         2019011605,
	URL:            "/api/v1/servers/localhost/zones/example.com.",
}

var testZone = []byte(`[{
	"account": "",
	"dnssec": false,
  "id": "example.com.",
  "kind": "Native",
  "last_check": 0,
  "masters": [],
  "name": "example.com.",
  "notified_serial": 0,
  "serial": 2019011605,
  "url": "/api/v1/servers/localhost/zones/example.com."
 }]`)

var testZonePostResp = []byte(`{
  "account": "",
  "api_rectify": false,
  "dnssec": false,
  "id": "example.com.",
  "kind": "Native",
  "last_check": 0,
  "master_tsig_key_ids": [],
  "masters": [],
  "name": "example.com.",
  "notified_serial": 0,
  "nsec3narrow": false,
  "nsec3param": "",
  "rrsets": [
    {
      "comments": [],
      "name": "example.com.",
      "records": [
        {
          "content": "a.misconfigured.powerdns.server. hostmaster.example.com. 2019012201 10800 3600 604800 3600",
          "disabled": false
        }
      ],
      "ttl": 300,
      "type": "SOA"
    },
    {
      "comments": [],
      "name": "example.com.",
      "records": [
        {
          "content": "ns2.example.com.",
          "disabled": false
        },
        {
          "content": "ns1.example.com.",
          "disabled": false
        }
      ],
      "ttl": 300,
      "type": "NS"
    }
  ],
  "serial": 2019012201,
  "slave_tsig_key_ids": [],
  "soa_edit": "",
  "soa_edit_api": "DEFAULT",
  "url": "/api/v1/servers/localhost/zones/example.com."
}`)

func TestZoneService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/v1/servers/localhost/zones", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		w.Write(testZone)
	})

	got, _, err := client.Zones.List(context.Background())
	if err != nil {
		t.Errorf("Zones.List returned error: %v", err)
	}

	want := []Zone{expectedZoneStruct}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Zones.List returned %+v,\n want %+v", got, want)
	}
}

func TestZoneService_Post(t *testing.T) {

	var wantRRSetSOA = RRSet{
		Comments: nil,
		Name:     "example.com.",
		Records: []Record{
			Record{
				Content:  "a.misconfigured.powerdns.server. hostmaster.example.com. 2019012201 10800 3600 604800 3600",
				Disabled: false},
		},
		TTL:    300,
		RRType: "SOA",
	}

	var wantRRSetNS = RRSet{
		Comments: nil,
		Name:     "example.com.",
		Records: []Record{
			Record{
				Content:  "ns2.example.com.",
				Disabled: false},
			Record{
				Content:  "ns1.example.com.",
				Disabled: false},
		},
		TTL:    300,
		RRType: "NS",
	}

	client, mux, _, teardown := setup()
	defer teardown()

	var testZoneRequest = ZoneRequest{
		Kind:    "Native",
		Masters: nil,
		Name:    "example.com.",
		Nameservers: []string{
			"ns1.example.com",
			"ns2.example.com",
		},
	}

	mux.HandleFunc("/api/v1/servers/localhost/zones", func(w http.ResponseWriter, r *http.Request) {
		v := new(ZoneRequest)
		json.NewDecoder(r.Body).Decode(v)

		testMethod(t, r, "POST")
		if !reflect.DeepEqual(v, &testZoneRequest) {
			t.Errorf("Request body = %+v, want %+v", v, testZoneRequest)
		}
		w.Write(testZonePostResp)
	})

	got, _, err := client.Zones.Post(context.Background(), testZoneRequest)
	if err != nil {
		t.Errorf("Zones.Post returned error: %v", err)
	}

	want := Zone{
		Account:          "",
		APIRectify:       false,
		DNSSec:           false,
		ID:               "example.com.",
		Kind:             "Native",
		LastCheck:        0,
		TSIGMasterKeyIDs: nil,
		Masters:          nil,
		Name:             "example.com.",
		NotifiedSerial:   0,
		NSEC3Narrow:      false,
		NSEC3Param:       "",
		RRSets:           []RRSet{wantRRSetSOA, wantRRSetNS},
		Serial:           2019012201,
		TSIGSlaveKeyIDs:  nil,
		SOAEdit:          "",
		SOAEditAPI:       "DEFAULT",
		URL:              "/api/v1/servers/localhost/zones/example.com.",
	}

	if !cmp.Equal(cmp.AllowUnexported(want), cmp.AllowUnexported(got)) {
		t.Errorf("Zones.Post returned %+v,\n                              want %+v", got, want)
	}

}
