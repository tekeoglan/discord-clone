export type MessageItem = {
  id: string;
  userName: string;
  timestamp: string;
  text: string;
};

export default function MessageItem(props: MessageItem) {
  return (
    <li className="relative outline-none">
      <div className="relative mt-4 px-12 min-h-11 break-words select-text">
        <div className="static">
          <h3 className="relative block min-h-5 leading-5 whitespace-break-spaces">
            <span className="mr-1 font-normal text-base text-white overflow-hidden">
              {props.userName}
            </span>
            <span className="ml-1 font-normal text-xs text-neutral-400 align-baseline">
              {props.timestamp}
            </span>
          </h3>
          <div className="relative font-normal text-base text-neutral-200 whitespace-break-spaces break-words select-text overflow-hidden">
            <span>{props.text}</span>
          </div>
        </div>
      </div>
    </li>
  );
}
