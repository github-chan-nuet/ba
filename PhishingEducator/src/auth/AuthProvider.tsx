import AuthContext from "./AuthContext";
import { loginAndReturnJwtToken } from "../api";
import useLocalStorage from "./useLocalStorage";
import { Outlet, useNavigate } from "react-router";
import { useToaster } from "../toaster/useToaster";
import { Toast, ToastTitle } from "@fluentui/react-components";

const AuthProvider = () => {
  const { dispatchToast } = useToaster();
  const navigate = useNavigate();
  const [token, setToken] = useLocalStorage('login-token', null);

  const handleLogin = async (email: string, password: string) => {
    const { data, error } = await loginAndReturnJwtToken({ body: { email, password } });
    if (!error) {
      setToken(data);
      navigate("/dashboard");
    } else if (error.title) {
      dispatchToast(
        <Toast>
          <ToastTitle>{error.title}</ToastTitle>
        </Toast>
      );
    }
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