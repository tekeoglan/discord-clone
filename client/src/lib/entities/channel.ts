import { BaseEntity } from "./base";
import { MessageEntity } from "./message";
import { UserEntity } from "./user";

export type ChannelEntity = {
  ChannelID: string;
  Messages: MessageEntity[] | null;
  Type: string;
} & BaseEntity;

export type FriendChannelEntity = {
  Users: string[];
} & ChannelEntity;

export type FriendChannelResponse = {
  FriendInfo: UserEntity;
} & ChannelEntity;

export type ChannelWsResponse = {
  Id: string;
  Name: string;
  CreatedAt: string;
  UpdatedAt: string;
  HasNotification: boolean;
};
