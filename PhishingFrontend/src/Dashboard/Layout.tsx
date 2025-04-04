import { FluentProvider, webLightTheme } from '@fluentui/react-components';
import NavBar from './components/Navbar.tsx';

function Layout() {
  return (
    <FluentProvider theme={webLightTheme}>
      <div
        style={{
          display: 'flex',
          height: '100vh',
          overflow: 'hidden'
        }}
      >
        <NavBar />
        <div></div>
      </div>
    </FluentProvider>
  )
}

export default Layout;