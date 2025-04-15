import { Link } from "react-router";

function Layout() {
  return (
    <div style={{
      backgroundColor: 'oklch(0.985 0.002 247.839)',
      color: 'oklch(0.278 0.033 256.848)',
      fontFamily: "ui-sans-serif, system-ui, sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol', 'Noto Color Emoji'"
    }}>
      <header style={{
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: 24,
      }}>
        <div>Securaware</div>
        <nav>
          <Link to="#">Solutions</Link>
          <Link to="#">Pricing</Link>
          <Link to="#">Resources</Link>
        </nav>
        <Link to="#contact">Contact Us</Link>
      </header>
    </div>
  )
}

export default Layout;