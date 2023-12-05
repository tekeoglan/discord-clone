export const endpoints = {
  WebSocket: `${process.env.NEXT_PUBLIC_WS_PATH}`,
  Login: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/login`,
  Logout: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/logout`,
  Register: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/register`,
  FetchUser: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/fetchUser`,
  GetFriendChannels: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/channel/fc/user`,
  GetConfirmedFriends: (cursorPos: number) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/getConfirmed?cursorPos=${cursorPos}`,
  GetPendingFriends: (cursorPos: number) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/getPending?cursorPos=${cursorPos}`,
  AddFriend: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/add`,
  AcceptFriendRequest: (friendId: string) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/confirm/${friendId}`,
  DeclineFriendRequest: (friendId: string) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/remove?id=${friendId}`,
  RemoveFriend: (userId: string) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/removeByUserIds?userId=${userId}`,
};
