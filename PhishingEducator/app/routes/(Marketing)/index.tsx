import {Body1Stronger, Button, Display, ToggleButton, tokens} from "@fluentui/react-components";
import MarketingStyles from "../../styles/Marketing.module.scss";
import lines from '../../assets/images/lines.svg';
import { FoodFish24Filled, Mail48Filled, Pen48Filled, ShieldTask28Filled, Trophy28Filled } from "@fluentui/react-icons";
import React, {useRef, useEffect, useState, type ReactElement} from 'react';
import {Chart, CategoryScale, LinearScale, BarElement, type ChartOptions,} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import AuthDrawer from "../../components/AuthDrawer";
import logo from "../../assets/images/securaware.png";

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

export default function MarketingLayout() {
    const [isAuthDrawerOpen, setIsAuthDrawerOpen] = useState(false);
    return (
        <div className={MarketingStyles.Marketing}>
            <div className={MarketingStyles.Marketing__container}>
                <AuthDrawer isOpen={isAuthDrawerOpen} setIsOpen={setIsAuthDrawerOpen} />
                <header style={{
                    display: 'flex',
                    justifyContent: 'space-between',
                    alignItems: 'center',
                    marginLeft: "1rem",
                    marginRight: "1rem",
                    paddingBlock: 24
                }}>
                    <div style={{
                        display: 'flex',
                        gap: '0.5rem',
                    }}>
                        <img src={logo} alt="" width={20} />
                        <Body1Stronger style={{ fontSize: '1.5rem' }}>
                            Securaware
                        </Body1Stronger>
                    </div>
                    <ToggleButton
                        appearance="primary"
                        onClick={() => setIsAuthDrawerOpen(!isAuthDrawerOpen)}
                        checked={isAuthDrawerOpen}
                    >
                        Login
                    </ToggleButton>
                </header>
                <Home setAuthOpen={() => setIsAuthDrawerOpen(true)} />
            </div>
        </div>
    )
}

type HomeProps = {setAuthOpen: () => void};
function Home({setAuthOpen}: HomeProps) {
    return (
        <main style={{display: "flex", flexDirection: "column", alignItems: "center", gap: "10rem", paddingBottom: "10rem"}}>
            <section className={MarketingStyles.CircleBackground}>
                <div className={MarketingStyles.CircleBackground__cirlce} />
            </section>
            <section
                style={{
                    maxWidth: "70rem",
                    textAlign: "center",
                    margin: "5rem auto 0 auto",
                    height: "80vh",
                }}
            >
                <h1 className={MarketingStyles.Title}>
                    Securaware
                </h1>
                <div>
                    <hr className={MarketingStyles.Line}/>
                </div>
                <Display style={{lineHeight: 1.1, fontSize: "2rem"}}>
                    Gemeinsam gegen <strong>Phishing und Cyberbetrug </strong>
                    für eine <strong>sichere digitale Zukunft</strong>
                </Display>
            </section>

            <section className={MarketingStyles.Section} style={{paddingTop: "3rem", backgroundColor: "#fff"}}>
                <h2 style={{textAlign: "center", fontSize: "2rem"}}>Was bietet Securaware?</h2>
                <FeatureGrid />
            </section>

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
                            // onClick={() => {
                            //     console.log("hoi");
                            //     console.log(setAuthOpen);
                            //     setAuthOpen();
                            // }}>
                            onClick={setAuthOpen}>
                            Schütze dich jetzt
                        </Button>
                    </div>
                </div>
            </section>
        </main>
    );
}

function FeatureGrid() {
    const gridRef = useRef(null);
    const [isInView, setIsInView] = useState(false);

    useEffect(() => {
        if (window.matchMedia("(orientation: portrait)").matches) {
            setIsInView(true);
        }
    }, []);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setIsInView(true);
                    observer.disconnect();
                }
            },
            { threshold: 0.9 }
        );
        if (gridRef.current) {
            observer.observe(gridRef.current);
        }
        return () => observer.disconnect();
    }, []);

    const className = isInView ? MarketingStyles.FeatureGrid + " " + MarketingStyles.FeatureGrid__active : MarketingStyles.FeatureGrid;

    return (
        <div className={className} ref={gridRef}>
            <FeatureCard title={"Online Kurse"}
                         description={"Über Securawares online Kurse lernst du alles rund um Phishing!"}
                         icon={<FoodFish24Filled />} />

            <FeatureCard title={"Phishing Simulation"}
                         description={"Identifiziere Phishing-E-Mails und vertiefe dein Wissen mit Hinweisen zu falsch identifizierten E-Mails"}
                         icon={<Mail48Filled />} />

            <FeatureCard title={"Teste dein Wissen"}
                         description={"Teste dein Wissen mit fortlaufend neuen Tests und sammle XP dabei"}
                         icon={<Pen48Filled />} />

            <FeatureCard title={"Spilerischer Vergleich"}
                         description={"Vergleiche dich mit Freunden und weiteren Benutzern über deine gesammelte XP"}
                         icon={<Trophy28Filled />} />
        </div>
    )
}

type FeatureCardProps = {title: string, description: string, icon: ReactElement};
function FeatureCard({title, description, icon}: FeatureCardProps) {
    const styledIcon = React.cloneElement(icon, {className: MarketingStyles.Icon})
    return (
        <section className={MarketingStyles.FeatureCard}>
            <div className={MarketingStyles.FeatureCard__icon}>
                {styledIcon}
            </div>
            <h3 className={MarketingStyles.FeatureCard__title}>{title}</h3>
            <p className={MarketingStyles.FeatureCard__text}>{description}</p>
        </section>
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
                <Bar ref={chartRef} options={options} data={data}/>
            </div>);
}


