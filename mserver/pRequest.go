package mserver

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PRequestInterface interface {
	BadMethod(r *http.Request, verb MVerb) (bool, interface{})
	Response(keyResp KeyResponse) interface{}
}

type PRequest struct {
	Request     *http.Request
	Response    http.ResponseWriter
	RefObServer PRequestInterface
}

func (p *PRequest) ValidMethod(verb MVerb) bool {
	notOk, any := p.RefObServer.BadMethod(p.Request, verb)

	if notOk {
		fmt.Println("Nao é ok2")
		p.Response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(p.Response).Encode(any)
		return false
	}

	return true
}

func (p *PRequest) Retreponse(keyResp KeyResponse) {
	any := p.RefObServer.Response(keyResp)
	p.Response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(p.Response).Encode(any)
}
