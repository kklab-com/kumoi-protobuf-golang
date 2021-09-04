package omega

import "encoding/json"

type TransitFrameData = isTransitFrame_Data

func (x *Error) error() string {
	bs, _ := json.Marshal(x)
	return string(bs)
}
