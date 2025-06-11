import type { CompletedExam, CompletedQuestion } from "@api/index";

export const getWrongAnswerCountForQuestion = (question: CompletedQuestion): number => {
  const userSet = new Set(question.userAnswers);
  const correctSet = new Set(
    question.answers.filter(a => a.isCorrect).map(a => a.id)
  );

  let diffCount = 0;
  for (const id of userSet) {
    if (!correctSet.has(id)) diffCount++;
  }
  for (const id of correctSet) {
    if (!userSet.has(id)) diffCount++;
  }

  return diffCount;
}

export const getWrongAnswerCount = (exam: CompletedExam): number => {
  return exam.questions.reduce((total, question) => total + getWrongAnswerCountForQuestion(question), 0);
}

export const getAchievedXP = (exam: CompletedExam): number => {
  const scorePerQuestion = 100 / exam.questions.length;

  let score = 0;
  for (const question of exam.questions) {
    const reductionPerWrongAnswer = 1 / question.answers.length;
    const wrongAnswerCount = getWrongAnswerCountForQuestion(question);
    score += (1 - wrongAnswerCount * reductionPerWrongAnswer) * scorePerQuestion;
  }
  score = Math.round(score);

  return Math.round(score / 100 * 1000);
}

export const getTotalAchievableXP = (): number => 1000;