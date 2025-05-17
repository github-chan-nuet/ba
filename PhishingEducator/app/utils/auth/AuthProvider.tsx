import { useNavigate } from "react-router";
import { useToaster } from "../toaster/useToaster";
import useLocalStorage from "../storage/useLocalStorage";
import { Toast, ToastTitle } from "@fluentui/react-components";
import AuthContext from "./AuthContext";
import { loginAndReturnJwtToken } from "../../api";

export default function AuthProvider({ children }: { children: React.ReactNode }) {
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
  };

  const handleLogout = () => {
    setToken(null);
    navigate("/");
  }

  const value = {
    token,
    onLogin: handleLogin,
    onLogout: handleLogout
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}