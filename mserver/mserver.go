package mserver

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	STR_ERROR_MESSAGE_LEN_DATA_STRUCT  string = "Len Data Struct MServer.MService == 0"
	STR_ERROR_MESSAGE_NOT_PATTERN_FUNC string = "Func %s error not pattern "
	STR_ERROR_MESSAGE_PATTERN_FUNC     string = "Func %s error not / pattern"
)

//
type Service struct {
	pattern string
	ifunc   func(w http.ResponseWriter, r *http.Request)
}

//
type MServer struct {
	mPort    string
	mService []Service
}

//
func NewService(strPatern string, refFunc func(w http.ResponseWriter, r *http.Request)) Service {
	return Service{
		pattern: strPatern,
		ifunc:   refFunc,
	}
}

//
func NewMServer(port string, mServices ...Service) *MServer {
	return &MServer{
		mPort:    port,
		mService: mServices,
	}
}

//
func (mserver *MServer) InitMServer() {

	if len(mserver.mService) == 0 {
		panic(errors.New(STR_ERROR_MESSAGE_LEN_DATA_STRUCT))
	}

	mserver.verService()
	panic(http.ListenAndServe(mserver.mPort, nil))

}

//
func (mserver *MServer) verService() {

	for i, data := range mserver.mService {

		if len(data.pattern) == 0 {

			msgError := fmt.Sprintf(STR_ERROR_MESSAGE_NOT_PATTERN_FUNC, strconv.Itoa(i+1))
			panic(errors.New(msgError))

		} else if !strings.Contains(data.pattern, "/") {

			msgError := fmt.Sprintf(STR_ERROR_MESSAGE_PATTERN_FUNC, strconv.Itoa(i+1))
			panic(errors.New(msgError))

		} else {

			http.HandleFunc(data.pattern, data.ifunc)

		}

	}

}
