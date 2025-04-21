import { Outlet } from "react-router";
import MarketingStyles from './Marketing.module.scss';
import { useState } from "react";
import { Body1Stronger, ToggleButton } from "@fluentui/react-components";
import AuthDrawer from "../../components/AuthDrawer";

function Layout() {
  const [isAuthDrawerOpen, setIsAuthDrawerOpen] = useState(false);

  return (
    <div className={MarketingStyles.Marketing}>
      <div className={MarketingStyles.Marketing__container}>
        <AuthDrawer isOpen={isAuthDrawerOpen} setIsOpen={setIsAuthDrawerOpen} />
        <header style={{
          display: 'flex',
          justifyContent: 'space-between',
          alignItems: 'center',
          paddingBlock: 24,
        }}>
          <Body1Stronger>Securaware</Body1Stronger>
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

export default Layout;