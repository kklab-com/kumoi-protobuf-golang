package omega

import (
	"encoding/json"
	"time"
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

func (x *TransitFrame) SetData(data isTransitFrame_Data) *TransitFrame {
	x.Data = data
	return x
}

func (x *TransitFrame) RenewTimestamp() *TransitFrame {
	x.Timestamp = time.Now().UnixNano()
	return x
}

func (x *TransitFrame) AsRequest() *TransitFrame {
	x.Class = TransitFrame_ClassRequest
	return x
}

func (x *TransitFrame) AsResponse() *TransitFrame {
	x.Class = TransitFrame_ClassResponse
	return x
}

func (x *TransitFrame) AsRequestReplay() *TransitFrame {
	x.Class = TransitFrame_ClassRequestReplay
	return x
}

func (x *TransitFrame) AsResponseReplay() *TransitFrame {
	x.Class = TransitFrame_ClassResponseReplay
	return x
}

func (x *TransitFrame) AsNotification() *TransitFrame {
	x.Class = TransitFrame_ClassNotification
	return x
}

func (x *TransitFrame) AsNotificationReplay() *TransitFrame {
	x.Class = TransitFrame_ClassNotificationReplay
	return x
}

func (x *TransitFrame) AsError() *TransitFrame {
	x.Class = TransitFrame_ClassError
	return x
}
