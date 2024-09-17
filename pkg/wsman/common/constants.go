/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package common

// TODO: Review if this file is still necessary.

const ValueNotFound string = "Value not found in map"

// OUTPUTS Lookups

// ReturnValuesToString is a map of return values to their string representation.
var ReturnValuesToString = map[int]string{
	0:    "PT_STATUS_SUCCESS",
	1:    "PT_STATUS_INTERNAL_ERROR",
	3:    "PT_STATUS_INVALID_PT_MODE",
	9:    "PT_STATUS_INVALID_REGISTRATION_DATA",
	10:   "PT_STATUS_APPLICATION_DOES_NOT_EXIST",
	11:   "PT_STATUS_NOT_ENOUGH_STORAGE",
	12:   "PT_STATUS_INVALID_NAME",
	13:   "PT_STATUS_BLOCK_DOES_NOT_EXIST",
	14:   "PT_STATUS_INVALID_BYTE_OFFSET",
	15:   "PT_STATUS_INVALID_BYTE_COUNT",
	16:   "PT_STATUS_NOT_PERMITTED",
	17:   "PT_STATUS_NOT_OWNER",
	18:   "PT_STATUS_BLOCK_LOCKED_BY_OTHER",
	19:   "PT_STATUS_BLOCK_NOT_LOCKED",
	20:   "PT_STATUS_INVALID_GROUP_PERMISSIONS",
	21:   "PT_STATUS_GROUP_DOES_NOT_EXIST",
	22:   "PT_STATUS_INVALID_MEMBER_COUNT",
	23:   "PT_STATUS_MAX_LIMIT_REACHED",
	24:   "PT_STATUS_INVALID_AUTH_TYPE",
	26:   "PT_STATUS_INVALID_DHCP_MODE",
	27:   "PT_STATUS_INVALID_IP_ADDRESS",
	28:   "PT_STATUS_INVALID_DOMAIN_NAME",
	30:   "PT_STATUS_REQUEST_UNEXPECTED",
	32:   "PT_STATUS_INVALID_PROVISIONING_STATE",
	34:   "PT_STATUS_INVALID_TIME",
	35:   "PT_STATUS_INVALID_INDEX",
	36:   "PT_STATUS_INVALID_PARAMETER",
	37:   "PT_STATUS_INVALID_NETMASK",
	38:   "PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED",
	2049: "PT_STATUS_UNSUPPORTED_OEM_NUMBER",
	2050: "PT_STATUS_UNSUPPORTED_BOOT_OPTION",
	2051: "PT_STATUS_INVALID_COMMAND",
	2052: "PT_STATUS_INVALID_SPECIAL_COMMAND",
	2053: "PT_STATUS_INVALID_HANDLE",
	2054: "PT_STATUS_INVALID_PASSWORD",
	2055: "PT_STATUS_INVALID_REALM",
	2056: "PT_STATUS_STORAGE_ACL_ENTRY_IN_USE",
	2057: "PT_STATUS_DATA_MISSING",
	2058: "PT_STATUS_DUPLICATE",
	2059: "PT_STATUS_EVENT_LOG_FROZEN",
	2060: "PT_STATUS_PKI_MISSING_KEYS",
	2061: "PT_STATUS_PKI_GENERATING_KEYS",
	2062: "PT_STATUS_INVALID_KEY",
	2063: "PT_STATUS_INVALID_CERT",
	2064: "PT_STATUS_CERT_KEY_NOT_MATCH",
	2065: "PT_STATUS_MAX_KERB_DOMAIN_REACHED",
	2066: "PT_STATUS_UNSUPPORTED",
	2067: "PT_STATUS_INVALID_PRIORITY",
	2068: "PT_STATUS_NOT_FOUND",
	2069: "PT_STATUS_INVALID_CREDENTIALS",
	2070: "PT_STATUS_INVALID_PASSPHRASE",
	2072: "PT_STATUS_NO_ASSOCIATION",
	2075: "PT_STATUS_AUDIT_FAIL",
	2076: "PT_STATUS_BLOCKING_COMPONENT",
	2081: "PT_STATUS_USER_CONSENT_REQUIRED",
	2082: "PT_STATUS_OPERATION_IN_PROGRESS",
}

