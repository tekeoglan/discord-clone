import { BaseEntity } from "./base";
import { UserEntity } from "./user";

export type FriendEntity = {
  Users: string[];
  Status: string;
} & BaseEntity;

export type FriendGetResponse = {
  Users: string[];
  FriendInfos: UserEntity[];
  Status: string;
} & BaseEntity;

type FriendAggragateResult = {
  FriendInfo: UserEntity;
  Status: string;
} & BaseEntity;

export type FriendGetAllResponse = {
  Friends: FriendAggragateResult[];
  CursorPos: number;
};

export type FriendRequestWs = {
  Id: string;
  UserName: string;
  Image: string;
  Type: 0 | 1;
};
