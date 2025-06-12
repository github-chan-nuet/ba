import useAuth from "@utils/auth/useAuth"
import { Title1 } from "@fluentui/react-components";
import WelcomeBanner from "@components/(Dashboard)/WelcomeBanner";

import DashboardStyles from '@styles/Dashboard.module.scss';

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
    <main className={DashboardStyles.DashboardHome}>
      <Title1>Schön dich zu sehen, {user?.firstname}!</Title1>
      <WelcomeBanner />
    </main>
  )
}