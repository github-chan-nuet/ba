import { Button, Display, tokens } from "@fluentui/react-components";
// import * as FluentIcons from "@fluentui/react-icons";
import MarketingStyles from "../../styles/Marketing.module.scss";
import hacker from '../../assets/images/hacker.png';
import phishing from '../../assets/images/phishing.png';
import lines from '../../assets/images/lines.svg';
import {ArrowRight24Regular} from "@fluentui/react-icons";
import {
    LineChart,
    DataVizPalette,
} from "@fluentui/react-charts";

export default function Home() {
    // const IconComponent = FluentIcons["LineHorizontal128Filled"];
  return (
    <>
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
          }} />
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
              <hr style={{height:"0.2rem", width: "35rem", backgroundColor: "black", color: "black", marginBottom: "4rem"}} />
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
                     backgroundImage: "linear-gradient(to bottom, oklch(0.985 0.002 247.839), #48d5ff)"}}
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
                <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing} />

                <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing} />

                <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing} />

                <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing} />

                <FunctionElement title={"Phishing"} description={"Das ist so"} imageSrc={phishing} />

            </div>
        </section>
        <section className={MarketingStyles.Section} style={{
            height: "80vh",
            backgroundColor: "#fff",
            display: "grid",
            gridTemplateColumns: "1fr 1fr",
            gridTemplateRows: "1fr",
            gap: "2rem",
            alignContent: "space-around",
            // border: "1px solid black",
        }}>
            <div style={{
                backgroundColor: tokens.colorPaletteBlueBorderActive,
                backgroundImage: "url(\""+lines+"\")",
                backgroundRepeat: "no-repeat",
                height: "100%",
                borderRadius: "2rem",
            }} />
            <div>
                <h3 style={{fontSize: "2rem"}}>Schütze dich jetzt</h3>
                <p>Lorem Impsum</p>
                <Button
                    appearance="primary"
                    icon={<ArrowRight24Regular />}
                    onClick={() => {}}>
                    Schütze dich jetzt
                </Button>
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
        </article>

        <Ah />
    </>
  );
}

function Ah() {
    const data: IChartProps = {
        chartTitle: 'Line Chart',
        lineChartData: [
            {
                legend: 'All',
                data: [
                    {
                        x: new Date('2020-01-01T00:00:00.000Z'),
                        y: 4_600_000_000,
                    },
                    {
                        x: new Date('2021-01-01T00:00:00.000Z'),
                        y: 7_000_000_000,
                    },
                    {
                        x: new Date('2022-01-01T00:00:00.000Z'),
                        y: 10_200_000_000,
                    },
                    {
                        x: new Date('2023-01-01T00:00:00.000Z'),
                        y: 12_600_000_000,
                    },
                    {
                        x: new Date('2024-01-01T00:00:00.000Z'),
                        y: 16_800_000_000,
                    },
                ],
                color: DataVizPalette.color4,
                lineOptions: {
                    lineBorderWidth: '4',
                },
            },
        ],
    };

    return (
        <LineChart
            culture={window.navigator.language}
            data={data}
            legendsOverflowText={'Overflow Items'}
            yMinValue={0}
            yMaxValue={18_000_000}
            // height={this.state.height}
            // width={this.state.width}
            xAxisTickCount={10}
            // allowMultipleShapesForPoints={this.state.allowMultipleShapes}
            enablePerfOptimization={true}
            yAxisTitle={'Phishing verursachte Schäden'}
            xAxisTitle={'Jahr'}
            // useUTC={this.state.useUTC}
        />
    );
}

// @ts-ignore
function FunctionElement({title, description, imageSrc} ) {
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
            }} />
            <b style={{marginTop: "1.3rem", fontSize: "2rem"}}>{title}</b>
            <p style={{marginTop: "1.7rem", fontSize: "1.4rem"}}>{description}</p>
        </div>
    );
}