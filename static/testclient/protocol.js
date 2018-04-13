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
        groupId: groupId
      }
    };
  }

  static createSetupReq(groupId) {
    return {
      type: 'setup_req',
      payload: {
        groupId: groupId
      }
    };
  }
}
