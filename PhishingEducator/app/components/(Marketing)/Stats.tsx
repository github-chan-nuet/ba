import Chart from 'react-apexcharts';

import StatsStyles from './Stats.module.scss';

export default function Stats() {
  return (
    <section className={StatsStyles.Stats}>
      <div className={StatsStyles.Stats__Grid}>
        <div className={StatsStyles.Stats__Text}>
          <h2 className={StatsStyles.Stats__Title}>Phishing-Angriffe nehmen <strong>rasant</strong> zu</h2>
          <p className={StatsStyles.Stats__Paragraph}>Phishing-Angriffe kommen immer häufiger vor, dies bestätigt die amerikanische Sicherheitsbehörde FBI. Die Graphik weist die aus Phishing entstandene Schäden in den USA auf. Der Trend ist klar - Phishing-Angriffe häufen sich.</p>
        </div>
        <div className={StatsStyles.Stats__Chart}>
          <CostDiagram />
        </div>
      </div>
    </section>
  )
}

function CostDiagram() {
  const options = {
    chart: {
      fontFamily: "Inter, sans-serif",
      dropShadow: {
        enabled: false,
      },
      toolbar: {
        show: false,
      },
    },
    tooltip: {
      shared: true,
      intersect: false,
      style: {
        fontFamily: "Inter, sans-serif",
      },
    },
    fill: {
      type: "gradient",
      gradient: {
        opacityFrom: 0.55,
        opacityTo: 0,
        shade: "#1C64F2",
        gradientToColors: ["#1C64F2"],
      },
    },
    dataLabels: {
      enabled: false,
    },
    stroke: {
      width: 6,
    },
    grid: {
      show: true,
      strokeDashArray: 4,
      padding: {
        left: 2,
        right: 2,
        top: -26
      },
    },
    legend: {
      show: false,
    },
    xaxis: {
      categories: ['2020', '2021', '2022', '2023', '2024'],
      floating: false,
      labels: {
        show: true,
      },
      axisBorder: {
        show: false,
      },
      axisTicks: {
        show: false,
      },
    },
    yaxis: {
      show: true,
      labels: {
        show: true,
        formatter: (value: string | number) => {
          if (value === 0 || typeof value === 'string') return '';
          return value / 1_000_000_000 + ' Mrd. $';
        }
      }
    },
  }
  const series = [
    {
      name: 'Aus Phishing entstandene Schäden in den USA',
      data: [4_600_000_000, 7_000_000_000, 10_200_000_000, 12_600_000_000, 16_800_000_000]
    }
  ]

  return (
    <Chart
      type="area"
      options={options}
      series={series}
      height="100%"
    />
  )
}