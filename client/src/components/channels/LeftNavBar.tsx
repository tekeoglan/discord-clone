"use client";

import { userStore } from "@/lib/stores/userStore";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function LeftNavBar() {
  const user = userStore((state) => state.current);
  const router = useRouter();

  useEffect(() => {
    if (!user) router.replace("/login");
  }, []);

  return (
    <nav className="flex flex-col shrink-0 w-[72px] bg-neutral-900 overflow-hidden">
      <ul>
        <li className="flex justify-center mb-2">
          <div className="rounded-xl bg-blue-500 cursor-pointer"></div>
        </li>
      </ul>
    </nav>
  );
}
