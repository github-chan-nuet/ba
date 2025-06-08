import { useState } from "react";
import { Link, Outlet } from "react-router";
import AuthDrawer from "@components/AuthDrawer";
import MarketingStyles from "@styles/Marketing.module.scss";
import logo from "@assets/images/securaware.png";
import { Body1Stronger, ToggleButton } from "@fluentui/react-components";

export default function MarketingLayout() {
  const [isAuthDrawerOpen, setIsAuthDrawerOpen] = useState(false);
  return (
    <div className={MarketingStyles.Marketing}>
      <AuthDrawer isOpen={isAuthDrawerOpen} setIsOpen={setIsAuthDrawerOpen} />
      <header className={MarketingStyles.Marketing__Header}>
        <Link to="/" className={MarketingStyles.Marketing__LogoContainer}>
          <img src={logo} alt="" width={20} />
          <Body1Stronger className={MarketingStyles.Marketing__LogoText}>Securaware</Body1Stronger>
        </Link>
        <ToggleButton
          appearance="primary"
          onClick={() => setIsAuthDrawerOpen(!isAuthDrawerOpen)}
          checked={isAuthDrawerOpen}
        >
          Login
        </ToggleButton>
      </header>
      <div className={MarketingStyles.Marketing__Container}>
        <main className={MarketingStyles.Marketing__Content}>
          <Outlet context={{ setAuthOpen: setIsAuthDrawerOpen }} />
        </main>
      </div>
    </div>
  )
}