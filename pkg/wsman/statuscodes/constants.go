package statuscodes

// The following values can be returned by Intel AMT as status codes. Methods in the Intel proprietary classes list the codes applicable to the method.
const (
	PT_STATUS_SUCCESS        = iota // Operation completed successfully.
	PT_STATUS_INTERNAL_ERROR        // An internal error occurred while performing the operation.
	_
	PT_STATUS_INVALID_PT_MODE // Specified mode of operation is invalid.
	_
	_
	_
	_
	_
	PT_STATUS_INVALID_REGISTRATION_DATA  // 1. Either an invalid name was entered or an “Enterprise” name was specified that was not pre-registered. 2. The current registration was attempted from an interface different from the one used for the initial registration of the application.
	PT_STATUS_APPLICATION_DOES_NOT_EXIST // The application handle provided in the request message is not valid.
	PT_STATUS_NOT_ENOUGH_STORAGE         // The number of bytes requested cannot be allocated in ISV storage.
	PT_STATUS_INVALID_NAME               // Specified name is invalid.
	PT_STATUS_BLOCK_DOES_NOT_EXIST       // The specified block does not exist.
	PT_STATUS_INVALID_BYTE_OFFSET        // The specified byte offset is invalid.
	PT_STATUS_INVALID_BYTE_COUNT         // The specified byte count is invalid.
	PT_STATUS_NOT_PERMITTED              // The requesting application is not permitted to request execution of the specified operation.
	PT_STATUS_NOT_OWNER                  // The requesting application is not the owner of the block as required for the requested operation.
	PT_STATUS_BLOCK_LOCKED_BY_OTHER      // The specified block is locked by another application.
	PT_STATUS_BLOCK_NOT_LOCKED           // The specified block is not locked.
	PT_STATUS_INVALID_GROUP_PERMISSIONS  // The specified group permission bits are invalid.
	PT_STATUS_GROUP_DOES_NOT_EXIST       // The specified group does not exist.
	PT_STATUS_INVALID_MEMBER_COUNT       // The specified member count is invalid.
	PT_STATUS_MAX_LIMIT_REACHED          // No available storage in the specified structure.
	PT_STATUS_INVALID_AUTH_TYPE          // Specified Key algorithm is invalid.
	_
	PT_STATUS_INVALID_DHCP_MODE   // Specified DHCP mode is invalid.
	PT_STATUS_INVALID_IP_ADDRESS  // Specified IP address is invalid.
	PT_STATUS_INVALID_DOMAIN_NAME // Specified Domain name is invalid.
	_
	PT_STATUS_REQUEST_UNEXPECTED // The requested operation cannot be performed because a prerequisite request message has not been received.
	_
	PT_STATUS_INVALID_PROVISIONING_STATE // Specified provisioning state is not valid.
	_
	PT_STATUS_INVALID_TIME               // Specified time is not valid.
	PT_STATUS_INVALID_INDEX              // Specified index is not valid.
	PT_STATUS_INVALID_PARAMETER          // Invalid input parameter.
	PT_STATUS_INVALID_NETMASK            // An invalid netmask was supplied (a valid netmask is an IP address in which all ‘1’s are before the ‘0’ – e.g. FFFC0000h is valid, FF0C0000h is invalid).
	PT_STATUS_FLASH_WRITE_LIMIT_EXCEEDED // The operation failed because the flash wear-out protection mechanism prevented a write to an NVRAM sector.
)

const (
	PT_STATUS_UNSUPPORTED_OEM_NUMBER   = iota + 2049 // The OEM number specified in the remote control command is not supported by the Intel AMT device.
	PT_STATUS_UNSUPPORTED_BOOT_OPTION                // The boot option specified in the remote control command is not supported by the Intel AMT device.
	PT_STATUS_INVALID_COMMAND                        // The command specified in the remote control command is not supported by the Intel AMT device.
	PT_STATUS_INVALID_SPECIAL_COMMAND                // The special command specified in the remote control command is not supported by the Intel AMT device.
	PT_STATUS_INVALID_HANDLE                         // The handle specified in the command is invalid.
	PT_STATUS_INVALID_PASSWORD                       // The password specified in the User ACL is invalid.
	PT_STATUS_INVALID_REALM                          // The realm specified in the User ACL is invalid.
	PT_STATUS_STORAGE_ACL_ENTRY_IN_USE               // The FPACL or EACL entry is used by an active registration and cannot be removed or modified.
	PT_STATUS_DATA_MISSING                           // Essential data is missing on CommitChanges() command.
	PT_STATUS_DUPLICATE                              // The parameter specified is a duplicate of an existing value.
	PT_STATUS_EVENTLOG_FROZEN                        // Event log is frozen.
	PT_STATUS_PKI_MISSING_KEYS                       // Reserved for future use.
	PT_STATUS_PKI_GENERATING_KEYS                    // Reserved for future use.
	PT_STATUS_INVALID_KEY                            // Invalid RSA Key.
	PT_STATUS_INVALID_CERT                           // Invalid X.509 Certificate or invalid certificate handle.
	PT_STATUS_CERT_KEY_NOT_MATCH                     // Key pair does not match.
	PT_STATUS_MAX_KERB_DOMAIN_REACHED                // The FW allows storing an SID from a limited number of domains. This SID domain does not exist and there is no space to store a new domain.
	PT_STATUS_UNSUPPORTED                            // Setting is not supported by this product.
	PT_STATUS_INVALID_PRIORITY                       // Priority setting is invalid.
	PT_STATUS_NOT_FOUND                              // Unable to find specified element.
	PT_STATUS_INVALID_CREDENTIALS                    // Invalid User credentials.
	PT_STATUS_INVALID_PASSPHRASE                     // Passphrase is invalid.
	_
	PT_STATUS_NO_ASSOCIATION // Current functionality requires association to a Key Pair.
	_
	_
	PT_STATUS_AUDIT_FAIL         // The command is defined in Audit Log policy as a critical event and cannot be logged.
	PT_STATUS_BLOCKING_COMPONENT // One of the ME components is not ready.
	_
	_
	_
	_
	PT_STATUS_USER_CONSENT_REQUIRED // User consent is required for this operation but was not received.
	PT_STATUS_OPERATION_IN_PROGRESS // Operation is not complete. This can occur when the Intel ME needs to generate a key pair (for example, when performing a full unprovision). Wait and then retry.
)
