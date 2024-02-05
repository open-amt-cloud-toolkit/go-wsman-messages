/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

// Package models provides a set of utility types, constants, and functions that are used broadly across amt, cim, and ips packages
package models

// HasSelector checks the SelectorSet and returns true if the SelectorSet contains a Selector
func (rp *ReferenceParmetersNoNamespace) HasSelector(name string, value string) bool {
	for _, selector := range rp.SelectorSet {
		if selector.Name == name && selector.Value == value {
			return true
		}
	}
	return false
}

// GetSelectorValue returns the Value property of a selector found in a SelectorSet based on the selector name
func (rp *ReferenceParmetersNoNamespace) GetSelectorValue(name string) string {
	for _, selector := range rp.SelectorSet {
		if selector.Name == name {
			return selector.Value
		}
	}
	return ""
}
