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
