"use client";

import FriendsContainer from "@/components/channels/me/friend/FriendsContainer";
import FriendsHeader, {
  FriendState,
} from "@/components/channels/me/friend/FriendsHeader";
import { useState } from "react";

export default function Me() {
  const [friendState, setFriendState] = useState<FriendState>(FriendState.ALL);

  return (
    <main className="flex flex-col h-full w-full bg-neutral-600">
      <FriendsHeader setState={setFriendState} />
      <FriendsContainer friendState={friendState} />
    </main>
  );
}
