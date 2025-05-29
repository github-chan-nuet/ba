import { Card, Title2 } from "@fluentui/react-components";
import { BuildingBank20Filled, ContactCard20Filled, Mail20Filled, Money20Filled, ShieldError20Filled } from "@fluentui/react-icons";

const Lesson4 = () => {
  return (
    <article style={{ fontSize: 16 }}>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>Konsequenzen von Phishing</Title2>
      <p style={{ color: '#4B5563', marginBottom: '24px' }}>
        Phishing kann schwerwiegende Folgen haben – sowohl für Einzelpersonen als auch für Unternehmen. 
        Die Konsequenzen reichen von finanziellen Verlusten über Identitätsdiebstahl bis hin zu erheblichen 
        Schäden für die Reputation und Sicherheit eines Unternehmens.
      </p>

      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(350px, 1fr))',
        gap: 16,
        marginBottom: 48
      }}>
        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <Money20Filled color="#EF4444" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Finanzielle Verluste</h2>
            </div>
            <p>
              Betroffene verlieren durch betrügerische Überweisungen, Einkäufe oder Kontoabbuchungen oft viel Geld.
              Unternehmen müssen nicht nur für entstandene Schäden aufkommen, sondern auch für Wiederherstellungsmaßnahmen.
            </p>
          </div>
        </Card>

        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <ContactCard20Filled color="#F59E0B" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Identitätsdiebstahl</h2>
            </div>
            <p>
              Persönliche Daten wie Name, Adresse oder Sozialversicherungsnummer können missbraucht werden,
              um falsche Identitäten anzulegen oder betrügerische Handlungen vorzunehmen.
            </p>
          </div>
        </Card>

        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <ShieldError20Filled color="#10B981" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Reputationsschäden</h2>
            </div>
            <p>
              Firmen, die von Phishing betroffen sind, verlieren oft das Vertrauen ihrer Kunden. 
              Ein einziger Vorfall kann langfristige Auswirkungen auf die Markenwahrnehmung haben.
            </p>
          </div>
        </Card>
      </div>

      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>Beispiele für schwerwiegende Phishing-Attacken</Title2>

      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(350px, 1fr))',
        gap: 16,
        marginBottom: 24
      }}>
        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <BuildingBank20Filled color="#6366F1" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>2016: Angriff auf Hillary Clintons Wahlkampfleiter</h2>
            </div>
            <p>
              Der Wahlkampfleiter von Hillary Clinton, John Podesta, wurde Opfer einer Phishing-E-Mail. 
              Sein E-Mail-Konto wurde kompromittiert und sensible Wahlkampfinhalte wurden veröffentlicht. 
              Der Vorfall hatte große politische Auswirkungen.
            </p>
          </div>
        </Card>

        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <Mail20Filled color="#EC4899" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>2020: Twitter-Hack</h2>
            </div>
            <p>
              Über Social Engineering und Phishing wurden Twitter-Mitarbeiter dazu gebracht, Zugangsdaten preiszugeben. 
              Dadurch konnten Angreifer Zugriff auf prominente Konten (z.B. Elon Musk, Barack Obama) erlangen und gefälschte Tweets verbreiten, 
              die zu Bitcoin-Betrug führten.
            </p>
          </div>
        </Card>
      </div>

      <p style={{ color: '#4B5563' }}>
        Diese Beispiele zeigen, wie gefährlich Phishing sein kann. Securaware hilft dir, solche Angriffe zu erkennen
        und dich effektiv zu schützen.
      </p>
    </article>
  );
}

export default Lesson4;