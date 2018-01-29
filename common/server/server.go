package server

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type Server struct {
	name    string
	val     reflect.Value
	typ     reflect.Type
	methods []reflect.Method
	lock    sync.Mutex
}

func NewServer() *Server {
	server := new(Server)
	server.methods = make([]reflect.Method, 0)
	return server
}

func (s *Server) Start(addr string) error {
	log.Println(fmt.Sprintf("server start on addr [%s]", addr))
	return http.ListenAndServe(addr, s)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, method := range s.methods {
		if strings.ToLower("/"+s.name+"/"+method.Name) == r.URL.Path {
			method.Func.Call([]reflect.Value{s.val, reflect.ValueOf(w), reflect.ValueOf(r)})
		}
	}
}

func (s *Server) Register(service interface{}) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.typ = reflect.TypeOf(service)
	s.val = reflect.ValueOf(service)
	s.name = reflect.Indirect(s.val).Type().Name()
	if s.name == "" {
		return fmt.Errorf("no service name for type %s", s.typ.String())
	}
	for m := 0; m < s.typ.NumMethod(); m++ {
		method := s.typ.Method(m)
		mtype := method.Type
		if mtype.NumIn() != 3 {
			return fmt.Errorf("method %s has wrong number of ins: %d", method.Name, mtype.NumIn())
		}
		reply := mtype.In(1)
		if reply.String() != "http.ResponseWriter" {
			return fmt.Errorf("%s argument type not exported: %s", method.Name, reply)
		}
		arg := mtype.In(2)
		if arg.String() != "*http.Request" {
			return fmt.Errorf("%s argument type not exported: %s", method.Name, arg)
		}
		s.methods = append(s.methods, method)
		log.Println("registry path:", strings.ToLower("/"+s.name+"/"+method.Name))
	}
	return nil
}
