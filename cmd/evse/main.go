package main

import (
	"crypto/ecdsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/DerAndereAndi/eebus-go/service"
	"github.com/DerAndereAndi/eebus-go/spine/model"
)

var myService *service.EEBUSService

func usage() {
	fmt.Println("Usage: go run /cmd/evse/main.go <serverport> <hems-ski> <hems-shipid> <crtfile> <keyfile>")
}

func main() {
	if len(os.Args) < 4 {
		usage()
		return
	}

	myService = &service.EEBUSService{
		DeviceBrand:        "Demo",
		DeviceModel:        "HEMS",
		DeviceSerialNumber: "234567890",
		DeviceIdentifier:   "Demo-EVSE-234567890",
		DeviceType:         model.DeviceTypeEnumTypeChargingStation,
		RemoteDeviceTypes: []model.DeviceTypeEnumType{
			model.DeviceTypeEnumTypeEnergyManagementSystem,
		},
	}

	var err error
	var certificate tls.Certificate

	myService.Port, err = strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
		log.Fatal(err)
	}

	remoteSki := os.Args[2]
	remoteShipID := os.Args[3]

	fmt.Println(os.Args)
	if len(os.Args) == 6 {
		certificate, err = tls.LoadX509KeyPair(os.Args[4], os.Args[5])
		if err != nil {
			usage()
			log.Fatal(err)
		}
	} else {
		certificate, err = myService.CreateCertificate("Demo", "Demo", "DE", "Demo-Unit-02")
		if err != nil {
			log.Fatal(err)
		}

		pemdata := pem.EncodeToMemory(&pem.Block{
			Type:  "CERTIFICATE",
			Bytes: certificate.Certificate[0],
		})
		fmt.Println(string(pemdata))

		b, err := x509.MarshalECPrivateKey(certificate.PrivateKey.(*ecdsa.PrivateKey))
		if err != nil {
			log.Fatal(err)
		}
		pemdata = pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: b})
		fmt.Println(string(pemdata))
	}

	myService.Certificate = certificate
	myService.Start()

	remoteService := service.ServiceDetails{
		SKI:    remoteSki,
		ShipID: remoteShipID,
	}
	myService.RegisterRemoteService(remoteService)

	if err = myService.MdnsAnnounce(); err != nil {
		log.Fatal(err)
	}
	defer myService.Shutdown()

	// Clean exit to make sure mdns shutdown is invoked
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sig:
		// User exit
	}
}