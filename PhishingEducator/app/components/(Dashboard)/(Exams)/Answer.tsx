import { tokens } from '@fluentui/react-components';
import { CheckmarkStarburst24Filled, DismissCircle24Filled } from '@fluentui/react-icons';

import ExamStyles from '@styles/Exam.module.scss';

type AnswerProps = {
  answer: string;
  isSelected: boolean;
  isCorrect: boolean;
  showResults: boolean;
  onClick: () => void;
};

export default function Answer({
  answer,
  isSelected,
  isCorrect,
  showResults,
  onClick
}: AnswerProps) {
  return (
    <button
      type="button"
      className={`${ExamStyles.Exam__Answer}`}
      data-selected={isSelected}
      data-showresults={showResults}
      data-correct={isCorrect}
      onClick={onClick}
    >
      { showResults && (
        isCorrect ? 
        <CheckmarkStarburst24Filled
          color={tokens.colorStatusSuccessBackground3}
          className={ExamStyles.Exam__AnswerCorrectionIndicator}
        /> :
        isSelected &&
        <DismissCircle24Filled
          color={tokens.colorStatusDangerBackground3}
          className={ExamStyles.Exam__AnswerCorrectionIndicator}
        />
      ) }
      { answer }
    </button>
  );
}