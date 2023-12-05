import { endpoints } from "@/lib/api";
import { useEffect, useRef, useState } from "react";

export default function useWebsocket(): [WebSocket | null, number] {
  const socket = useRef<WebSocket | null>(null);
  const [refCount, setRefCount] = useState<number>(0);

  useEffect(() => {
    if (!socket.current) {
      const ws = new WebSocket(endpoints.WebSocket);

      socket.current = ws;
    }

    setRefCount((prev) => prev + 1);

    return () => {
      if (!socket.current || refCount != 1) {
        setRefCount((prev) => (prev != 0 ? prev - 1 : 0));
        return;
      }

      socket.current.close;
    };
  }, []);

  return [socket.current, refCount];
}
