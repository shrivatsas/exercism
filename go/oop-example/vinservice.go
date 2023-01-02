package vin

type VINService struct {
	client VINAPIClient
}

type VINServiceConfig struct {
	APIURL string
	APIKey string
	// more configuration values
}

func NewVINService(config *VINServiceConfig, apiClient VINAPIClient) *VINService {

	return &VINService{apiClient}
}

func (s *VINService) CreateFromCode(code string) (VIN, error) {

	if s.client.IsEuropean(code) {
		return NewEUVIN(code)
	}

	return NewVIN(code)
}