// ConvertReturnValueToString returns the string representation of the return value.
func ConvertReturnValueToString(value int) string {
	if value, exists := ReturnValuesToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// EnabledStateToString is a map of EnabledState values to their string representations.
var EnabledStateToString = map[int]string{
	0:  "Unknown",
	1:  "Other",
	2:  "Enabled",
	3:  "Disabled",
	4:  "ShuttingDown",
	5:  "NotApplicable",
	6:  "EnabledbutOffline",
	7:  "InTest",
	8:  "Deferred",
	9:  "Quiesce",
	10: "Starting",
}

// ConvertEnabledStateToString returns the string representation of an EnabledState value.
func ConvertEnabledStateToString(value int) string {
	if value, exists := EnabledStateToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// RequestedStateToString is a map of RequestedState values to their string representations.
var RequestedStateToString = map[int]string{
	0:  "Unknown",
	2:  "Enabled",
	3:  "Disabled",
	4:  "ShutDown",
	5:  "NoChange",
	6:  "Offline",
	7:  "Test",
	8:  "Deferred",
	9:  "Quiesce",
	10: "Reboot",
	11: "Reset",
	12: "NotApplicable",
}

// ConvertRequestedStateToString returns the string representation of a RequestedState value.
func ConvertRequestedStateToString(value int) string {
	if value, exists := RequestedStateToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// OperationalStatusToString is a map of OperationalStatus value to string.
var OperationalStatusToString = map[int]string{
	0:  "Unknown",
	1:  "Other",
	2:  "OK",
	3:  "Degraded",
	4:  "Stressed",
	5:  "Predictive Failure",
	6:  "Error",
	7:  "Non-Recoverable Error",
	8:  "Starting",
	9:  "Stopping",
	10: "Stopped",
	11: "In Service",
	12: "No Contact",
	13: "Lost Communication",
	14: "Aborted",
	15: "Dormant",
	16: "Supporting Entity in Error",
	17: "Completed",
	18: "Power Mode",
	19: "Relocating",
}

// ConvertOperationalStatusToString returns the string representation of OperationalStatus.
func ConvertOperationalStatusToString(value int) string {
	if value, exists := OperationalStatusToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// EnabledDefaultToString is a map of EnabledDefault value to string.
var EnabledDefaultToString = map[int]string{
	2: "Enabled",
	3: "Disabled",
	5: "Not Applicable",
	6: "Enabled but Offline",
	7: "No Default",
	9: "Quiesce",
}

// ConvertEnabledDefaultToString returns the string representation of EnabledDefault.
func ConvertEnabledDefaultToString(value int) string {
	if value, exists := EnabledDefaultToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// HealthStateToString indicates the current health of the element.
var HealthStateToString = map[int]string{
	0:  "Unknown",
	5:  "OK",
	10: "Degraded",
	15: "Minor Failure",
	20: "Major Failure",
	25: "Critical Failure",
	30: "Non-recoverable Error",
}

// ConvertHealthStateToString returns the string representation of the health state.
func ConvertHealthStateToString(value int) string {
	if value, exists := HealthStateToString[value]; exists {
		return value
	}

	return ValueNotFound
}

// PackageTypeToString is a mapping of the PackageType value to a string.
var PackageTypeToString = map[int]string{
	0:  "Unknown",
	1:  "Other",
	2:  "Rack",
	3:  "Chassis Frame",
	4:  "Cross Connect Back Plane",
	5:  "Container Frame Slot",
	6:  "Power Supply",
	7:  "Fan",
	8:  "Sensor",
	9:  "Module Card",
	10: "Port Connector",
	11: "Battery",
	12: "Processor",
	13: "Memory",
	14: "Power Source Generator",
	15: "Storage Media Package",
	16: "Blade",
	17: "Blade Expansion",
}

// ConvertPackageTypeToString returns the string representation of the PackageType value.
func ConvertPackageTypeToString(value int) string {
	if value, exists := PackageTypeToString[value]; exists {
		return value
	}

	return ValueNotFound
}
