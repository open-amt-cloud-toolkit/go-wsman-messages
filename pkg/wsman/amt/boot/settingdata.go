/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package boot

import (
	"encoding/xml"
	"fmt"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

// Instantiates a new Boot Setting Data service.
func NewBootSettingDataWithClient(wsmanMessageCreator *message.WSManMessageCreator, client client.WSMan) SettingData {
	return SettingData{
		base: message.NewBaseWithClient(wsmanMessageCreator, AMTBootSettingData, client),
	}
}

// Get retrieves the representation of the instance.
func (settingData SettingData) Get() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Get(nil),
		},
	}

	// send the message to AMT
	err = settingData.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Enumerate returns an enumeration context which is used in a subsequent Pull call.
func (settingData SettingData) Enumerate() (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Enumerate(),
		},
	}

	// send the message to AMT
	err = settingData.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

// Pull returns the instances of this class.  An enumeration context provided by the Enumerate call is used as input.
func (settingData SettingData) Pull(enumerationContext string) (response Response, err error) {
	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.Pull(enumerationContext),
		},
	}

	// send the message to AMT
	err = settingData.base.Execute(response.Message)
	if err != nil {
		return
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return
	}

	return
}

func (settingData SettingData) Put(bootSettingData BootSettingDataRequest) (response Response, err error) {
	header := settingData.base.WSManMessageCreator.CreateHeader(message.BaseActionsPut, AMTBootSettingData, nil, "", "")
	body := fmt.Sprintf(
		`<Body><h:AMT_BootSettingData xmlns:h="%sAMT_BootSettingData"><h:BIOSPause>%t</h:BIOSPause><h:BIOSSetup>%t</h:BIOSSetup><h:BootMediaIndex>%d</h:BootMediaIndex><h:ConfigurationDataReset>%t</h:ConfigurationDataReset><h:ElementName>%s</h:ElementName><h:EnforceSecureBoot>%t</h:EnforceSecureBoot><h:FirmwareVerbosity>%d</h:FirmwareVerbosity><h:ForcedProgressEvents>%t</h:ForcedProgressEvents><h:IDERBootDevice>%d</h:IDERBootDevice><h:InstanceID>%s</h:InstanceID><h:LockKeyboard>%t</h:LockKeyboard><h:LockPowerButton>%t</h:LockPowerButton><h:LockResetButton>%t</h:LockResetButton><h:LockSleepButton>%t</h:LockSleepButton><h:OwningEntity>%s</h:OwningEntity><h:PlatformErase>%t</h:PlatformErase><h:RSEPassword>%s</h:RSEPassword><h:ReflashBIOS>%t</h:ReflashBIOS><h:SecureErase>%t</h:SecureErase><h:UefiBootParametersArray>%s</h:UefiBootParametersArray><h:UefiBootNumberOfParams>%d</h:UefiBootNumberOfParams><h:UseIDER>%t</h:UseIDER><h:UseSOL>%t</h:UseSOL><h:UseSafeMode>%t</h:UseSafeMode><h:UserPasswordBypass>%t</h:UserPasswordBypass></h:AMT_BootSettingData></Body>`,
		settingData.base.WSManMessageCreator.ResourceURIBase,
		bootSettingData.BIOSPause,
		bootSettingData.BIOSSetup,
		bootSettingData.BootMediaIndex,
		bootSettingData.ConfigurationDataReset,
		bootSettingData.ElementName,
		bootSettingData.EnforceSecureBoot,
		bootSettingData.FirmwareVerbosity,
		bootSettingData.ForcedProgressEvents,
		bootSettingData.IDERBootDevice,
		bootSettingData.InstanceID,
		bootSettingData.LockKeyboard,
		bootSettingData.LockPowerButton,
		bootSettingData.LockResetButton,
		bootSettingData.LockSleepButton,
		bootSettingData.OwningEntity,
		bootSettingData.PlatformErase,
		bootSettingData.RSEPassword,
		bootSettingData.ReflashBIOS,
		bootSettingData.SecureErase,
		bootSettingData.UefiBootParametersArray,
		bootSettingData.UefiBootNumberOfParams,
		bootSettingData.UseIDER,
		bootSettingData.UseSOL,
		bootSettingData.UseSafeMode,
		bootSettingData.UserPasswordBypass)

	response = Response{
		Message: &client.Message{
			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
		},
	}

	// send the message to AMT
	err = settingData.base.Execute(response.Message)
	if err != nil {
		return response, err
	}

	// put the xml response into the go struct
	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
	if err != nil {
		return response, err
	}

	return response, err
}


