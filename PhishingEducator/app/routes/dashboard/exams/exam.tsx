import type { Route } from "./+types/exam";
import { useState } from "react";
import { Helmet } from "react-helmet-async";
import { Title1 } from "@fluentui/react-components";
import useAuth from "@utils/auth/useAuth";
import { completeExam, getCompletedExam, getExamsByExamId, type CompletedExam } from "@api/index";
import ExamProgress from "@components/(Dashboard)/(Exams)/ExamProgress";
import ExamResults from "@components/(Dashboard)/(Exams)/ExamResults";
import Question from "@components/(Dashboard)/(Exams)/Question";

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
  const { onExperienceGain } = useAuth();

  const [selectedAnswers, setSelectedAnswers] = useState<UserAnswers>(completedExam?.questions?.map(q => ({ questionId: q.id, answers: q.userAnswers })) ?? []);
  const [examCompletion, setExamCompletion] = useState<CompletedExam|null>(completedExam ?? null);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  
  if (!exam) {
    throw new Response("Not Found", { status: 404 });
  }

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
    const { data: xpGain } = await completeExam({
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
    <>
      <Helmet>
        <title>Securaware - { exam.title } Prüfung</title>
      </Helmet>
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
    </>
  );
}
