/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package auditlog

import (
	"testing"
	"time"
)

func TestOverwritePolicy_String(t *testing.T) {
	tests := []struct {
		state    OverwritePolicy
		expected string
	}{
		{OverwritePolicyUnknown, "Unknown"},
		{OverwritePolicyWrapsWhenFull, "WrapsWhenFull"},
		{OverwritePolicyNeverOverwrites, "NeverOverwrites"},
		{OverwritePolicyPartialRestrictedRollover, "PartialRestrictedRollover"},
		{OverwritePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestStoragePolicy_String(t *testing.T) {
	tests := []struct {
		state    StoragePolicy
		expected string
	}{
		{StoragePolicyNoRollOver, "NoRollOver"},
		{StoragePolicyRollOver, "RollOver"},
		{StoragePolicyRestrictedRollOver, "RestrictedRollOver"},
		{StoragePolicy(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestEnabledState_String(t *testing.T) {
	tests := []struct {
		state    EnabledState
		expected string
	}{
		{EnabledStateUnknown, "Unknown"},
		{EnabledStateOther, "Other"},
		{EnabledStateEnabled, "Enabled"},
		{EnabledStateDisabled, "Disabled"},
		{EnabledStateShuttingDown, "ShuttingDown"},
		{EnabledStateNotApplicable, "NotApplicable"},
		{EnabledStateEnabledButOffline, "EnabledButOffline"},
		{EnabledStateInTest, "InTest"},
		{EnabledStateDeferred, "Deferred"},
		{EnabledStateQuiesce, "Quiesce"},
		{EnabledStateStarting, "Starting"},
		{EnabledState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestRequestedState_String(t *testing.T) {
	tests := []struct {
		state    RequestedState
		expected string
	}{
		{RequestedStateUnknown, "Unknown"},
		{RequestedStateEnabled, "Enabled"},
		{RequestedStateDisabled, "Disabled"},
		{RequestedStateShutDown, "ShutDown"},
		{RequestedStateNoChange, "NoChange"},
		{RequestedStateOffline, "Offline"},
		{RequestedStateTest, "Test"},
		{RequestedStateDeferred, "Deferred"},
		{RequestedStateQuiesce, "Quiesce"},
		{RequestedStateReboot, "Reboot"},
		{RequestedStateReset, "Reset"},
		{RequestedStateNotApplicable, "NotApplicable"},
		{RequestedState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestGetAuditLogExtendedDataString(t *testing.T) {
	tests := []struct {
		name         string
		auditEventId int
		data         string
		expected     string
	}{
		{"ACLEntryAdded", ACLEntryAdded, "\x00\x05Hello World", "Hello"},
		{"ACLEntryRemoved", ACLEntryRemoved, "\x00\x05Hello World", "Hello"},
		{"ACLEntryModified", ACLEntryModified, "\x01\x00Hello World", "Hello World"},
		{"ACLAccessWithInvalidCredentials", ACLAccessWithInvalidCredentials, "\x00", "Invalid ME access"},
		{"ACLAccessWithInvalidCredentials", ACLAccessWithInvalidCredentials, "\x01", "Invalid MEBx access"},
		{"ACLEntryStateChanged", ACLEntryStateChanged, "\x00\x00Hello World", "Disabled, Hello World"},
		{"ACLEntryStateChanged", ACLEntryStateChanged, "\x01\x01", "Enabled"},
		{"TLSStateChanged", TLSStateChanged, "\x01\x02", "Remote ServerAuth, Local MutualAuth"},
		{"SetRealmAuthenticationMode", SetRealmAuthenticationMode, "\x00\x00\x00\x00\x02", "Redirection, Disabled"},
		{"AMTUnprovisioningStarted", AMTUnprovisioningStarted, "\x03", "Local WSMAN"},
		{"FirmwareUpdate", FirmwareUpdate, "\x00\x01\x00\x02\x00\x03\x00\x04\x00\x05\x00\x06\x00\x07\x00\x08", "From 1.2.3.4 to 5.6.7.8"},
		{"AMTTimeSet", AMTTimeSet, "\x00\x00\x00\x00", time.Unix(0, 0).Local().Format(time.RFC1123)},
		{"OptInPolicyChange", OptInPolicyChange, "\x00\x01", "From None to KVM"},
		{"SendConsentCode", SendConsentCode, "\x00", "Success"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetAuditLogExtendedDataString(tt.auditEventId, tt.data)
			if result != tt.expected {
				t.Errorf("GetAuditLogExtendedDataString(%d, %q) = %v; want %v", tt.auditEventId, tt.data, result, tt.expected)
			}
		})
	}
}
