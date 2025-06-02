import { Button, Divider, Subtitle1, Subtitle2, Title1, tokens } from "@fluentui/react-components";
import { Link } from "react-router";
import { DismissCircle28Regular, Sparkle28Filled } from "@fluentui/react-icons";
import { getAchievedXP, getTotalAchievableXP, getWrongAnswerCount } from "../utils/exam";

import ExamStyles from "../styles/Exam.module.scss";
import type { CompletedExam } from "../api";

type ExamProgressProps = {
  completedExam: CompletedExam;
}

export default function ExamResults({ completedExam }: ExamProgressProps) {
  return (
    <div className={ExamStyles.Exam__Aside}>
      <div className={ExamStyles.Exam__AsideContainer}>
        <Subtitle1>Ergebnisse</Subtitle1>
        <div className={ExamStyles.Exam__ResultStatistics}>
          <div className={ExamStyles.Exam__ResultStatistic}>
            <div
              style={{
                display: 'flex',
                alignItems: 'center',
                gap: 6
              }}
            >
              <Sparkle28Filled color={tokens.colorStatusSuccessBackground3} />
              <div>
                <Title1>{ getAchievedXP(completedExam) }</Title1>
                <Subtitle2> / { getTotalAchievableXP() }</Subtitle2>
              </div>
            </div>
            <span
              style={{
                marginLeft: 30
              }}
            >
              Erreiche Punktzahl
            </span>
          </div>
          <div className={ExamStyles.Exam__ResultStatistic}>
            <div
              style={{
                display: 'flex',
                alignItems: 'center',
                gap: 6
              }}
            >
              <DismissCircle28Regular color={tokens.colorStatusDangerBackground3} />
              <div>
                <Title1>{ getWrongAnswerCount(completedExam) }</Title1>
              </div>
            </div>
            <span
              style={{
                marginLeft: 30
              }}
            >
              Falsche Antworten
            </span>
          </div>
        </div>
        <Divider />
        <div
          style={{
            display: 'flex',
            gap: tokens.spacingHorizontalS,
            justifyContent: 'space-between'
          }}
        >
          <Link to="/dashboard/exams" tabIndex={-1}>
            <Button appearance="secondary">
              Zurück
            </Button>
          </Link>
        </div>
      </div>
    </div>
  )
}