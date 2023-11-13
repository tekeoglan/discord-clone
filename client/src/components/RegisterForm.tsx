"use client";

import { useRouter } from "next/navigation";
import {
  MouseEventHandler,
  FocusEventHandler,
  FormEventHandler,
  useReducer,
} from "react";
import { z } from "zod";

import { endpoints } from "@/lib/api";

type Data = {
  value: string;
  err: string;
};

enum FetchingStatus {
  INITIAL = "INITIAL",
  PENDING = "PENDING",
  FULLFILLED = "FULLFILLED",
  REJECTED = "REJECTED",
}

type FetchState = {
  status: FetchingStatus;
  message: string;
};

type State = {
  userName: Data;
  email: Data;
  password: Data;
  fetching: FetchState;
};

enum ActionType {
  UPDATE_USER_NAME = "UPDATE_USER_NAME",
  UPDATE_EMAIL = "UPDATE_EMAIL",
  UPDATE_PASSWORD = "UPDATE_PASSWORD",
  UPDATE_FETCHING_STATUS = "UPDATE_FETCHING_STATUS",
}

type InputPayload = Partial<Data>;
type FetchPayload = FetchState;

type Action =
  | {
      type: ActionType.UPDATE_USER_NAME;
      payload: InputPayload;
    }
  | {
      type: ActionType.UPDATE_EMAIL;
      payload: InputPayload;
    }
  | {
      type: ActionType.UPDATE_PASSWORD;
      payload: InputPayload;
    }
  | {
      type: ActionType.UPDATE_FETCHING_STATUS;
      payload: FetchPayload;
    };

const initialState = {
  userName: { value: "", err: "" },
  email: { value: "", err: "" },
  password: { value: "", err: "" },
  fetching: { status: FetchingStatus.INITIAL, message: "" },
};

const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case ActionType.UPDATE_USER_NAME:
      return {
        ...state,
        userName: {
          value: action.payload.value ?? state.userName.value,
          err: action.payload.err ?? "",
        },
      };
    case ActionType.UPDATE_EMAIL:
      return {
        ...state,
        email: {
          value: action.payload.value ?? state.email.value,
          err: action.payload.err ?? "",
        },
      };
    case ActionType.UPDATE_PASSWORD:
      return {
        ...state,
        password: {
          value: action.payload.value ?? state.password.value,
          err: action.payload.err ?? "",
        },
      };
    case ActionType.UPDATE_FETCHING_STATUS:
      return {
        ...state,
        fetching: {
          status: action.payload.status,
          message: action.payload.message,
        },
      };
  }
};

