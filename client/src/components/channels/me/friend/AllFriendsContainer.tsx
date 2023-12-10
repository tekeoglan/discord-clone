"use client";

import { endpoints } from "@/lib/api";
import { FriendGetAllResponse } from "@/lib/entities/friend";
import { useEffect, useState } from "react";
import ConfirmedFriendItem from "./ConfirmedFriendItem";
import { WebsocketAction, WebsocketRequest } from "@/lib/websocket";
import { userStore } from "@/lib/stores/userStore";
import { UserEntity } from "@/lib/entities/user";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import { ReadyState } from "react-use-websocket";

type WSMessage =
  | { action: WebsocketAction.RemoveFriendAction; data: string }
  | { action: WebsocketAction.AddFriendAction; data: UserEntity };

export default function AllFriendsContainer() {
  const user = userStore((state) => state.current);
  const [friends, setFriends] = useState<FriendGetAllResponse | null>(null);
  const [newFriends, setNewFriends] = useState<UserEntity[]>([]);
  const { readyState, sendJsonMessage, lastJsonMessage } =
    useWebSocket<WSMessage>(endpoints.WebSocket, {
      share: true,
    });

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
      case WebsocketAction.AddFriendAction:
        console.log("add_friend:", lastJsonMessage.data);
        setNewFriends((prev) => [lastJsonMessage.data, ...prev]);
        break;

      case WebsocketAction.RemoveFriendAction:
        console.log("remove_friend:", lastJsonMessage.data);
        setNewFriends((prev) =>
          prev.filter((item) => item.ID != lastJsonMessage.data)
        );
        setFriends((prev) =>
          prev
            ? {
                ...prev,
                Friends: prev.Friends.filter(
                  (item) => item.ID == lastJsonMessage.data
                ),
              }
            : prev
        );
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    (async function () {
      try {
        const response = await fetch(endpoints.GetConfirmedFriends(0), {
          mode: "cors",
          cache: "no-cache",
          credentials: "include",
        });

        if (response.ok) {
          const data = await response.json();
          setFriends(data);
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  return (
    <div className="flex flex-col grow shrink overflow-hidden">
      <div className="flex items-center justify-between">
        <h2 className="mt-4 mr-5 mb-2 ml-[30px] box-border text-ellipsis whitespace-nowrap uppercase overflow-hidden font-semibold text-neutral-300 grow shrink">
          {`all friends - ${friends?.Friends?.length ?? 0}`}
        </h2>
      </div>
      <div className="pb-2 mt-2 overflow-x-hidden overflow-y-auto box-border grow shrink">
        <div>
          <div>
            {newFriends.length > 0
              ? newFriends.map((item) => (
                  <ConfirmedFriendItem
                    key={item.ID}
                    userName={item.UserName}
                    userId={item.ID}
                  />
                ))
              : null}
          </div>
          <div>
            {friends?.Friends?.map((item) => (
              <ConfirmedFriendItem
                key={item.ID}
                userName={item.FriendInfo.UserName}
                userId={item.FriendInfo.ID}
              />
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
