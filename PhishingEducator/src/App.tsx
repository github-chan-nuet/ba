import { createBrowserRouter, RouterProvider } from "react-router";
import MarketingLayout from './pages/Marketing/_Layout.tsx'
import DashboardLayout from './pages/Dashboard/_Layout.tsx'
import Courses from "./pages/Dashboard/Courses";
import Home from "./pages/Dashboard/Home";
import AuthProvider from "./auth/AuthProvider.tsx";
import { client } from "./api/client.gen.ts";
import { FluentProvider, webLightTheme } from "@fluentui/react-components";
import LandingPage from "./pages/Marketing/LandingPage.tsx";
import GlobalToaster from "./toaster/GlobalToaster.tsx";
import ProtectedRoute from "./auth/ProtectedRoute.tsx";
import CourseLesson from "./pages/Dashboard/Courses/CourseLesson.tsx";
import { courseData } from "./data/courses.tsx";

client.setConfig({
  baseUrl: import.meta.env.VITE_API_BASE_URL
})

const courseRoutes = courseData.flatMap(course =>
  course.lessons.map(lesson => ({
    element: <CourseLesson course={course} lesson={lesson} />,
    handle: course.label,
    path: `${course.handle}/${lesson.handle}`,
  })))

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
              ...courseRoutes
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