import { tokens } from "@fluentui/react-components";
import * as FluentIcons from "@fluentui/react-icons";
import { NavLink } from "react-router";

import NavBarStyles from './NavBar.module.scss';
import useAuth from "../auth/useAuth";

function NavBar() { 
  const { onLogout } = useAuth();

  return (
    <div className={NavBarStyles.NavBar}>
      <div className={NavBarStyles.NavBar__items}>
        <NavBarItem
          label="Securaware"
          iconBaseName="Home"
          href="/dashboard"
        />
        <NavBarItem
          label="Phishing-Simulation"
          iconBaseName="MailWarning"
          href="/dashboard/phishing-simulation"
        />
        <NavBarItem
          label="Leaderboard"
          iconBaseName="Trophy"
          href="/dashboard/leaderboard"
        />
        <NavBarItem
          label="Online-Kurse"
          iconBaseName="Lightbulb"
          href="/dashboard/courses"
        />
      </div>
      <div className={NavBarStyles.NavBar__items}>
        <button className={NavBarStyles.NavBar__item} onClick={onLogout}>
          <FluentIcons.SignOut24Regular />
          <span className={NavBarStyles.NavBar__itemLabel}>Abmelden</span>
        </button>
      </div>
    </div>
  )
}

type FluentIconKey = keyof typeof FluentIcons;
type Regular24Keys = Extract<FluentIconKey, `${string}24Regular`>;
type Filled24Keys = Extract<FluentIconKey, `${string}24Filled`>;
type IconBaseName = Regular24Keys extends `${infer Base}24Regular` ? Base : never;

type NavBarItemProps = {
  href: string,
  label: string,
  iconBaseName: IconBaseName
}

function NavBarItem({ href, label, iconBaseName }: NavBarItemProps) {
  const iconKey = `${iconBaseName}24Regular` as Regular24Keys;
  const iconKeyActive = `${iconBaseName}24Filled` as Filled24Keys;
  const IconComponent = FluentIcons[iconKey];
  const IconActiveComponent = FluentIcons[iconKeyActive];

  return (
    <NavLink
      to={href}
      end
      className={NavBarStyles.NavBar__item}
    >
      {({ isActive }) => (
        <>
          {
            isActive ?
            <IconActiveComponent color={tokens.colorNeutralForeground2BrandSelected}  /> :
            <IconComponent  />
          }
          <span className={NavBarStyles.NavBar__itemLabel}>{label}</span>
        </>
      )}
    </NavLink>
  )
}

export default NavBar;