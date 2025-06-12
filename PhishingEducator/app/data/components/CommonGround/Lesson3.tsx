import { Card, Subtitle1, Title2 } from "@fluentui/react-components";
import { Globe20Filled, MailAlert20Filled, People20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson3 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>Wie verbreitet ist Phishing?</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Phishing ist eine der am weitesten verbreiteten Formen der Internetkriminalität. 
        Jeden Tag werden Millionen von Phishing-E-Mails, SMS und gefälschten Webseiten weltweit verschickt. 
        Sowohl Privatpersonen als auch Unternehmen sind betroffen.
      </p>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Der Grund für die hohe Verbreitung ist einfach: Phishing ist für Angreifer leicht durchzuführen, 
        kostengünstig und oft sehr erfolgreich. Schon ein einziger Klick kann zu einem finanziellen Schaden 
        oder Datenverlust führen.
      </p>

      <div className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <MailAlert20Filled color="#0EA5E9" />
              <Subtitle1 as="h3">Geringe Einstiegshürde für Angreifer</Subtitle1>
            </div>
            <p>
              Phishing-Kampagnen lassen sich mit wenig technischem Wissen und einfachen Mitteln durchführen.
              Es gibt sogar Baukastensysteme und Foren, die kriminelle Akteure unterstützen.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <People20Filled color="#D97706" />
              <Subtitle1 as="h3">Erfolgsquote durch menschliches Verhalten</Subtitle1>
            </div>
            <p>
              Phishing nutzt menschliche Schwächen wie Neugier, Stress oder Unwissenheit aus. 
              Viele Empfänger reagieren reflexartig auf vermeintlich wichtige oder dringende Nachrichten.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Globe20Filled color="#6366F1" />
              <Subtitle1 as="h3">Globale Reichweite durch digitale Kanäle</Subtitle1>
            </div>
            <p>
              Phishing kennt keine geografischen Grenzen. Eine einzige Kampagne kann in Sekunden weltweit
              verbreitet werden - per E-Mail, Social Media oder Messaging-Apps.
            </p>
          </div>
        </Card>
      </div>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Die hohe Verbreitung macht es umso wichtiger, dass sich jeder mit dem Thema Phishing auseinandersetzt.
        Durch Schulungen wie bei Securaware kann man lernen, Gefahren zu erkennen und sich effektiv zu schützen.
      </p>
    </article>
  );
}

export default Lesson3;