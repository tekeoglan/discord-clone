"use client";

import { endpoints } from "@/lib/api";
import { FriendGetAllResponse } from "@/lib/entities/friend";
import { useEffect, useState } from "react";
import PendingFriendItem from "./PendingFriendItem";
import { userStore } from "@/lib/stores/userStore";
import {
  WebsocketAction,
  WebsocketRequest,
  WsFriendRequestResponse,
} from "@/lib/websocket";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import { ReadyState } from "react-use-websocket";

type WSMessage =
  | { action: WebsocketAction.RemoveFriendAction; data: string }
  | { action: WebsocketAction.AddRequestAction; data: WsFriendRequestResponse };

export default function PendingFriendsContainer() {
  const user = userStore((state) => state.current);
  const [friends, setFriends] = useState<FriendGetAllResponse | null>(null);
  const [newFriendRequests, setNewFriendRequests] = useState<
    WsFriendRequestResponse[]
  >([]);
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
      case WebsocketAction.AddRequestAction:
        console.log("add_request:", lastJsonMessage.data);
        setNewFriendRequests((prev) => [lastJsonMessage.data, ...prev]);
        break;
      case WebsocketAction.RemoveFriendAction:
        console.log("remove_friend:", lastJsonMessage.data);
        setNewFriendRequests((prev) =>
          prev.filter((item) => item.id != lastJsonMessage.data)
        );

        setFriends((prev) =>
          prev
            ? {
                ...prev,
                Friends: prev.Friends.filter(
                  (item) => item.ID != lastJsonMessage.data
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
        const response = await fetch(endpoints.GetPendingFriends(0), {
          mode: "cors",
          cache: "no-cache",
          credentials: "include",
        });

        if (response.ok) {
          const data = await response.json();
          setFriends(data);
        }
      } catch (e) {
        console.log(e);
      }
    })();
  }, []);

  return (
    <div className="flex flex-col grow shrink overflow-hidden">
      <div className="flex items-center justify-between">
        <h2 className="mt-4 mr-5 mb-2 ml-[30px] box-border text-ellipsis whitespace-nowrap uppercase overflow-hidden font-semibold text-neutral-300 grow shrink">
          {`pending friends - ${friends?.Friends?.length ?? 0}`}
        </h2>
      </div>
      <div className="pb-2 mt-2 overflow-x-hidden overflow-y-auto box-border grow shrink">
        <div>
          {newFriendRequests.length > 0
            ? newFriendRequests.map((item) => (
                <PendingFriendItem
                  key={item.id}
                  userName={item.userName}
                  itemId={item.id}
                />
              ))
            : null}
        </div>
        <div>
          {friends?.Friends?.map((item) => (
            <PendingFriendItem
              key={item.ID}
              userName={item.FriendInfo.UserName}
              itemId={item.ID}
            />
          ))}
        </div>
      </div>
    </div>
  );
}
