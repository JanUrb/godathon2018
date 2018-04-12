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
  static createGroupAttachReq(groupId) {
    return {
      type: 'groupAttach_req',
      payload: {
        id: '9876'
      }
    };
  }
}
