CREATE TABLE exams
(
    id uuid NOT NULL,
    CONSTRAINT exams_pkey PRIMARY KEY (id)
);

CREATE TABLE exam_questions
(
    id       uuid NOT NULL,
    exam_fk  uuid NOT NULL,
    question text NOT NULL,
    CONSTRAINT exam_questions_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exam_questions_exam FOREIGN KEY (exam_fk) REFERENCES public.exams (id)
);

CREATE TABLE exam_question_answers
(
    id          uuid    NOT NULL,
    question_fk uuid    NOT NULL,
    answer      text    NOT NULL,
    is_correct  boolean NOT NULL,
    CONSTRAINT exam_question_answers_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exam_question_answers_exam_question FOREIGN KEY (question_fk) REFERENCES public.exam_questions (id)
);

CREATE TABLE exam_completions
(
    id           uuid    NOT NULL,
    exam_fk      uuid    NOT NULL,
    user_fk      uuid    NOT NULL,
    completed_at date    NOT NULL,
    score        integer NOT NULL,
    CONSTRAINT exam_completions_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exam_completions_exam FOREIGN KEY (exam_fk) REFERENCES public.exams (id),
    CONSTRAINT fk_exam_completions_user FOREIGN KEY (user_fk) REFERENCES public.users (id),
    -- Composite Unique Constraint
    CONSTRAINT unique_exam_completion_per_usr UNIQUE (exam_fk, user_fk)
);

CREATE TABLE exam_completion_answers
(
    id           uuid NOT NULL,
    exam_comp_fk uuid NOT NULL,
    answer_fk    uuid NOT NULL,
    CONSTRAINT exam_completion_answers_pkey PRIMARY KEY (id),
    CONSTRAINT fk_exam_completion_answers_exam_completion FOREIGN KEY (exam_comp_fk) REFERENCES public.exam_completions (id),
    CONSTRAINT fk_exam_completion_answers_answer FOREIGN KEY (answer_fk) REFERENCES public.exam_question_answers (id)
);