package protocol

import "encoding/json"

const (
	// register_req    = "register_req"
	// register_ack    = "register_ack"
	MessageType_groupAttach_req = "groupAttach_req"
	MessageType_groupAttach_ack = "groupAttach_ack"
	MessageType_setup_req       = "setup_req"
	MessageType_setup_ack       = "setup_ack"
	MessageType_setup_ind       = "setup_ind"
	MessageType_setup_res       = "setup_res"
	MessageType_connect_req     = "connect_req"
	MessageType_connect_ack     = "connect_ack"
	MessageType_txDemand_req    = "txDemand_req"
	MessageType_txDemand_ack    = "txDemand_ack"
	MessageType_txCeased_req    = "txCeased_req"
	MessageType_txCeased_ack    = "txCeased_ack"
	MessageType_txInfo_ind      = "txInfo_ind"
)

type Generic_message struct {
	Msg_type string          `json:"type"`
	Payload  json.RawMessage `json:"payload"`
}

type Register_req_payload struct {
	User  int    `json:"user"`
	Token string `json:"token"`
}

type Register_ack_payload struct {
	Result int `json:"result"`
}

type Group_attach_req struct {
	Id int `json:"id"`
}

type Group_attach_ack struct {
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

type Connect_req struct {
	Call_id int `json:"callId"`
}

type Connect_ack struct {
	Result           int `json:"result"`
	Call_id          int `json:"callId"`
	Talking_party_id int `json:"talkingPartyId"`
}

type Tx_demand_req struct {
	Call_id int `json:"callId"`
}

type Tx_demand_ack struct {
	Call_id int `json:"callId"`
}

type Tx_ceased_req struct {
	Call_id int `json:"callId"`
}

type Tx_ceased_ack struct {
	Result  int `json:"result"`
	Call_id int `json:"callId"`
}

type Tx_info_ind struct {
	Call_id          int `json:"callId"`
	Talking_party_id int `json:"talkingPartyId"`
}

func EncodeRegisterReqPayload(data Register_req_payload) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "register_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeRegisterAckPayload(data Register_ack_payload) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "register_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeGroupAttachReq(data Group_attach_req) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "groupAttach_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeGroupAttachAck(data Group_attach_ack) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "groupAttach_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeSetupReq(data Setup_req) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "setup_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeSetupAck(data Setup_ack) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "setup_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeSetupInd(data Setup_ind) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "setup_ind"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeSetupRes(data Setup_res) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "setup_res"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeConnectReq(data Connect_req) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "connect_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeConnectAck(data Connect_ack) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "connect_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeTxDemandReq(data Tx_demand_req) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "txDemand_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeTxDemandAck(data Tx_demand_ack) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "txDemand_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeTxCeasedReq(data Tx_ceased_req) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "txCeased_req"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeTxCeasedAck(data Tx_ceased_ack) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "txCeased_ack"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func EncodeTxInfoInd(data Tx_info_ind) ([]byte, error, error) {
	var msgStruct Generic_message
	msgStruct.Msg_type = "txInfo_ind"
	payload, err1 := json.Marshal(data)
	msgStruct.Payload = payload
	msg, err2 := json.Marshal(msgStruct)
	return msg, err1, err2
}

func DecodeRegisterReqPayload(payload []byte) (Register_req_payload, error) {
	var result Register_req_payload
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeRegisterAckPayload(payload []byte) (Register_ack_payload, error) {
	var result Register_ack_payload
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

func DecodeConnectReq(payload []byte) (Connect_req, error) {
	var result Connect_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeConnectAck(payload []byte) (Connect_ack, error) {
	var result Connect_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeTxDemandReq(payload []byte) (Tx_demand_req, error) {
	var result Tx_demand_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeTxDemandAck(payload []byte) (Tx_demand_ack, error) {
	var result Tx_demand_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeTxCeasedReq(payload []byte) (Tx_ceased_req, error) {
	var result Tx_ceased_req
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeTxCeasedAck(payload []byte) (Tx_ceased_ack, error) {
	var result Tx_ceased_ack
	err := json.Unmarshal(payload, &result)
	return result, err
}

func DecodeTxInfoInd(payload []byte) (Tx_info_ind, error) {
	var result Tx_info_ind
	err := json.Unmarshal(payload, &result)
	return result, err
}
