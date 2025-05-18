import { useEffect, useState } from "react";
import { useNavigate } from "react-router";
import { useToaster } from "../toaster/useToaster";
import useLocalStorage from "../storage/useLocalStorage";
import { Toast, ToastTitle } from "@fluentui/react-components";
import AuthContext from "./AuthContext";
import { getUser, loginAndReturnJwtToken, type User } from "../../api";

export default function AuthProvider({ children }: { children: React.ReactNode }) {
  const { dispatchToast } = useToaster();
  const navigate = useNavigate();
  const [token, setToken] = useLocalStorage('login-token', null);
  const [user, setUser] = useState<User | null>(null);

  const addExperienceGain = (xpGain: number, newLevel: number|undefined) => {
    setUser(prev => {
      return {
        ...prev,
        totalExperience: (prev?.totalExperience ?? 0) + xpGain,
        level: newLevel ?? prev?.level
      };
    });
  };

  useEffect(() => {
    const fetchUser = async () => {
      if (!token) {
        setUser(null);
        return;
      }

      try {
        const userId = parseJwt(token).id;
        const response = await getUser({
          path: {
            userId,
          },
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        setUser(response.data || null);
      } catch (e) {
        console.error("Failed to load user profile", e);
        setUser(null);
        dispatchToast(
          <Toast>
            <ToastTitle>Failed to load user profile</ToastTitle>
          </Toast>
        );
      }
    };

    fetchUser();
  }, [token, dispatchToast]);

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
    setUser(null);
    navigate("/");
  }

  const value = {
    token,
    user,
    onExperienceGain: addExperienceGain,
    onLogin: handleLogin,
    onLogout: handleLogout
  };

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
}

// Helper to decode JWT
function parseJwt(token: string): { id: string } {
  const base64Url = token.split('.')[1];
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  const jsonPayload = decodeURIComponent(window.atob(base64).split('').map(c =>
    `%${('00' + c.charCodeAt(0).toString(16)).slice(-2)}`
  ).join(''));

  return JSON.parse(jsonPayload);
}