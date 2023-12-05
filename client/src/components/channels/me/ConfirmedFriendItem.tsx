"use client";

import { MouseEventHandler } from "react";

export default function ConfirmedFriendItem({
  userName,
}: {
  userName: string;
}) {
  const sendMessage: MouseEventHandler<HTMLDivElement> = () => {};

  return (
    <div className="h-[60px] ml-[30px] mr-[20px] flex font-normal text-base overflow-hidden box-border cursor-pointer border-t border-solid border-neutral-500 hover:px-[10px] hover:py-4 hover:my-0 hover:mr-[10px] hover:ml-5 hover:bg-neutral-500 hover:rounded-lg hover:border-transparent">
      <div className="flex grow items-center justify-between max-w-full">
        <div className="flex items-center overflow-hidden">
          <div className="w-8 h-8 mr-3 bg-blue-500 rounded-full"></div>
          <div className="flex grow">
            <span className="whitespace-nowrap overflow-hidden text-ellipsis text-white font-semibold">
              {userName}
            </span>
          </div>
        </div>
        <div className="ml-2 flex">
          <div
            className="w-9 h-9 flex items-center justify-center bg-neutral-700 rounded-full"
            onClick={sendMessage}
          >
            <svg
              className="fill-neutral-400 hover:fill-white"
              xmlns="http://www.w3.org/2000/svg"
              height="18"
              viewBox="0 -960 960 960"
              width="18"
            >
              <path d="M80-80v-720q0-33 23.5-56.5T160-880h640q33 0 56.5 23.5T880-800v480q0 33-23.5 56.5T800-240H240L80-80Zm126-240h594v-480H160v525l46-45Zm-46 0v-480 480Z" />
            </svg>
          </div>
        </div>
      </div>
    </div>
  );
}
