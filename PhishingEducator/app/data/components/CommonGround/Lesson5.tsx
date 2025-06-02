import { Card, Title2 } from "@fluentui/react-components"
import { Book20Filled, Briefcase20Filled, ContactCardGroup20Filled, Globe20Filled, Shield20Filled } from "@fluentui/react-icons";

const Lesson5 = () => {
  return (
    <article style={{ fontSize: 16 }}>
      <Title2 as={"h2"} style={{ display: 'block', marginBottom: 16 }}>Warum Securaware?</Title2>
      <p style={{ color: '#4B5563', marginBottom: '24px' }}>
        Die digitale Welt veränder sich rasant - wer sicher unterwegs sein möchte, muss sich laufend weiterbilden.
        Securaware bietet dir die Möglichkeit, dein Wissen rund um Cyber-Sicherheit gezielt auszubauen.
      </p>

      <div style={{
        display: 'grid',
        gridTemplateColumns: 'repeat(auto-fit, minmax(350px, 1fr))',
        gap: 16,
        marginBottom: 24
      }}>
        <Card style={{
          flex: '1 0 min(100%, 350px)'
        }}>
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
              <Shield20Filled color="#2563EB" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Selbstschutz vor finanziellen Verlusten</h2>
            </div>
            <p>
              Wer Phishing erkennt, kann sich und sein Geld schützen. Securaware hilft dir, Risiken zu erkennen und richtig zu handeln.
            </p>
          </div>
        </Card>
        <Card style={{
          flex: '1 0 min(100%, 350px)'
        }}>
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
              <Briefcase20Filled color="#10B981" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Berufliche Verantwortung</h2>
            </div>
            <p>
              Im Berufsalltag kann ein Klick auf einen Phishing-Link großen Schaden anrichten. Securaware sensibilisiert und stärkt dich im Arbeitsumfeld.
            </p>
          </div>
        </Card>
        <Card style={{
          flex: '1 0 min(100%, 350px)'
        }}>
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
              <ContactCardGroup20Filled color="#F59E0B" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Schutz der Familie und Angehörige</h2>
            </div>
            <p>
              Ältere Menschen oder Kinder sind besonders gefährdet. Dein Wissen kann deine Familie schützen - Securaware zeigt dir wie.
            </p>
          </div>
        </Card>
        <Card style={{
          flex: '1 0 min(100%, 350px)'
        }}>
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
              <Book20Filled color="#6366F1" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Karrierevorteile durch digitale Kompetenz</h2>
            </div>
            <p>
              Digitale Mündigkeit wird immer wichtiger. Wer sich mit Sicherheit im Netz auskennt, sammelt Pluspunkte im Beruf.
            </p>
          </div>
        </Card>
        <Card style={{
          flex: '1 0 min(100%, 350px)'
        }}>
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
              <Globe20Filled color="#EC4899" />
              <h2 style={{ fontSize: 20, fontWeight: 600 }}>Mithelfen, die Gesellschaft zu schützen</h2>
            </div>
            <p>
              Phishing ist ein gesellschaftliches Problem. Wer sich weiterbildet, kann andere aufklären und so zur digitalen Sicherheit aller beitragen.
            </p>
          </div>
        </Card>
      </div>
    </article>
  )
}

export default Lesson5;