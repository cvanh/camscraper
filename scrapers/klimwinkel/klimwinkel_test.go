package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetproduct(t *testing.T){
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://www.klimwinkel.nl/api/products/69",
    httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	httpmock.GetTotalCallCount()

	f := getProduct(69);

	t.Logf("%v",f)
}