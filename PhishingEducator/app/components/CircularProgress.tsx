import { Subtitle2Stronger, tokens } from "@fluentui/react-components";

type CircularProgressProps = {
  value: number;
  max?: number;
  size?: number;
  strokeWidth?: number;
};

export default function CircularProgress({ value, max = 100, size = 120, strokeWidth = 12 }: CircularProgressProps) {
  const radius = (size - strokeWidth) / 2;
  const circumference = 2 * Math.PI * radius;
  const offset = circumference - (value / max) * circumference;
  
  return (
    <div
      role="progressbar"
      aria-valuenow={value}
      aria-valuemin={0}
      aria-valuemax={max}
      style={{
        position: 'relative',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        width: size,
        height: size,
      }}
    >
      <svg
        width={size}
        height={size}
        style={{
          transform: 'rotate(-90deg)',
        }}
      >
        { /* Background circle */ }
        <circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          stroke="#e5e7eb"
          strokeWidth={strokeWidth}
          fill="transparent"
        />
        { /* Foreground circle */ }
        <circle
          cx={size / 2}
          cy={size / 2}
          r={radius}
          stroke={tokens.colorBrandBackground}
          strokeWidth={strokeWidth}
          fill="transparent"
          strokeDasharray={circumference}
          strokeDashoffset={offset - 0.00001}
          strokeLinecap="round"
          style={{
            transition: 'stroke-dashoffset 0.5s ease',
          }}
        />
      </svg>
      {/* Percentage Label */}
      <Subtitle2Stronger
        style={{
          position: 'absolute',
        }}
      >
        {value} / {max}
      </Subtitle2Stronger>
    </div>
  )
}