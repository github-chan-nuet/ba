import { Outlet } from "react-router";
import UserAvatar from "@components/(Dashboard)/UserAvatar";
import Breadcrumbs from "@components/(Dashboard)/Breadcrumbs";
import NavBar from "@components/NavBar";

import DashboardStyles from '@styles/Dashboard.module.scss';

export const handle = "Securaware";

export default function DashboardLayout() {
  return (
    <div className={DashboardStyles.Dashboard__Container}>
      <NavBar />
      <div className={DashboardStyles.Dashboard__Main}>
        <div className={DashboardStyles.Dashboard__Header}>
          <Breadcrumbs />
          <UserAvatar />
        </div>
        <div className={DashboardStyles.Dashboard__Content}>
          <Outlet />
        </div>
      </div>
    </div>
  )
}