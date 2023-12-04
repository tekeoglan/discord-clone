import { BaseEntity } from "./base";

export type UserEntity = {
  UserName: string;
  Email: string;
  Password: string;
  Image: string;
} & BaseEntity;

export type UserMeta = Omit<UserEntity, "Password">;
