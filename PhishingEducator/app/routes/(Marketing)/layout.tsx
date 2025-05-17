import { useState } from "react";

import MarketingStyles from "../../styles/Marketing.module.scss";
import AuthDrawer from "../../components/AuthDrawer";
import { Body1Stronger, ToggleButton } from "@fluentui/react-components";
import { Outlet } from "react-router";

export default function MarketingLayout() {
  const [isAuthDrawerOpen, setIsAuthDrawerOpen] = useState(false);

  return (
    <div className={MarketingStyles.Marketing}>
      <div className={MarketingStyles.Marketing__container}>
        <AuthDrawer isOpen={isAuthDrawerOpen} setIsOpen={setIsAuthDrawerOpen} />
        <header style={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          paddingBlock: 24
        }}>
          <Body1Stronger style={{ fontSize: '1.5rem' }}>Securaware</Body1Stronger>
          <ToggleButton
            appearance="primary"
            onClick={() => setIsAuthDrawerOpen(!isAuthDrawerOpen)}
            checked={isAuthDrawerOpen}
          >
            Login
          </ToggleButton>
        </header>
        <main>
          <Outlet />
        </main>
      </div>
    </div>
  )
}