"use client";

import LoginForm from "@/components/LoginForm";
import { endpoints } from "@/lib/api";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function Login() {
  const router = useRouter();

  useEffect(() => {
    const init = async () => {
      const response = await fetch(endpoints.FetchUser, {
        method: "GET",
        credentials: "include",
        mode: "cors",
      });

      if (response.ok) {
        router.replace("/channels/me");
      }
    };
    init();
  }, []);

  return (
    <main>
      <LoginForm />
    </main>
  );
}
