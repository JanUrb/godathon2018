let register_req = {
  type    : 'register_req',
  payload : {
    user  : 1234,
    token : 'secret_token'
  }
};

let register_ack = {
  type    : 'register_ack',
  payload : {
    result : 200
  }
};

let groupAttach_req = {
  type    : 'groupAttach_req',
  payload : {
    id : 9876
  }
};


let groupAttach_ack = {
  type    : 'groupAttach_ack',
  payload : {
    result : 200
  }
};


let setup_req = {
  type    : 'setup_req',
  payload : {
    callType : 'p2m',
    calledId : 1234
  }
};

let setup_ack = {
  type    : 'setup_ack',
  payload : {
    result : 200,
    callId : 546
  }
};

let setup_ind = {
  type    : 'setup_ind',
  payload : {
    callType  : 'p2m',
    calledId  : 1234,
    callingId : 345,
    callId    : 546
  }
};

let setup_res = {
  type    : 'setup_res',
  payload : {
    result : 200,
    callId : 546
  }
};

let connect_req = {
  type    : 'connect_req',
  payload : {
    callId : 546,
  }
};

let connect_ack = {
  type    : 'connect_ack',
  payload : {
    result : 200,
    callId : 546,
    talkingPartyId: 1234,
  }
};

let txDemand_req = {
  type    : 'txDemand_req',
  payload : {
    callId : 546,
  }
};

let txDemand_ack = {
  type    : 'txDemand_ack',
  payload : {
    result : 200,
    callId : 546,
  }
};

let txCeased_req = {
  type    : 'txCeased_req',
  payload : {
    callId : 546,
  }
};

let txCeased_ack = {
  type    : 'txCeased_ack',
  payload : {
    result: 200,
    callId : 546,
  }
};


let txInfo_ind = {
  type: 'txInfo_ind',
  payload: {
    callId: 546,
    talkingPartyId: 1234,
  }
};
