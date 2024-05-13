/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package chip

const CIMChip string = "CIM_Chip"

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityInError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)

// operationalStatusMap is a map of the OperationalStatus enumeration.
var operationalStatusMap = map[OperationalStatus]string{
	OperationalStatusUnknown:                 "Unknown",
	OperationalStatusOther:                   "Other",
	OperationalStatusOK:                      "OK",
	OperationalStatusDegraded:                "Degraded",
	OperationalStatusStressed:                "Stressed",
	OperationalStatusPredictiveFailure:       "Predictive Failure",
	OperationalStatusError:                   "Error",
	OperationalStatusNonRecoverableError:     "Non-Recoverable Error",
	OperationalStatusStarting:                "Starting",
	OperationalStatusStopping:                "Stopping",
	OperationalStatusStopped:                 "Stopped",
	OperationalStatusInService:               "In Service",
	OperationalStatusNoContact:               "No Contact",
	OperationalStatusLostCommunication:       "Lost Communication",
	OperationalStatusAborted:                 "Aborted",
	OperationalStatusDormant:                 "Dormant",
	OperationalStatusSupportingEntityInError: "Supporting Entity In Error",
	OperationalStatusCompleted:               "Completed",
	OperationalStatusPowerMode:               "Power Mode",
	OperationalStatusRelocating:              "Relocating",
}

// String returns a human-readable string representation of the OperationalStatus enumeration.
func (e OperationalStatus) String() string {
	if s, ok := operationalStatusMap[e]; ok {
		return s
	}

	return "Value not found in map"
}
