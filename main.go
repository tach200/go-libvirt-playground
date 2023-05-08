package main

import (
	_ "embed" // needed for go:embed
	"log"
	"net"
	"net/url"
	"time"

	libvirt "github.com/digitalocean/go-libvirt"
)

//go:embed ubuntu-netboot.xml
var xmlString string

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
	client, err := newLibvirtClient()
	if err != nil {
		log.Fatalf("error: couldn't create libvirt client : %s", err.Error())
	}
	defer client.Disconnect()

	domain, err := client.DomainDefineXML(xmlString)
	if err != nil {
		log.Fatalf("error: couldnt define xml domain : %s", err.Error())
	}

	// err = client.DomainSetAutostart(domain, 1)
	// if err != nil {
	// 	log.Fatal("error: couldnt set auto start : %s", err.Error())
	// }

	err = client.DomainCreate(domain)
	if err != nil {
		log.Fatalf("error: couldnt create domain : %s", err.Error())
	}

	// client.DomainCreate(libvirt.Domain{
	// 	Name: "",
	// 	UUID: libvirt.UUID{},
	// 	ID:   1,
	// })
}

func newLibvirtClient() (*libvirt.Libvirt, error) {
	u, err := url.Parse("unix:///var/run/libvirt/libvirt-sock")
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTimeout(u.Scheme, u.Host+u.Path, 2*time.Second)
	if err != nil {
		return nil, err
	}

	client := libvirt.New(conn)
	if err := client.Connect(); err != nil {
		return nil, err
	}

	return client, nil
}
