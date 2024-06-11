/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/common"
)

const (
	AMTAuditLog   string = "AMT_AuditLog"
	ReadRecords   string = "ReadRecords"
	ValueNotFound string = "Value not found in map"
)

const (
	OverwritePolicyUnknown                   OverwritePolicy = 0
	OverwritePolicyWrapsWhenFull             OverwritePolicy = 2
	OverwritePolicyNeverOverwrites           OverwritePolicy = 7
	OverwritePolicyPartialRestrictedRollover OverwritePolicy = 32768
)

var OverwritePolicyToString = map[OverwritePolicy]string{
	OverwritePolicyUnknown:                   "Unknown",
	OverwritePolicyWrapsWhenFull:             "WrapsWhenFull",
	OverwritePolicyNeverOverwrites:           "NeverOverwrites",
	OverwritePolicyPartialRestrictedRollover: "PartialRestrictedRollover",
}

// OverwritePolicyToString returns a string representation of a OverwritePolicy.
func (r OverwritePolicy) String() string {
	if value, exists := OverwritePolicyToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	EnabledStateUnknown EnabledState = iota
	EnabledStateOther
	EnabledStateEnabled
	EnabledStateDisabled
	EnabledStateShuttingDown
	EnabledStateNotApplicable
	EnabledStateEnabledButOffline
	EnabledStateInTest
	EnabledStateDeferred
	EnabledStateQuiesce
	EnabledStateStarting
)

var EnabledStateToString = map[EnabledState]string{
	EnabledStateUnknown:           "Unknown",
	EnabledStateOther:             "Other",
	EnabledStateEnabled:           "Enabled",
	EnabledStateDisabled:          "Disabled",
	EnabledStateShuttingDown:      "ShuttingDown",
	EnabledStateNotApplicable:     "NotApplicable",
	EnabledStateEnabledButOffline: "EnabledButOffline",
	EnabledStateInTest:            "InTest",
	EnabledStateDeferred:          "Deferred",
	EnabledStateQuiesce:           "Quiesce",
	EnabledStateStarting:          "Starting",
}

// EnabledStateToString returns a string representation of a EnabledState.
func (r EnabledState) String() string {
	if value, exists := EnabledStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	RequestedStateUnknown       RequestedState = 0
	RequestedStateEnabled       RequestedState = 2
	RequestedStateDisabled      RequestedState = 3
	RequestedStateShutDown      RequestedState = 4
	RequestedStateNoChange      RequestedState = 5
	RequestedStateOffline       RequestedState = 6
	RequestedStateTest          RequestedState = 7
	RequestedStateDeferred      RequestedState = 8
	RequestedStateQuiesce       RequestedState = 9
	RequestedStateReboot        RequestedState = 10
	RequestedStateReset         RequestedState = 11
	RequestedStateNotApplicable RequestedState = 12
)

var RequestedStateToString = map[RequestedState]string{
	RequestedStateUnknown:       "Unknown",
	RequestedStateEnabled:       "Enabled",
	RequestedStateDisabled:      "Disabled",
	RequestedStateShutDown:      "ShutDown",
	RequestedStateNoChange:      "NoChange",
	RequestedStateOffline:       "Offline",
	RequestedStateTest:          "Test",
	RequestedStateDeferred:      "Deferred",
	RequestedStateQuiesce:       "Quiesce",
	RequestedStateReboot:        "Reboot",
	RequestedStateReset:         "Reset",
	RequestedStateNotApplicable: "NotApplicable",
}

// RequestedStateToString returns a string representation of a RequestedState.
func (r RequestedState) String() string {
	if value, exists := RequestedStateToString[r]; exists {
		return value
	}

	return ValueNotFound
}

const (
	StoragePolicyNoRollOver StoragePolicy = iota
	StoragePolicyRollOver
	StoragePolicyRestrictedRollOver
)

var StoragePolicyToString = map[StoragePolicy]string{
	StoragePolicyNoRollOver:         "NoRollOver",
	StoragePolicyRollOver:           "RollOver",
	StoragePolicyRestrictedRollOver: "RestrictedRollOver",
}

