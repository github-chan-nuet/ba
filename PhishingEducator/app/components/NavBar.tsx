import useAuth from "../utils/auth/useAuth";
import * as FluentIcons from "@fluentui/react-icons";

import NavBarStyles from "../styles/NavBar.module.scss";
import { NavLink } from "react-router";
import { createFocusOutlineStyle, makeStyles, tokens } from "@fluentui/react-components";

export default function NavBar() {
  const { onLogout } = useAuth();

  return (
    <div className={NavBarStyles.NavBar}>
      <div className={NavBarStyles.NavBar__items}>
        <NavBarItem
          label="Securaware"
          iconBaseName="Home"
          href="/dashboard"
          matchOnSubPages={false}
        />
        { /*
        <NavBarItem
          label="Phishing-Simulation"
          iconBaseName="MailWarning"
          href="/dashboard/phishing-simulation"
        />
        */ }
        { /*
        <NavBarItem
          label="Leaderboard"
          iconBaseName="Trophy"
          href="/dashboard/leaderboard"
        />
        */ }
        <NavBarItem
          label="Online-Kurse"
          iconBaseName="Lightbulb"
          href="/dashboard/courses"
        />
        <NavBarItem
          label="Prüfungen"
          iconBaseName="BookQuestionMark"
          href="/dashboard/exams"
        />
      </div>
      <div className={NavBarStyles.NavBar__items}>
        <button className={NavBarStyles.NavBar__item} onClick={onLogout}>
          <FluentIcons.SignOut24Regular />
          <span className={NavBarStyles.NavBar__itemLabel}>Abmelden Test</span>
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
  iconBaseName: IconBaseName,
  matchOnSubPages?: boolean
};

const useStyles = makeStyles({
  focusIndicator: createFocusOutlineStyle({
    selector: 'focus-within'
  })
});

function NavBarItem({ href, label, iconBaseName, matchOnSubPages = true }: NavBarItemProps) {
  const iconKey = `${iconBaseName}24Regular` as Regular24Keys;
  const iconKeyActive = `${iconBaseName}24Filled` as Filled24Keys;
  const IconComponent = FluentIcons[iconKey];
  const IconActiveComponent = FluentIcons[iconKeyActive];

  const styles = useStyles();

  return (
    <NavLink
      to={href}
      end={!matchOnSubPages}
      className={`${NavBarStyles.NavBar__item} ${styles.focusIndicator}`}
    >
      {({ isActive }) => (
        <>
          { isActive ? (
            <IconActiveComponent color={tokens.colorNeutralForeground2BrandSelected} />
          ) : (
            <IconComponent />
          )}
          <span className={NavBarStyles.NavBar__itemLabel}>{label}</span>
        </>
      )}
    </NavLink>
  )
}