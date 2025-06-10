import { FluentProvider, webLightTheme } from "@fluentui/react-components"
import { Links, Meta, Outlet, Scripts, ScrollRestoration } from "react-router"
import GlobalToaster from "./utils/toaster/GlobalToaster"
import AuthProvider from "./utils/auth/AuthProvider"
import { client } from "./api/client.gen";

import "./styles/reset.scss";
import "./styles/app.scss";
import Loading from "@components/(Marketing)/Loading";

const token = typeof window !== "undefined" ? (window.localStorage.getItem('login-token') ?? undefined) : undefined;
client.setConfig({
  baseUrl: import.meta.env.VITE_API_BASE_URL,
  auth: token ? JSON.parse(token) : undefined
});

export default function App() {
  return (
    <AuthProvider>
      <Outlet />
      <GlobalToaster />
    </AuthProvider>
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
        <meta name="fluentui-insertion-point" content="fluentui-insertion-point" />
      </head>
      <body>
        <FluentProvider theme={webLightTheme}>
          {children}
        </FluentProvider>
        <ScrollRestoration />
        <Scripts />
      </body>
    </html>
  )
}

export function HydrateFallback() {
  return (
    <Loading />
  )
}