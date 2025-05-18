import { Title1 } from "@fluentui/react-components";
import useAuth from "../../utils/auth/useAuth"

export default function DashboardHome() {
  const { user } = useAuth();

  return (
    <Title1>Schön dich zu sehen, {user?.firstname}!</Title1>
  )
}