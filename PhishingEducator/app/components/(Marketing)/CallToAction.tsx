import { useOutletContext } from 'react-router';
import { Button } from '@fluentui/react-components';
import { ShieldTask28Filled } from '@fluentui/react-icons';

import CallToActionStyles from './CallToAction.module.scss';

type MarketingContext = {setAuthOpen: () => void};

export default function CallToAction() {
  const context = useOutletContext<MarketingContext>();

  return (
    <section className={CallToActionStyles.CallToAction}>
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
            className={CallToActionStyles.CallToAction__Button}
          >
            Jetzt loslegen
          </Button>
        </div>
      </div>
    </section>
  )
}