// StoragePolicyToString returns a string representation of a StoragePolicy.
func (r StoragePolicy) String() string {
	if value, exists := StoragePolicyToString[r]; exists {
		return value
	}

	return ValueNotFound
}

func convertToAuditLogResult(auditlogdata []string) []AuditLogRecord {
	records := []AuditLogRecord{}

	for _, eventRecord := range auditlogdata {
		ptr := 0

		decodedEventRecord, err := base64.StdEncoding.DecodeString(eventRecord)
		if err != nil {
			continue
		}

		decodedEventRecordStr := string(decodedEventRecord)
		auditLogRecord := AuditLogRecord{}

		auditLogRecord.AuditAppID = common.ReadShort(decodedEventRecordStr, 0)
		auditLogRecord.EventID = common.ReadShort(decodedEventRecordStr, 2)
		auditLogRecord.AuditApp = AMTAuditStringTable[auditLogRecord.AuditAppID]
		// auditLogRecord.InitiatorType = decodedEventRecordStr[:4]
		auditLogRecord.Event = AMTAuditStringTable[(auditLogRecord.AuditAppID*100)+auditLogRecord.EventID]

		// if auditLogRecord.Event {
		// 	auditLogRecord.Event = '#' + auditLogRecord.EventID
		// }

		initiatorType, initiator, pointer := getInitiatorInfo(decodedEventRecordStr)
		auditLogRecord.InitiatorType = initiatorType
		auditLogRecord.Initiator = initiator
		ptr = pointer

		// Read timestamp
		timeStamp := common.ReadInt(decodedEventRecordStr, ptr)
		auditLogRecord.Time = time.Unix(int64(timeStamp), 0)
		ptr += 4

		// Read network access

		auditLogRecord.MCLocationType = []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++
		netlen := []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++
		auditLogRecord.NetAddress = strings.ReplaceAll(decodedEventRecordStr[ptr:ptr+int(netlen)], "0000:0000:0000:0000:0000:0000:0000:0001", "::1")

		// Read extended data
		ptr += int(netlen)
		exlen := []byte(decodedEventRecordStr[ptr : ptr+1])[0]
		ptr++
		auditLogRecord.Ex = decodedEventRecordStr[ptr : ptr+int(exlen)]
		auditLogRecord.ExStr = GetAuditLogExtendedDataString((auditLogRecord.AuditAppID*100)+auditLogRecord.EventID, auditLogRecord.Ex)

		records = append([]AuditLogRecord{auditLogRecord}, records...)
	}

	return records
}

const (
	ACLEntryAdded                   = 1602
	ACLEntryModified                = 1603
	ACLEntryRemoved                 = 1604
	ACLAccessWithInvalidCredentials = 1605
	ACLEntryStateChanged            = 1606
	TLSStateChanged                 = 1607
	SetRealmAuthenticationMode      = 1617
	AMTUnprovisioningStarted        = 1619
	FirmwareUpdate                  = 1900
	AMTTimeSet                      = 2100
	OptInPolicyChange               = 3000
	SendConsentCode                 = 3001
)

var RealmNames = []string{
	"Redirection",
	"PT Administration",
	"Hardware Asset",
	"Remote Control",
	"Storage",
	"Event Manager",
	"Storage Admin",
	"Agent Presence Local",
	"Agent Presence Remote",
	"Circuit Breaker",
	"Network Time",
	"General Information",
	"Firmware Update",
	"EIT",
	"LocalUN",
	"Endpoint Access Control",
	"Endpoint Access Control Admin",
	"Event Log Reader",
	"Audit Log",
	"ACL Realm",
	"",
	"",
	"Local System",
	// Add more as needed
}

