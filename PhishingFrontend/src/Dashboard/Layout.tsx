import { FluentProvider, webLightTheme } from '@fluentui/react-components';
import NavBar from './components/Navbar.tsx';

function Layout() {
  return (
    <FluentProvider theme={webLightTheme}>
      <div className="Dashboard__container">
        <NavBar />
        <div></div>
      </div>
    </FluentProvider>
  )
}

export default Layout;