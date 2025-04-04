import { Button, FluentProvider, webLightTheme } from '@fluentui/react-components';

function Layout() {
  return (
    <FluentProvider theme={webLightTheme}>
      { /* Navbar */ }
      { /* Content */ }
      <Button appearance='primary'>Hallo</Button>
    </FluentProvider>
  )
}

export default Layout;