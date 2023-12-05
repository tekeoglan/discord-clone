"use client";

import Loading from "@/components/common/Loading";
import LoginForm from "@/components/LoginForm";
import { endpoints } from "@/lib/api";
import { userStore } from "@/lib/stores/userStore";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";

type LoginState = "initial" | "loading" | "loaded";

export default function Login() {
  const router = useRouter();
  const setUser = userStore((state) => state.setUser);

  const [state, setState] = useState<LoginState>("initial");

  useEffect(() => {
    if (state != "initial") return;
    setState("loading");

    const init = async () => {
      try {
        const response = await fetch(endpoints.FetchUser, {
          credentials: "include",
          mode: "cors",
          cache: "no-store",
        });

        if (response.ok) {
          const data = await response.json();
          setUser(data);
          return router.replace("/channels/me");
        } else {
          setState("loaded");
        }
      } catch (e) {
        console.log(e);
        setState("loaded");
      }
    };
    init();
  }, [state]);

  return <main>{state === "loaded" ? <LoginForm /> : <Loading />}</main>;
}
