import { FluentProvider, webLightTheme } from "@fluentui/react-components"
import { Links, Meta, Outlet, Scripts, ScrollRestoration } from "react-router"
import GlobalToaster from "./utils/toaster/GlobalToaster"
import AuthProvider from "./utils/auth/AuthProvider"
import { client } from "./api/client.gen";

import "./styles/reset.scss";
import "./styles/app.scss";

client.setConfig({
  baseUrl: import.meta.env.VITE_API_BASE_URL
})

export default function App() {
  return (
    <FluentProvider theme={webLightTheme}>
      <AuthProvider>
        <Outlet />
        <GlobalToaster />
      </AuthProvider>
    </FluentProvider>
  )
}

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="de">
      <head>
        <meta charSet="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <Meta />
        <Links />
      </head>
      <body>
        {children}
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  )
}

export function HydrateFallback() {
  return (
    <p>Lädt, bitte warten...</p>
  )
}