/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package messagelog

const (
	AMT_MessageLog        string = "AMT_MessageLog"
	GetRecords            string = "GetRecords"
	PositionToFirstRecord string = "PositionToFirstRecord"
)

const (
	CharacterSetUnknown CharacterSet = iota
	CharacterSetOther
	CharacterSetASCII
	CharacterSetUnicode
	CharacterSetISO2022
	CharacterSetISO8859
	CharacterSetExtendedUNIXCode
	CharacterSetUTF8
	CharacterSetUCS2
	CharacterSetBitmappedData
	CharacterSetOctetString
	CharacterSetDefinedbyIndividualRecords
)

const (
	LastChangeUnknown LastChange = iota
	LastChangeAdd
	LastChangeDelete
	LastChangeModify
	LastChangeLogCleared
)

const (
	OverwritePolicyUnknown OverwritePolicy = iota
	OverwritePolicyOther
	OverwritePolicyWrapsWhenFull
	OverwritePolicyClearWhenNearFull
	OverwritePolicyOverwriteOutdatedWhenNeeded
	OverwritePolicyRemoveOutdatedRecords
	OverwritePolicyOverwriteSpecificRecords
	OverwritePolicyNeverOverwrite
)

const (
	CapabilitiesUnknown Capabilities = iota
	CapabilitiesOther
	CapabilitiesWriteRecordSupported
	CapabilitiesDeleteRecordSupported
	CapabilitiesCanMoveBackwardinLog
	CapabilitiesFreezeLogSupported
	CapabilitiesClearLogSupported
	CapabilitiesSupportsAddressingbyOrdinalRecordNumber
	CapabilitiesVariableLengthRecordsSupported
	CapabilitiesVariableFormatsforRecords
	CapabilitiesCanFlagRecordsforOverwrite
)

const (
	LogStateUnknown       LogState = 0
	LogStateNormal        LogState = 2
	LogStateErasing       LogState = 3
	LogStateNotApplicable LogState = 4
)
