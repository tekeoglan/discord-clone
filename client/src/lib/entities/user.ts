import { BaseEntity } from "./base";

export type UserEntity = {
  UserName: string;
  Email: string;
  Password: string;
  Image: string;
} & BaseEntity;
