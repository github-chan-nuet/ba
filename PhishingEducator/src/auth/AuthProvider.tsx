import AuthContext from "./AuthContext";
import { loginAndReturnJwtToken } from "../api";
import useLocalStorage from "./useLocalStorage";

type AuthProviderProps = {
  children: React.ReactNode;
};

const AuthProvider = ({ children }: AuthProviderProps) => {
  const [token, setToken] = useLocalStorage('login-token', null);

  const handleLogin = async (email: string, password: string) => {
    const result = await loginAndReturnJwtToken({ body: { email, password } });
    setToken(result.data);
  }

  const handleLogout = () => {
    setToken(null);
  }
  
  const value = {
    token,
    onLogin: handleLogin,
    onLogout: handleLogout
  }
  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
}

export default AuthProvider;