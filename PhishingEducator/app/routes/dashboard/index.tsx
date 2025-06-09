import { Title1, tokens } from "@fluentui/react-components";
import useAuth from "@utils/auth/useAuth"
import WelcomeBanner from "@components/WelcomeBanner";

export function meta() {
  return [
    { title: 'Securaware - Dashboard' },
    {
      name: 'description',
      content: 'Behalte mit dem Securaware-Dashboard deinen Fortschritt, empfohlene Kurse und Prüfungen im Blick - alles Wichtige rund um deine Sicherheit an einem Ort.'
    },
    {
      name: 'keywords',
      content: 'Securaware, Dashboard, Fortschritt, Sicherheitstraining, Phishing Schutz, Online Sicherheit, Lernstatus, Kursübersicht, Prüfungsübersicht, Cybertraining'
    }
  ]
}

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