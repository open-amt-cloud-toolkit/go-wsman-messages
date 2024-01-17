/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package methods

import "fmt"

func RequestStateChange(className string) string {
	return fmt.Sprintf("http://intel.com/wbem/wscim/1/ips-schema/1/%s/RequestStateChange", className)
}

func GenerateAction(className string, methodName string) string {
	return fmt.Sprintf("http://intel.com/wbem/wscim/1/ips-schema/1/%s/%s", className, methodName)
}

func GenerateInputMethod(methodName string) string {
	return fmt.Sprintf("%s_INPUT", methodName)
}
