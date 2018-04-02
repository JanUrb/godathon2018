package protocol

import "encoding/json"

const (
	MessageTypeReqisterReq    = "register_req"
	MessageTypeReqisterAck    = "register_ack"
	MessageTypeGroupAttachReq = "groupAttach_req"
	MessageTypeGroupAttachAck = "groupAttach_ack"
	MessageTypeSetupReq       = "setup_req"
	MessageTypeSetupAck       = "setup_ack"
	MessageTypeSetupInd       = "setup_ind"
	MessageTypeSetupRes       = "setup_res"
	MessageTypeDisconnectReq  = "disconnect_req"
	MessageTypeDisconnectAck  = "disconnect_ack"
	MessageTypeDisconnectInd  = "disconnect_ind"
)

type GenericMessage struct {
	Msg_type string          `json:"type"`
	Payload  json.RawMessage `json:"payload"`
}

type RegisterReq struct {
	User string `json:"user"`
}

type RegisterAck struct {
	Result int `json:"result"`
}

type GroupAttachReq struct {
	GroupID int `json:"groupId"`
}

type GroupAttachAck struct {
	GroupID int `json:"groupId"`
	Result  int `json:"result"`
}

type SetupReq struct {
	GroupID int `json:"groupId"`
}

type SetupAck struct {
	Result  int `json:"result"`
	GroupID int `json:"groupId"`
}

type SetupInd struct {
	CallingID int `json:"callingId"`
	GroupID   int `json:"groupId"`
}

type SetupRes struct {
	Result  int `json:"result"`
	GroupID int `json:"callId"`
}

type DisconnectReq struct {
	GroupID int `json:"callId"`
}

type DisconnectAck struct {
	Result  int `json:"result"`
	GroupID int `json:"callId"`
}

type DisconnectInd struct {
	GroupID int `json:"callId"`
}

func EncodeRegisterReq(data RegisterReq) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeReqisterReq
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeRegisterAck(data RegisterAck) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeReqisterAck
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeGroupAttachReq(data GroupAttachReq) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeGroupAttachReq
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeGroupAttachAck(data GroupAttachAck) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeGroupAttachAck
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupReq(data SetupReq) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeSetupReq
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupAck(data SetupAck) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeSetupAck
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupInd(data SetupInd) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeSetupInd
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupRes(data SetupRes) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeSetupRes
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectReq(data DisconnectReq) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeDisconnectReq
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectAck(data DisconnectAck) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeDisconnectAck
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectInd(data DisconnectInd) ([]byte, error) {
	var msgStruct GenericMessage
	msgStruct.Msg_type = MessageTypeDisconnectInd
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)

}

func DecodeRegisterReq(payload []byte) (RegisterReq, error) {
	var result RegisterReq
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeRegisterAck(payload []byte) (RegisterAck, error) {
	var result RegisterAck
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeGroupAttachReq(payload []byte) (GroupAttachReq, error) {
	var result GroupAttachReq
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeGroupAttachAck(payload []byte) (GroupAttachAck, error) {
	var result GroupAttachAck
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupReq(payload []byte) (SetupReq, error) {
	var result SetupReq
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupAck(payload []byte) (SetupAck, error) {
	var result SetupAck
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupInd(payload []byte) (SetupInd, error) {
	var result SetupInd
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupRes(payload []byte) (SetupRes, error) {
	var result SetupRes
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectReq(payload []byte) (DisconnectReq, error) {
	var result DisconnectReq
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectAck(payload []byte) (DisconnectAck, error) {
	var result DisconnectAck
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectInd(payload []byte) (DisconnectInd, error) {
	var result DisconnectInd
	err := json.Unmarshal(payload, &result)
	return result, err
}
