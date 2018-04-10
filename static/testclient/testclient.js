const ADDRESS = 'ws://localhost:4242/ws';

class Client {
  constructor() {
    this.name = 'hey';
    this.socket = undefined;
  }

  connect() {
    this.socket = new WebSocket(ADDRESS);
    this.socket.onopen = ev => {
      console.log('socket.onopen: ', ev);
    };
    this.socket.onmessage = ev => {
      console.log('socket.onmessage: ', ev);
    };
    this.socket.onerror = ev => {
      console.log('socket.onerror: ', ev);
    };
  }

  register() {
    this.socket.send(
      JSON.stringify(ProtocolGenerator.createRegisterReq('testUser'))
    );
  }
}
