/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

import "testing"

func TestTargetOperatingSystem_String(t *testing.T) {
	tests := []struct {
		state    TargetOperatingSystem
		expected string
	}{
		{Unknown, "Unknown"},
		{Other, "Other"},
		{MACOS, "MACOS"},
		{ATTUNIX, "ATTUNIX"},
		{DGUX, "DGUX"},
		{DECNT, "DECNT"},
		{Tru64UNIX, "Tru64UNIX"},
		{OpenVMS, "OpenVMS"},
		{HPUX, "HPUX"},
		{AIX, "AIX"},
		{MVS, "MVS"},
		{OS400, "OS400"},
		{OS2, "OS2"},
		{JavaVM, "JavaVM"},
		{MSDOS, "MSDOS"},
		{WIN3x, "WIN3x"},
		{WIN95, "WIN95"},
		{WIN98, "WIN98"},
		{WINNT, "WINNT"},
		{WINCE, "WINCE"},
		{NCR3000, "NCR3000"},
		{NetWare, "NetWare"},
		{OSF, "OSF"},
		{DCOS, "DCOS"},
		{ReliantUNIX, "ReliantUNIX"},
		{SCOUnixWare, "SCOUnixWare"},
		{SCOOpenServer, "SCOOpenServer"},
		{Sequent, "Sequent"},
		{IRIX, "IRIX"},
		{Solaris, "Solaris"},
		{SunOS, "SunOS"},
		{U6000, "U6000"},
		{ASERIES, "ASERIES"},
		{HPNonStopOS, "HPNonStopOS"},
		{HPNonStopOSS, "HPNonStopOSS"},
		{BS2000, "BS2000"},
		{LINUX, "LINUX"},
		{Lynx, "Lynx"},
		{XENIX, "XENIX"},
		{VM, "VM"},
		{InteractiveUNIX, "InteractiveUNIX"},
		{BSDUNIX, "BSDUNIX"},
		{FreeBSD, "FreeBSD"},
		{NetBSD, "NetBSD"},
		{GNUHurd, "GNUHurd"},
		{OS9, "OS9"},
		{MACHKernel, "MACHKernel"},
		{Inferno, "Inferno"},
		{QNX, "QNX"},
		{EPOC, "EPOC"},
		{IxWorks, "IxWorks"},
		{VxWorks, "VxWorks"},
		{MiNT, "MiNT"},
		{BeOS, "BeOS"},
		{HPMPE, "HPMPE"},
		{NextStep, "NextStep"},
		{PalmPilot, "PalmPilot"},
		{Rhapsody, "Rhapsody"},
		{Windows2000, "Windows2000"},
		{Dedicated, "Dedicated"},
		{OS390, "OS390"},
		{VSE, "VSE"},
		{TPF, "TPF"},
		{WindowsMe, "WindowsMe"},
		{CalderaOpenUNIX, "CalderaOpenUNIX"},
		{OpenBSD, "OpenBSD"},
		{NotApplicable, "NotApplicable"},
		{WindowsXP, "WindowsXP"},
		{ZOS, "ZOS"},
		{MicrosoftWindowsServer2003, "MicrosoftWindowsServer2003"},
		{MicrosoftWindowsServer200364Bit, "MicrosoftWindowsServer200364Bit"},
		{WindowsXP64Bit, "WindowsXP64Bit"},
		{WindowsXPEmbedded, "WindowsXPEmbedded"},
		{WindowsVista, "WindowsVista"},
		{WindowsVista64Bit, "WindowsVista64Bit"},
		{WindowsEmbeddedForPointOfService, "WindowsEmbeddedForPointOfService"},
		{MicrosoftWindowsServer2008, "MicrosoftWindowsServer2008"},
		{MicrosoftWindowsServer200864Bit, "MicrosoftWindowsServer200864Bit"},
		{FreeBSD64Bit, "FreeBSD64Bit"},
		{RedHatEnterpriseLinux, "RedHatEnterpriseLinux"},
		{RedHatEnterpriseLinux64Bit, "RedHatEnterpriseLinux64Bit"},
		{Solaris64Bit, "Solaris64Bit"},
		{SUSE, "SUSE"},
		{SUSE64Bit, "SUSE64Bit"},
		{SLES, "SLES"},
		{SLES64Bit, "SLES64Bit"},
		{NovellOES, "NovellOES"},
		{NovellLinuxDesktop, "NovellLinuxDesktop"},
		{SunJavaDesktopSystem, "SunJavaDesktopSystem"},
		{Mandriva, "Mandriva"},
		{Mandriva64Bit, "Mandriva64Bit"},
		{TurboLinux, "TurboLinux"},
		{TurboLinux64Bit, "TurboLinux64Bit"},
		{Ubuntu, "Ubuntu"},
		{Ubuntu64Bit, "Ubuntu64Bit"},
		{Debian, "Debian"},
		{Debian64Bit, "Debian64Bit"},
		{Linux24x, "Linux24x"},
		{Linux24x64Bit, "Linux24x64Bit"},
		{Linux26x, "Linux26x"},
		{Linux26x64Bit, "Linux26x64Bit"},
		{Linux64Bit, "Linux64Bit"},
		{Other64Bit, "Other64Bit"},
		{MicrosoftWindowsServer2008R2, "MicrosoftWindowsServer2008R2"},
		{VMwareESXi, "VMwareESXi"},
		{MicrosoftWindows7, "MicrosoftWindows7"},
		{CentOS32bit, "CentOS32bit"},
		{CentOS64bit, "CentOS64bit"},
		{OracleEnterpriseLinux32bit, "OracleEnterpriseLinux32bit"},
		{OracleEnterpriseLinux64bit, "OracleEnterpriseLinux64bit"},
		{EComStation32BitX, "EComStation32BitX"},
		{MicrosoftWindowsServer2011, "MicrosoftWindowsServer2011"},
		{MicrosoftWindowsServer201164Bit, "MicrosoftWindowsServer201164Bit"},
		{MicrosoftWindowsServer8, "MicrosoftWindowsServer8"},
		{TargetOperatingSystem(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestSoftwareElementState_String(t *testing.T) {
	tests := []struct {
		state    SoftwareElementState
		expected string
	}{
		{SoftwareElementStateDeployable, "Deployable"},
		{SoftwareElementStateInstallable, "Installable"},
		{SoftwareElementStateExecutable, "Executable"},
		{SoftwareElementStateRunning, "Running"},
		{SoftwareElementState(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestOperationalStatus_String(t *testing.T) {
	tests := []struct {
		state    OperationalStatus
		expected string
	}{
		{OperationalStatusUnknown, "Unknown"},
		{OperationalStatusOther, "Other"},
		{OperationalStatusOK, "OK"},
		{OperationalStatusDegraded, "Degraded"},
		{OperationalStatusStressed, "Stressed"},
		{OperationalStatusPredictiveFailure, "PredictiveFailure"},
		{OperationalStatusError, "Error"},
		{OperationalStatusNonRecoverableError, "NonRecoverableError"},
		{OperationalStatusStarting, "Starting"},
		{OperationalStatusStopping, "Stopping"},
		{OperationalStatusStopped, "Stopped"},
		{OperationalStatusInService, "InService"},
		{OperationalStatusNoContact, "NoContact"},
		{OperationalStatusLostCommunication, "LostCommunication"},
		{OperationalStatusAborted, "Aborted"},
		{OperationalStatusDormant, "Dormant"},
		{OperationalStatusSupportingEntityInError, "SupportingEntityInError"},
		{OperationalStatusCompleted, "Completed"},
		{OperationalStatusPowerMode, "PowerMode"},
		{OperationalStatusRelocating, "Relocating"},
		{OperationalStatus(999), "Value not found in map"},
	}

	for _, test := range tests {
		result := test.state.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}
