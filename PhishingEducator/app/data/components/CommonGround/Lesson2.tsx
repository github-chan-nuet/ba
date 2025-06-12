import { Card, Subtitle1, Title2 } from "@fluentui/react-components"
import { Link20Filled, PersonAlert20Filled, TextField20Filled, Warning20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson2 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>Was ist Phishing?</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Unter dem Begriff <strong>Phishing</strong> versteht man eine Methode des digitalen Betrugs.
        Angreifer geben sich als vertrauenswürdige Quelle aus - zum Beispiel eine Bank, ein Kollege oder eine Behörde -
        um an vertrauliche Informationen wie Passwörter, Kreditkartendaten oder Zugangsdaten zu gelangen.
      </p>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Ziel ist es, durch Täuschung Personen dazu zu bringen, sensible Daten preiszugeben. Phishing zählt zu den
        häufigsten Cyberangriffsformen weltweit.
      </p>

      <div className={`${Styles.Lesson__Section} ${Styles.Lesson__InfoCardGrid}`}>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <PersonAlert20Filled color="#0EA5E9" />
              <Subtitle1 as="h3">Täuschung</Subtitle1>
            </div>
            <p>
              Die Täter wirken oft täuschend echt. Sie imitieren bekannte Absender wie Banken, Online-Shops oder Kollegen. 
              Die Kontaktaufnahme erfolgt meist per E-Mail, SMS oder über gefälschte Webseiten.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Warning20Filled color="#F59E0B" />
              <Subtitle1 as="h3">Dringlichkeit</Subtitle1>
            </div>
            <p>
              Häufig setzen Angreifer unter Druck - z.B. durch angebliche Kontosperrungen oder Sicherheitswarnungen,
              um schnelles Handeln zu erzwingen.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <Link20Filled color="#EF4444" />
              <Subtitle1 as="h3">Falsche Links oder Anhänge</Subtitle1>
            </div>
            <p>
              In der Nachricht befinden sich oft gefälschte Links oder Anhänge. Diese führen zu manipulierten Webseiten 
              oder installieren Schadsoftware auf dem Gerät.
            </p>
          </div>
        </Card>
        <Card>
          <div className={Styles.Lesson__InfoCardContent}>
            <div className={Styles.Lesson__InfoCardHead}>
              <TextField20Filled color="#8B5CF6" />
              <Subtitle1 as="h3">Dateneingabe</Subtitle1>
            </div>
            <p>
              Auf gefälschten Seiten wird der Nutzer dazu aufgefordert, vertrauliche Daten wie Login-Informationen
              oder Kreditkartendaten einzugeben. Diese landen direkt beim Angreifer.
            </p>
          </div>
        </Card>
      </div>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Phishing ist deshalb so gefährlich, weil es oft sehr gut gemacht ist. Nur wer gut informiert ist,
        kann diese Angriffe erkennen und verhindern. Schulungen wie bei Securaware helfen dabei,
        die Merkmale frühzeitig zu identifizieren und sich selbst sowie andere zu schützen.
      </p>
    </article>
  )
}

export default Lesson2;