import { Body1, Title1, Title3, tokens } from "@fluentui/react-components";
import { completeExam, getCompletedExam, getExamsByExamId, type Answer, type CompletedExam, type Question } from "@api/index";
import type { Route } from "./+types/exam";
import { useState } from "react";
import useAuth from "@utils/auth/useAuth";
import { CheckmarkStarburst24Filled, DismissCircle24Filled } from "@fluentui/react-icons";
import ExamProgress from "@components/(Dashboard)/ExamProgress";
import ExamResults from "@components/(Dashboard)/ExamResults";

import ExamStyles from "@styles/Exam.module.scss";

type UserAnswer = {
  questionId: string;
  answers: Array<string>;
}
type UserAnswers = Array<UserAnswer>;

export const handle = async ({ params }: Route.ClientLoaderArgs) => {
  const { data: exam } = await getExamsByExamId({
    path: {
      examId: params.examId ?? ''
    }
  });
  if (exam) {
    return exam.title;
  }
}

export async function clientLoader({ params }: Route.ClientLoaderArgs) {
  const { data: completedExam } = await getCompletedExam({
    path: {
      examId: params.examId ?? ''
    }
  });
  const { data: exam } = await getExamsByExamId({
    path: {
      examId: params.examId ?? ''
    }
  });

  return { completedExam, exam };
}

export default function Exam({ loaderData }: Route.ComponentProps) {
  const { exam, completedExam } = loaderData;
  if (!exam) {
    throw new Response("Not Found", { status: 404 });
  }

  const { onExperienceGain } = useAuth();

  const [selectedAnswers, setSelectedAnswers] = useState<UserAnswers>(completedExam?.questions?.map(q => ({ questionId: q.id, answers: q.userAnswers })) ?? []);
  const [examCompletion, setExamCompletion] = useState<CompletedExam|null>(completedExam ?? null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const readyForSubmission = exam.questions.every(q => selectedAnswers.find(sa => sa.questionId === q.id && sa.answers.length > 0));

  const toggleAnswer = (questionId: string, answerId: string) => {
    if (examCompletion) return;
    setSelectedAnswers((prev) => {
      const existing = prev.find((ua) => ua.questionId === questionId);

      const userAnswer = existing
        ? {
          ...existing,
          answers: existing.answers.includes(answerId)
            ? existing.answers.filter((a) => a !== answerId)
            : [...existing.answers, answerId],
        }
        : {
          questionId,
          answers: [answerId],
        };

      return [...prev.filter((ua) => ua.questionId !== questionId), userAnswer];
    });
  };

   const handleSubmit = async () => {
    if (!readyForSubmission) return;

    setIsLoading(true);
    const { data: xpGain, error } = await completeExam({
      path: {
        examId: exam.id
      },
      body: selectedAnswers,
    });

    if (xpGain) {
      onExperienceGain(xpGain.newExperienceGained, xpGain.newLevel);

      const { data: completedExam, error } = await getCompletedExam({
        path: {
          examId: exam.id
        }
      });
      if (!error) {
        setExamCompletion(completedExam);
      }
    }
    setIsLoading(false);
  }

  return (
    <div>
      <Title1>{ exam.title }</Title1>
      <div className={`${ExamStyles.Exam__Container} ${ExamStyles.Exam__Container_inverted}`}>
        <div className={ExamStyles.Exam__Questions}>
          { exam.questions.map((question, idx) => (
            <Question
              key={question.id}
              question={question}
              questionNr={idx + 1}
              totalQuestions={exam.questions.length}
              answers={question.answers}
              selectedAnswers={
                examCompletion?.questions?.find(q => q.id === question.id)?.userAnswers ??
                selectedAnswers.find(ua => ua.questionId === question.id)?.answers ?? []
              }
              correctAnswers={examCompletion?.questions?.find(q => q.id === question.id)?.answers.filter(a => a.isCorrect)?.map(a => a.id) ?? []}
              showResults={!!examCompletion}
              onSetAnswers={toggleAnswer}
            />
          )) }
        </div>
        { examCompletion ? (
          <ExamResults
            completedExam={examCompletion}
          />
        ) : (
          <ExamProgress
            exam={exam}
            selectedAnswers={selectedAnswers}
            submissionDisabled={!readyForSubmission}
            showSpinner={isLoading}
            onSubmit={handleSubmit}
          />
        ) }
      </div>
    </div>
  );
}

type QuestionProps = {
  question: Question;
  answers: Answer[];
  selectedAnswers: string[];
  questionNr: number;
  totalQuestions: number;
  correctAnswers: string[];
  showResults: boolean;
  onSetAnswers: (questionId: string, answerId: string) => void;
};

function Question({
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
            isDisabled={false}
            showResults={showResults}
            onClick={() => onSetAnswers(question.id, answer.id)}
          />
        )) }
      </div>
    </div>
  );
}

type AnswerProps = {
  answer: string;
  isSelected: boolean;
  isCorrect: boolean;
  isDisabled: boolean;
  showResults: boolean;
  onClick: () => void;
};

function Answer({
  answer,
  isSelected,
  isCorrect,
  isDisabled,
  showResults,
  onClick
}: AnswerProps) {
  return (
    <div
      className={ExamStyles.Exam__Answer}
      style={{
        cursor: isDisabled ? 'not-allowed' : 'pointer',
        border: `1px solid ${isSelected ? (showResults ? (isCorrect ? tokens.colorStatusSuccessBackground3 : tokens.colorStatusDangerBackground3) : tokens.colorBrandBackground) : '#d2d2d7'}`,
        boxShadow: isSelected ? `0 0 0 4px ${showResults ? (isCorrect ? tokens.colorStatusSuccessBackground2 : tokens.colorStatusDangerBackground2) : tokens.colorBrandBackground2Pressed}` : 'none',
      }}
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
    </div>
  );
}