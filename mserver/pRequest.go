package mserver

import (
	"encoding/json"
	"net/http"
)

type PRequestInterface interface {
	BadMethod(r *http.Request) (bool, interface{})
	Response() interface{}
}

type PRequest struct {
	Request     *http.Request
	Response    http.ResponseWriter
	RefObServer PRequestInterface
}

func (p *PRequest) ValidMethod() (bool, interface{}) {
	notOk, any := p.RefObServer.BadMethod(p.Request)

	if notOk {
		p.Response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(p.Response).Encode(any)
		return false, nil
	}

	return true, any
}

func (p *PRequest) Retreponse() {
	any := p.RefObServer.Response()
	p.Response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(p.Response).Encode(any)
}
