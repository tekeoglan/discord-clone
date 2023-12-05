import AddFriendsContainer from "./AddFriendsContainer";
import AllFriendsContainer from "./AllFriendsContainer";
import { FriendState } from "./FriendsHeader";
import PendingFriendsContainer from "./PendingFriendsContainer";

export default function FriendsContainer({
  friendState,
}: {
  friendState: FriendState;
}) {
  return (
    <div className="flex h-full  max-w-[720px] overflow-hidden">
      {friendState === FriendState.ALL ? <AllFriendsContainer /> : null}
      {friendState === FriendState.PENDING ? <PendingFriendsContainer /> : null}
      {friendState === FriendState.ADDING ? <AddFriendsContainer /> : null}
    </div>
  );
}
