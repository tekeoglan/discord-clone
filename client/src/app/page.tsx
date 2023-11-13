"use client";

import { endpoints } from "@/lib/api";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Home() {
  const router = useRouter();

  useEffect(() => {
    const fetchUser = async () => {
      const user = await fetch(endpoints.FetchUser, {
        credentials: "include",
        mode: "cors",
      });

      if (!user.ok) {
        router.replace("/login");
      }
    };
    fetchUser();
  }, []);

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <h1>hello, world</h1>
    </main>
  );
}
