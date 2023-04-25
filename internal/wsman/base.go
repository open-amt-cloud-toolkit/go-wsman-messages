package wsman

import "fmt"

type Base struct {
	WSManMessageCreator *WSManMessageCreator
	className           string
}

func NewBase(wsmanMessageCreator *WSManMessageCreator, className string) Base {
	return Base{
		WSManMessageCreator: wsmanMessageCreator,
		className:           className,
	}
}

func (b *Base) Enumerate() string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsEnumerate, b.className, nil, "", "")
	return b.WSManMessageCreator.CreateXML(header, EnumerateBody)
}

func (b *Base) Get(selector *Selector) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsGet, b.className, selector, "", "")
	return b.WSManMessageCreator.CreateXML(header, GetBody)
}

func (b *Base) Pull(enumerationContext string) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsPull, b.className, nil, "", "")
	body := createCommonBodyPull(enumerationContext, 0, 0)
	return b.WSManMessageCreator.CreateXML(header, body)
}
func (b *Base) Delete(selector *Selector) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsDelete, b.className, selector, "", "")
	return b.WSManMessageCreator.CreateXML(header, DeleteBody)
}

func (b *Base) Put(data interface{}, useHeaderSelector bool, customSelector *Selector) string {
	if customSelector == nil {
		customSelector = &Selector{Name: "InstanceID", Value: fmt.Sprintf("%v", data)}
	}
	var header string
	if useHeaderSelector {
		header = b.WSManMessageCreator.CreateHeader(BaseActionsPut, b.className, customSelector, "", "")
	} else {
		header = b.WSManMessageCreator.CreateHeader(BaseActionsPut, b.className, nil, "", "")
	}
	body := b.WSManMessageCreator.createCommonBodyCreateOrPut(b.className, data)
	return b.WSManMessageCreator.CreateXML(header, body)
}

func (b *Base) Create(data interface{}, selector *Selector) string {
	header := b.WSManMessageCreator.CreateHeader(BaseActionsCreate, b.className, selector, "", "")
	body := b.WSManMessageCreator.createCommonBodyCreateOrPut(b.className, data)
	return b.WSManMessageCreator.CreateXML(header, body)
}

func (b *Base) RequestStateChange(actionName string, requestedState int) string {

	header := b.WSManMessageCreator.CreateHeader(actionName, b.className, nil, "", "")
	body := createCommonBodyRequestStateChange(fmt.Sprintf("%s%s", b.WSManMessageCreator.ResourceURIBase, b.className), requestedState)
	return b.WSManMessageCreator.CreateXML(header, body)
}
