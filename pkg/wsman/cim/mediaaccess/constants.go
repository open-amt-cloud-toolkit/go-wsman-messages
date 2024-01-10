/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package mediaaccess

const CIM_MediaAccessDevice string = "CIM_MediaAccessDevice"

const (
	CapabilitiesValuesUnknown CapabilitiesValues = iota
	CapabilitiesValuesOther
	CapabilitiesValuesSequentialAccess
	CapabilitiesValuesRandomAccess
	CapabilitiesValuesSupportsWriting
	CapabilitiesValuesEncryption
	CapabilitiesValuesCompression
	CapabilitiesValuesSupportsRemoveableMedia
	CapabilitiesValuesManualCleaning
	CapabilitiesValuesAutomaticCleaning
	CapabilitiesValuesSMARTNotification
	CapabilitiesValuesSupportsDualSidedMedia
	CapabilitiesValuesPredismountEjectNotRequired
)

const (
	SecurityValuesOther SecurityValues = iota + 1
	SecurityValuesUnknown
	SecurityValuesNone
	SecurityValuesReadOnly
	SecurityValuesLockedOut
	SecurityValuesBootBypass
	SecurityValuesBootBypassAndReadOnly
)
