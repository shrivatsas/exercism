package vin

import (
	"testing"
)

const euSmallVIN = "W09000051T2123456"

type mockAPIClient struct {
	apiCalls int
}

func NewMockAPIClient() *mockAPIClient {

	return &mockAPIClient{}
}

func (client *mockAPIClient) IsEuropean(code string) bool {

	client.apiCalls++
	return true
}

func TestVIN_EU_SmallManufacturer(t *testing.T) {

	testVIN, _ := NewEUVIN(euSmallVIN)
	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}
}

// this fails with an error
func TestVIN_EU_SmallManufacturer_Polymorphism(t *testing.T) {

	var testVINs []VIN
	testVIN, _ := NewEUVIN(euSmallVIN)
	// having to cast testVIN already hints something is odd
	testVINs = append(testVINs, testVIN)

	for _, vin := range testVINs {
		manufacturer := vin.Manufacturer()
		if manufacturer != "W09123" {
			t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
		}
	}
}

func TestVIN_EU_SmallManufacturer_Service(t *testing.T) {

	apiClient := NewMockAPIClient()
	service := NewVINService(&VINServiceConfig{}, apiClient)
	testVIN, _ := service.CreateFromCode(euSmallVIN)

	manufacturer := testVIN.Manufacturer()
	if manufacturer != "W09123" {
		t.Errorf("unexpected manufacturer %s for VIN %s", manufacturer, testVIN)
	}

	if apiClient.apiCalls != 1 {
		t.Errorf("unexpected number of API calls: %d", apiClient.apiCalls)
	}
}
