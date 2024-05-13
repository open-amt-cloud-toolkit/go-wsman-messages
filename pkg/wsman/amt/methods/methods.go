/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package methods provides utility functions used across the amt packages
package methods

import "fmt"

// RequestStateChange creates an amt specific request state change action.
func RequestStateChange(className string) string {
	return fmt.Sprintf("http://intel.com/wbem/wscim/1/amt-schema/1/%s/RequestStateChange", className)
}

// GenerateAction creates an amt specific action for custom methods.
func GenerateAction(className, methodName string) string {
	return fmt.Sprintf("http://intel.com/wbem/wscim/1/amt-schema/1/%s/%s", className, methodName)
}

// GenerateInputMethod creates the string used to populate the XML tag for INPUT calls.
func GenerateInputMethod(methodName string) string {
	return fmt.Sprintf("%s_INPUT", methodName)
}
