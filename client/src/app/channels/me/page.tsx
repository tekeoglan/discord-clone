"use client";

import FriendsContainer from "@/components/channels/me/FriendsContainer";
import FriendsHeader, {
  FriendState,
} from "@/components/channels/me/FriendsHeader";
import { useState } from "react";

export default function Me() {
  const [friendState, setFriendState] = useState<FriendState>(FriendState.ALL);

  return (
    <main className="flex flex-col">
      <FriendsHeader setState={setFriendState} />
      <FriendsContainer friendState={friendState} />
    </main>
  );
}
