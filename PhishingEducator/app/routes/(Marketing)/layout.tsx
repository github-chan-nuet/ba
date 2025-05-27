import { useState } from "react";

import MarketingStyles from "../../styles/Marketing.module.scss";
import AuthDrawer from "../../components/AuthDrawer";
import { Body1Stronger, ToggleButton } from "@fluentui/react-components";
import { Outlet } from "react-router";
import logo from '../../assets/images/securaware.png';

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
          marginLeft: "1rem",
          marginRight: "1rem",
          paddingBlock: 24
        }}>
          <div style={{
            display: 'flex',
            gap: '0.5rem',
          }}>
            <img src={logo} alt="" width={20} />
            <Body1Stronger style={{ fontSize: '1.5rem' }}>
              Securaware
            </Body1Stronger>
          </div>
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