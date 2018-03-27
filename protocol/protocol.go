package protocol

import "encoding/json"

const (
	MessageType_register_req    = "register_req"
	MessageType_register_ack    = "register_ack"
	MessageType_groupAttach_req = "groupAttach_req"
	MessageType_groupAttach_ack = "groupAttach_ack"
	MessageType_setup_req       = "setup_req"
	MessageType_setup_ack       = "setup_ack"
	MessageType_setup_ind       = "setup_ind"
	MessageType_setup_res       = "setup_res"
	MessageType_disconnect_req  = "disconnect_req"
	MessageType_disconnect_ack  = "disconnect_ack"
	MessageType_disconnect_ind  = "disconnect_ind"
)

type Generic_message struct {
	Msg_type string          `json:"type"`
	Payload  json.RawMessage `json:"payload"`
}

type Register_req struct {
	User string `json:"user"`
}

type Register_ack struct {
	Result int `json:"result"`
}

type Group_attach_req struct {
	ID int `json:"id"`
}

type Group_attach_ack struct {
	Id     int `json:"id"`
	Result int `json:"result"`
}

type Setup_req struct {
	Call_type string `json:"callType"`
	Called_id int    `json:"calledId"`
}

type Setup_ack struct {
	Result  int `json:"result"`
	Call_id int `json:"callId"`
}

type Setup_ind struct {
	Call_type  string `json:"callType"`
	Called_id  int    `json:"calledId"`
	Calling_id int    `json:"callingId"`
	Call_id    int    `json:"callId"`
}

type Setup_res struct {
	Result  int `json:"result"`
	Call_id int `json:"callId"`
}

type Disconnect_req struct {
	Call_id int `json:"callId"`
}

type Disconnect_ack struct {
	Result  int `json:"result"`
	Call_id int `json:"callId"`
}

type Disconnect_ind struct {
	Call_id int `json:"callId"`
}

func EncodeRegisterReq(data Register_req) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_register_req
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeRegisterAck(data Register_ack) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_register_ack
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeGroupAttachReq(data Group_attach_req) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_groupAttach_req
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeGroupAttachAck(data Group_attach_ack) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_groupAttach_ack
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupReq(data Setup_req) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_setup_req
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupAck(data Setup_ack) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_setup_ack
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupInd(data Setup_ind) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_setup_ind
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeSetupRes(data Setup_res) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_setup_res
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectReq(data Disconnect_req) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_disconnect_req
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectAck(data Disconnect_ack) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_disconnect_ack
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)
}

func EncodeDisconnectInd(data Disconnect_ind) ([]byte, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = MessageType_disconnect_ind
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	msgStruct.Payload = payload
	return json.Marshal(msgStruct)

}

func DecodeRegisterReq(payload []byte) (Register_req, error) {
	var result Register_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeRegisterAck(payload []byte) (Register_ack, error) {
	var result Register_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeGroupAttachReq(payload []byte) (Group_attach_req, error) {
	var result Group_attach_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeGroupAttachAck(payload []byte) (Group_attach_ack, error) {
	var result Group_attach_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupReq(payload []byte) (Setup_req, error) {
	var result Setup_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupAck(payload []byte) (Setup_ack, error) {
	var result Setup_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupInd(payload []byte) (Setup_ind, error) {
	var result Setup_ind
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeSetupRes(payload []byte) (Setup_res, error) {
	var result Setup_res
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectReq(payload []byte) (Disconnect_req, error) {
	var result Disconnect_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectAck(payload []byte) (Disconnect_ack, error) {
	var result Disconnect_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeDisconnectInd(payload []byte) (Disconnect_ind, error) {
	var result Disconnect_ind
	err := json.Unmarshal(payload, &result)
	return result, err
}
