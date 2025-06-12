import { Subtitle2Stronger, tokens } from "@fluentui/react-components";

import CircularProgressStyles from './CircularProgress.module.scss';

type CircularProgressProps = {
  value: number;
  max?: number;
  size?: number;
  strokeWidth?: number;
  ariaLabel: string;
};

export default function CircularProgress({ value, max = 100, size = 120, strokeWidth = 12, ariaLabel }: CircularProgressProps) {
  const radius = (size - strokeWidth) / 2;
  const circumference = 2 * Math.PI * radius;
  const offset = circumference - (value / max) * circumference;
  
  return (
    <div
      className={CircularProgressStyles.CircularProgress}
      role="progressbar"
      aria-valuenow={value}
      aria-valuemin={0}
      aria-valuemax={max}
      aria-label={ariaLabel}
      style={{
        width: size,
        height: size,
      }}
    >
      <svg
        className={CircularProgressStyles.CircularProgress__Vector}
        width={size}
        height={size}
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
          className={CircularProgressStyles.CircularProgress__VectorCircle}
        />
      </svg>
      {/* Percentage Label */}
      <Subtitle2Stronger className={CircularProgressStyles.CircularProgress__Label}>
        {value} / {max}
      </Subtitle2Stronger>
    </div>
  )
}