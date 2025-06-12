import { Body1, Title3 } from "@fluentui/react-components";
import type { Answer as AnswerType, Question } from "@api/index";
import Answer from "./Answer";

import ExamStyles from '@styles/Exam.module.scss';

type QuestionProps = {
  question: Question;
  answers: AnswerType[];
  selectedAnswers: string[];
  questionNr: number;
  totalQuestions: number;
  correctAnswers: string[];
  showResults: boolean;
  onSetAnswers: (questionId: string, answerId: string) => void;
};

export default function Question({
  question,
  answers,
  selectedAnswers,
  questionNr,
  totalQuestions,
  correctAnswers,
  showResults,
  onSetAnswers
}: QuestionProps) {
  return (
    <div className={ExamStyles.Exam__Question}>
      <Body1>
        Frage {questionNr} von {totalQuestions}
      </Body1>
      <Title3 className={ExamStyles.Exam__QuestionTitle}>
        { question.question }
      </Title3>
      <div className={ExamStyles.Exam__Answers}>
        { answers.map(answer => (
          <Answer
            key={answer.id}
            answer={answer.answer}
            isSelected={selectedAnswers.includes(answer.id)}
            isCorrect={correctAnswers.includes(answer.id)}
            showResults={showResults}
            onClick={() => onSetAnswers(question.id, answer.id)}
          />
        )) }
      </div>
    </div>
  );
}