import { createContext } from "react"

export default createContext<{
  token: string,
  onLogin: (email: string, password: string) => void;
  onLogout: () => void;
}>({
  token: "",
  onLogin: () => {},
  onLogout: () => {}
});