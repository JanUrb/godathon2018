let register_req = {
  type: 'register_req',
  payload: {
    user: 'name'
  }
};

let register_ack = {
  type: 'register_ack',
  payload: {
    result: 200
  }
};

let groupAttach_req = {
  type: 'groupAttach_req',
  payload: {
    id: '9876'
  }
};

let groupAttach_ack = {
  type: 'groupAttach_ack',
  payload: {
    id: 9876,
    result: 200
  }
};

let setup_req = {
  type: 'setup_req',
  payload: {
    calledId: 1234
  }
};

let setup_ack = {
  type: 'setup_ack',
  payload: {
    result: 200,
    callId: 546
  }
};

let setup_ind = {
  type: 'setup_ind',
  payload: {
    callType: 'p2m',
    calledId: 1234,
    callingId: 345,
    callId: 546
  }
};

let setup_res = {
  type: 'setup_res',
  payload: {
    result: 200,
    callId: 546
  }
};

let disconnect_req = {
  type: '',
  callId: 1243
};

let disconnect_ack = {
  type: '',
  callId: 1234,
  result: 200
};

let disconnect_ind = {
  type: '',
  callId: 1224
};
