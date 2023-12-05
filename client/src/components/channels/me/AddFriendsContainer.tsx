"use client";

import { endpoints } from "@/lib/api";
import { FormEventHandler, useState } from "react";
import { z } from "zod";

export default function AddFriendsContainer() {
  const [fetching, setFetching] = useState(false);
  const emailSchema = z.string().trim().email();

  const sendFrienRequest: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();
    if (fetching) return;
    setFetching(true);

    const entries = new FormData(e.target as HTMLFormElement);
    const data = Object.fromEntries(entries);

    const parsedVal = emailSchema.safeParse(data.email);

    if (parsedVal.success) {
      try {
        const response = await fetch(endpoints.AddFriend, {
          method: "POST",
          mode: "cors",
          credentials: "include",
          cache: "no-store",
          body: entries,
        });

        if (!response.ok) {
          const err = await response.json();
          alert(err.message);
        } else {
          alert("friend request sent");
        }
      } catch (e) {
        alert(e);
        setFetching(false);
      }
    } else {
      alert("invalid email");
    }
    setFetching(false);
  };

  return (
    <div className="flex flex-col grow shrink overflow-hidden">
      <header className="shrink-0 border-solid border-b border-neutral-500 py-5 px-8">
        <h2 className="mb-2 font-semibold text-base text-white uppercase">
          Add Friend
        </h2>
        <form autoComplete="off" onSubmit={sendFrienRequest}>
          <div className="font-normal text-sm text-neutral-300">
            You can add friends with their Email address.
          </div>
          <div className="relative flex items-center mt-4 px-3 bg-neutral-900 rounded-lg border-solid border border-neutral-900 hover:border-blue-500 focus-within:border-blue-500">
            <div className="py-1 mr-4 grow shrink whitespace-pre box-border z-10">
              <input
                className="p-0 h-10 w-full box-border bg-inherit text-white rounded outline-none"
                name="email"
                maxLength={99}
              ></input>
            </div>
            <button
              className="relative h-8 min-w-15 min-h-8 flex items-center justify-center box-border bg-blue-600 rounded py-[2px] px-4 select-none"
              style={{ cursor: fetching ? "wait" : "pointer" }}
              type="submit"
            >
              <div className="font-normal text-sm text-white">
                Sen Friend Request
              </div>
            </button>
          </div>
        </form>
      </header>
    </div>
  );
}
