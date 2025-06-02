import { createContext } from "react"
import type { User, UserPatchModel } from "../../api";

export default createContext<{
  token: string,
  user: (User & { totalXpForNextLevel: number; id?: string|undefined } ) | null,
  updateUser: (userPatch: UserPatchModel) => Promise<void>,
  onExperienceGain: (xpGain: number, newLevel: number|undefined) => void;
  onLogin: (email: string, password: string) => Promise<void>;
  onLogout: () => void;
}>({
  token: "",
  user: null,
  updateUser: () => new Promise(() => {}),
  onExperienceGain: () => {},
  onLogin: () => new Promise(() => {}),
  onLogout: () => {}
});