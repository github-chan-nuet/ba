import type { Route } from "./+types/root";
import { isRouteErrorResponse, Links, Meta, Outlet, Scripts, ScrollRestoration } from "react-router"
import { HelmetProvider } from "react-helmet-async";
import GlobalToaster from "@utils/toaster/GlobalToaster"
import AuthProvider from "@utils/auth/AuthProvider"
import { client } from "@api/client.gen";
import { FluentProvider, webLightTheme } from "@fluentui/react-components"
import Loading from "@components/(Marketing)/Loading";

import "./styles/reset.scss";
import "./styles/app.scss";

const token = typeof window !== "undefined" ? (window.localStorage.getItem('login-token') ?? undefined) : undefined;
client.setConfig({
  baseUrl: import.meta.env.VITE_API_BASE_URL,
  auth: token ? JSON.parse(token) : undefined
});

export default function App() {
  return (
    <HelmetProvider>
      <AuthProvider>
        <Outlet />
        <GlobalToaster />
      </AuthProvider>
    </HelmetProvider>
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

export function ErrorBoundary({
  error
}: Route.ErrorBoundaryProps) {
  if (isRouteErrorResponse(error)) {
    return (
      <>
        <h1>{error.status} {error.statusText}</h1>
        <p>{error.data}</p>
      </>
    )
  } else if (error instanceof Error) {
    return (
      <>
        <h1>Error</h1>
        <p>{error.message}</p>
        <p>The stack trace is:</p>
        <pre>{error.stack}</pre>
      </>
    )
  } else {
    return <h1>Unknown Error</h1>;
  }
}