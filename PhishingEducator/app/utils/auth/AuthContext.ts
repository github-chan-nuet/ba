import { createContext } from "react"
import type { User } from "../../api";

export default createContext<{
  token: string,
  user: User | null,
  onExperienceGain: (xpGain: number, newLevel: number|undefined) => void;
  onLogin: (email: string, password: string) => void;
  onLogout: () => void;
}>({
  token: "",
  user: null,
  onExperienceGain: () => {},
  onLogin: () => {},
  onLogout: () => {}
});