package auditlog

import (
	"encoding/xml"

	"github.com/open-amt-cloud-toolkit/go-wsman-messages/internal/wsman"
	"github.com/open-amt-cloud-toolkit/go-wsman-messages/pkg/amt/actions"
)

const AMT_AuditLog = "AMT_AuditLog"

type AuditLog struct {
	base wsman.Base
}

type ReadRecords_INPUT struct {
	XMLName    xml.Name `xml:"h:ReadRecords_INPUT"`
	H          string   `xml:"xmlns:h,attr"`
	StartIndex int      `xml:"h:StartIndex" json:"StartIndex"`
}

func NewAuditLog(wsmanMessageCreator *wsman.WSManMessageCreator) AuditLog {
	return AuditLog{base: wsman.NewBase(wsmanMessageCreator, AMT_AuditLog)}
}

func (AuditLog AuditLog) Get() string {
	return AuditLog.base.Get(nil)
}
func (AuditLog AuditLog) Enumerate() string {
	return AuditLog.base.Enumerate()
}
func (AuditLog AuditLog) Pull(enumerationContext string) string {
	return AuditLog.base.Pull(enumerationContext)
}
func (a AuditLog) ReadRecords(startIndex int) string {
	if startIndex < 1 {
		startIndex = 0
	}
	header := a.base.WSManMessageCreator.CreateHeader(string(actions.ReadRecords), AMT_AuditLog, nil, "", "")
	body := a.base.WSManMessageCreator.CreateBody("ReadRecords_INPUT", AMT_AuditLog, &ReadRecords_INPUT{StartIndex: startIndex})

	return a.base.WSManMessageCreator.CreateXML(header, body)
}
