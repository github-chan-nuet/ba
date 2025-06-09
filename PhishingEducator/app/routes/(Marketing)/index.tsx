import {Button, tokens} from "@fluentui/react-components";
import MarketingStyles from "@styles/Marketing.module.scss";
import lines from '@assets/images/lines.svg';
import { ShieldTask28Filled } from "@fluentui/react-icons";
import {useRef, useEffect, useState} from 'react';
import {Chart, CategoryScale, LinearScale, BarElement, type ChartOptions,} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import Hero from "@components/(Marketing)/Hero";
import FeatureGrid from "@components/(Marketing)/FeatureGrid";
import { useOutletContext } from "react-router";

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

            <section className={MarketingStyles.Section + " " + MarketingStyles.PhishingAttacks}>
                <div className={MarketingStyles.PhishingAttacks__text_container}>
                    <h3 className={MarketingStyles.Section__title}>Phishing-Angriffe <br/> nehmen rasant zu</h3>
                    <p className={MarketingStyles.Section__text}>Phishing-Angriffe kommen immer häufiger vor, dies bestätigt die amerikanische Sicherheitsbehörde FBI.
                        Die Graphik weist die aus Phishing entstandene Schäden in den USA auf. Der Trend ist klar - Phishing-Angriffe häufen sich.</p>
                </div>
                <PhishingCostDiagram />
            </section>

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


Chart.register(CategoryScale, LinearScale, BarElement,);

function PhishingCostDiagram() {
    const chartRef = useRef(null);
    const containerRef = useRef(null);
    const [showChart, setShowChart] = useState(false);
    const [isPortrait, setIsPortrait] = useState(false);

    useEffect(() => {
        setIsPortrait(window.matchMedia("(orientation: portrait)").matches)
    }, []);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setShowChart(true);
                    observer.disconnect();
                }
            },
            { threshold: 1 }
        );
        if (containerRef.current) {
            observer.observe(containerRef.current);
        }
        return () => observer.disconnect();
    }, []);

    const fontFamily = 'ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"';

    const options: ChartOptions<'bar'> = {
        responsive: true,
        maintainAspectRatio: false,
        animation: {
            duration: 2000,
            easing: 'easeOutSine',
        },
        scales: {
            y: {
                min: 0,
                max: 18_000_000_000,
                ticks: {
                    stepSize: 4_000_000_000,
                    display: true,
                    callback: function (tickValue: string | number) {
                        if (tickValue === 0 || typeof tickValue === 'string') return '';
                        return tickValue / 1_000_000_000 + ' Mrd. $';
                    },
                    font: {
                        size: isPortrait ? 12 : 16,
                        family: fontFamily,
                    },
                    color: 'black',
                },
            },
            x: {
                ticks: {
                    font: {
                        size: isPortrait ? 14 : 26,
                        family: fontFamily,
                    },
                    color: 'black',
                },
            }
        },
        plugins: {
            legend: {
                display: false,
            },
            title: {
                display: false,
            },
            tooltip: {
                enabled: false,
            },
        },
    };

    const filledBar = [4_600_000_000, 7_000_000_000, 10_200_000_000, 12_600_000_000, 16_800_000_000];
    const emptyBar = filledBar.map(() => 0);
    const data = {
        labels: ['2020', '2021', '2022', '2023', '2024'],
        datasets: [
            {
                label: 'Aus Phishing entstandene Schäden in den USA',
                data: showChart ? filledBar : emptyBar,
                backgroundColor: '#d83b01',
            },
        ],
    };
    return (<div ref={containerRef} style={{maxWidth: "100%"}}>
                <Bar ref={chartRef} options={options} data={data} style={{ width: "100%", height: "100%" }}/>
            </div>);
}


