"use client";

import { MouseEventHandler, useEffect, useState } from "react";
import { endpoints } from "@/lib/api";
import { FriendChannelResponse } from "@/lib/entities/channel";

export default function DMList() {
  const [data, setData] = useState<FriendChannelResponse[] | null>(null);

  const createDmHandler: MouseEventHandler<HTMLSpanElement> = () => {};

  useEffect(() => {
    (async function () {
      try {
        const data = await fetch(endpoints.GetFriendChannels, {
          mode: "cors",
          credentials: "include",
          cache: "no-store",
        });

        if (data.ok) {
          const list = await data.json();
          setData(list);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <div className="flex flex-col grow overflow-hidden">
      <div className="flex mb-1 text-neutral-400 hover:text-neutral-200 px-3 pt-3">
        <span className="grow font-semibold text-xs uppercase cursor-default">
          Direct Messages
        </span>
        <span
          className="ml-1 cursor-pointer"
          role="button"
          onClick={createDmHandler}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            height="16"
            viewBox="0 -960 960 960"
            width="16"
            fill="white"
          >
            <path d="M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z" />
          </svg>
        </span>
      </div>
      <nav className="flex flex-col px-3">
        <ul className="text-semibold text-sm">
          {data?.map((val) => {
            return (
              <li key={val.FriendInfo.ID}>
                <div
                  className="flex p-1 items-center rounded hover:bg-neutral-600"
                  role="button"
                >
                  <div className="w-7 h-7 mr-2 rounded-full bg-blue-500"></div>
                  <span className="overflow-hidden text-ellipsis">
                    {val.FriendInfo.UserName}
                  </span>
                </div>
              </li>
            );
          })}
        </ul>
      </nav>
    </div>
  );
}
