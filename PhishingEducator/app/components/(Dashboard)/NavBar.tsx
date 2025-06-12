import { NavLink } from "react-router";
import { createFocusOutlineStyle, makeStyles, tokens } from "@fluentui/react-components";
import * as FluentIcons from "@fluentui/react-icons";
import useAuth from "@utils/auth/useAuth";

import NavBarStyles from "./NavBar.module.scss";

const useStyles = makeStyles({
  focusIndicator: createFocusOutlineStyle()
});

export default function NavBar() {
  const { onLogout } = useAuth();
  const styles = useStyles();

  return (
    <aside className={NavBarStyles.NavBar}>
      <nav className={NavBarStyles.NavBar__Items}>
        <ul>
          <NavBarItem
            label="Securaware"
            iconBaseName="Home"
            href="/dashboard"
            matchOnSubPages={false}
          />
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
        </ul>
      </nav>
      <div className={NavBarStyles.NavBar__Items}>
        <button className={`${NavBarStyles.NavBar__Item} ${styles.focusIndicator}`} onClick={onLogout}>
          <FluentIcons.SignOut24Regular />
          <span className={NavBarStyles.NavBar__ItemLabel}>Abmelden</span>
        </button>
      </div>
    </aside>
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

function NavBarItem({ href, label, iconBaseName, matchOnSubPages = true }: NavBarItemProps) {
  const iconKey = `${iconBaseName}24Regular` as Regular24Keys;
  const iconKeyActive = `${iconBaseName}24Filled` as Filled24Keys;
  const IconComponent = FluentIcons[iconKey];
  const IconActiveComponent = FluentIcons[iconKeyActive];

  const styles = useStyles();

  return (
    <li>
      <NavLink
        to={href}
        end={!matchOnSubPages}
        className={`${NavBarStyles.NavBar__Item} ${styles.focusIndicator}`}
      >
        {({ isActive }) => (
          <>
            { isActive ? (
              <IconActiveComponent color={tokens.colorNeutralForeground2BrandSelected} />
            ) : (
              <IconComponent />
            )}
            <span className={NavBarStyles.NavBar__ItemLabel}>{label}</span>
          </>
        )}
      </NavLink>
    </li>
  )
}