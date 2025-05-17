import { Breadcrumb, BreadcrumbButton, BreadcrumbDivider, BreadcrumbItem } from "@fluentui/react-components";
import React, { useEffect, useState } from "react";
import { useMatches, useNavigate } from "react-router"

export default function Breadcrumbs() {
  const matches = useMatches();
  const navigate = useNavigate();

  const [breadcrumbItems, setBreadcrumbItems] = useState<
    { pathname: string; label: string }[]
  >([]);

  useEffect(() => {
    let isMounted = true;

    const resolveBreadcrumbs = async () => {
      const resolved = await Promise.all(
        matches.map(async (match) => {
          let label = "";
          if (typeof match.handle === "function") {
            label = await match.handle({ params: match.params });
          } else if (typeof match.handle === "string") {
            label = match.handle;
          }
          return { pathname: match.pathname, label };
        })
      );

      // Filter out empty labels
      const filtered = resolved.filter((item) => !!item.label);

      if (isMounted) {
        setBreadcrumbItems(filtered);
      }
    };

    resolveBreadcrumbs();

    return () => {
      isMounted = false;
    }
  }, [matches])

  return (
    <Breadcrumb>
      { breadcrumbItems.map((item, idx) => {
        return (
          <React.Fragment key={idx}>
            { idx > 0 && <BreadcrumbDivider /> }
            <BreadcrumbItem>
              <BreadcrumbButton
                as="button"
                onClick={() => navigate(item.pathname)}
                current={idx === breadcrumbItems.length - 1}
              >
                {item.label}
              </BreadcrumbButton>
            </BreadcrumbItem>
          </React.Fragment>
        )
      })}
    </Breadcrumb>
  )
}