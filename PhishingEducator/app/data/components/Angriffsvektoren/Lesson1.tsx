import { Subtitle1, Title2 } from "@fluentui/react-components";
import { Shield20Filled, Warning20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson1 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>E-Mail</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        E-Mail ist einer der ältesten und gleichzeitig effektivsten Angriffsvektoren im Bereich Phishing. 
        Cyberkriminelle nutzen sie, um massenhaft gefälschte Nachrichten zu versenden, die scheinbar von 
        vertrauenswürdigen Absendern wie Banken, Online-Shops oder Behörden stammen.
      </p>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__TextBox}`}>
        <Subtitle1 as="h3" className={Styles.Lesson__Subtitle}>Wie E-Mail-Phishing funktioniert</Subtitle1>
        <ul>
          <li>
            <strong>Dringlichkeit erzeugen:</strong> Aufforderung zur Konto-Verifizierung oder Warnungen über angebliche Sicherheitsprobleme.
          </li>
          <li>
            <strong>Professionelles Design:</strong> Offizielle Logos, gefälschte Absender und realistisch wirkende Texte.
          </li>
          <li>
            <strong>Gefälschte Links:</strong> Weiterleitungen auf täuschend echte Websites zur Dateneingabe.
          </li>
          <li>
            <strong>Schädliche Anhänge:</strong> Infizierte Dateien, die beim Öffnen Malware installierne können.
          </li>
        </ul>
      </section>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__MessageBox} `} data-type="example">
        <div className={Styles.Lesson__MessageBoxInner}>
          <Warning20Filled className={`shrink-0 ${Styles.Lesson__MessageBoxIcon}`} />
          <strong>Beispiel:</strong> Eine angebliche Nachricht deiner Bank fordert dich auf, deine Zugangsdaten zu aktualisieren - über einen Link, der auf eine gefälsche Login-Seite führt.
        </div>
      </section>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__MessageBox}`} data-type="notice">
        <div className={Styles.Lesson__MessageBoxInner}>
          <Shield20Filled className={`shrink-0 ${Styles.Lesson__MessageBoxIcon}`} />
          <strong>Hinweis:</strong> Echte Unternehmen fordern niemals per E-Mail zur Eingabe sensibler Daten auf.
        </div>
      </section>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Aufgrund der hohen Reichweite, der geringen Kosten und der potenziellen Erfolgsquote bleibt die E-Mail 
        eines der beliebtesten Werkzeuge von Phishing-Angreifern. Aufmerksamkeit, technisches Grundwissen und gesunde 
        Skepsis sind entscheidend, um sich zu schützen.
      </p>
    </article>
  );
}

export default Lesson1;