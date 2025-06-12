import { Link } from "react-router";
import type { Exam } from "@api/index";
import { Button, Divider, Spinner, Subtitle1 } from "@fluentui/react-components";
import CircularProgress from "@components/(Dashboard)/CircularProgress";

import ExamStyles from "@styles/Exam.module.scss";
import ExamProgressStyles from './ExamProgress.module.scss';

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
        <div className={ExamProgressStyles.ExamProgress__Actions}>
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