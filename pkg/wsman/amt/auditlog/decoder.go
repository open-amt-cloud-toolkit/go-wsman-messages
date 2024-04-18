/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

// INPUTS Constants
const (
	AMT_AuditLog string = "AMT_AuditLog"
	ReadRecords  string = "ReadRecords"
)

const (
	Unknown                   OverwritePolicy = 0
	WrapsWhenFull             OverwritePolicy = 2
	NeverOverwrites           OverwritePolicy = 7
	PartialRestrictedRollover OverwritePolicy = 32768
)

// overwritePolicyToString is a map of OverwritePolicy values to their string representations
var overwritePolicyToString = map[OverwritePolicy]string{
	Unknown:                   "Unknown",
	WrapsWhenFull:             "WrapsWhenFull",
	NeverOverwrites:           "NeverOverwrites",
	PartialRestrictedRollover: "PartialRestrictedRollover",
}

// String returns the string representation of the OverwritePolicy value
func (o OverwritePolicy) String() string {
	if value, exists := overwritePolicyToString[o]; exists {
		return value
	}

	return "Value not found in map"
}

const (
	NoRollOver StoragePolicy = iota
	RollOver
	RestrictedRollOver
)

// storagePolicyToString is a map of StoragePolicy values to their string representations
var storagePolicyToString = map[StoragePolicy]string{
	NoRollOver:         "NoRollOver",
	RollOver:           "RollOver",
	RestrictedRollOver: "RestrictedRollOver",
}

// String returns the string representation of the StoragePolicy value
func (s StoragePolicy) String() string {
	if value, exists := storagePolicyToString[s]; exists {
		return value
	}

	return "Value not found in map"
}
