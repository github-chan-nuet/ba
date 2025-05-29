import { Card, Title2 } from "@fluentui/react-components"
import { Link20Filled, PersonAlert20Filled, TextField20Filled, Warning20Filled } from "@fluentui/react-icons";

const Lesson2 = () => {
  return (
    <>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>Was ist Phishing?</Title2>
      <p style={{ fontSize: '16px', color: '#4B5563', marginBottom: '16px' }}>
        Unter dem Begriff <strong>Phishing</strong> versteht man eine Methode des digitalen Betrugs.
        Angreifer geben sich als vertrauenswürdige Quelle aus – zum Beispiel eine Bank, ein Kollege oder eine Behörde –
        um an vertrauliche Informationen wie Passwörter, Kreditkartendaten oder Zugangsdaten zu gelangen.
      </p>

      <p style={{ fontSize: '16px', color: '#4B5563', marginBottom: '24px' }}>
        Ziel ist es, durch Täuschung Personen dazu zu bringen, sensible Daten preiszugeben. Phishing zählt zu den
        häufigsten Cyberangriffsformen weltweit.
      </p>

      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(350px, 1fr))',
        gap: 16,
        marginBottom: 24
      }}>
        <Card>
          <div style={{
            padding: 16,
            marginBottom: 16
          }}>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              gap: 8,
              marginBottom: 8,
            }}>
              <PersonAlert20Filled color="#0EA5E9" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Täuschung</h2>
            </div>
            <p>
              Die Täter wirken oft täuschend echt. Sie imitieren bekannte Absender wie Banken, Online-Shops oder Kollegen. 
              Die Kontaktaufnahme erfolgt meist per E-Mail, SMS oder über gefälschte Webseiten.
            </p>
          </div>
        </Card>
        <Card>
          <div style={{
            padding: 16,
            marginBottom: 16
          }}>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              gap: 8,
              marginBottom: 8,
            }}>
              <Warning20Filled color="#F59E0B" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Dringlichkeit</h2>
            </div>
            <p>
              Häufig setzen Angreifer unter Druck - z.B. durch angebliche Kontosperrungen oder Sicherheitswarnungen,
              um schnelles Handeln zu erzwingen.
            </p>
          </div>
        </Card>
        <Card>
          <div style={{
            padding: 16,
            marginBottom: 16
          }}>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              gap: 8,
              marginBottom: 8,
            }}>
              <Link20Filled color="#EF4444" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Falsche Links oder Anhänge</h2>
            </div>
            <p>
              In der Nachricht befinden sich oft gefälschte Links oder Anhänge. Diese führen zu manipulierten Webseiten 
              oder installieren Schadsoftware auf dem Gerät.
            </p>
          </div>
        </Card>
        <Card>
          <div style={{
            padding: 16,
            marginBottom: 16
          }}>
            <div style={{
              display: 'flex',
              alignItems: 'center',
              gap: 8,
              marginBottom: 8,
            }}>
              <TextField20Filled color="#8B5CF6" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Dateneingabe</h2>
            </div>
            <p>
              Auf gefälschten Seiten wird der Nutzer dazu aufgefordert, vertrauliche Daten wie Login-Informationen
              oder Kreditkartendaten einzugeben. Diese landen direkt beim Angreifer.
            </p>
          </div>
        </Card>
      </div>

      <p style={{ fontSize: '16px', color: '#4B5563' }}>
        Phishing ist deshalb so gefährlich, weil es oft sehr gut gemacht ist. Nur wer gut informiert ist,
        kann diese Angriffe erkennen und verhindern. Schulungen wie bei Securaware helfen dabei,
        die Merkmale frühzeitig zu identifizieren und sich selbst sowie andere zu schützen.
      </p>
    </>
  )
}

export default Lesson2;