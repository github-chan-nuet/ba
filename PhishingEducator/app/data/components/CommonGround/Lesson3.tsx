import { Card, Title2 } from "@fluentui/react-components";
import { Globe20Filled, MailAlert20Filled, People20Filled } from "@fluentui/react-icons";

const Lesson3 = () => {
  return (
    <>
      <Title2 as="h2" style={{ display: 'block', marginBottom: 16 }}>Wie verbreitet ist Phishing?</Title2>
      <p style={{ fontSize: '16px', color: '#4B5563', marginBottom: '16px' }}>
        Phishing ist eine der am weitesten verbreiteten Formen der Internetkriminalität. 
        Jeden Tag werden Millionen von Phishing-E-Mails, SMS und gefälschten Webseiten weltweit verschickt. 
        Sowohl Privatpersonen als auch Unternehmen sind betroffen.
      </p>

      <p style={{ fontSize: '16px', color: '#4B5563', marginBottom: '24px' }}>
        Der Grund für die hohe Verbreitung ist einfach: Phishing ist für Angreifer leicht durchzuführen, 
        kostengünstig und oft sehr erfolgreich. Schon ein einziger Klick kann zu einem finanziellen Schaden 
        oder Datenverlust führen.
      </p>

      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(350px, 1fr))',
        gap: 16,
        marginBottom: 24
      }}>
        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <MailAlert20Filled color="#0EA5E9" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Geringe Einstiegshürde für Angreifer</h2>
            </div>
            <p>
              Phishing-Kampagnen lassen sich mit wenig technischem Wissen und einfachen Mitteln durchführen.
              Es gibt sogar Baukastensysteme und Foren, die kriminelle Akteure unterstützen.
            </p>
          </div>
        </Card>

        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <People20Filled color="#D97706" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Erfolgsquote durch menschliches Verhalten</h2>
            </div>
            <p>
              Phishing nutzt menschliche Schwächen wie Neugier, Stress oder Unwissenheit aus. 
              Viele Empfänger reagieren reflexartig auf vermeintlich wichtige oder dringende Nachrichten.
            </p>
          </div>
        </Card>

        <Card>
          <div style={{ padding: '16px', marginBottom: '16px' }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '8px', marginBottom: '8px' }}>
              <Globe20Filled color="#6366F1" />
              <h2 style={{ fontSize: '20px', fontWeight: '600' }}>Globale Reichweite durch digitale Kanäle</h2>
            </div>
            <p>
              Phishing kennt keine geografischen Grenzen. Eine einzige Kampagne kann in Sekunden weltweit
              verbreitet werden - per E-Mail, Social Media oder Messaging-Apps.
            </p>
          </div>
        </Card>
      </div>

      <p style={{ fontSize: '16px', color: '#4B5563' }}>
        Die hohe Verbreitung macht es umso wichtiger, dass sich jeder mit dem Thema Phishing auseinandersetzt.
        Durch Schulungen wie bei Securaware kann man lernen, Gefahren zu erkennen und sich effektiv zu schützen.
      </p>
    </>
  );
}

export default Lesson3;