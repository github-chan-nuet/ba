import {Button, Display, tokens} from "@fluentui/react-components";
import MarketingStyles from "../../styles/Marketing.module.scss";
import lines from '../../assets/images/lines.svg';
import { FoodFish24Filled, Mail48Filled, Pen48Filled, ShieldTask28Filled, Trophy28Filled } from "@fluentui/react-icons";
import React, {useRef, useEffect, useState, type ReactElement} from 'react';
import {Chart, CategoryScale, LinearScale, BarElement, type ChartOptions,} from 'chart.js';

import { Bar } from 'react-chartjs-2';

export default function Home() {
    useEffect(() => {
        document.title = 'Securaware';
    }, []);
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
                <Display style={{lineHeight: 1.1}}>
                    Gemeinsam gegen <strong>Phishing und Cyberbetrug </strong>
                    für eine <strong>sichere digitale Zukunft</strong>
                </Display>
            </section>
            <section className={MarketingStyles.Section}
                     style={{
                         paddingTop: "3rem",
                         // backgroundImage: "linear-gradient(to bottom, oklch(0.985 0.002 247.839), #48d5ff)",
                         // backgroundImage: "radial-gradient(oklch(0.985 0.002 247.839), #48d5ff)",
                         backgroundColor: "#fff",
                     }}
            >
                <h2 style={{textAlign: "center", fontSize: "2rem"}}>Was bietet Securaware?</h2>

                <div className={MarketingStyles.FeatureGrid}>
                    <FeatureCard title={"Phishing"} description={"Phishing"}  icon={<FoodFish24Filled />} />

                    <FeatureCard title={"Phishing Simulation"} description={"Phishing Simulation"} icon={<Mail48Filled />} />

                    <FeatureCard title={"Teste dein Wissen"} description={"Test"}  icon={<Pen48Filled />} />

                    <FeatureCard title={"Spilerischer Vergleich"} description={"Spilerischer Vergleich"} icon={<Trophy28Filled />} />
                </div>
            </section>

            <section className={MarketingStyles.Section} style={{
                backgroundColor: "#fff",
                display: "grid",
                gridTemplateColumns: "1fr 2fr",
                gridTemplateRows: "1fr",
                gap: "2rem",
                alignContent: "space-around",
                height: "80vh",
            }}>
                <div style={{
                    textAlign: "center",
                    display: "flex",
                    justifyContent: "space-around",
                    flexDirection: "column",
                }}>
                    <h3 className={MarketingStyles.Section__title}>Phishing-Angriffe <br/> nehmen rasant zu</h3>
                    <p className={MarketingStyles.Section__text}>Phishing-Angriffe kommen immer häufiger vor, dies bestätigt die amerikanische Sicherheitsbehörde FBI.
                        Die Graphik weist die aus Phishing entstandene Schäden in den USA auf. Der Trend ist klar - Phishing-Angriffe häufen sich.</p>
                </div>
                <PhishingCostDiagram />
            </section>

            <section className={MarketingStyles.Section} style={{
                height: "80vh",
                backgroundColor: "#fff",
                display: "grid",
                gridTemplateColumns: "2fr 3fr",
                gridTemplateRows: "1fr",
                gap: "2rem",
                alignContent: "space-around",
                // border: "1px solid black",
            }}>
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
                            shape={"circular"} // shape: 'rounded' | 'circular' | 'square';
                            iconPosition="after"
                            appearance="primary"
                            icon={<ShieldTask28Filled/>}
                            style={{width: "15rem", height: "4rem"}}
                            onClick={() => {
                            }}>
                            Schütze dich jetzt
                        </Button>
                    </div>
                </div>
            </section>
        </main>
    );
}

type FeatureCardProps = {title: string, description: string, icon: ReactElement};
function FeatureCard({title, description, icon}: FeatureCardProps) {
    const styledIcon = React.cloneElement(icon, {className: MarketingStyles.Icon})
    return (
        <section className={MarketingStyles.FeatureCard}>
            <div className={MarketingStyles.FeatureCard__icon}>
                {styledIcon}
            </div>
            <h3 style={{marginTop: "1.3rem", fontSize: "2rem"}}>{title}</h3>
            <p style={{marginTop: "1.7rem", fontSize: "1.4rem"}}>{description}</p>
        </section>
    );
}

Chart.register(CategoryScale, LinearScale, BarElement,);

function PhishingCostDiagram() {
    const chartRef = useRef(null);
    const containerRef = useRef(null);
    const [showChart, setShowChart] = useState(false);

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
                        size: 16,
                        family: fontFamily,
                    },
                    color: 'black',
                },
            },
            x: {
                ticks: {
                    font: {
                        size: 26,
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
                text: "Aus Phishing entstandene Schäden in den USA",
                font: {
                    size: 30,
                    family: fontFamily,
                },
                color: 'black',
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
    return (<div ref={containerRef}>
                <Bar ref={chartRef} options={options} data={data}/>
            </div>);
}


