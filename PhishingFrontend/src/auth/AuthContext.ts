import { createContext } from "react";

const AuthContext = createContext<{
  token: string,
  onLogin: () => void;
  onLogout: () => void;
}>({
  token: "",
  onLogin: () => {},
  onLogout: () => {}
});

export default AuthContext;