import { Title1, tokens } from "@fluentui/react-components";
import useAuth from "../../utils/auth/useAuth"
import WelcomeBanner from "../../components/WelcomeBanner";

export default function DashboardHome() {
  const { user } = useAuth();

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        gap: tokens.spacingVerticalXXL
      }}
    >
      <Title1>Schön dich zu sehen, {user?.firstname}!</Title1>
      <WelcomeBanner />
    </div>
  )
}