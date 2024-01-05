/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package timesynchronization

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/message"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_TimeSynchronizationService = "AMT_TimeSynchronizationService"

type (
	Response struct {
		XMLName xml.Name       `xml:"Envelope"`
		Header  message.Header `xml:"Header"`
		Body    Body           `xml:"Body"`
	}
	Body struct {
		XMLName                         xml.Name `xml:"Body"`
		GetLowAccuracyTimeSynch_OUTPUT  GetLowAccuracyTimeSynch_OUTPUT
		SetHighAccuracyTimeSynch_OUTPUT message.ReturnValue
	}
	GetLowAccuracyTimeSynch_OUTPUT struct {
		Ta0         int64
		ReturnValue int
	}
)

type Service struct {
	base message.Base
}
type SetHighAccuracyTimeSynch_INPUT struct {
	XMLName xml.Name `xml:"h:SetHighAccuracyTimeSynch_INPUT"`
	H       string   `xml:"xmlns:h,attr"`
	Ta0     int64    `xml:"h:Ta0"`
	Tm1     int64    `xml:"h:Tm1"`
	Tm2     int64    `xml:"h:Tm2"`
}

func NewTimeSynchronizationService(wsmanMessageCreator *message.WSManMessageCreator) Service {
	return Service{
		base: message.NewBase(wsmanMessageCreator, AMT_TimeSynchronizationService),
	}
}

// Get retrieves the representation of the instance
func (s Service) Get() string {
	return s.base.Get(nil)
}

// Enumerates the instances of this class
func (s Service) Enumerate() string {
	return s.base.Enumerate()
}

// Pulls instances of this class, following an Enumerate operation
func (s Service) Pull(enumerationContext string) string {
	return s.base.Pull(enumerationContext)
}
func (s Service) SetHighAccuracyTimeSynch(ta0, tm1, tm2 int64) string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.SetHighAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody("SetHighAccuracyTimeSynch_INPUT", AMT_TimeSynchronizationService, &SetHighAccuracyTimeSynch_INPUT{
		Ta0: ta0,
		Tm1: tm1,
		Tm2: tm2,
	})
	return s.base.WSManMessageCreator.CreateXML(header, body)
}

func (s Service) GetLowAccuracyTimeSynch() string {
	header := s.base.WSManMessageCreator.CreateHeader(string(actions.GetLowAccuracyTimeSynch), AMT_TimeSynchronizationService, nil, "", "")
	body := s.base.WSManMessageCreator.CreateBody("GetLowAccuracyTimeSynch_INPUT", AMT_TimeSynchronizationService, nil)
	return s.base.WSManMessageCreator.CreateXML(header, body)
}
