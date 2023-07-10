package main

import "log"

// This is a place holder of a function that will return a full configurtion object after decorating it
type ConfigFunction func(*Configurations) 

func withMaxConn(n int) ConfigFunction{
	return func(config *Configurations){
		config.this_MaxConn = n
	}
}

func withTLS(config *Configurations){
	config.this_TLS = true
}

func withId(id string) ConfigFunction{
	return func(config *Configurations){
		config.this_Id = id
	}
}

// This is the configurations that will be available to the clients to create a server
type Configurations struct{
	this_MaxConn int
	this_Id string
	this_TLS bool
}

type ConfigurationsByPointer struct{
	this_MaxConn *int
	this_Id *string
	this_TLS *bool
}

// Just a default set of configurations is the clients does not provide any config
func defaultConfig() Configurations {
	return Configurations{
		this_MaxConn: 10,
		this_Id: "default",
		this_TLS: false,
	}
}

// This is a server with some configurations
type Server struct {
	this_Configurations Configurations
}

// Construct new server with the given configurations, Configurations can be of any number
func newServer(configs ...ConfigFunction) *Server {
	defaultConfig := defaultConfig()

	for _,configFunc := range configs {
		configFunc(&defaultConfig)
	}
	return &Server{
		this_Configurations: defaultConfig,
	}
}


// Value receiver method for server
func (server Server) printServerInfo1() {
	log.Println(server.this_Configurations.this_Id)
}

func (server *Server) printServerInfo2() {
	log.Println(server.this_Configurations.this_Id)
}

func (server *Server) printServerInfo3() {
	a := &server.this_Configurations.this_MaxConn
	b := &server.this_Configurations.this_MaxConn
	c := server.this_Configurations.this_MaxConn
	log.Println(&server.this_Configurations.this_MaxConn)
	log.Println(*a)
	*b += 1
	log.Println(*b)
	log.Println(*a)
	c += 1
	log.Println(&c)
	log.Println(c)
}

func main() {
	myServer := newServer(withTLS, withMaxConn(3),withId("NPADHYA"))
	myServer2 := newServer(withTLS)
	myServer3 := newServer(withId("SOME ID"))
	myServer4 := newServer(withMaxConn(-1))

	myServer.printServerInfo1()
	myServer2.printServerInfo1()
	myServer3.printServerInfo1()
	myServer4.printServerInfo1()

	myServer.printServerInfo2()
	myServer2.printServerInfo2()
	myServer3.printServerInfo2()
	myServer4.printServerInfo2()

	myServer.printServerInfo3()
	myServer2.printServerInfo3()
	myServer3.printServerInfo3()
	myServer4.printServerInfo3()

	cbp := ConfigurationsByPointer{}
	a := 1
	b:= "stringValue"
	c := false

	cbp2 := ConfigurationsByPointer{&a,&b,&c }

	log.Println(cbp)
	log.Println(*cbp2.this_Id)
}