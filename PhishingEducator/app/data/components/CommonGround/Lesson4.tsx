import { Card, Subtitle1, Title2 } from "@fluentui/react-components";
import { BuildingBank20Filled, ContactCard20Filled, Mail20Filled, Money20Filled, ShieldError20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson4 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>Konsequenzen von Phishing</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Phishing kann schwerwiegende Folgen haben – sowohl für Einzelpersonen als auch für Unternehmen. 
        Die Konsequenzen reichen von finanziellen Verlusten über Identitätsdiebstahl bis hin zu erheblichen 
        Schäden für die Reputation und Sicherheit eines Unternehmens.
      </p>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Money20Filled color="#EF4444" />
              <Subtitle1 as="h3">Finanzielle Verluste</Subtitle1>
            </div>
            <p>
              Betroffene verlieren durch betrügerische Überweisungen, Einkäufe oder Kontoabbuchungen oft viel Geld.
              Unternehmen müssen nicht nur für entstandene Schäden aufkommen, sondern auch für Wiederherstellungsmaßnahmen.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <ContactCard20Filled color="#F59E0B" />
              <Subtitle1 as="h3">Identitätsdiebstahl</Subtitle1>
            </div>
            <p>
              Persönliche Daten wie Name, Adresse oder Sozialversicherungsnummer können missbraucht werden,
              um falsche Identitäten anzulegen oder betrügerische Handlungen vorzunehmen.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <ShieldError20Filled color="#10B981" />
              <Subtitle1 as="h3">Reputationsschäden</Subtitle1>
            </div>
            <p>
              Firmen, die von Phishing betroffen sind, verlieren oft das Vertrauen ihrer Kunden. 
              Ein einziger Vorfall kann langfristige Auswirkungen auf die Markenwahrnehmung haben.
            </p>
          </div>
        </Card>
      </section>

      <Title2 as="h2" className={Styles.Lesson__Title}>Beispiele für schwerwiegende Phishing-Attacken</Title2>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <BuildingBank20Filled color="#6366F1" />
              <Subtitle1 as="h3">2016: Angriff auf Hillary Clintons Wahlkampfleiter</Subtitle1>
            </div>
            <p>
              Der Wahlkampfleiter von Hillary Clinton, John Podesta, wurde Opfer einer Phishing-E-Mail. 
              Sein E-Mail-Konto wurde kompromittiert und sensible Wahlkampfinhalte wurden veröffentlicht. 
              Der Vorfall hatte große politische Auswirkungen.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Mail20Filled color="#EC4899" />
              <Subtitle1 as="h3">2020: Twitter-Hack</Subtitle1>
            </div>
            <p>
              Über Social Engineering und Phishing wurden Twitter-Mitarbeiter dazu gebracht, Zugangsdaten preiszugeben. 
              Dadurch konnten Angreifer Zugriff auf prominente Konten (z.B. Elon Musk, Barack Obama) erlangen und gefälschte Tweets verbreiten, 
              die zu Bitcoin-Betrug führten.
            </p>
          </div>
        </Card>
      </section>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Diese Beispiele zeigen, wie gefährlich Phishing sein kann. Securaware hilft dir, solche Angriffe zu erkennen
        und dich effektiv zu schützen.
      </p>
    </article>
  );
}

export default Lesson4;