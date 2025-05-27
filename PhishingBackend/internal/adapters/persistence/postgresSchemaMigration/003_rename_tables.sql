ALTER TABLE users RENAME TO "user";
ALTER TABLE lesson_completions RENAME TO lesson_completion;
ALTER TABLE exams RENAME TO exam;
ALTER TABLE exam_questions RENAME TO exam_question;
ALTER TABLE exam_question_answers RENAME TO exam_question_answer;
ALTER TABLE exam_completions RENAME TO exam_completion;
ALTER TABLE exam_completion_answers RENAME TO exam_completion_answer;

ALTER TABLE user
    ALTER COLUMN firstname SET NOT NULL,
    ALTER COLUMN lastname SET NOT NULL,
    ALTER COLUMN "password" SET NOT NULL,
    ALTER COLUMN email SET NOT NULL;

ALTER TABLE lesson_completion
    ALTER COLUMN course_id SET NOT NULL,
    ALTER COLUMN lesson_id SET NOT NULL,
    ALTER COLUMN user_fk SET NOT NULL,
    ALTER COLUMN "time" SET NOT NULL;

