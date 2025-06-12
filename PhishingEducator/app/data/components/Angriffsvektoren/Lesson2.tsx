import { Subtitle1, Title2 } from "@fluentui/react-components";
import { Clock20Filled, Warning20Filled } from "@fluentui/react-icons";

import Styles from '@data/courses.module.scss'

const Lesson2 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>SMS (Smishing)</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Smishing - also Phishing per SMS - ist eine immer häufiger genutzte Methode von Cyberkriminellen. 
        Dabei werden SMS-Nachrichten mit täuschend echten Inhalten verschickt, um persönliche Daten abzugreifen 
        oder Schadsoftware auf mobilen Geräten zu installieren. Diese Angriffe wirken oft besonders glaubwürdig, 
        da SMS als direkter und persönlicher Kommunikationsweg wahrgenommen werden.
      </p>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__TextBox}`}>
        <Subtitle1 as="h3" className={Styles.Lesson__Subtitle}>Typische Merkmale von Smishing</Subtitle1>
        <ul>
          <li>
            <strong>Kurz und eindringlich:</strong> Nachrichten setzen auf Angst oder Neugier, z.B. "Ihr Konto wurde gesperrt" oder "Sie haben ein Paket verpasst".
          </li>
          <li>
            <strong>Schädliche Links:</strong> Weiterleitungen zu gefälschten Webseiten oder automatischer Malware-Download.
          </li>
          <li>
            <strong>Vertrauenswürdige Absendernamen:</strong> Die SMS erscheinen von bekannten Namen oder Nummern zu kommen.
          </li>
        </ul>
      </section>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__MessageBox}`} data-type="example">
        <div className={Styles.Lesson__MessageBoxInner}>
          <Clock20Filled className={`shrink-0 ${Styles.Lesson__MessageBoxIcon}`} />
          <strong>Beispiel:</strong> "Ihr Paket konnte nicht zugestellt werden. Prüfen Sie Ihre Lieferdetails hier: [gefälschter Link]"
        </div>
      </section>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__MessageBox}`} data-type="warning">
        <div className={Styles.Lesson__MessageBoxInner}>
          <Warning20Filled className={`shrink-0 ${Styles.Lesson__MessageBoxIcon}`} />
          <strong>Achtung:</strong> Das Klicken auf solche Links kann zur Infektion des Geräts oder zur Preisgabe sensibler Informationen führen.
        </div>
      </section>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Smishing ist deshalb besonders gefährlich, weil mobile Geräte oft schlechter geschützt sind als klassische Computer. 
        Zudem verleiten die Kürze der Nachricht und die scheinbare Dringlichkeit viele dazu, vorschnell zu handeln.
      </p>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Achte darauf, niemals auf verdächtige Links in SMS zu klicken und überprüfe bei Unsicherheit den Absender oder 
        die Quelle über offizielle Wege - niemals direkt über die empfangene Nachricht.
      </p>
    </article>
  );
}

export default Lesson2;