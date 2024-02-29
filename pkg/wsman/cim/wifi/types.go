/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package wifi

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/cim/models"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/common"
)

type Port struct {
	base   message.Base
	client client.WSMan
}

type EndpointSettings struct {
	base message.Base
}

// OUTPUT
// Response Types
type (
	Response struct {
		*client.Message
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}

	Body struct {
		XMLName                   xml.Name `xml:"Body"`
		WiFiPortGetResponse       WiFiPort
		EnumerateResponse         common.EnumerateResponse
		PullResponse              PullResponse
		RequestStateChange_OUTPUT common.ReturnValue
	}

	PullResponse struct {
		XMLName               xml.Name                       `xml:"PullResponse"`
		EndpointSettingsItems []WiFiEndpointSettingsResponse `xml:"Items>CIM_WiFiEndpointSettings"`
		WiFiPortItems         []WiFiPort                     `xml:"Items>CIM_WiFiPort"`
	}

	WiFiEndpointSettingsResponse struct {
		XMLName              xml.Name             `xml:"CIM_WiFiEndpointSettings"`
		AuthenticationMethod AuthenticationMethod // AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied.
		BSSType              BSSType              // BSSType shall indicate the Basic Service Set (BSS) Type that shall be used when the settings are applied.
		ElementName          string               // The user-friendly name for this instance of SettingData. In addition, the user-friendly name can be used as an index property for a search or query. (Note: The name does not have to be unique within a namespace.)
		EncryptionMethod     EncryptionMethod     // EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied.
		InstanceID           string               // Within the scope of the instantiating Namespace, InstanceID opaquely and uniquely identifies an instance of this class.
		Priority             int                  // Priority shall indicate the priority of the instance among all WiFiEndpointSettings instances.
		SSID                 string               // SSID shall indicate the Service Set Identifier (SSID) that shall be used when the settings are applied to a WiFiEndpoint. An SSID identifies a wireless network.
	}

	WiFiPort struct {
		XMLName                 xml.Name              `xml:"CIM_WiFiPort"`
		LinkTechnology          LinkTechnology        // An enumeration of the types of links. When set to 1 ("Other"), the related property OtherLinkTechnology contains a string description of the type of link.
		DeviceID                string                // An address or other identifying information to uniquely name the LogicalDevice.
		CreationClassName       string                // CreationClassName indicates the name of the class or the subclass used in the creation of an instance.
		SystemName              string                // The scoping System's Name.
		SystemCreationClassName string                // The scoping System's CreationClassName.
		ElementName             string                // A user-friendly name for the object. This property allows each instance to define a user-friendly name in addition to its key properties, identity data, and description information.
		HealthState             models.HealthState    // Indicates the current health of the element.
		EnabledState            models.EnabledState   // EnabledState is an integer enumeration that indicates the enabled and disabled states of an element.
		RequestedState          models.RequestedState // RequestedState is an integer enumeration that indicates the last requested or desired state for the element, irrespective of the mechanism through which it was requested.
		PortType                int                   // PortType shall contain the specific 802.11 operating mode that is currently enabled on the Port.
		PermanentAddress        string                // IEEE 802 EUI-48 MAC address, formatted as twelve hexadecimal digits (for example, "010203040506"), with each pair representing one of the six octets of the MAC address in "canonical" bit order.
	}
)

// INPUT
// Request Types
type (
	WiFiEndpointSettings_INPUT struct {
		XMLName              xml.Name `xml:"CIM_WiFiEndpointSettings"`
		H                    string   `xml:"xmlns:q,attr"`
		AuthenticationMethod AuthenticationMethod
		BSSType              BSSType
		ElementName          string
		EncryptionMethod     EncryptionMethod
		InstanceID           string
		Priority             int
		SSID                 string
	}
	WiFiEndpointSettingsRequest struct {
		XMLName xml.Name `xml:"h:WiFiEndpointSettingsInput"`
		H       string   `xml:"xmlns:q,attr"`
		// SettingData
		ElementName          string               `xml:"q:ElementName,omitempty"`
		InstanceID           string               `xml:"q:InstanceID,omitempty"`
		AuthenticationMethod AuthenticationMethod `xml:"q:AuthenticationMethod,omitempty"`
		EncryptionMethod     EncryptionMethod     `xml:"q:EncryptionMethod,omitempty"`
		SSID                 string               `xml:"q:SSID,omitempty"` // Max Length 32
		Priority             int                  `xml:"q:Priority,omitempty"`
		PSKPassPhrase        string               `xml:"q:PSKPassPhrase,omitempty"` // Min Length 8 Max Length 63
		BSSType              BSSType              `xml:"q:BSSType,omitempty"`
		Keys                 []string             `xml:"q:Keys,omitempty"` // OctetString ArrayType=Indexed Max Length 256
		KeyIndex             int                  `xml:"q:KeyIndex,omitempty"`
		PSKValue             int                  `xml:"q:PSKValue,omitempty"` // OctetString
	}
)

