import { FluentProvider, webLightTheme } from '@fluentui/react-components';
import NavBar from './components/Navbar.tsx';
import { Outlet, useMatches } from 'react-router';

function Layout() {
  const matches = useMatches();
  console.log(matches);

  return (
    <FluentProvider theme={webLightTheme}>
      <div className="Dashboard__container">
        <NavBar />
        <div className="Dashboard__content">
          <Outlet />
        </div>
      </div>
    </FluentProvider>
  )
}

export default Layout;