const baseUrl = process.env.NEXT_PUBLIC_API_BASE_PATH;

const wsUrl = baseUrl ? baseUrl.replace(/^http/, "ws") : "";

export const endpoints = {
  WebSocket: `${wsUrl}/ws`,
  Login: `${baseUrl}/login`,
  Logout: `${baseUrl}/logout`,
  Register: `${baseUrl}/register`,
  FetchUser: `${baseUrl}/fetchUser`,
  GetFriendChannels: `${baseUrl}/channel/fc/user`,
  GetConfirmedFriends: (cursorPos: number) =>
    `${baseUrl}/friend/getConfirmed?cursorPos=${cursorPos}`,
  GetPendingFriends: (cursorPos: number) =>
    `${baseUrl}/friend/getPending?cursorPos=${cursorPos}`,
  GetChannelMessages: (channelId: string, cursorPos: number) =>
    `${baseUrl}/message/channel?channelId=${channelId}&cursorPos=${cursorPos}`,
  GetFcByUserIds: (friendId: string) =>
    `${baseUrl}/channel/fc/byUserIds?friendId=${friendId}`,
  CreateFriendChannel: (friendId: string) =>
    `${baseUrl}/channel/fc?friendId=${friendId}`,
  AddFriend: `${baseUrl}/friend/add`,
  AcceptFriendRequest: (friendId: string) =>
    `${baseUrl}/friend/confirm/${friendId}`,
  DeclineFriendRequest: (friendId: string) =>
    `${baseUrl}/friend/remove?id=${friendId}`,
  RemoveFriend: (userId: string) =>
    `${baseUrl}/friend/removeByUserIds?userId=${userId}`,
  SendMessage: `${baseUrl}/message/add`,
};
