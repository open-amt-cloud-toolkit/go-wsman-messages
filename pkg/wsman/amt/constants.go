/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package amt

// Realms represents a set of enumerated constants for managing various aspects of the system.
type Realms int

const (
	ADMINISTRATION                Realms = 3  // ADMINISTRATION manages security control data, power saving options, Intel AMT setup and configuration, and local network options.(
	AGENT_PRESENCE_LOCAL          Realms = 9  // AGENT_PRESENCE_LOCAL is used by an application designed to run on the local platform to report that it is running and to send heartbeats periodically.(
	AGENT_PRESENCE_REMOTE         Realms = 10 // AGENT_PRESENCE_REMOTE is used to register Local Agent applications and to specify the behavior of Intel AMT when an application is running or stops running unexpectedly.(
	AUDIT_LOG                     Realms = 20 // AUDIT_LOG configures the Audit Log.(
	CIRCUIT_BREAKER               Realms = 11 // CIRCUIT_BREAKER is used to define filters, counters, and policies to monitor incoming and outgoing network traffic and to block traffic when a suspicious condition is detected.(
	ENDPOINT_ACCESS_CONTROL       Realms = 17 // ENDPOINT_ACCESS_CONTROL is deprecated and not supported starting CSME release 18.0. Returns settings associated with NAC/NAP posture.(
	ENDPOINT_ACCESS_CONTROL_ADMIN Realms = 18 // ENDPOINT_ACCESS_CONTROL_ADMIN is deprecated and not supported starting CSME release 18.0. Configures and enables the NAC/NAP posture.(
	EVENT_LOG_READER              Realms = 19 // EVENT_LOG_READER controls access for reading the Intel AMT event log.(
	EVENT_MANAGER                 Realms = 7  // EVENT_MANAGER allows configuring hardware and software events to generate alerts and to send them to a remote console and/or log them locally.(
	GENERAL_INFO                  Realms = 13 // GENERAL_INFO returns general setting and status information.(
	HARDWARE_ASSET                Realms = 4  // HARDWARE_ASSET is used to retrieve information about the hardware inventory of the platform.(
	LOCAL_APPS                    Realms = 24 // LOCAL_APPS provides alerts to a user on the local interface.(
	NETWORK_TIME                  Realms = 12 // NETWORK_TIME is used to set the clock in the Intel AMT device and synchronize it to network time.(
	REDIRECTION                   Realms = 2  // REDIRECTION enables and disables the redirection capability and retrieves the redirection log.(
	REMOTE_CONTROL                Realms = 5  // REMOTE_CONTROL enables powering a platform up or down remotely.(
	STORAGE                       Realms = 6  // STORAGE is used to access, configure, manage, write to and read from non-volatile user storage.(
	STORAGE_ADMIN                 Realms = 8  // STORAGE_ADMIN is used to configure the global parameters that govern the allocation and use of non-volatile storage.(
	USER_ACCESS_CONTROL           Realms = 21 // USER_ACCESS_CONTROL allows users to control the properties of their own ACL entries.
)
