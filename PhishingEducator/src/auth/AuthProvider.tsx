import AuthContext from "./AuthContext";
import { loginAndReturnJwtToken } from "../api";
import useLocalStorage from "./useLocalStorage";
import { Outlet, useNavigate } from "react-router";

const AuthProvider = () => {
  const navigate = useNavigate();
  const [token, setToken] = useLocalStorage('login-token', null);

  const handleLogin = async (email: string, password: string) => {
    const result = await loginAndReturnJwtToken({ body: { email, password } });
    setToken(result.data);
    navigate("/dashboard");
  }

  const handleLogout = () => {
    setToken(null);
    navigate("/");
  }
  
  const value = {
    token,
    onLogin: handleLogin,
    onLogout: handleLogout
  }
  return (
    <AuthContext.Provider value={value}>
      <Outlet />
    </AuthContext.Provider>
  );
}

export default AuthProvider;