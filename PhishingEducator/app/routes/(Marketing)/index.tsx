import {Button, tokens} from "@fluentui/react-components";
import MarketingStyles from "@styles/Marketing.module.scss";
import lines from '@assets/images/lines.svg';
import { ShieldTask28Filled } from "@fluentui/react-icons";
import Hero from "@components/(Marketing)/Hero";
import FeatureGrid from "@components/(Marketing)/FeatureGrid";
import { useOutletContext } from "react-router";
import Stats from "@components/(Marketing)/Stats";

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

type HomeContext = {setAuthOpen: () => void};
export default function Home() {
    const context = useOutletContext<HomeContext>();

    return (
        <>
            <Hero
                title={<>Gemeinsam gegen <strong>Phishing und Cyberbetrug</strong> für eine <strong>sichere digitale Zukunft</strong></>}
                subtitle={<>Scrolle nach unten und erfahre mehr darüber, wie du dich effektiv vor Phishing schützen kannst!</>}
                showScrollIcon={true}
            />

            <FeatureGrid />

            <Stats />

            <section className={MarketingStyles.Section + " " + MarketingStyles.SecurawareSolution}>
                <div style={{
                    backgroundColor: tokens.colorPaletteBlueBorderActive,
                    backgroundImage: "url(\"" + lines + "\")",
                    backgroundRepeat: "no-repeat",
                    height: "100%",
                    borderRadius: "2rem",
                }}/>
                <div style={{
                    textAlign: "center",
                    display: "flex",
                    justifyContent: "space-around",
                    flexDirection: "column",
                }}>
                    <h3 className={MarketingStyles.Section__title}>Securaware als Lösung</h3>
                    <p className={MarketingStyles.Section__text}>Wie schütze ich mich vor Phishing-Angriffen? Wie identifiziere ich diese? Securaware schützt dich und hilft dir.</p>
                    <div>
                        <Button
                            size="large"
                            shape={"circular"}
                            iconPosition="after"
                            appearance="primary"
                            icon={<ShieldTask28Filled/>}
                            style={{width: "15rem", height: "4rem"}}
                            onClick={context.setAuthOpen}>
                            Schütze dich jetzt
                        </Button>
                    </div>
                </div>
            </section>
        </>
    );
}
