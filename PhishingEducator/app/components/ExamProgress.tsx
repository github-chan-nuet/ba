import { Button, Divider, Spinner, Subtitle1, tokens } from "@fluentui/react-components";
import { Link } from "react-router";
import CircularProgress from "./(Dashboard)/CircularProgress";
import type { Exam } from "@api/index";

import ExamStyles from "@styles/Exam.module.scss";

type UserAnswer = {
  questionId: string;
  answers: Array<string>;
}
type UserAnswers = Array<UserAnswer>;

type ExamProgressProps = {
  exam: Exam;
  selectedAnswers: UserAnswers;
  submissionDisabled: boolean;
  showSpinner: boolean;
  onSubmit: () => void;
};

export default function ExamProgress({ exam, selectedAnswers, submissionDisabled, showSpinner, onSubmit }: ExamProgressProps) {
  const answeredQuestions = exam.questions.filter((q) => (selectedAnswers.find(q2 => q2.questionId === q.id)?.answers.length ?? 0) > 0).length;
  
  return (
    <div className={ExamStyles.Exam__Aside}>
      <div className={ExamStyles.Exam__AsideContainer}>
        <Subtitle1>Fortschritt</Subtitle1>
        <CircularProgress
          value={answeredQuestions}
          max={exam.questions.length}
          ariaLabel="Fortschritt in der Prüfung"
        />
        <Divider />
        <div
          style={{
            display: 'flex',
            gap: tokens.spacingHorizontalS,
            justifyContent: 'space-between'
          }}
        >
          <Link to="/dashboard/exams">
            <Button appearance="secondary">
              Zurück
            </Button>
          </Link>
          <Button
            appearance="primary"
            onClick={onSubmit}
            disabled={submissionDisabled}
          >
            { showSpinner ? <Spinner size="extra-small" /> : "Absenden" }
          </Button>
        </div>
      </div>
    </div>
  );
}