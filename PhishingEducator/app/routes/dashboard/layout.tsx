import UserAvatar from "../../components/UserAvatar";
import Breadcrumbs from "../../components/Breadcrumbs";
import NavBar from "../../components/NavBar";
import { Outlet } from "react-router";

export const handle = "Securaware";

export default function DashboardLayout() {
  return (
    <div className="Dashboard__container | container">
      <NavBar />
      <div className="Dashboard__content">
        <div className="Dashboard__header">
          <Breadcrumbs />
          <UserAvatar />
        </div>
        <Outlet />
      </div>
    </div>
  )
}