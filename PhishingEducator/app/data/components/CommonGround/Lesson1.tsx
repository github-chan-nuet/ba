import { Card, Subtitle1, Title2 } from "@fluentui/react-components";
import { Shield20Filled, Warning20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson1 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>Schadensausmass von Phishing</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Phishing ist eine Form des Cyberangriffs, bei dem Angreifer versuchen,
        sensible Informationen wie Passwörter, Kreditkartendaten oder Zugangsdaten
        zu Online-Konten zu stehlen. Die Auswirkungen solcher Angriffe können erheblich sein.
      </p>

      <div className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Warning20Filled color="#DC2626" />
              <Subtitle1 as="h3">Finanzieller Schaden</Subtitle1>
            </div>
            <p>
              Betroffene können direkt Geld verlieren, etwa durch betrügerische Überweisungen
              oder Kreditkartenmissbrauch. Auch Unternehmen erleiden oft hohe Kosten
              durch Betrug, Wiederherstellung und Imageverlust.
            </p>
          </div>
        </Card>

        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Shield20Filled color="#CA8A04" />
              <Subtitle1 as="h3">Verlust von Daten und Vertrauen</Subtitle1>
            </div>
            <p>
              Neben finanziellen Verlusten kann auch ein erheblicher Reputationsschaden entstehen.
              Kunden verlieren Vertrauen in ein Unternehmen, wenn deren Daten nicht sicher sind.
              In manchen Fällen ist auch der dauerhafte Verlust von Daten möglich.
            </p>
          </div>
        </Card>
      </div>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Die Bekämpfung von Phishing erfordert Aufklärung, technische Schutzmaßnahmen
        sowie ein wachsames Verhalten der Nutzer.
      </p>
    </article>
  );
}

export default Lesson1;