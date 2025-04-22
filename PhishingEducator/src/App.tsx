import { createBrowserRouter, RouterProvider } from "react-router";
import MarketingLayout from './pages/Marketing/_Layout.tsx'
import DashboardLayout from './pages/Dashboard/_Layout.tsx'
import Courses from "./pages/Dashboard/Courses";
import Home from "./pages/Dashboard/Home";
import Lesson from "./pages/Dashboard/Lesson";
import AuthProvider from "./auth/AuthProvider.tsx";
import { client } from "./api/client.gen.ts";
import { FluentProvider, webLightTheme } from "@fluentui/react-components";
import LandingPage from "./pages/Marketing/LandingPage.tsx";
import GlobalToaster from "./toaster/GlobalToaster.tsx";
import ProtectedRoute from "./auth/ProtectedRoute.tsx";

client.setConfig({
  baseUrl: "http://localhost:8080/api"
})

const router = createBrowserRouter([
  {
    element: <AuthProvider />,
    children: [
      {
        path: '/',
        element: <MarketingLayout />,
        children: [
          { index: true, element: <LandingPage /> },
        ]
      },
      {
        path: 'dashboard',
        element:
          <ProtectedRoute>
            <DashboardLayout />
          </ProtectedRoute>,
        handle: 'Securaware',
        children: [
          { index: true, element: <Home /> },
          {
            path: 'courses',
            handle: 'Online-Kurse',
            children: [
              { index: true, element: <Courses /> },
              { path: ':course/lectures/:lecture', element: <Lesson /> }
            ]
          }
        ]
      }
    ]
  }
])

const App = () => {
  return (
    <FluentProvider theme={webLightTheme}>
      <RouterProvider router={router} />
      <GlobalToaster />
    </FluentProvider>
  )
}

export default App;