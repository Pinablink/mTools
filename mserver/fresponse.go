package mserver

// KeyResponse
type KeyResponse string

// ImpResponse
type ImpResponse interface {
	AssembyResponse() interface{}
}

// FacResponse
type FacResponse struct {
	cResponse map[KeyResponse]ImpResponse
}

// NewFacResponse
func NewFacResponse() *FacResponse {
	var mmap map[KeyResponse]ImpResponse = make(map[KeyResponse]ImpResponse)
	return &FacResponse{cResponse: mmap}
}

// SetImpl
func (facResponse *FacResponse) SetImpl(key KeyResponse, impReponse ImpResponse) {
	facResponse.cResponse[key] = impReponse
}

// GetImpl
func (facResponse *FacResponse) GetImpl(key KeyResponse) ImpResponse {
	return facResponse.cResponse[key]
}
