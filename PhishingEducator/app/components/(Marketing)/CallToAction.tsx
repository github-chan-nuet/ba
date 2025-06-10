import { Button, tokens } from '@fluentui/react-components';
import CallToActionStyles from './CallToAction.module.scss';
import { ShieldTask28Filled } from '@fluentui/react-icons';
import { useOutletContext } from 'react-router';

type MarketingContext = {setAuthOpen: () => void};

export default function CallToAction() {
  const context = useOutletContext<MarketingContext>();

  return (
    <section
      className={CallToActionStyles.CallToAction}
      style={{
        backgroundColor: tokens.colorBrandBackground2,
      }}
    >
      <div className={CallToActionStyles.CallToAction__Container}>
        <h2 className={CallToActionStyles.CallToAction__Title}>Bereit dich zu Schützen?<br />Registriere dich jetzt kostenlos.</h2>
        <div className={CallToActionStyles.CallToAction__Actions}>
          <Button
            size="large"
            shape="circular"
            iconPosition="after"
            appearance="primary"
            icon={<ShieldTask28Filled />}
            onClick={context.setAuthOpen}
            style={{
              width: "15rem",
              height: "4rem"
            }}
          >
            Jetzt loslegen
          </Button>
        </div>
      </div>
    </section>
  )
}