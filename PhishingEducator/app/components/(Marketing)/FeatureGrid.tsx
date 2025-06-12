import { cloneElement, type ReactElement } from 'react';
import { tokens } from '@fluentui/react-components';
import { FoodFishRegular, MailRegular, PenRegular, TrophyRegular } from '@fluentui/react-icons';

import FeatureGridStyles from './FeatureGrid.module.scss';

export default function FeatureGrid() {
  return (
    <section className={FeatureGridStyles.FeatureGrid}>
      <div className={FeatureGridStyles.FeatureGrid__Grid}>
        <h2 className={FeatureGridStyles.FeatureGrid__Title}>
          <strong>Alles</strong> was du für deine Sicherheit benötigst
        </h2>
        <dl className={FeatureGridStyles.FeatureGrid__Features}>
          <Feature
            icon={<FoodFishRegular />}
            title="Online Kurse"
            description="Eigne dir ein Grundwissen an, um die Eigenschaften und Erkennungsmerkmale von Phishing kennenzulernen."
          />
          <Feature
            icon={<MailRegular />}
            title="Phishing Simulation"
            description="Identifiziere Phishing E-Mails anhand praxisnaher Simulationen und vertiefe dein Wissen mit Tipps rund um die Erkennungsmerkmale."
          />
          <Feature
            icon={<PenRegular />}
            title="Teste dein Wissen"
            description="Stelle dein Wissen mit interaktiven Tests auf die Probe und sammle Erfahrung um Phishing E-Mails korrekt erkennen zu können."
          />
          <Feature
            icon={<TrophyRegular />}
            title="Spielerischer Vergleich"
            description="Vergleiche dich mit deinen Freunden und fordere Sie heraus. Sammle am meisten Erfahrungspunkte um an der Spitze zu bleiben."
          />
        </dl>
      </div>
    </section>
  )
}

type FeatureProps = {
  icon: ReactElement;
  title: string;
  description: string;
}

function Feature({ icon, title, description }: FeatureProps) {
  const coloredIcon = cloneElement(icon, {
    fontSize: "32px",
    color: tokens.colorBrandBackgroundInverted
  })

  return (
    <div>
        <dt className={FeatureGridStyles.Feature__Top}>
          <div className={FeatureGridStyles.Feature__Icon}>
            { coloredIcon }
          </div>
          { title }
        </dt>
        <dd className={FeatureGridStyles.Feature__Content}>
          { description }
        </dd>
    </div>
  )
}