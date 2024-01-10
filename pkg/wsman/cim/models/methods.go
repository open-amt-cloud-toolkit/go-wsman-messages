/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package models

func (rp *ReferenceParmetersNoNamespace) HasSelector(name string, value string) bool {
	for _, selector := range rp.SelectorSet {
		if selector.Name == name && selector.Value == value {
			return true
		}
	}
	return false
}

func (rp *ReferenceParmetersNoNamespace) GetSelectorValue(name string) string {
	for _, selector := range rp.SelectorSet {
		if selector.Name == name {
			return selector.Value
		}
	}
	return ""
}