// Return human readable extended audit log data
// TODO: Just put some of them here, but many more still need to be added, helpful link here:
// https://software.intel.com/sites/manageability/AMT_Implementation_and_Reference_Guide/default.htm?turl=WordDocuments%2Fsecurityadminevents.htm
func GetAuditLogExtendedDataString(auditEventId int, data string) string {
	var extendedDataString string

	switch auditEventId {
	case ACLEntryAdded, ACLEntryRemoved:
		if data[0] == 0 {
			extendedDataString = data[2 : 2+data[1]]
		}
	case ACLEntryModified:
		if data[1] == 0 {
			extendedDataString = data[2:]
		}
	case ACLAccessWithInvalidCredentials:
		extendedDataString = []string{"Invalid ME access", "Invalid MEBx access"}[data[0]]
	case ACLEntryStateChanged:
		r := []string{"Disabled", "Enabled"}[data[0]]
		if data[1] == 0 {
			r += ", " + data[2:]
		}

		extendedDataString = r
	case TLSStateChanged:
		extendedDataString = "Remote " + []string{"NoAuth", "ServerAuth", "MutualAuth"}[data[0]] + ", Local " + []string{"NoAuth", "ServerAuth", "MutualAuth"}[data[1]]
	case SetRealmAuthenticationMode:
		extendedDataString = RealmNames[common.ReadInt(data, 0)] + ", " + []string{"NoAuth", "Auth", "Disabled"}[data[4]]
	case AMTUnprovisioningStarted:
		extendedDataString = []string{"BIOS", "MEBx", "Local MEI", "Local WSMAN", "Remote WSMAN"}[data[0]]
	case FirmwareUpdate:
		extendedDataString = "From " + strconv.Itoa(common.ReadShort(data, 0)) + "." + strconv.Itoa(common.ReadShort(data, 2)) + "." + strconv.Itoa(common.ReadShort(data, 4)) + "." + strconv.Itoa(common.ReadShort(data, 6)) + " to " + strconv.Itoa(common.ReadShort(data, 8)) + "." + strconv.Itoa(common.ReadShort(data, 10)) + "." + strconv.Itoa(common.ReadShort(data, 12)) + "." + strconv.Itoa(common.ReadShort(data, 14))
	case AMTTimeSet:
		t := time.Unix(int64(common.ReadInt(data, 0)), 0).Local()
		extendedDataString = t.Format(time.RFC1123)
	case OptInPolicyChange:
		extendedDataString = "From " + []string{"None", "KVM", "All"}[data[0]] + " to " + []string{"None", "KVM", "All"}[data[1]]
	case SendConsentCode:
		extendedDataString = []string{"Success", "Failed 3 times"}[data[0]]
	default:
		extendedDataString = ""
	}

	return extendedDataString
}

const (
	HTTPDigest     byte = 0
	Kerberos       byte = 1
	Local          byte = 2
	KvmDefaultPort byte = 3
)

// [initiatorType: number, initiator: string, ptr: number] need to type it.
func getInitiatorInfo(decodedEventRecord string) (initatorType byte, initiator string, ptr int) {
	var userlen uint8

	initiatorType := []byte(decodedEventRecord[4:5])[0]

	switch initiatorType {
	case HTTPDigest:
		userlen = []byte(decodedEventRecord[5:6])[0]
		initiator = decodedEventRecord[6 : 6+userlen]
		ptr = 6 + int(userlen)
	case Kerberos:
		userlen = []byte(decodedEventRecord[9:10])[0]
		initiator = common.GetSidString(decodedEventRecord[10 : 10+userlen])
		ptr = 10 + int(userlen)
	case Local:
		initiator = "Local"
		ptr = 5
	case KvmDefaultPort:
		initiator = "KVM Default Port"
		ptr = 5
	default:
		initiator = ""
		ptr = 0
	}

	return initiatorType, initiator, ptr
}

var ExtendedDataMap = map[int]string{
	0: "Invalid ME access",
	1: "Invalid MEBx access",
}

