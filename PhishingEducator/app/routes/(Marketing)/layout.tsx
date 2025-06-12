import { useState } from "react";
import { Link, Outlet } from "react-router";
import { Body1Stronger, ToggleButton } from "@fluentui/react-components";
import AuthDrawer from "@components/(Marketing)/AuthDrawer";
import Footer from "@components/(Marketing)/Footer";

import MarketingStyles from "@styles/Marketing.module.scss";
import logo from "@assets/images/securaware.png";

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
      <main className={MarketingStyles.Marketing__Container}>
        <Outlet context={{ setAuthOpen: setIsAuthDrawerOpen }} />
      </main>
      <Footer />
    </div>
  )
}