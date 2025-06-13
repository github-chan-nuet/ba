import { index, layout, route, type RouteConfig } from "@react-router/dev/routes";

export default [
  layout("routes/(Marketing)/layout.tsx", [
    index("routes/(Marketing)/index.tsx"),
    route("phishing-simulation", "routes/(Marketing)/phishing-simulation.tsx"),

    // Catch-all route (404)
    route("*", "routes/(Marketing)/not-found.tsx")
  ]),
  layout("routes/(Utils)/protected.tsx", [
    route("dashboard", "routes/dashboard/layout.tsx", [
      index("routes/dashboard/index.tsx"),
      route("courses", "routes/dashboard/courses/layout.tsx", [
        index("routes/dashboard/courses/index.tsx"),
        route(":courseHandle/:lessonHandle", "routes/dashboard/courses/lesson.tsx")
      ]),
      route("exams", "routes/dashboard/exams/layout.tsx", [
        index("routes/dashboard/exams/index.tsx"),
        route(":examId", "routes/dashboard/exams/exam.tsx")
      ])
    ])
  ])
] satisfies RouteConfig;