export default function RegisterForm() {
  const router = useRouter();

  const [state, dispatch] = useReducer(reducer, initialState);

  const userNameSchema = z
    .string()
    .trim()
    .min(4, "*musn't be less than 4 character")
    .max(12, "*musn't be greater than 12 character");

  const emailSchema = z.string().trim().email("*unvalid email");

  const passSchema = z
    .string()
    .min(6, "*musn't be less than 6 character")
    .max(14, "*musn't be greater than 14 character");

  const userNameBlurHandler: FocusEventHandler<HTMLInputElement> = (e) => {
    const res = userNameSchema.safeParse(e.target.value);
    if (!res.success) {
      dispatch({
        type: ActionType.UPDATE_USER_NAME,
        payload: { err: res.error.issues[0].message },
      });
      return;
    }

    dispatch({
      type: ActionType.UPDATE_USER_NAME,
      payload: { value: res.data },
    });
  };

  const emailBlurHandler: FocusEventHandler<HTMLInputElement> = (e) => {
    const res = emailSchema.safeParse(e.target.value);
    if (!res.success) {
      dispatch({
        type: ActionType.UPDATE_EMAIL,
        payload: { err: res.error.issues[0].message },
      });
      return;
    }

    dispatch({ type: ActionType.UPDATE_EMAIL, payload: { value: res.data } });
  };

  const passBlurHandler: FocusEventHandler<HTMLInputElement> = (e) => {
    const res = passSchema.safeParse(e.target.value);
    if (!res.success) {
      dispatch({
        type: ActionType.UPDATE_PASSWORD,
        payload: { err: res.error.issues[0].message },
      });
      return;
    }

    dispatch({
      type: ActionType.UPDATE_PASSWORD,
      payload: { value: res.data },
    });
  };

  const confirmPassBlurHandler: FocusEventHandler<HTMLInputElement> = (e) => {
    const res = passSchema.safeParse(e.target.value);
    if (res.success && res.data === state.password.value) {
      dispatch({
        type: ActionType.UPDATE_PASSWORD,
        payload: { err: "" },
      });
      return;
    }

    dispatch({
      type: ActionType.UPDATE_PASSWORD,
      payload: { err: "*confirmation password not valid" },
    });
  };

  const registerHandler: FormEventHandler<HTMLFormElement> = async (e) => {
    e.preventDefault();

    if (state.fetching.status == FetchingStatus.PENDING) return;

    if (!state.userName.value || !state.email.value || !state.password.value)
      return;

    const entries = new FormData(e.target as HTMLFormElement);

    dispatch({
      type: ActionType.UPDATE_FETCHING_STATUS,
      payload: {
        status: FetchingStatus.PENDING,
        message: "processing the input...",
      },
    });

    const response = await fetch(endpoints.Register, {
      method: "POST",
      credentials: "omit",
      mode: "cors",
      cache: "no-store",
      body: entries,
    });

    if (response.ok) {
      dispatch({
        type: ActionType.UPDATE_FETCHING_STATUS,
        payload: {
          status: FetchingStatus.FULLFILLED,
          message: "user successfully created",
        },
      });
    } else {
      if (response.status == 409) {
        dispatch({
          type: ActionType.UPDATE_FETCHING_STATUS,
          payload: {
            status: FetchingStatus.REJECTED,
            message: "this email address is already in use",
          },
        });
      } else {
        dispatch({
          type: ActionType.UPDATE_FETCHING_STATUS,
          payload: {
            status: FetchingStatus.REJECTED,
            message: "something went wrong",
          },
        });
      }
    }
  };

  const loginHandler: MouseEventHandler<HTMLButtonElement> = () => {
    router.push("/login");
  };

  return (
    <section className="w-96 p-8 bg-zinc-800 font-normal text-base text-white rounded">
      <form
        className="w-full flex flex-col grow justify-center items-center"
        onSubmit={registerHandler}
      >
        <div className="w-full flex flex-col">
          <label className="mb-1 text-sm font-semibold text-neutral-300 uppercase">
            username
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            name="userName"
            type="text"
            maxLength={99}
            onBlur={userNameBlurHandler}
          ></input>
          <span className="inline-block h-4 w-full font-light text-xs text-red-500">
            {state.userName.err}
          </span>
        </div>
        <div className="w-full flex flex-col mt-4">
          <label className="mb-1 text-sm font-semibold text-neutral-300 uppercase">
            email
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            name="email"
            type="text"
            maxLength={99}
            onBlur={emailBlurHandler}
          ></input>
          <span className="inline-block h-4 w-full font-light text-xs text-red-500">
            {state.email.err}
          </span>
        </div>
        <div className="w-full flex flex-col mt-4">
          <label className="mb-1 text-sm font-semibold text-neutral-300 uppercase">
            password
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            name="password"
            type="password"
            maxLength={99}
            onBlur={passBlurHandler}
          ></input>
          <span className="inline-block h-4 w-full font-light text-xs text-red-500">
            {state.password.err}
          </span>
        </div>
        <div className="w-full flex flex-col mt-4">
          <label className="mb-1 text-sm font-semibold text-neutral-300 uppercase">
            password
          </label>
          <input
            className="h-10 p-1 bg-zinc-900 rounded caret-white outline-none"
            type="password"
            maxLength={99}
            onBlur={confirmPassBlurHandler}
          ></input>
          <span className="inline-block h-4 w-full font-light text-xs text-red-500">
            {state.password.err}
          </span>
        </div>
        <div className="w-full flex justify-center mt-8">
          <button
            className="min-w-[134px] min-h-[44px] w-full bg-blue-600 rounded"
            type="submit"
          >
            Register
          </button>
        </div>
        <div className="w-full flex mt-1 text-left">
          <span className="text-xs text-zinc-400">
            Already have an account?
          </span>
          <button
            className="inline-block ml-1 text-xs text-sky-500 align-bottom"
            type="button"
            onClick={loginHandler}
          >
            Login
          </button>
        </div>
      </form>
      <div
        className={`flex items-center w-full pl-2 h-8 mt-4 rounded ${
          state.fetching.status === FetchingStatus.FULLFILLED
            ? "bg-green-400"
            : state.fetching.status === FetchingStatus.REJECTED
            ? "bg-red-400"
            : "hidden"
        }`}
      >
        <span className="text-left text-xs text-white">
          {state.fetching.message}
        </span>
      </div>
    </section>
  );
}
