/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package environmentdetection

// INPUTS Constants
const (
	AMT_EnvironmentDetectionSettingData string = "AMT_EnvironmentDetectionSettingData"
)

const (
	LocalDomains DetectionAlgorithm = iota
	RemoteURLs
)

// DetectionAlgorithmToString is a map of DetectionAlgorithm values to their string representations
var detectionAlgorithmToString = map[DetectionAlgorithm]string{
	LocalDomains: "LocalDomains",
	RemoteURLs:   "RemoteURLs",
}

// String returns the string representation of the DetectionAlgorithm value
func (d DetectionAlgorithm) String() string {
	if value, exists := detectionAlgorithmToString[d]; exists {
		return value
	}

	return "Value not found in map"
}
