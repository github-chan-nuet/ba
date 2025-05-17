import { Card } from "@fluentui/react-components";
import { Shield20Filled, Warning20Filled } from "@fluentui/react-icons";

const Lesson1 = () => {
  return (
    <>
      <p style={{ fontSize: '16px', color: '#4B5563', marginBottom: '24px' }}>
        Phishing ist eine Form des Cyberangriffs, bei dem Angreifer versuchen,
        sensible Informationen wie Passwörter, Kreditkartendaten oder Zugangsdaten
        zu Online-Konten zu stehlen. Die Auswirkungen solcher Angriffe können erheblich sein.
      </p>

      <div style={{
        display: 'flex',
        gap: 16,
        marginBottom: 24
      }}>
        <Card style={{
          flex: 1
        }}>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <Warning20Filled color="#EF4444" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Finanzieller Schaden</h2>
            </div>
            <p>
              Betroffene können direkt Geld verlieren, etwa durch betrügerische Überweisungen
              oder Kreditkartenmissbrauch. Auch Unternehmen erleiden oft hohe Kosten
              durch Betrug, Wiederherstellung und Imageverlust.
            </p>
          </div>
        </Card>

        <Card style={{
          flex: 1
        }}>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <Shield20Filled color="#D97706" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Verlust von Daten und Vertrauen</h2>
            </div>
            <p>
              Neben finanziellen Verlusten kann auch ein erheblicher Reputationsschaden entstehen.
              Kunden verlieren Vertrauen in ein Unternehmen, wenn deren Daten nicht sicher sind.
              In manchen Fällen ist auch der dauerhafte Verlust von Daten möglich.
            </p>
          </div>
        </Card>
      </div>

      <p style={{ fontSize: '16px', color: '#4B5563' }}>
        Die Bekämpfung von Phishing erfordert Aufklärung, technische Schutzmaßnahmen
        sowie ein wachsames Verhalten der Nutzer.
      </p>
    </>
  );
}

export default Lesson1;