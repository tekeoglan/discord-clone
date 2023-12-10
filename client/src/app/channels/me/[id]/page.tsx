"use client";

import MessageItem from "@/components/channels/me/chat/MessageItem";
import { endpoints } from "@/lib/api";
import { MessageEntity, MessageGetAllResult } from "@/lib/entities/message";
import { userStore } from "@/lib/stores/userStore";
import { FormEventHandler, useEffect, useState } from "react";
import { useInView } from "react-intersection-observer";
import { WebsocketAction, WebsocketRequest } from "@/lib/websocket";
import useChatScroll from "@/hooks/useChatScroll";
import { useWebSocket } from "react-use-websocket/dist/lib/use-websocket";
import { ReadyState } from "react-use-websocket";

type WSMessage = {
  action: WebsocketAction.NewMessageAction;
  data: MessageEntity;
};

export default function Page({ params }: { params: { id: string } }) {
  const { ref, inView } = useInView({ threshold: 0 });
  const user = userStore((state) => state.current);
  const [initialized, setInitialized] = useState(false);
  const [messages, setMessages] = useState<MessageEntity[]>([]);
  const [newMessages, setNewMessages] = useState<MessageEntity[]>([]);
  const [cursor, setCursor] = useState(0);
  const [loadingMore, setLoadingMore] = useState(false);
  const scrollableRef = useChatScroll([newMessages, initialized]);
  const { readyState, sendJsonMessage, lastJsonMessage } =
    useWebSocket<WSMessage>(endpoints.WebSocket, {
      share: true,
    });

  const sendText: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    if (!user) return;

    const request = new FormData(e.target as HTMLFormElement);
    request.append("userId", user.ID);
    request.append("userName", user.UserName);
    request.append("channelId", params.id);

    try {
      const response = await fetch(endpoints.SendMessage, {
        method: "Post",
        mode: "cors",
        credentials: "include",
        body: request,
        cache: "no-store",
      });

      if (!response.ok) {
        alert("couldn't send message");
      }
    } catch (error) {
      alert("couldn't send message");
      console.log(error);
    }

    (e.target as HTMLFormElement).reset();
  };

  useEffect(() => {
    if (!(params.id && readyState == ReadyState.OPEN)) return;

    const wsRequest: WebsocketRequest = {
      action: WebsocketAction.JoinUserAction,
      room: params.id,
    };

    sendJsonMessage(wsRequest);

    return () => {
      sendJsonMessage({
        action: WebsocketAction.LeaveRoomAction,
        room: params.id,
      });
    };
  }, [params.id, readyState]);

  useEffect(() => {
    if (!(lastJsonMessage && Object.keys(lastJsonMessage).length > 0)) return;

    switch (lastJsonMessage.action) {
      case WebsocketAction.NewMessageAction:
        console.log("new_message", lastJsonMessage.data);
        setNewMessages((prev) => [...prev, lastJsonMessage.data]);
        break;
      default:
        break;
    }
  }, [lastJsonMessage]);

  useEffect(() => {
    (async function () {
      try {
        const response = await fetch(
          endpoints.GetChannelMessages(params.id, 0),
          {
            mode: "cors",
            credentials: "include",
            cache: "no-store",
          }
        );

        if (response.ok) {
          const data = (await response.json()) as MessageGetAllResult;
          if (data.Messages) {
            setMessages(data.Messages);
            setCursor(data.CursorPos);
            setInitialized(true);
            console.log("fresh data: ", data);
          }
        }
      } catch (error) {
        console.log(error);
      }
    })();
  }, []);

  useEffect(() => {
    if (!loadingMore && inView && cursor != 0) {
      console.log("paginating...");
      setLoadingMore(true);
      const fetchMoreMessages = async () => {
        try {
          const response = await fetch(
            endpoints.GetChannelMessages(
              params.id,
              cursor + newMessages.length
            ),
            {
              mode: "cors",
              credentials: "include",
              cache: "no-store",
            }
          );
          if (response.ok) {
            const data = (await response.json()) as MessageGetAllResult;
            if (data.Messages) {
              setMessages([...messages, ...data.Messages]);
              setCursor(data.CursorPos);
            } else {
              setCursor(0);
            }
            setLoadingMore(false);
          }
        } catch (error) {
          setLoadingMore(false);
          console.log(error);
        }
      };
      fetchMoreMessages();
    }
  }, [inView, loadingMore, cursor]);

  return (
    <div className="relative flex flex-col overflow-hidden grow shrink bg-neutral-600">
      <section className="relative flex grow-0 shrink-0 min-h-[43px] z-50 w-full border-b border-neutral-900"></section>
      <div className="max-w-[1024px] min-w-0 min-h-0 grow shrink flex justify-stretch items-stretch relative border-r border-neutral-800">
        <div className="relative flex h-full w-full">
          <div className="absolute flex h-full w-full">
            <main className="relative flex flex-col grow shrink">
              <div className="relative flex grow shrink z-0">
                <div
                  ref={scrollableRef}
                  className="absolute inset-0 overflow-y-auto overflox-x-hidden box-border"
                >
                  <div className="relative min-h-full flex flex-col justify-end items-stretch">
                    <ol className="relative min-h-0 overflow-hidden">
                      <div ref={ref} className="text-transparent">
                        observer
                      </div>
                      <div className="relative flex flex-col-reverse">
                        {messages.length > 0
                          ? messages.map((message) => (
                              <MessageItem
                                key={message.ID}
                                id={message.ID}
                                userName={message.UserName}
                                timestamp={message.UpdatedAt}
                                text={message.Text}
                              />
                            ))
                          : null}
                      </div>
                      <div>
                        {newMessages.length > 0
                          ? newMessages.map((message) => (
                              <MessageItem
                                key={message.ID}
                                id={message.ID}
                                userName={message.UserName}
                                timestamp={message.UpdatedAt}
                                text={message.Text}
                              />
                            ))
                          : null}
                      </div>
                      <div className="h-[30px] w-[1px] pointer-events-none"></div>
                    </ol>
                  </div>
                </div>
              </div>
              <form
                className="relative shrink-0 px-4 -mt-2"
                onSubmit={sendText}
              >
                <div className="relative mb-6 w-full bg-neutral-500 rounded-lg">
                  <div className="relative max-h-[50vh] overflow-x-hidden overflow-y-auto">
                    <div className="relative pl-4 flex">
                      <div className="sticky grow-0 shrink-0 self-stretch">
                        <button
                          type="button"
                          className="sticky flex justify-center items-center cursor-pointer -ml-4 h-11 py-[10px] px-[16px] top-0 select-none"
                        >
                          <svg
                            aria-hidden="true"
                            role="img"
                            xmlns="http://www.w3.org/2000/svg"
                            width="24"
                            height="24"
                            fill="none"
                            viewBox="0 0 24 24"
                          >
                            <circle
                              cx="12"
                              cy="12"
                              r="10"
                              fill="transparent"
                            ></circle>
                            <path
                              fill="currentColor"
                              fillRule="evenodd"
                              d="M12 23a11 11 0 1 0 0-22 11 11 0 0 0 0 22Zm0-17a1 1 0 0 1 1 1v4h4a1 1 0 1 1 0 2h-4v4a1 1 0 1 1-2 0v-4H7a1 1 0 1 1 0-2h4V7a1 1 0 0 1 1-1Z"
                              clipRule="evenodd"
                            ></path>
                          </svg>
                        </button>
                      </div>
                      <div className="relative w-full h-11 min-h-11 font-sm text-base text-white box-border">
                        <input
                          autoComplete="off"
                          name="text"
                          className="relative w-full left-0 right-[10px] outline-none break-words py-[11px] pr-[10px] whitespace-break-spaces text-left bg-inherit"
                        ></input>
                      </div>
                    </div>
                  </div>
                </div>
              </form>
            </main>
          </div>
        </div>
      </div>
    </div>
  );
}
