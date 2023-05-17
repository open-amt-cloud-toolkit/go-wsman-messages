/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_TimeSynchronizationService = "AMT_TimeSynchronizationService"

type Service struct {
	base wsman.Base
}
type SetHighAccuracyTimeSynch_INPUT struct {
	XMLName xml.Name `xml:"h:SetHighAccuracyTimeSynch_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Ta0     int64    `xml:"h:Ta0"`
	Tm1     int64    `xml:"h:Tm1"`
	Tm2     int64    `xml:"h:Tm2"`
}

func NewTimeSynchronizationService(wsmanMessageCreator *wsman.WSManMessageCreator) Service {
	return Service{
		base: wsman.NewBase(wsmanMessageCreator, AMT_TimeSynchronizationService),
	}
}

// Get retrieves the representation of the instance
func (TimeSynchronizationService Service) Get() string {
	return TimeSynchronizationService.base.Get(nil)
}

// Enumerates the instances of this class
func (TimeSynchronizationService Service) Enumerate() string {
	return TimeSynchronizationService.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (TimeSynchronizationService Service) Pull(enumerationContext string) string {
	return TimeSynchronizationService.base.Pull(enumerationContext)
}
func (TLSCredentialContext Service) SetHighAccuracyTimeSynch(ta0, tm1, tm2 int64) string {
	header := TLSCredentialContext.base.WSManMessageCreator.CreateHeader(string(actions.SetHighAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := TLSCredentialContext.base.WSManMessageCreator.CreateBody("SetHighAccuracyTimeSynch_INPUT", AMT_TimeSynchronizationService, &SetHighAccuracyTimeSynch_INPUT{
		Ta0: ta0,
		Tm1: tm1,
		Tm2: tm2,
	})
	return TLSCredentialContext.base.WSManMessageCreator.CreateXML(header, body)
}

func (TLSCredentialContext Service) GetLowAccuracyTimeSynch() string {
	header := TLSCredentialContext.base.WSManMessageCreator.CreateHeader(string(actions.GetLowAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := TLSCredentialContext.base.WSManMessageCreator.CreateBody("GetLowAccuracyTimeSynch_INPUT", AMT_TimeSynchronizationService, nil)
	return TLSCredentialContext.base.WSManMessageCreator.CreateXML(header, body)
}
