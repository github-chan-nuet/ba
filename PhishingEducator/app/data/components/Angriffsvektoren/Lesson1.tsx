import { Subtitle1, Title2 } from "@fluentui/react-components";
import { Shield20Filled, Warning20Filled } from "@fluentui/react-icons";

const Lesson1 = () => {
  return (
    <article style={{ fontSize: 16 }}>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>E-Mail</Title2>
      <p style={{ color: '#4B5563', marginBottom: '24px' }}>
        E-Mail ist einer der ältesten und gleichzeitig effektivsten Angriffsvektoren im Bereich Phishing. 
        Cyberkriminelle nutzen sie, um massenhaft gefälschte Nachrichten zu versenden, die scheinbar von 
        vertrauenswürdigen Absendern wie Banken, Online-Shops oder Behörden stammen.
      </p>

      <section
        style={{
          background: "#F3F4F6",
          borderRadius: 8,
          padding: 16,
          marginBottom: 24
        }}
      >
        <Subtitle1 as="h3" style={{ display: 'block' }}>Wie E-Mail-Phishing funktioniert</Subtitle1>
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

      <section
        style={{
          display: 'flex',
          flexDirection: 'column',
          gap: 16,
          marginBottom: 32
        }}
      >
        <div
          style={{
            background: '#E0F2FE',
            padding: '12px 16px',
            borderRadius: 6
          }}
        >
          <div
            style={{
              display: 'flex',
              gap: 8
            }}
          >
            <Warning20Filled color="#0284C7" />
            <strong>Beispiel:</strong> Eine angebliche Nachricht deiner Bank fordert dich auf, deine Zugangsdaten zu aktualisieren - über einen Link, der auf eine gefälsche Login-Seite führt.
          </div>
        </div>
        <div
          style={{
            background: '#FEF9C3',
            padding: '12px 16px',
            borderRadius: 6
          }}
        >
          <div
            style={{
              display: 'flex',
              gap: 8
            }}
          >
            <Shield20Filled color="#CA8A04" />
            <strong>Hinweis:</strong> Echte Unternehmen fordern niemals per E-Mail zur Eingabe sensibler Daten auf.
          </div>
        </div>
      </section>

      <p style={{ color: '#4B5563' }}>
        Aufgrund der hohen Reichweite, der geringen Kosten und der potenziellen Erfolgsquote bleibt die E-Mail 
        eines der beliebtesten Werkzeuge von Phishing-Angreifern. Aufmerksamkeit, technisches Grundwissen und gesunde 
        Skepsis sind entscheidend, um sich zu schützen.
      </p>
    </article>
  );
}

export default Lesson1;