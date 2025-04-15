import {
  createBrowserRouter,
  RouterProvider
} from 'react-router'

import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './reset.scss'
import './index.scss'
import MarketingLayout from './Marketing/Layout.tsx'
import DashboardLayout from './Dashboard/Layout.tsx'
import Home from './Dashboard/Home.tsx'
import Courses from './Dashboard/Courses.tsx'
import Lesson from './Dashboard/Lesson.tsx'

const router = createBrowserRouter([
  {
    path: '',
    element: <MarketingLayout />
  },
  {
    path: 'dashboard',
    element: <DashboardLayout />,
    handle: 'Securaware',
    children: [
      { index: true, element: <Home /> },
      {
        path: "courses",
        handle: 'Online-Kurse',
        children: [
          { index: true, element: <Courses /> },
          { path: ":course/lectures/:lecture", element: <Lesson /> }
        ]
      }
    ]
  },
])

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
