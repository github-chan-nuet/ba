import { Subtitle1, Title2 } from "@fluentui/react-components";

import Styles from '@data/courses.module.scss'

const Lesson3 = () => {
  return (
    <article className={Styles.Lesson}>
      <Title2 as="h2" className={Styles.Lesson__Title}>Telefonanruf</Title2>
      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Phishing per Telefonanruf - auch als "Voice Phishing" oder "Vishing" bekannt - ist eine besonders persönliche Methode, um Opfer zu täuschen. 
        Angreifer rufen gezielt Personen an und geben sich als vertrauenswürdige Institutionen wie Banken, IT-Support, Polizei oder Unternehmen aus.
      </p>

      <section className={`${Styles.Lesson__Section} ${Styles.Lesson__TextBox}`}>
        <Subtitle1 as="h3" className={Styles.Lesson__Subtitle}>Typische Szenarien</Subtitle1>
        <ul>
          <li>Ein angeblicher Bankmitarbeiter meldet eine verdächtige Kontobewegung.</li>
          <li>Der Anrufer behauptet, es liege ein technisches Problem vor, das dringend gelöst werden müsse.</li>
          <li>Eine vermeintliche Strafverfolgungsbehörde warnt vor angeblichem Identitätsdiebstahl.</li>
        </ul>
      </section>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        Die Betrüger stellen dabei gezielte Fragen nach PINs, Passwörtern, TANs oder persönlichen Daten. Ziel ist es, Zugang zu Konten zu erhalten oder Folgeangriffe vorzubereiten.
        Das direkte Gespräch baut Vertrauen auf - Opfer fühlen sich unter Druck gesetzt, schnell zu handeln und geben dadurch Informationen preis, die sie schriftlich niemals teilen würden.
      </p>

      <p className={`${Styles.Lesson__Section} ${Styles.Lesson__Paragraph}`}>
        <strong>Es ist deshalb wichtig, stets wachsam zu sein.</strong> Echte Institutionen fordern <strong>nie</strong> telefonisch die Herausgabe sensibler Zugangsdaten!
      </p>
    </article>
  );
}

export default Lesson3;