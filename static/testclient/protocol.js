/**
 * Helper class to create the protocol messages
 */
class ProtocolGenerator {
  static createRegisterReq(username) {
    return {
      type: 'register_req',
      payload: {
        user: username
      }
    };
  }
}
