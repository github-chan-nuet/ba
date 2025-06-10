import Hero from "@components/(Marketing)/Hero";
import FeatureGrid from "@components/(Marketing)/FeatureGrid";
import Stats from "@components/(Marketing)/Stats";
import CallToAction from "@components/(Marketing)/CallToAction";

export function meta() {
    return [
        { title: 'Securaware - Phishing erkennen und sicher im Netz unterwegs sein' },
        {
            name: 'description',
            content: 'Securaware hilft dir, Phishing-Angriffe zu erkennen und sicher im Internet zu surfen. Lerne, wie du dich vor Betrug und Datendiebstahl schützt - einfach, effektiv und verständlich.'
        },
        {
            name: 'keywords',
            content: 'Phishing, Phishing erkennen, Online Sicherheit, Internet Schutz, Betrug verhindern, Cybersecurity für Privatpersonen, Securaware, Phishing Schutz, sicher im Netz, Online Betrug'
        }
    ]
}

export default function Home() {
    return (
        <>
            <Hero
                title={<>Gemeinsam gegen <strong>Phishing und Cyberbetrug</strong> für eine <strong>sichere digitale Zukunft</strong></>}
                subtitle={<>Scrolle nach unten und erfahre mehr darüber, wie du dich effektiv vor Phishing schützen kannst!</>}
                showScrollIcon={true}
            />
            <FeatureGrid />
            <Stats />
            <CallToAction />
        </>
    );
}
