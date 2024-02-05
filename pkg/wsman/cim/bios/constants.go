/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package bios

const (
	CIM_BIOSElement string = "CIM_BIOSElement"
)

const (
	Unknown TargetOperatingSystem = iota
	Other
	MACOS
	ATTUNIX
	DGUX
	DECNT
	Tru64UNIX
	OpenVMS
	HPUX
	AIX
	MVS
	OS400
	OS2
	JavaVM
	MSDOS
	WIN3x
	WIN95
	WIN98
	WINNT
	WINCE
	NCR3000
	NetWare
	OSF
	DCOS
	ReliantUNIX
	SCOUnixWare
	SCOOpenServer
	Sequent
	IRIX
	Solaris
	SunOS
	U6000
	ASERIES
	HPNonStopOS
	HPNonStopOSS
	BS2000
	LINUX
	Lynx
	XENIX
	VM
	InteractiveUNIX
	BSDUNIX
	FreeBSD
	NetBSD
	GNUHurd
	OS9
	MACHKernel
	Inferno
	QNX
	EPOC
	IxWorks
	VxWorks
	MiNT
	BeOS
	HPMPE
	NextStep
	PalmPilot
	Rhapsody
	Windows2000
	Dedicated
	OS390
	VSE
	TPF
	WindowsMe
	CalderaOpenUNIX
	OpenBSD
	NotApplicable
	WindowsXP
	ZOS
	MicrosoftWindowsServer2003
	MicrosoftWindowsServer200364Bit
	WindowsXP64Bit
	WindowsXPEmbedded
	WindowsVista
	WindowsVista64Bit
	WindowsEmbeddedforPointofService
	MicrosoftWindowsServer2008
	MicrosoftWindowsServer200864Bit
	FreeBSD64Bit
	RedHatEnterpriseLinux
	RedHatEnterpriseLinux64Bit
	Solaris64Bit
	SUSE
	SUSE64Bit
	SLES
	SLES64Bit
	NovellOES
	NovellLinuxDesktop
	SunJavaDesktopSystem
	Mandriva
	Mandriva64Bit
	TurboLinux
	TurboLinux64Bit
	Ubuntu
	Ubuntu64Bit
	Debian
	Debian64Bit
	Linux24x
	Linux24x64Bit
	Linux26x
	Linux26x64Bit
	Linux64Bit
	Other64Bit
	MicrosoftWindowsServer2008R2
	VMwareESXi
	MicrosoftWindows7
	CentOS32bit
	CentOS64bit
	OracleEnterpriseLinux32bit
	OracleEnterpriseLinux64bit
	EComStation32bitx
	MicrosoftWindowsServer2011
	MicrosoftWindowsServer201164Bit
	MicrosoftWindowsServer8
)

const (
	Deployable SoftwareElementState = iota
	Installable
	Executable
	Running
)

const (
	OperationalStatusUnknown OperationalStatus = iota
	OperationalStatusOther
	OperationalStatusOK
	OperationalStatusDegraded
	OperationalStatusStressed
	OperationalStatusPredictiveFailure
	OperationalStatusError
	OperationalStatusNonRecoverableError
	OperationalStatusStarting
	OperationalStatusStopping
	OperationalStatusStopped
	OperationalStatusInService
	OperationalStatusNoContact
	OperationalStatusLostCommunication
	OperationalStatusAborted
	OperationalStatusDormant
	OperationalStatusSupportingEntityinError
	OperationalStatusCompleted
	OperationalStatusPowerMode
	OperationalStatusRelocating
)
