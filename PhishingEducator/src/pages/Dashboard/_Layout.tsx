import NavBar from '../../components/NavBar.tsx';
import { Outlet } from 'react-router';
import Breadcrumbs from '../../components/Breadrumbs.tsx';
import UserAvatar from '../../components/UserAvatar.tsx';

function Layout() {
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

export default Layout;