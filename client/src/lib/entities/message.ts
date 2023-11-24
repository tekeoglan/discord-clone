import { BaseEntity } from "./base";
import { UserEntity } from "./user";

export type MessageEntity = {
  UserID: string;
  ChannelID: string;
  Text: string;
} & BaseEntity;

export type DirectMessage = {
  ChannelID: string;
  User: UserEntity;
};

export type UpdateMessageRequest = {
  ID: string;
  Text: string;
};
