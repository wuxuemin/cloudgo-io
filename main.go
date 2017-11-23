package main

import (
	"cloudgo-io/service"
	"os"

	flag "github.com/spf13/pflag"
)

const (
	//PORT default:9090
	PORT string = "9090"
)

func main() {
	//get port
	portnum := os.Getenv("PORT")
	if len(portnum) == 0 {
		portnum = PORT
	}
	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		portnum = *pPort
	}
	m := service.NewServer(portnum)
	service.RunServer(m)
}
