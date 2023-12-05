import { create } from "zustand";
import { UserMeta } from "../entities/user";
import { createJSONStorage, persist } from "zustand/middleware";

type AcountState = {
  current: UserMeta | null;
  setUser: (account: UserMeta) => void;
  logout: () => void;
};

export const userStore = create<AcountState>()(
  persist(
    (set) => ({
      current: null,
      setUser: (account) => set({ current: account }),
      logout: () => set({ current: null }),
    }),
    {
      name: "user-storage",
      storage: createJSONStorage(() => sessionStorage),
      partialize: (state) => ({ current: state.current }),
    }
  )
);
