"use client";

import { endpoints } from "@/lib/api";
import { FriendGetAllResponse } from "@/lib/entities/friend";
import { useEffect, useState } from "react";
import ConfirmedFriendItem from "./ConfirmedFriendItem";
import {
  WebsocketAction,
  WebsocketReadyType,
  WebsocketRequest,
} from "@/lib/websocket";
import { userStore } from "@/lib/stores/userStore";
import { UserEntity } from "@/lib/entities/user";
import useWebsocket from "@/hooks/useWebSocket";

type WSMessage =
  | { action: "remove_friend"; data: string }
  | { action: "add_friend"; data: UserEntity };

export default function AllFriendsContainer() {
  const [ws, _] = useWebsocket();
  const user = userStore((state) => state.current);
  const [friends, setFriends] = useState<FriendGetAllResponse | null>(null);
  const [newFriends, setNewFriends] = useState<UserEntity[]>([]);

  useEffect(() => {
    if (!(ws && user && ws.readyState == WebsocketReadyType.OPEN)) return;

    const wsRequest: WebsocketRequest = {
      action: WebsocketAction.JoinUserAction,
      room: user.ID,
    };

    ws.send(JSON.stringify(wsRequest));

    ws.addEventListener("message", (e) => {
      const response: WSMessage = JSON.parse(e.data);

      switch (response.action) {
        case "add_friend":
          console.log("add_friend:", response.data);
          setNewFriends((prev) => [response.data, ...prev]);
          break;

        case "remove_friend":
          console.log("remove_friend:", response.data);
          setNewFriends((prev) =>
            prev.filter((item) => item.ID == response.data)
          );
          setFriends((prev) =>
            prev
              ? {
                  ...prev,
                  Friends: prev.Friends.filter(
                    (item) => item.ID == response.data
                  ),
                }
              : prev
          );
          break;

        default:
          break;
      }
    });

    return () => {
      ws.send(
        JSON.stringify({
          action: WebsocketAction.LeaveRoomAction,
          room: user.ID,
        })
      );

      ws.close();
    };
  }, [ws, ws?.readyState]);

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
        <h2 className="mt-4 mr-5 mb-2 ml-[30px] box-border text-ellipsis whitespace-nowrap uppercase overflow-hidden font-semibold font-xs grow shrink">
          {`all friends -- ${friends?.Friends?.length ?? 0}`}
        </h2>
      </div>
      <div className="pb-2 mt-2 overflow-x-hidden overflow-y-scroll box-border grow shrink">
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
