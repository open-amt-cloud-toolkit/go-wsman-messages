/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package userinitiatedconnection

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func (r *Response) JSON() string {
	jsonOutput, err := json.Marshal(r.Body)
	if err != nil {
		return ""
	}
	return string(jsonOutput)
}

func (r *Response) YAML() string {
	yamlOutput, err := yaml.Marshal(r.Body)
	if err != nil {
		return ""
	}
	return string(yamlOutput)
}
