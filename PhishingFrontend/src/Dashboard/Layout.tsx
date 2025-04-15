import { FluentProvider, webLightTheme } from '@fluentui/react-components';
import NavBar from './components/Navbar.tsx';
import { Outlet } from 'react-router';
import Breadcrumbs from './components/Breadrumbs.tsx';

function Layout() {
  return (
    <FluentProvider theme={webLightTheme}>
      <div className="Dashboard__container">
        <NavBar />
        <div className="Dashboard__content">
          <div className="Dashboard__header">
            <Breadcrumbs />
          </div>
          <Outlet />
        </div>
      </div>
    </FluentProvider>
  )
}

export default Layout;