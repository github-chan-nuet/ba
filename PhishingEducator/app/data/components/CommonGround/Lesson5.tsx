import { Card, Subtitle1, Title2 } from "@fluentui/react-components"
import { Book20Filled, Briefcase20Filled, ContactCardGroup20Filled, Globe20Filled, Shield20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson5 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as={"h2"} className={Styles.Lesson__Title}>Warum Securaware?</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Die digitale Welt veränder sich rasant - wer sicher unterwegs sein möchte, muss sich laufend weiterbilden.
        Securaware bietet dir die Möglichkeit, dein Wissen rund um Cyber-Sicherheit gezielt auszubauen.
      </p>

      <div className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Shield20Filled color="#2563EB" />
              <Subtitle1 as="h3">Selbstschutz vor finanziellen Verlusten</Subtitle1>
            </div>
            <p>
              Wer Phishing erkennt, kann sich und sein Geld schützen. Securaware hilft dir, Risiken zu erkennen und richtig zu handeln.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Briefcase20Filled color="#10B981" />
              <Subtitle1 as="h3">Berufliche Verantwortung</Subtitle1>
            </div>
            <p>
              Im Berufsalltag kann ein Klick auf einen Phishing-Link großen Schaden anrichten. Securaware sensibilisiert und stärkt dich im Arbeitsumfeld.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <ContactCardGroup20Filled color="#F59E0B" />
              <Subtitle1 as="h3">Schutz der Familie und Angehörige</Subtitle1>
            </div>
            <p>
              Ältere Menschen oder Kinder sind besonders gefährdet. Dein Wissen kann deine Familie schützen - Securaware zeigt dir wie.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Book20Filled color="#6366F1" />
              <Subtitle1 as="h3">Karrierevorteile durch digitale Kompetenz</Subtitle1>
            </div>
            <p>
              Digitale Mündigkeit wird immer wichtiger. Wer sich mit Sicherheit im Netz auskennt, sammelt Pluspunkte im Beruf.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Globe20Filled color="#EC4899" />
              <Subtitle1 as="h3">Mithelfen, die Gesellschaft zu schützen</Subtitle1>
            </div>
            <p>
              Phishing ist ein gesellschaftliches Problem. Wer sich weiterbildet, kann andere aufklären und so zur digitalen Sicherheit aller beitragen.
            </p>
          </div>
        </Card>
      </div>
    </article>
  )
}

export default Lesson5;