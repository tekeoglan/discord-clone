"use client";

import { endpoints } from "@/lib/api";
import { MouseEventHandler, useState } from "react";

const enum FetchState {
  INITIAL = "INITIAL",
  FETCHING = "FETCHING",
  FETCHED = "FETCHED",
  FAILED = "FAILED",
}

export default function PendingFriendItem({
  userName,
  itemId,
}: {
  userName: string;
  itemId: string;
}) {
  const [fetchState, setFetchState] = useState<FetchState>(FetchState.INITIAL);

  const acceptFriend: MouseEventHandler<HTMLDivElement> = async () => {
    if (fetchState == FetchState.FETCHING) return;

    try {
      const response = await fetch(endpoints.AcceptFriendRequest(itemId), {
        method: "POST",
        mode: "cors",
        credentials: "include",
        cache: "no-store",
      });

      if (!response.ok) {
        setFetchState(FetchState.FAILED);
        alert("something went wrong");
        return;
      }

      setFetchState(FetchState.FETCHED);
    } catch (e) {
      setFetchState(FetchState.FAILED);
      alert("something went wrong");
    }
  };

  const declineFriend: MouseEventHandler<HTMLDivElement> = async () => {
    if (fetchState == FetchState.FETCHING) return;

    try {
      const response = await fetch(endpoints.DeclineFriendRequest(itemId), {
        method: "POST",
        mode: "cors",
        credentials: "include",
        cache: "no-store",
      });

      if (!response.ok) {
        setFetchState(FetchState.FAILED);
        alert("something went wrong");
        return;
      }

      setFetchState(FetchState.FETCHED);
    } catch (e) {
      setFetchState(FetchState.FAILED);
      alert("something went wrong");
    }
  };

  return (
    <div
      className="h-[60px] ml-[30px] mr-[20px] flex font-normal text-base overflow-hidden box-border cursor-pointer border-t border-solid border-neutral-500 hover:px-[10px] hover:py-4 hover:my-0 hover:mr-[10px] hover:ml-5 hover:bg-neutral-500 hover:rounded-lg hover:border-transparent"
      style={{
        display: fetchState === FetchState.FETCHED ? "none" : "initial",
      }}
    >
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
            onClick={acceptFriend}
          >
            <svg
              className="fill-neutral-400 hover:fill-white"
              xmlns="http://www.w3.org/2000/svg"
              height="18"
              viewBox="0 -960 960 960"
              width="18"
            >
              <path d="M382-240 154-468l57-57 171 171 367-367 57 57-424 424Z" />
            </svg>
          </div>
          <div
            className="w-9 h-9 ml-3 flex items-center justify-center bg-neutral-700 rounded-full"
            onClick={declineFriend}
          >
            <svg
              className="fill-neutral-400 hover:fill-white"
              xmlns="http://www.w3.org/2000/svg"
              height="18"
              viewBox="0 -960 960 960"
              width="18"
            >
              <path d="m256-200-56-56 224-224-224-224 56-56 224 224 224-224 56 56-224 224 224 224-56 56-224-224-224 224Z" />
            </svg>
          </div>
        </div>
      </div>
    </div>
  );
}
