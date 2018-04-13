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
    groupId: 1234
  }
};

let setup_ack = {
  type: 'setup_ack',
  payload: {
    result: 200,
    groupId: 123
  }
};

let setup_ind = {
  type: 'setup_ind',
  payload: {
    callingId: 345,
    groupId: 123
  }
};

let disconnect_req = {
  type: 'disconnect_req',
  groupId: 1243
};

let disconnect_ack = {
  type: 'disconnect_ack',
  groupId: 1234,
  result: 200
};

let disconnect_ind = {
  type: 'disconnect_ind',
  groupId: 1224
};
