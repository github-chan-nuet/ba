import useAuth from "../../utils/auth/useAuth";
import { Navigate, Outlet } from "react-router";

export default function ProtectedRoute() {
  const { token } = useAuth();

  if (!token) {
    return <Navigate to="/" replace />
  }

  return <Outlet />
}