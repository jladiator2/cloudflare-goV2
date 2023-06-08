package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountCustomNameserver_Get(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": [
				{
					"ns_name": "ns1.example.com",
					"ns_set": 1,
					"dns_records": [
						{
							"type": "A",
							"value": "192.0.2.1"
						},
						{
							"type": "AAAA",
							"value": "2400:cb00:2049:1::ffff:ffee"
						}
					]
				},
				{
					"ns_name": "ns2.example.com",
					"ns_set": 1,
					"dns_records": [
						{
							"type": "A",
							"value": "192.0.2.2"
						},
						{
							"type": "AAAA",
							"value": "2400:cb00:2049:1::ffff:fffe"
						}
					]
				}
			]
		}`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/custom_ns", handler)
	want := []CustomNameserverResult{
		{
			DNSRecords: []CustomNameserverRecord{
				{
					Type:  "A",
					Value: "192.0.2.1",
				},
				{
					Type:  "AAAA",
					Value: "2400:cb00:2049:1::ffff:ffee",
				},
			},
			NSName: "ns1.example.com",
			NSSet:  1,
		},
		{
			DNSRecords: []CustomNameserverRecord{
				{
					Type:  "A",
					Value: "192.0.2.2",
				},
				{
					Type:  "AAAA",
					Value: "2400:cb00:2049:1::ffff:fffe",
				},
			},
			NSName: "ns2.example.com",
			NSSet:  1,
		},
	}

	actual, err := client.GetAccountCustomNameservers(context.Background(), AccountIdentifier("01a7362d577a6c3019a474fd6f485823"))

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccountCustomNameserver_Create(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"success": true,
			"errors": [],
			"messages": [],
			"result": {
				"ns_name": "ns1.example.com",
				"ns_set": 1,
				"dns_records": [
					{
						"type": "A",
						"value": "192.0.2.1"
					},
					{
						"type": "AAAA",
						"value": "2400:cb00:2049:1::ffff:ffee"
					}
				]
			}
		}`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/custom_ns", handler)
	want := CustomNameserverResult{
		DNSRecords: []CustomNameserverRecord{
			{
				Type:  "A",
				Value: "192.0.2.1",
			},
			{
				Type:  "AAAA",
				Value: "2400:cb00:2049:1::ffff:ffee",
			},
		},
		NSName: "ns1.example.com",
		NSSet:  1,
	}

	actual, err := client.CreateAccountCustomNameserver(
		context.Background(),
		AccountIdentifier("01a7362d577a6c3019a474fd6f485823"),
		CustomNameserver{
			NSName: "ns1.example.com",
			NSSet:  1,
		},
	)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccountCustomNameserver_GetEligibleZones(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
    "result": [
        "example.com",
        "example2.com",
        "example3.com"
    ],
    "success": true,
    "errors": [],
    "messages": []
}`)
	}

	mux.HandleFunc("/accounts/01a7362d577a6c3019a474fd6f485823/custom_ns/availability", handler)
	want := []string{
		"example.com",
		"example2.com",
		"example3.com",
	}

	actual, err := client.GetEligibleZonesAccountCustomNameservers(
		context.Background(),
		AccountIdentifier("01a7362d577a6c3019a474fd6f485823"),
	)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}

func TestAccountCustomNameserver_GetAccountCustomNameserverZoneMetadata(t *testing.T) {
	setup()
	defer teardown()

	handler := func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method, "Expected method 'GET', got %s", r.Method)
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, `{
			"result": {
				"type": "account",
				"ns_set": "1",
				"enabled": true
			},
			"success": true,
			"errors": [],
			"messages": []
		}`)
	}

	mux.HandleFunc("/zones/023e105f4ecef8ad9ca31a8372d0c353/custom_ns", handler)
	want := CustomNameserverZoneMetadata{
		Type:    "account",
		NSSet:   "1",
		Enabled: true,
	}

	actual, err := client.GetAccountCustomNameserverZoneMetadata(
		context.Background(),
		ZoneIdentifier("023e105f4ecef8ad9ca31a8372d0c353"),
	)

	if assert.NoError(t, err) {
		assert.Equal(t, want, actual)
	}
}
