package main

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt"
)

func main() {
	amt := amt.NewMessages()
	amt.AlarmClockService.Get()
	// fmt.Println(cim.BIOSElement.Get())
	// fmt.Println(cim.BIOSElement.Get())
	// fmt.Println(cim.BIOSElement.Get())

	// fmt.Println(cim.BootService.Get())
	// fmt.Println(cim.MediaAccessDevice.Enumerate())

	// amt := amt.NewMessages()
	// acs := models.AlarmClockService{
	// 	Service: cimModels.Service{
	// 		SystemName:          "AlarmClockService",
	// 		PrimaryOwnerName:    "Intel(R) AMT",
	// 		PrimaryOwnerContact: "http://www.intel.com",
	// 	},
	// }
	// cert := ips.Certificate{
	// 	H:                       "http://intel.com/wbem/wscim/1/ips-schema/1/IPS_IEEE8021xSettings",
	// 	ServerCertificateIssuer: "issuer",
	// 	ClientCertificate:       "clientCertificate",
	// }
	// xmlString, err := xml.Marshal(cert)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(xmlString))

}
