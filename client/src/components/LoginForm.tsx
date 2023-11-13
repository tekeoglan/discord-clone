"use client";

import { useRouter } from "next/navigation";
import { MouseEventHandler, FormEventHandler, useState } from "react";
import { z } from "zod";
import { endpoints } from "@/lib/api";

export default function LoginForm() {
  const router = useRouter();

  const [valid, setValid] = useState(true);
  const [fetching, setFetching] = useState(false);

  const errMessage = "email or password is not valid";

  const formSchema = z.object({
    email: z.string().trim().email(),
    password: z.string().min(6).max(14),
  });

  const loginHandler: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    if (fetching) return;

    const entries = new FormData(e.target as HTMLFormElement);
    const loginData = Object.fromEntries(entries);

    const res = formSchema.safeParse(loginData);
    if (!res.success) {
      setValid(false);
      return;
    }

    setFetching(true);
    const response = await fetch(endpoints.Login, {
      method: "POST",
      credentials: "include",
      mode: "cors",
      cache: "no-store",
      body: entries,
    });

    if (!response.ok) {
      setValid(false);
      setFetching(false);
      return;
    }

    setValid(true);
    setFetching(false);
    router.replace("/");
  };

  const registerHandler: MouseEventHandler<HTMLButtonElement> = () => {
    router.push("/register");
  };

  return (
    <section className="w-96 p-8 bg-zinc-800 font-normal text-base text-white rounded">
      <form
        className="w-full flex flex-col grow justify-center items-center"
        onSubmit={loginHandler}
      >
        <div className="w-full flex flex-col mt-8">
          <label className="flex items-center mb-1 text-sm font-semibold text-neutral-300 uppercase">
            email
            <span
              className={`ml-1 font-light text-xs normal-case ${
                valid ? "hidden" : "text-red-500"
              }`}
            >
              {"-"}
              <span className="ml-1">{errMessage}</span>
            </span>
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            name="email"
            type="text"
            maxLength={99}
          ></input>
        </div>
        <div className="w-full flex flex-col mt-4 mb-8">
          <label className="flex items-center mb-1 text-sm font-semibold text-neutral-300 uppercase">
            password
            <span
              className={`ml-1 font-light text-xs normal-case ${
                valid ? "hidden" : "text-red-500"
              }`}
            >
              {"-"}
              <span className="ml-1">{errMessage}</span>
            </span>
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            name="password"
            type="password"
            maxLength={99}
          ></input>
        </div>
        <div className="w-full flex justify-center">
          <button
            className="min-w-[134px] min-h-[44px] w-full bg-blue-600 rounded"
            type="submit"
          >
            Log In
          </button>
        </div>
        <div className="w-full flex mt-1 text-left">
          <span className="text-xs text-zinc-400">Need an account?</span>
          <button
            className="inline-block ml-1 text-xs text-sky-500 align-bottom"
            type="button"
            onClick={registerHandler}
          >
            Register
          </button>
        </div>
      </form>
    </section>
  );
}
