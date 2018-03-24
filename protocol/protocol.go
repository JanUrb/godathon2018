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
