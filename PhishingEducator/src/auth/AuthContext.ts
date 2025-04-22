import { createContext } from "react";

const AuthContext = createContext<{
  token: string,
  onLogin: (email: string, password: string) => void;
  onLogout: () => void;
}>({
  token: "",
  onLogin: () => {},
  onLogout: () => {}
});

export default AuthContext;