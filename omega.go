package omega

import (
	"encoding/json"
)

type TransitFrameData = isTransitFrame_Data

func (x *Error) Error() string {
	bs, _ := json.Marshal(x)
	return string(bs)
}

func (x *TransitFrame) Clone() *TransitFrame {
	return &TransitFrame{
		TransitId:        x.GetTransitId(),
		Class:            x.GetClass(),
		Version:          x.GetVersion(),
		Timestamp:        x.GetTimestamp(),
		Error:            x.GetError(),
		MessageId:        x.GetMessageId(),
		RefererMessageId: x.GetRefererMessageId(),
		Data:             x.GetData(),
	}
}
