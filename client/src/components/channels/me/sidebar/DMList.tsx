"use client";

import { MouseEventHandler, useEffect, useRef, useState } from "react";
import { endpoints } from "@/lib/api";
import { FriendChannelResponse } from "@/lib/entities/channel";
import Link from "next/link";
import { useParams } from "next/navigation";
import {
  WebsocketAction,
  WebsocketChannelResponse,
  WebsocketRequest,
} from "@/lib/websocket";
import { userStore } from "@/lib/stores/userStore";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import { ReadyState } from "react-use-websocket";

type WSMessage = {
  action: WebsocketAction.AddChannelAction;
  data: WebsocketChannelResponse;
};

export default function DMList() {
  const user = userStore((state) => state.current);
  const [data, setData] = useState<FriendChannelResponse[]>([]);
  const [newDms, setNewDms] = useState<WebsocketChannelResponse[]>([]);
  const listRef = useRef<HTMLUListElement | null>(null);
  const selectedRef = useRef<HTMLElement | null>(null);
  const params = useParams();
  const { readyState, sendJsonMessage, lastJsonMessage } =
    useWebSocket<WSMessage>(endpoints.WebSocket, {
      share: true,
    });
  const createDmHandler: MouseEventHandler<HTMLSpanElement> = () => {};

  useEffect(() => {
    if (!(user && readyState == ReadyState.OPEN)) return;

    const wsRequest: WebsocketRequest = {
      action: WebsocketAction.JoinUserAction,
      room: user.ID,
    };

    sendJsonMessage(wsRequest);

    return () => {
      sendJsonMessage({
        action: WebsocketAction.LeaveRoomAction,
        room: user.ID,
      });
    };
  }, [user, readyState]);

  useEffect(() => {
    if (!(lastJsonMessage && Object.keys(lastJsonMessage).length > 0)) return;

    switch (lastJsonMessage.action) {
      case WebsocketAction.AddChannelAction:
        console.log("add_channel:", lastJsonMessage.data);
        setNewDms((prev) => [lastJsonMessage.data, ...prev]);
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

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
          if (list) {
            setData(list);
          }
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  useEffect(() => {
    if (params.id && listRef.current) {
      if (selectedRef.current) {
        selectedRef.current.style.background = "initial";
      }

      const el = listRef.current.querySelector(
        `[data-fc-id="${params.id}"]`
      ) as HTMLElement;
      if (el) {
        el.style.background = "#57534e";
        selectedRef.current = el;
      }
    }
  }, [params.id, listRef.current]);

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
        <ul className="text-semibold text-sm text-white" ref={listRef}>
          <div>
            {newDms.length > 0
              ? newDms.map((val) => {
                  return (
                    <li key={val.id}>
                      <Link href={`/channels/me/${val.id}`}>
                        <div
                          className="flex p-1 items-center rounded hover:bg-neutral-600"
                          data-fc-id={val.id}
                        >
                          <div className="w-7 h-7 mr-2 rounded-full bg-blue-500"></div>
                          <span className="overflow-hidden text-ellipsis">
                            {val.name}
                          </span>
                        </div>
                      </Link>
                    </li>
                  );
                })
              : null}
          </div>
          <div>
            {data.length > 0
              ? data.map((val) => {
                  return (
                    <li key={val.ChannelID}>
                      <Link href={`/channels/me/${val.ChannelID}`}>
                        <div
                          className="flex p-1 items-center rounded hover:bg-neutral-600"
                          data-fc-id={val.ChannelID}
                        >
                          <div className="w-7 h-7 mr-2 rounded-full bg-blue-500"></div>
                          <span className="overflow-hidden text-ellipsis">
                            {val.FriendInfo.UserName}
                          </span>
                        </div>
                      </Link>
                    </li>
                  );
                })
              : null}
          </div>
        </ul>
      </nav>
    </div>
  );
}
