import { tokens } from "@fluentui/react-components";
import { Home24Regular, Lightbulb24Filled, MailWarning24Regular, Trophy24Regular } from "@fluentui/react-icons";
import { ReactElement } from "react";
import { Link } from "react-router";

import NavBarStyles from './NavBar.module.scss';

function NavBar() {  
  return (
    <div className={NavBarStyles.NavBar}>
      <div className={NavBarStyles.NavBar__items}>
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
      className={NavBarStyles.NavBar__item}
    >
      {icon}
      <span className={NavBarStyles.NavBar__itemLabel}>{label}</span>
    </Link>
  )
}

export default NavBar;