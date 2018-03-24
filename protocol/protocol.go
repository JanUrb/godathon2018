package protocol

import "encoding/json"

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

func DecodeRegisterReqPayload(payload []byte) Register_req_payload {
	var result Register_req_payload
	json.Unmarshal(payload, &result)
	return result
}

func DecodeRegisterAckPayload(payload []byte) Register_ack_payload {
	var result Register_ack_payload
	json.Unmarshal(payload, &result)
	return result
}

func DecodeGroupAttachReq(payload []byte) Group_attach_req {
	var result Group_attach_req
	json.Unmarshal(payload, &result)
	return result
}

func DecodeGroupAttachAck(payload []byte) Group_attach_ack {
	var result Group_attach_ack
	json.Unmarshal(payload, &result)
	return result
}

func DecodeSetupReq(payload []byte) Setup_req {
	var result Setup_req
	json.Unmarshal(payload, &result)
	return result
}

func DecodeSetupAck(payload []byte) Setup_ack {
	var result Setup_ack
	json.Unmarshal(payload, &result)
	return result
}

func DecodeSetupInd(payload []byte) Setup_ind {
	var result Setup_ind
	json.Unmarshal(payload, &result)
	return result
}

func DecodeSetupRes(payload []byte) Setup_res {
	var result Setup_res
	json.Unmarshal(payload, &result)
	return result
}

func DecodeConnectReq(payload []byte) Connect_req {
	var result Connect_req
	json.Unmarshal(payload, &result)
	return result
}

func DecodeConnectAck(payload []byte) Connect_ack {
	var result Connect_ack
	json.Unmarshal(payload, &result)
	return result
}

func DecodeTxDemandReq(payload []byte) Tx_demand_req {
	var result Tx_demand_req
	json.Unmarshal(payload, &result)
	return result
}

func DecodeTxDemandAck(payload []byte) Tx_demand_ack {
	var result Tx_demand_ack
	json.Unmarshal(payload, &result)
	return result
}

func DecodeTxCeasedReq(payload []byte) Tx_ceased_req {
	var result Tx_ceased_req
	json.Unmarshal(payload, &result)
	return result
}

func DecodeTxCeasedAck(payload []byte) Tx_ceased_ack {
	var result Tx_ceased_ack
	json.Unmarshal(payload, &result)
	return result
}

func DecodeTxInfoInd(payload []byte) Tx_info_ind {
	var result Tx_info_ind
	json.Unmarshal(payload, &result)
	return result
}
