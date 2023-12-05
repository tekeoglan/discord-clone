export const enum WebsocketReadyType {
  CONNECTING = 0,
  OPEN,
  CLOSING,
  CLOSED,
}

export type WebsocketRequest = {
  action: string;
  room: string;
  message?: any;
};

export type WebsocketResponse = {
  action: string;
  data: any;
};

export type WsFriendRequestResponse = {
  id: string;
  userId: string;
  userName: string;
  image: string;
};

export const enum WebsocketAction {
  JoinUserAction = "joinUser",
  JoinChannelAction = "joinChannel",
  LeaveRoomAction = "leaveRoom",

  // emitted messages

  NewMessageAction = "new_message",
  EditMessageAction = "edit_message",
  DeleteMessageAction = "delete_message",
  AddChannelAction = "add_channel",
  AddPrivateChannelAction = "add_private_channel",
  EditChannelAction = "edit_channel",
  DeleteChannelAction = "delete_channel",
  NewDirectMessageAction = "new_direct_message",
  NewDMNotificationAction = "new_dm_notification",
  ToggleOnlineEmission = "toggle_online",
  ToggleOfflineEmission = "toggle_offline",
  SendRequestAction = "send_request",
  AddRequestAction = "add_request",
  AddFriendAction = "add_friend",
  RemoveFriendAction = "remove_friend",
}
