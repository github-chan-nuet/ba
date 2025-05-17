import { useEffect } from "react";
import useAuth from "../../utils/auth/useAuth";
import { Outlet, useNavigate } from "react-router";

export default function ProtectedRoute() {
  const { token } = useAuth();
  const navigate = useNavigate();

  useEffect(() => {
    if (!token) {
      navigate("/");
    }
  }, [navigate, token])

  return token && (
    <Outlet />
  )
}