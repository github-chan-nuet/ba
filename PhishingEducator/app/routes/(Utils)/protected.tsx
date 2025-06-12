import { Navigate, Outlet } from "react-router";
import useAuth from "@utils/auth/useAuth";

export default function ProtectedRoute() {
  const { token } = useAuth();

  if (!token) {
    return <Navigate to="/" replace />
  }

  return <Outlet />
}