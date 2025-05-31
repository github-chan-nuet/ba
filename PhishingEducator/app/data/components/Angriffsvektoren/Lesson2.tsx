import { Subtitle1, Title2 } from "@fluentui/react-components";
import { Clock20Filled, Warning20Filled } from "@fluentui/react-icons";

const Lesson2 = () => {
  return (
    <article style={{ fontSize: 16 }}>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>SMS (Smishing)</Title2>
      <p style={{ color: '#4B5563', marginBottom: '24px' }}>
        Smishing - also Phishing per SMS - ist eine immer häufiger genutzte Methode von Cyberkriminellen. 
        Dabei werden SMS-Nachrichten mit täuschend echten Inhalten verschickt, um persönliche Daten abzugreifen 
        oder Schadsoftware auf mobilen Geräten zu installieren. Diese Angriffe wirken oft besonders glaubwürdig, 
        da SMS als direkter und persönlicher Kommunikationsweg wahrgenommen werden.
      </p>

      <section
        style={{
          background: "#FEF3C7",
          borderRadius: 8,
          padding: 16,
          marginBottom: 24
        }}
      >
        <Subtitle1 as="h3" style={{ display: 'block' }}>Typische Merkmale von Smishing</Subtitle1>
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
            background: '#DBEAFE',
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
            <Clock20Filled color="#2563EB" />
            <strong>Beispiel:</strong> "Ihr Paket konnte nicht zugestellt werden. Prüfen Sie Ihre Lieferdetails hier: [gefälschter Link]"
          </div>
        </div>
        <div
          style={{
            background: '#FECACA',
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
            <Warning20Filled color="#DC2626" />
            <strong>Achtung:</strong> Das Klicken auf solche Links kann zur Infektion des Geräts oder zur Preisgabe sensibler Informationen führen.
          </div>
        </div>
      </section>

      <p style={{ color: '#4B5563', marginBottom: 16 }}>
        Smishing ist deshalb besonders gefährlich, weil mobile Geräte oft schlechter geschützt sind als klassische Computer. 
        Zudem verleiten die Kürze der Nachricht und die scheinbare Dringlichkeit viele dazu, vorschnell zu handeln.
      </p>

      <p style={{ color: '#4B5563' }}>
        Achte darauf, niemals auf verdächtige Links in SMS zu klicken und überprüfe bei Unsicherheit den Absender oder 
        die Quelle über offizielle Wege - niemals direkt über die empfangene Nachricht.
      </p>
    </article>
  );
}

export default Lesson2;