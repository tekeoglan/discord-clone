export const endpoints = {
  Login: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/login`,
  Register: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/register`,
  FetchUser: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/fetchUser`,
  GetFriendChannels: `${process.env.NEXT_PUBLIC_API_BASE_PATH}/channel/fc/user`,
  GetConfirmedFriends: (cursorPos: number) =>
    `${process.env.NEXT_PUBLIC_API_BASE_PATH}/friend/getConfirmed?cursorPos=${cursorPos}`,
};
