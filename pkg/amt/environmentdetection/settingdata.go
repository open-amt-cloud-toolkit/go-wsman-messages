package environmentdetection

import (
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/cim/models"
)

const AMT_EnvironmentDetectionSettingData = "AMT_EnvironmentDetectionSettingData"

type EnvironmentDetectionSettingData struct {
	models.SettingData
	DetectionAlgorithm         DetectionAlgorithm
	DetectionStrings           []string
	DetectionIPv6LocalPrefixes []string
}

type DetectionAlgorithm uint8

const (
	LocalDomains DetectionAlgorithm = iota
	RemoteURLs
)

type SettingData struct {
	base wsman.Base
}

func NewEnvironmentDetectionSettingData(wsmanMessageCreator *wsman.WSManMessageCreator) SettingData {
	return SettingData{
		base: wsman.NewBase(wsmanMessageCreator, AMT_EnvironmentDetectionSettingData),
	}
}
func (EnvironmentDetectionSettingData SettingData) Get() string {
	return EnvironmentDetectionSettingData.base.Get(nil)
}
func (EnvironmentDetectionSettingData SettingData) Enumerate() string {
	return EnvironmentDetectionSettingData.base.Enumerate()
}
func (EnvironmentDetectionSettingData SettingData) Pull(enumerationContext string) string {
	return EnvironmentDetectionSettingData.base.Pull(enumerationContext)
}
func (EnvironmentDetectionSettingData SettingData) Put(environmentDetectionSettingData EnvironmentDetectionSettingData) string {
	return EnvironmentDetectionSettingData.base.Put(environmentDetectionSettingData, false, nil)
}
