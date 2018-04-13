const ADDRESS = 'ws://localhost:4242/ws';

class Client {
  constructor(statemachine) {
    this.socket = undefined;
    this.statemachine = statemachine;
  }

  connect() {
    this.socket = new WebSocket(ADDRESS);
    this.socket.onopen = ev => {
      console.log('socket.onopen: ', ev);
    };
    this.socket.onmessage = ev => {
      console.log('socket.onmessage: ', ev);
      this.statemachine.handleMessage(ev.data);
    };
    this.socket.onerror = ev => {
      console.log('socket.onerror: ', ev);
    };
  }

  getState() {
    this.statemachine.getState();
  }

  register() {
    this.socket.send(
      JSON.stringify(ProtocolGenerator.createRegisterReq('testUser'))
    );
  }

  attachGroup(groupId) {
    this.socket.send(
      JSON.stringify(ProtocolGenerator.createGroupAttachReq(groupId))
    );
  }

  requestSetup() {
    this.socket.send(
      JSON.stringify(
        ProtocolGenerator.createSetupReq(this.statemachine._connectedGroupId)
      )
    );
  }
}

const StateEnum = Object.freeze({
  NotRegistered: 'Not registered',
  Registered: 'Registered',
  AttachedToGroup: 'Attached to group',
  Calling: 'Calling',
  GettingCalled: 'Getting Called'
});

class Statemachine {
  constructor() {
    this._state = StateEnum.NotRegistered;
    this._connectedGroupId = undefined;
    this._inCall = false;
  }

  getState() {
    console.log(this._state);
    return this._state;
  }

  getGroup() {
    console.log(this._connectedGroupId);
    return this._connectedGroupId;
  }

  handleMessage(msg) {
    console.log('Handling message ', msg);
    let m = JSON.parse(msg);
    switch (m.type) {
      case 'register_ack':
        this.onRegister(m.payload.result);
        break;
      case 'groupAttach_ack':
        this.onGroupAttach(m.payload.groupId, m.payload.result);
        break;
      case 'setup_ack':
        this.onSetupAck(m.payload.groupId, m.payload.result);
        break;
      default:
        console.log('Unknown type', m.type);
    }
  }

  onRegister(result) {
    if (this._state !== StateEnum.NotRegistered) {
      console.log('Register in wrong state!');
      return;
    }
    if (result !== 200) {
      console.log('Not registering due to wrong result code');
      this._state = StateEnum.NotRegistered;
      return;
    }
    this._state = StateEnum.Registered;
  }

  onGroupAttach(groupId, resultCode) {
    if (this._state !== StateEnum.Registered) {
      console.log('GroupAttach in wrong state!');
      return;
    }
    if (resultCode !== 200) {
      console.log(
        'GroupAttach not happening due to wrong result code: ',
        resultCode
      );
      return;
    }
    this._connectedGroupId = groupId;
    this._state = StateEnum.AttachedToGroup;
  }

  onSetupAck(groupId, resultCode) {
    if (this._state !== StateEnum.AttachedToGroup) {
      console.log('SetupAck in wrong state');
      return;
    }

    if (resultCode !== 200) {
      console.log(
        'SetupAck not happening because of the wrong result code: ',
        resultCode
      );
      return;
    }
    this._inCall = true;
    this._state = StateEnum.Calling;
  }
}

function bootstrapTestClient() {
  const statemachine = new Statemachine();
  return new Client(statemachine);
}
