import { index, layout, route, type RouteConfig } from "@react-router/dev/routes";

export default [
  layout("routes/(Marketing)/layout.tsx", [
    index("routes/(Marketing)/index.tsx")
  ]),
  layout("routes/(Utils)/protected.tsx", [
    route("dashboard", "routes/dashboard/layout.tsx", [
      index("routes/dashboard/index.tsx"),
      route("courses", "routes/dashboard/courses/layout.tsx", [
        index("routes/dashboard/courses/index.tsx"),
        route(":courseHandle/:lessonHandle", "routes/dashboard/courses/lesson.tsx")
      ])
    ])
  ])
] satisfies RouteConfig;