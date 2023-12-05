"use client";

import { Dispatch, SetStateAction, MouseEventHandler } from "react";

export const enum FriendState {
  ALL = "ALL",
  PENDING = "PENDING",
  ADDING = "ADDING",
}

export default function FriendsHeader({
  setState,
}: {
  setState: Dispatch<SetStateAction<FriendState>>;
}) {
  const getAll: MouseEventHandler<HTMLDivElement> = () => {
    setState(FriendState.ALL);
  };

  const getPending: MouseEventHandler<HTMLDivElement> = () => {
    setState(FriendState.PENDING);
  };

  const addFriend: MouseEventHandler<HTMLDivElement> = () => {
    setState(FriendState.ADDING);
  };

  return (
    <section className="w-full p-2 border-b border-neutral-900">
      <div className="w-full flex items-center overflow-hidden">
        <div className="h-6 mx-2 grow-0 shrink-0">
          <div className="flex">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="24"
              viewBox="0 -960 960 960"
              width="24"
              fill="white"
            >
              <path d="M40-160v-112q0-34 17.5-62.5T104-378q62-31 126-46.5T360-440q66 0 130 15.5T616-378q29 15 46.5 43.5T680-272v112H40Zm720 0v-120q0-44-24.5-84.5T666-434q51 6 96 20.5t84 35.5q36 20 55 44.5t19 53.5v120H760ZM360-480q-66 0-113-47t-47-113q0-66 47-113t113-47q66 0 113 47t47 113q0 66-47 113t-113 47Zm400-160q0 66-47 113t-113 47q-11 0-28-2.5t-28-5.5q27-32 41.5-71t14.5-81q0-42-14.5-81T544-792q14-5 28-6.5t28-1.5q66 0 113 47t47 113ZM120-240h480v-32q0-11-5.5-20T580-306q-54-27-109-40.5T360-360q-56 0-111 13.5T140-306q-9 5-14.5 14t-5.5 20v32Zm240-320q33 0 56.5-23.5T440-640q0-33-23.5-56.5T360-720q-33 0-56.5 23.5T280-640q0 33 23.5 56.5T360-560Zm0 320Zm0-400Z" />
            </svg>
            <div className="mx-2 text-base font-semibold text-white">
              <span>Friends</span>
            </div>
          </div>
        </div>
        <div className="flex font-normal text-base text-center text-white overflow-hidden text-ellipsis ">
          <div
            className="min-w-10 mx-2 py-px px-2 flex justify-center shrink-0 rounded items-center whitespace-nowrap cursor-pointer hover:bg-neutral-400 select-none"
            onClick={getAll}
          >
            <span>All</span>
          </div>
          <div
            className="min-w-10 mx-2 py-px px-2 flex rounded justify-center items-center shrink-0 whitespace-nowrap cursor-pointer hover:bg-neutral-400 select-none"
            onClick={getPending}
          >
            <span>Pending</span>
          </div>
          <div
            className="min-w-10 mx-2 py-px px-2 flex rounded justify-center items-center bg-green-700 shrink-0 whitespace-nowrap cursor-pointer hover:bg-green-600 select-none"
            onClick={addFriend}
          >
            <span>Add Friend</span>
          </div>
        </div>
      </div>
    </section>
  );
}
