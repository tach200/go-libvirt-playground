package main

import (
	"log"
	"net"
	"net/url"
	"time"

	libvirt "github.com/digitalocean/go-libvirt"
)

// type instance struct {
// 	ID        string
// 	Name      string
// 	Dir       string
// 	RootFS    string
// 	VNCPort   int
// 	Metadata  map[string]string
// 	RAM       int
// 	VCPU      int
// 	CPUShares int
// 	Sockets   int
// 	Cores     int
// 	Threads   int
// }

func main() {
	client, err := createLibvirtClient()
	if err != nil {
		log.Fatalf("error: couldn't create libvirt client : %s", err.Error())
	}
	defer client.Disconnect()

	// client.DomainCreateXML()

	// client.DomainCreate(libvirt.Domain{
	// 	Name: "",
	// 	UUID: libvirt.UUID{},
	// 	ID:   1,
	// })
}

// func createResource() {

// }

func createLibvirtClient() (*libvirt.Libvirt, error) {
	u, err := url.Parse("unix:///var/run/libvirt/libvirt-sock")
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTimeout(u.Scheme, u.Host+u.Path, 2*time.Second)
	if err != nil {
		return nil, err
	}

	l := libvirt.New(conn)
	if err := l.Connect(); err != nil {
		return nil, err
	}

	return l, nil
}
