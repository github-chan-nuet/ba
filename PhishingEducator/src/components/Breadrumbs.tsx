import { Breadcrumb, BreadcrumbButton, BreadcrumbItem, BreadcrumbDivider } from "@fluentui/react-components";
import { useMatches } from "react-router";

function Breadcrumbs() {
  const matches  = useMatches();
  
  let breadcrumbs = 0;

  return (
    <Breadcrumb>
      {
        matches.map(match => {
          if (!match.handle) return;

          breadcrumbs++;
          return (
            <>
              { breadcrumbs > 1 ? <BreadcrumbDivider /> : '' }
              <BreadcrumbItem>
                <BreadcrumbButton href={match.pathname}>{match.handle as string}</BreadcrumbButton>
              </BreadcrumbItem>
            </>
          )
        })
      }
    </Breadcrumb>
  )
}

export default Breadcrumbs;