package iqzone

import (
	"encoding/json"
	"github.com/prebid/prebid-server/openrtb_ext"
	"testing"
)

func TestValidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, validParam := range validParams {
		if err := validator.Validate(openrtb_ext.BidderIQZone, json.RawMessage(validParam)); err != nil {
			t.Errorf("Schema rejected IQZone params: %s", validParam)
		}
	}
}

func TestInvalidParams(t *testing.T) {
	validator, err := openrtb_ext.NewBidderParamsValidator("../../static/bidder-params")
	if err != nil {
		t.Fatalf("Failed to fetch the json-schemas. %v", err)
	}

	for _, invalidParam := range invalidParams {
		if err := validator.Validate(openrtb_ext.BidderIQZone, json.RawMessage(invalidParam)); err == nil {
			t.Errorf("Schema allowed unexpected params: %s", invalidParam)
		}
	}
}

var validParams = []string{
	`{"pubid": "19f1b372c7548ec1fe734d2c9f8dc688"}`,
}

var invalidParams = []string{
	`{"publisher": "19f1b372c7548ec1fe734d2c9f8dc688"}`,
	`nil`,
	``,
	`[]`,
	`true`,
	`{"pubid": 42}`,
}
