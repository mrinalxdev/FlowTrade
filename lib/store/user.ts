import { create } from "zustand";
import { User } from "@supabase/supabase-js";

interface UserState {
  users: User | undefined;
  setUser: (user: User | undefined) => void;
}

export const useUser = create<UserState>()((set) => ({
  users: undefined,
  setUser: (users) => set(() => ({ users })),
}));
