CREATE TABLE users
(
    id         uuid NOT NULL,
    firstname  text NULL,
    lastname   text NULL,
    "password" bytea NULL,
    email      text NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email);
);


CREATE TABLE lesson_completions
(
    id        uuid NOT NULL,
    course_id uuid NULL,
    lesson_id uuid NULL,
    user_fk   uuid NULL,
    "time"    date NULL,
    CONSTRAINT lesson_completions_pkey PRIMARY KEY (id),
    CONSTRAINT fk_lesson_completions_user FOREIGN KEY (user_fk) REFERENCES public.users (id),
    -- Composite Unique Constraint
    CONSTRAINT unique_lesson_completion_per_usr UNIQUE (course_id, lesson_id, user_fk);
);