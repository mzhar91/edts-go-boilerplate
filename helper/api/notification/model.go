package notification

import (
	_model "sg-edts.com/edts-go-boilerplate/model"
	"google.golang.org/genproto/googleapis/type/datetime"
)

type Recipient struct {
	From string   `json:"from"`
	To   []string `json:"to"`
	Cc   []string `json:"cc"`
	Bcc  []string `json:"bcc"`
}

type EmailRequest struct {
	ActionRefId int64  `json:"action_ref_id"`
	TypeRefId   int64  `json:"type_ref_id"`
	Data        string `json:"data"`
	Params      map[string]interface{}
	Recipient   Recipient
	RequestInfo _model.RequestInfo
}

type EmailResponse struct {
	Status      int               `json:"status"`
	Recipient   string            `json:"recipient"`
	DeliveredAt datetime.DateTime `json:"delivered_at"`
}
