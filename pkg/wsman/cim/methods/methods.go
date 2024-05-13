/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package methods provides utility functions used across the cim packages
package methods

import (
	"fmt"
)

func RequestStateChange(className string) string {
	return fmt.Sprintf("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/%s/RequestStateChange", className)
}

func GenerateAction(className, methodName string) string {
	return fmt.Sprintf("http://schemas.dmtf.org/wbem/wscim/1/cim-schema/2/%s/%s", className, methodName)
}

func GenerateInputMethod(methodName string) string {
	return fmt.Sprintf("%s_INPUT", methodName)
}