type (

	// An enumeration of the types of links. When set to 1 ("Other"), the related property OtherLinkTechnology contains a string description of the type of link.
	//
	// ValueMap={0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//
	// Values={Unknown, Other, Ethernet, IB, FC, FDDI, ATM, Token Ring, Frame Relay, Infrared, BlueTooth, Wireless LAN}
	LinkTechnology int

	//AuthenticationMethod shall specify the 802.11 authentication method used when the settings are applied. * Other (1): shall indicate that the desired authentication method is not specified in the list below. If AuthenticationMethod contains 1, OtherAuthenticationMethod should not be NULL and should not be empty.
	//
	// * Open System (2): shall indicate that the desired authentication method is Open System. AuthenticationMethod should contain 2 only if EncryptionMethod contains 2 ("WEP") or 5 ("None").
	//
	// * Shared Key (3): shall indicate that the desired authentication method is Shared Key. AuthenticationMethod should contain 3 only if EncryptionMethod contains 2 ("WEP") or 5 ("None").
	//
	// * WPA PSK (4): shall indicate that the desired authentication method is WPA (Wi-Fi Protected Access) PSK (Pre-Shared Key). AuthenticationMethod should contain 4 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	//
	// * WPA IEEE 802.1x (5): shall indicate that the desired authentication method is WPA (Wi-Fi Protected Access) IEEE 802.1x. AuthenticationMethod should contain 5 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	//
	// * WPA2 PSK (6): shall indicate that the desired authentication method is WPA2 (Wi-Fi Protected Access Version 2) PSK (Pre-Shared Key). AuthenticationMethod should containt 6 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	//
	// * WPA2 IEEE 802.1x (7): shall indicate that the desired authentication method is WPA2 (Wi-Fi Protected Access Version 2) IEEE 802.1x. AuthenticationMethod should contain 7 only if EncryptionMethod contains 3 ("TKIP") or 4 ("CCMP").
	//
	// * WPA3 IEEE 802.1x (32768)(Supported in Intel CSME 12.0.80.1708, Intel CSME 14.0 and later): shall indicate that WPA3 (Wi-Fi Protected Access Version 3) SAE IEEE 802.1x authentication is supported. SupportedAuthenticationMethods shall contain 32768 only if SupportedEncryptionMethods contains 4 (CCMP).
	//
	// * WPA3 OWE IEEE 802.1x (32769) (Supported in Intel CSME 12.0.80.1708, Intel CSME 14.0 and later): shall indicate that WPA3 (Wi-Fi Protected Access Version 3) OWE (Opportunistic Wireless Encryption) IEEE 802.1x authentication is supported. SupportedAuthenticationMethods shall contain 32769 only if SupportedEncryptionMethods contains 4 (CCMP).
	//
	// ValueMap={1, 2, 3, 4, 5, 6, 7, 8..32767, 32768, 32769, 32770..}
	//
	// Values={Other, Open System, Shared Key, WPA PSK, WPA IEEE 802.1x, WPA2 PSK, WPA2 IEEE 802.1x, DMTF Reserved, WPA3 SAE, WPA3 OWE, Vendor Reserved}
	AuthenticationMethod int

	// BSSType shall indicate the Basic Service Set (BSS) Type that shall be used when the settings are applied. A Basic Service Set is a set of stations controlled by a single coordination function.
	//
	// * Independent: the WiFiEndpoint is associated directly to another client station.
	//
	// * Infrastructure: the WiFiEndpoint is associated to a network via an access point.
	BSSType int

	// EncryptionMethod shall specify the 802.11 encryption method used when the settings are applied. * Other (1): shall indicate that the desired encryption method is not specified in the list below. If this value is used, OtherEncryptionMethod should not be NULL and should not be empty.
	//
	// * WEP (2): shall indicate that the desired encryption method is Wired Equivalency Privacy (WEP). This value should be used only if AuthenticationMethod contains 2 ("Open System") or 3 ("Shared Key").
	//
	// * TKIP (3): shall indicate that the desired encryption method is Temporal Key Integrity Protocol (TKIP). This value should be used only if AuthenticationMethod contains 4 ("WPA PSK"), 5 ("WPA IEEE 802.1x"), 6 ("WPA2 PSK"), or 7 ("WPA2 IEEE 802.1x").
	//
	// * CCMP (4): shall indicate that the desired encryption method is Counter Mode with Cipher Block Chaining Message Authentication Code Protocol (CCMP). This value should be used only if AuthenticationMethod contains 4 ("WPA PSK"), 5 ("WPA IEEE 802.1x"), 6 ("WPA2 PSK"), or 7 ("WPA2 IEEE 802.1x").
	//
	// * None (5): shall indicate that no encryption is desired. This value should be used only if AuthenticationMethod contains 2 ("Open System") or 3 ("Shared Key").
	//
	// ValueMap={1, 2, 3, 4, 5, 6..}
	//
	// Values={Other, WEP, TKIP, CCMP, None, DMTF Reserved}
	EncryptionMethod int

	ReturnValue int
)