var AMTAuditStringTable = map[int]string{
	16:   "Security Admin",
	17:   "RCO",
	18:   "Redirection Manager",
	19:   "Firmware Update Manager",
	20:   "Security Audit Log",
	21:   "Network Time",
	22:   "Network Administration",
	23:   "Storage Administration",
	24:   "Event Manager",
	25:   "Circuit Breaker Manager",
	26:   "Agent Presence Manager",
	27:   "Wireless Configuration",
	28:   "EAC",
	29:   "KVM",
	30:   "User Opt-In Events",
	32:   "Screen Blanking",
	33:   "Watchdog Events",
	1600: "Provisioning Started",
	1601: "Provisioning Completed",
	1602: "ACL Entry Added",
	1603: "ACL Entry Modified",
	1604: "ACL Entry Removed",
	1605: "ACL Access with Invalid Credentials",
	1606: "ACL Entry State",
	1607: "TLS State Changed",
	1608: "TLS Server Certificate Set",
	1609: "TLS Server Certificate Remove",
	1610: "TLS Trusted Root Certificate Added",
	1611: "TLS Trusted Root Certificate Removed",
	1612: "TLS Preshared Key Set",
	1613: "Kerberos Settings Modified",
	1614: "Kerberos Master Key Modified",
	1615: "Flash Wear out Counters Reset",
	1616: "Power Package Modified",
	1617: "Set Realm Authentication Mode",
	1618: "Upgrade Client to Admin Control Mode",
	1619: "Unprovisioning Started",
	1700: "Performed Power Up",
	1701: "Performed Power Down",
	1702: "Performed Power Cycle",
	1703: "Performed Reset",
	1704: "Set Boot Options",
	1705: "Remote graceful power down initiated",
	1706: "Remote graceful reset initiated",
	1707: "Remote Standby initiated",
	1708: "Remote Hiberate initiated",
	1709: "Remote NMI initiated",
	1800: "IDER Session Opened",
	1801: "IDER Session Closed",
	1802: "IDER Enabled",
	1803: "IDER Disabled",
	1804: "SoL Session Opened",
	1805: "SoL Session Closed",
	1806: "SoL Enabled",
	1807: "SoL Disabled",
	1808: "KVM Session Started",
	1809: "KVM Session Ended",
	1810: "KVM Enabled",
	1811: "KVM Disabled",
	1812: "VNC Password Failed 3 Times",
	1900: "Firmware Updated",
	1901: "Firmware Update Failed",
	2000: "Security Audit Log Cleared",
	2001: "Security Audit Policy Modified",
	2002: "Security Audit Log Disabled",
	2003: "Security Audit Log Enabled",
	2004: "Security Audit Log Exported",
	2005: "Security Audit Log Recovered",
	2100: "Intel(R) ME Time Set",
	2200: "TCPIP Parameters Set",
	2201: "Host Name Set",
	2202: "Domain Name Set",
	2203: "VLAN Parameters Set",
	2204: "Link Policy Set",
	2205: "IPv6 Parameters Set",
	2300: "Global Storage Attributes Set",
	2301: "Storage EACL Modified",
	2302: "Storage FPACL Modified",
	2303: "Storage Write Operation",
	2400: "Alert Subscribed",
	2401: "Alert Unsubscribed",
	2402: "Event Log Cleared",
	2403: "Event Log Frozen",
	2500: "CB Filter Added",
	2501: "CB Filter Removed",
	2502: "CB Policy Added",
	2503: "CB Policy Removed",
	2504: "CB Default Policy Set",
	2505: "CB Heuristics Option Set",
	2506: "CB Heuristics State Cleared",
	2600: "Agent Watchdog Added",
	2601: "Agent Watchdog Removed",
	2602: "Agent Watchdog Action Set",
	2700: "Wireless Profile Added",
	2701: "Wireless Profile Removed",
	2702: "Wireless Profile Updated",
	2703: "An existing profile sync was modified",
	2704: "An existing profile link preference was changed",
	2705: "Wireless profile share with UEFI enabled setting was changed",
	2800: "EAC Posture Signer SET",
	2801: "EAC Enabled",
	2802: "EAC Disabled",
	2803: "EAC Posture State",
	2804: "EAC Set Options",
	2900: "KVM Opt-in Enabled",
	2901: "KVM Opt-in Disabled",
	2902: "KVM Password Changed",
	2903: "KVM Consent Succeeded",
	2904: "KVM Consent Failed",
	3000: "Opt-In Policy Change",
	3001: "Send Consent Code Event",
	3002: "Start Opt-In Blocked Event",
	3301: "A user has modified the Watchdog Action settings",
	3302: "A user has modified a Watchdog to add, remove, or alter the Watchdog Action connected to it",
}
