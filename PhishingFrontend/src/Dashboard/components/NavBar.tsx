import { tokens } from "@fluentui/react-components";
import { Home24Regular, Lightbulb24Filled, MailWarning24Regular, Trophy24Regular } from "@fluentui/react-icons";
import { ReactElement } from "react";
import { Link } from "react-router";

function NavBar() {  
  return (
    <div
      style={{
        backgroundColor: tokens.colorNeutralBackground3,
        width: 95,
        height: '100%',
        paddingBlock: 24,
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'space-between'
      }}
    >
      <div
        style={{
          display: 'flex',
          flexDirection: 'column'
        }}
      >
        <NavBarItem
          label="Securaware"
          icon={<Home24Regular />}
          href="/dashboard"
        />
        <NavBarItem
          label="Phishing-Simulation"
          icon={<MailWarning24Regular />}
          href="/dashboard/phishing-simulation"
        />
        <NavBarItem
          label="Leaderboard"
          icon={<Trophy24Regular />}
          href="/dashboard/leaderboard"
        />
        <NavBarItem
          label="Online-Kurse"
          icon={<Lightbulb24Filled color={tokens.colorNeutralForeground2BrandSelected} />}
          href="/dashboard/courses"
        />
      </div>
      <div></div>
    </div>
  )
}

type NavBarItemProps = {
  href: string,
  label: string,
  icon: ReactElement
}

function NavBarItem({ href, label, icon }: NavBarItemProps) {
  return (
    <Link
      to={href}
      style={{
        display: 'flex',
        flexDirection: 'column',
        alignItems: 'center',
        padding: '8px 4px',
        gap: 2
      }}
    >
      {icon}
      <span style={{ textAlign: 'center' }}>{label}</span>
    </Link>
  )
}

export default NavBar;