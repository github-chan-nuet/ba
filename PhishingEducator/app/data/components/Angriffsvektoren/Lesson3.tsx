import { Subtitle1, Title2 } from "@fluentui/react-components";

const Lesson3 = () => {
  return (
    <article style={{ fontSize: 16 }}>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>Telefonanruf</Title2>
      <p style={{ color: '#4B5563', marginBottom: '24px' }}>
        Phishing per Telefonanruf - auch als "Voice Phishing" oder "Vishing" bekannt - ist eine besonders persönliche Methode, um Opfer zu täuschen. 
        Angreifer rufen gezielt Personen an und geben sich als vertrauenswürdige Institutionen wie Banken, IT-Support, Polizei oder Unternehmen aus.
      </p>

      <section
        style={{
          background: "#DBEAFE",
          borderRadius: 8,
          padding: 16,
          marginBottom: 24
        }}
      >
        <Subtitle1 as="h3" style={{ display: 'block' }}>Typische Szenarien</Subtitle1>
        <ul>
          <li>Ein angeblicher Bankmitarbeiter meldet eine verdächtige Kontobewegung.</li>
          <li>Der Anrufer behauptet, es liege ein technisches Problem vor, das dringend gelöst werden müsse.</li>
          <li>Eine vermeintliche Strafverfolgungsbehörde warnt vor angeblichem Identitätsdiebstahl.</li>
        </ul>
      </section>

      <p style={{ color: '#4B5563', marginBottom: 24 }}>
        Die Betrüger stellen dabei gezielte Fragen nach PINs, Passwörtern, TANs oder persönlichen Daten. Ziel ist es, Zugang zu Konten zu erhalten oder Folgeangriffe vorzubereiten.
        Das direkte Gespräch baut Vertrauen auf - Opfer fühlen sich unter Druck gesetzt, schnell zu handeln und geben dadurch Informationen preis, die sie schriftlich niemals teilen würden.
      </p>

      <p style={{ color: '#4B5563' }}>
        <strong>Es ist deshalb wichtig, stets wachsam zu sein.</strong> Echte Institutionen fordern <strong>nie</strong> telefonisch die Herausgabe sensibler Zugangsdaten!
      </p>
    </article>
  );
}

export default Lesson3;