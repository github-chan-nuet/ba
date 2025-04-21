import { useState } from "react";
import AuthContext from "./AuthContext";

type AuthProviderProps = {
  children: React.ReactNode;
};

const fakeAuth = (): Promise<string> =>
  new Promise((resolve) => {
    setTimeout(() => resolve("2342f2f1d131rf12"), 250);
  });

const AuthProvider = ({ children }: AuthProviderProps) => {
  const [token, setToken] = useState("");

  const handleLogin = async () => {
    const token = await fakeAuth();

    setToken(token);
  }

  const handleLogout = () => {
    setToken("");
  }
  
  const value = {
    token,
    onLogin: handleLogin,
    onLogout: handleLogout
  }
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export default AuthProvider;