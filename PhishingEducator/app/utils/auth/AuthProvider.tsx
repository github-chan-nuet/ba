import { useCallback, useEffect, useState } from "react";
import { useNavigate } from "react-router";
import useLocalStorage from "../storage/useLocalStorage";
import AuthContext from "./AuthContext";
import { getUser, loginAndReturnJwtToken, updateUser as patchUser, type User, type UserPatchModel } from "../../api";
import { client } from "../../api/client.gen";

export default function AuthProvider({ children }: { children: React.ReactNode }) {
  const navigate = useNavigate();
  const [token, setToken] = useLocalStorage('login-token', null);
  const [user, setUser] = useState<(User & { totalXpForNextLevel: number; id?: string | undefined; } ) | null>(null);

  const refreshUser = useCallback(async () => {
    if (!token) return;

    try {
      const userId = parseJwt(token).id;

      const { data } = await getUser({
        path: {
          userId
        }
      });

      if (!data) throw new Error('Kein Profil gefunden');
      setUser({
        ...data,
        totalXpForNextLevel: calculateTotalXpForNextLevel(data.level ?? 0),
        id: userId
      });
    } catch (e) {
      setToken(null);
      setUser(null);
      console.error(e);
    }
  }, [setToken, token]);

  useEffect(() => {
    client.setConfig({
      auth: token
    });
    if (token) void refreshUser();
  }, [token, refreshUser]);

  const login = async (email: string, password: string) => {
    const { data, error } = await loginAndReturnJwtToken({ body: { email, password } });
    if (error) {
      throw error;
    }
    if (data) {
      setToken(data);
      navigate("/dashboard");
    }
  };

  const logout = () => {
    setToken(null);
    setUser(null);
    navigate("/");
  }


  const updateUser = async (userPatch: UserPatchModel) => {
    if (!token) return;

    const userId = parseJwt(token).id;

    const { error } = await patchUser({
      path: {
        userId
      },
      body: userPatch
    });
    if (error) throw error;
    await refreshUser();
  }

  const addExperienceGain = (xpGain: number, newLevel: number|undefined) => {
    setUser(prev => {
      const level = newLevel ?? prev?.level ?? 0;

      return {
        ...prev,
        totalExperience: (prev?.totalExperience ?? 0) + xpGain,
        level,

        // level is calculated by the formula: 1 + ln(x/200 + 1) / ln(1.5) = level
        totalXpForNextLevel: calculateTotalXpForNextLevel(level)
      };
    });
  };

  const calculateTotalXpForNextLevel = (currentLevel: number) => {
    return Math.ceil(200 * ((1.5 ** currentLevel) - 1));
  }

  const value = {
    token,
    user,
    updateUser,
    onExperienceGain: addExperienceGain,
    onLogin: login,
    onLogout: logout
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