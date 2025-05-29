import {Button, Display, tokens} from "@fluentui/react-components";
// import * as FluentIcons from "@fluentui/react-icons";
import MarketingStyles from "../../styles/Marketing.module.scss";
import hacker from '../../assets/images/hacker.png';
import phishing from '../../assets/images/phishing.png';
import lines from '../../assets/images/lines.svg';
import {
    ArrowRight24Regular, FoodFish24Filled, Mail48Filled, Pen48Filled,
    Shield24Filled,
    Shield24Regular,
    ShieldTask28Filled,
    Trophy28Filled
} from "@fluentui/react-icons";
import { useRef, useEffect, useState } from 'react';

export default function Home() {
    // const IconComponent = FluentIcons["LineHorizontal128Filled"];
    useEffect(() => {
        document.title = 'Securaware';
    }, []);
    return (

        <article style={{display: "flex", flexDirection: "column", alignItems: "center", gap: "10rem"}}>
            <div
                style={{
                    position: "absolute",
                    top: 0,
                    left: 0,
                    zIndex: -10,
                    height: "100vh",
                    width: "100%",
                    backgroundColor: "oklch(0.985 0.002 247.839)",
                }}
            >
                <div
                    style={{
                        position: "absolute",
                        bottom: "auto",
                        left: "auto",
                        right: 0,
                        top: 0,
                        height: "500px",
                        width: "500px",
                        transform: "translate(-70%, 25%)",
                        borderRadius: "9999px", // full rounding
                        backgroundColor: "rgba(0, 120, 212, 0.5)", // "rgba(173, 109, 244, 0.5)",
                        opacity: 0.5,
                        filter: "blur(80px)",
                    }}/>
            </div>
            <div
                style={{
                    maxWidth: "70rem",
                    textAlign: "center",
                    margin: "5rem auto 0 auto",
                    height: "100vh",
                }}
            >
                <h1 className={MarketingStyles.Title}>
                    Securaware
                </h1>
                <div>
                    <hr style={{
                        height: "0.2rem",
                        width: "35rem",
                        backgroundColor: "black",
                        color: "black",
                        marginBottom: "4rem"
                    }}/>
                </div>
                {/*<IconComponent style={{fontSize: "2rem", width: "10rem"}} />*/}
                <Display
                    style={{
                        lineHeight: 1.1
                    }}
                >
                    Gemeinsam gegen <strong>Phishing und Cyberbetrug </strong>
                    für eine <strong>sichere digitale Zukunft</strong>
                </Display>
            </div>
            <section className={MarketingStyles.Section}
                     style={{
                         paddingTop: "10rem",
                         // backgroundImage: "linear-gradient(to bottom, oklch(0.985 0.002 247.839), #48d5ff)",
                         backgroundImage: "radial-gradient(oklch(0.985 0.002 247.839), #48d5ff)",
                     }}
            >
                Test

                <div style={{
                    display: "grid",
                    gap: "3rem",
                    gridTemplateColumns: "1fr 1fr 1fr",
                    gridTemplateRows: "1fr 1fr",
                    gridAutoFlow: "row",
                    justifyItems: "center",
                    margin: "10rem",
                }}>
                    <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing}/>

                    <FoodFish24Filled />

                    <FunctionElement title={"Phishing Simulation"} description={"Das ist so"} imageSrc={phishing}/>

                    <Mail48Filled width="1000" height="1000" style={{fontSize: "1rem", width: "10rem", height: "10rem"}} />

                    <FunctionElement title={"Test"} description={"Das ist so"} imageSrc={phishing}/>

                    <Pen48Filled style={{fontSize: "1rem", width: "10rem", height: "10rem"}} />

                    <FunctionElement title={"Spilerischer Vergleich"} description={"Das ist so"} imageSrc={phishing}/>

                    <Trophy28Filled />

                    <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing}/>

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

            <div style={{
                backgroundColor: tokens.colorPaletteBlueBackground2,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteBlueBackground2
            </div>
            <div style={{
                backgroundColor: tokens.colorPaletteBlueForeground2,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteBlueForeground2
            </div>
            <div style={{
                backgroundColor: tokens.colorPaletteBlueBorderActive,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteBlueBorderActive
            </div>
            <div style={{
                backgroundColor: tokens.colorPaletteRoyalBlueBackground2,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteRoyalBlueBackground2
            </div>
            <div style={{
                backgroundColor: tokens.colorPaletteRoyalBlueForeground2,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteRoyalBlueForeground2
            </div>
            <div style={{
                backgroundColor: tokens.colorPaletteRoyalBlueBorderActive,
                width: "100%",
                height: "100%",
            }}>
                colorPaletteRoyalBlueBorderActive
            </div>
            {tokens.colorPaletteRedBorder1}
        </article>
    );
}

// @ts-ignore
function FunctionElement({title, description, imageSrc}) {
    return (
        <div style={{
            width: "100%",
            height: "45rem",
            borderRadius: "1.5rem",
            backgroundColor: "white",
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
        }}>
            <img src={imageSrc} alt="fds" style={{
                width: "17.4rem",
                backgroundColor: "oklch(0.985 0.002 247.839)",
                borderRadius: "9999rem",
                marginTop: "2.5rem",
            }}/>
            <b style={{marginTop: "1.3rem", fontSize: "2rem"}}>{title}</b>
            <p style={{marginTop: "1.7rem", fontSize: "1.4rem"}}>{description}</p>
        </div>
    );
}

import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
} from 'chart.js';

import { Bar } from 'react-chartjs-2';
import { Pen48Regular } from "@fluentui/react-icons/fonts";


function PhishingCostDiagram() {
    const chartRef = useRef(null);
    const containerRef = useRef(null);
    const [showChart, setShowChart] = useState(false);

    useEffect(() => {
        const observer = new IntersectionObserver(
            ([entry]) => {
                if (entry.isIntersecting) {
                    setShowChart(true);
                    observer.disconnect(); // Only trigger once
                }
            },
            { threshold: 1 }
        );

        if (containerRef.current) {
            observer.observe(containerRef.current);
        }

        return () => observer.disconnect();
    }, []);


    ChartJS.register(
        CategoryScale,
        LinearScale,
        BarElement,
        Title,
        Tooltip,
        Legend
    );

    const fontFamily = 'ui-sans-serif, system-ui, sans-serif, "Apple Color Emoji", "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji"';

    const options = {
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
                    callback: function (value: number) {
                        if (value === 0) return '';
                        return value / 1_000_000_000 + ' Mrd. $';
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

// 'colorPaletteRedBackground1' | 'colorPaletteRedBackground2' | 'colorPaletteRedBackground3' | 'colorPaletteRedForeground1' |
// 'colorPaletteRedForeground2' | 'colorPaletteRedForeground3' | 'colorPaletteRedForegroundInverted' | 'colorPaletteRedBorderActive'
// | 'colorPaletteRedBorder1' | 'colorPaletteRedBorder2';

    const data = {
        labels: ['2020', '2021', '2022', '2023', '2024'],
        datasets: [
            {
                label: 'Aus Phishing entstandene Schäden in den USA',
                data: showChart ? [
                    4_600_000_000,
                    7_000_000_000,
                    10_200_000_000,
                    12_600_000_000,
                    16_800_000_000,
                ] : [0,0,0,0,0],
                backgroundColor: '#d83b01',
            },
        ],
    };
    return (<div ref={containerRef}>
                <Bar ref={chartRef} options={options} data={data}/>
            </div>);
}


