import MarketingFooterStyles from './Footer.module.scss';

export default function MarketingFooter() {
  const startYear = 2025;
  const currentYear = new Date().getFullYear();
  const yearDisplay = currentYear > startYear ? `${startYear} - ${currentYear}` : `${startYear}`;

  return (
    <footer className={MarketingFooterStyles.Footer}>
      <div className={MarketingFooterStyles.Footer__Container}>
        <p className={MarketingFooterStyles.Footer__Copyright}>
          &copy; { yearDisplay } Patrick Scheidegger & Mischa Binder - Alle Rechte vorbehalten.
        </p>
      </div>
    </footer>
  )
}