// Put will change properties of the selected instance.
// func (settingData SettingData) Put(bootSettingData BootSettingDataRequest) (response Response, err error) {
// 	header := settingData.base.WSManMessageCreator.CreateHeader(message.BaseActionsPut, AMTBootSettingData, nil, "", "")
// 	body := fmt.Sprintf(
// 		`<Body><h:AMT_BootSettingData xmlns:h="%sAMT_BootSettingData"><h:BIOSLastStatus>%d</h:BIOSLastStatus><h:BIOSLastStatus>%d</h:BIOSLastStatus><h:BIOSPause>%t</h:BIOSPause><h:BIOSSetup>%t</h:BIOSSetup><h:BootMediaIndex>%d</h:BootMediaIndex><h:BootguardStatus>%d</h:BootguardStatus><h:ConfigurationDataReset>%t</h:ConfigurationDataReset><h:ElementName>%s</h:ElementName><h:EnforceSecureBoot>%t</h:EnforceSecureBoot><h:FirmwareVerbosity>%d</h:FirmwareVerbosity><h:ForcedProgressEvents>%t</h:ForcedProgressEvents><h:IDERBootDevice>%d</h:IDERBootDevice><h:InstanceID>%s</h:InstanceID><h:LockKeyboard>%t</h:LockKeyboard><h:LockPowerButton>%t</h:LockPowerButton><h:LockResetButton>%t</h:LockResetButton><h:LockSleepButton>%t</h:LockSleepButton><h:OptionsCleared>%t</h:OptionsCleared><h:OwningEntity>%s</h:OwningEntity><h:PlatformErase>%t</h:PlatformErase><h:RPEEnabled>%t</h:RPEEnabled><h:RSEPassword>%s</h:RSEPassword><h:ReflashBIOS>%t</h:ReflashBIOS><h:SecureBootControlEnabled>%t</h:SecureBootControlEnabled><h:SecureErase>%t</h:SecureErase><h:UEFIHTTPSBootEnabled>%t</h:UEFIHTTPSBootEnabled><h:UefiBootParametersArray>%s</h:UefiBootParametersArray><h:UEFILocalPBABootEnabled>%t</h:UEFILocalPBABootEnabled><h:UefiBootNumberOfParams>%d</h:UefiBootNumberOfParams><h:UseIDER>%t</h:UseIDER><h:UseSOL>%t</h:UseSOL><h:UseSafeMode>%t</h:UseSafeMode><h:UserPasswordBypass>%t</h:UserPasswordBypass><h:WinREBootEnabled>%t</h:WinREBootEnabled></h:AMT_BootSettingData></Body>`,
// 		settingData.base.WSManMessageCreator.ResourceURIBase,
// 		bootSettingData.BIOSLastStatus[0],
// 		bootSettingData.BIOSLastStatus[1],
// 		bootSettingData.BIOSPause,
// 		bootSettingData.BIOSSetup,
// 		bootSettingData.BootMediaIndex,
// 		bootSettingData.BootguardStatus,
// 		bootSettingData.ConfigurationDataReset,
// 		bootSettingData.ElementName,
// 		bootSettingData.EnforceSecureBoot,
// 		bootSettingData.FirmwareVerbosity,
// 		bootSettingData.ForcedProgressEvents,
// 		bootSettingData.IDERBootDevice,
// 		bootSettingData.InstanceID,
// 		bootSettingData.LockKeyboard,
// 		bootSettingData.LockPowerButton,
// 		bootSettingData.LockResetButton,
// 		bootSettingData.LockSleepButton,
// 		bootSettingData.OptionsCleared,
// 		bootSettingData.OwningEntity,
// 		bootSettingData.PlatformErase,
// 		bootSettingData.RPEEnabled,
// 		bootSettingData.RSEPassword,
// 		bootSettingData.ReflashBIOS,
// 		bootSettingData.SecureBootControlEnabled,
// 		bootSettingData.SecureErase,
// 		bootSettingData.UEFIHTTPSBootEnabled,
// 		bootSettingData.UefiBootParametersArray,
// 		bootSettingData.UEFILocalPBABootEnabled,
// 		bootSettingData.UefiBootNumberOfParams,
// 		bootSettingData.UseIDER,
// 		bootSettingData.UseSOL,
// 		bootSettingData.UseSafeMode,
// 		bootSettingData.UserPasswordBypass,
// 		bootSettingData.WinREBootEnabled)

// 	response = Response{
// 		Message: &client.Message{
// 			XMLInput: settingData.base.WSManMessageCreator.CreateXML(header, body),
// 		},
// 	}

// 	// send the message to AMT
// 	err = settingData.base.Execute(response.Message)
// 	if err != nil {
// 		return response, err
// 	}

// 	// put the xml response into the go struct
// 	err = xml.Unmarshal([]byte(response.XMLOutput), &response)
// 	if err != nil {
// 		return response, err
// 	}

// 	return response, err
// }
