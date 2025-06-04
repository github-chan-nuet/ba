CREATE TABLE phishing_simulation_content_category
(
  id     uuid NOT NULL,
  "name" text NULL,
  
  CONSTRAINT phishing_simulation_content_category_pkey PRIMARY KEY (id)
);

CREATE TABLE phishing_simulation_recognition_feature
(
  id                   uuid NOT NULL,
  "name"               text NOT NULL,
  is_always_applicable bool NOT NULL,
  user_instruction     text NULL,

  CONSTRAINT phishing_simulation_recognition_feature_pkey PRIMARY KEY (id)
);

CREATE TABLE phishing_simulation_recognition_feature_value
(
  id                     uuid NOT NULL,
  "value"                text NOT NULL,
  "level"                int  NOT NULL,
  recognition_feature_fk uuid NOT NULL,
  content_category_fk    uuid NOT NULL,

  CONSTRAINT phishing_simulation_recognition_feature_value_pkey PRIMARY KEY (id),
  CONSTRAINT fk_phishing_simulation_recognition_feature_value_recognition_feature FOREIGN KEY (recognition_feature_fk) REFERENCES public.phishing_simulation_recognition_feature (id),
  CONSTRAINT fk_phishing_simulation_recognition_feature_value_content_category FOREIGN KEY (content_category_fk) REFERENCES public.phishing_simulation_content_category (id)
);

CREATE TABLE phishing_simulation_content_template
(
  id                  uuid NOT NULL,
  "subject"           text NOT NULL,
  content             text NOT NULL,
  content_category_fk uuid NOT NULL,

  CONSTRAINT phishing_simulation_content_template_pkey PRIMARY KEY (id),
  CONSTRAINT fk_phishing_simulation_content_template_content_category FOREIGN KEY (content_category_fk) REFERENCES public.phishing_simulation_content_category (id)
);

CREATE TABLE phishing_simulation_run
(
  id          uuid NOT NULL,
  user_fk     uuid NOT NULL,
  template_fk uuid NOT NULL,
  sender_addr text NOT NULL,
  sender_name text NOT NULL,
  "subject"   text NOT NULL,
  content     text NOT NULL,
  sent_at     timestamptz NULL,
  opened_at   timestamptz NULL,

  CONSTRAINT phishing_simulation_run_pkey PRIMARY KEY (id),
  CONSTRAINT fk_phishing_simulation_run_user FOREIGN KEY (user_fk) REFERENCES public.user (id),
  CONSTRAINT fk_phishing_simulation_run_template FOREIGN KEY (template_fk) REFERENCES public.phishing_simulation_content_template (id)
);

CREATE TABLE phishing_simulation_user_vulnerability
(
  id                     uuid NOT NULL,
  user_fk                uuid NOT NULL,
  score                  float NOT NULL,
  content_category_fk    uuid NOT NULL,
  recognition_feature_fk uuid NOT NULL,

  CONSTRAINT phishing_simulation_user_vulnerability_pkey PRIMARY KEY (id),
  CONSTRAINT fk_phishing_simulation_user_vulnerability_user FOREIGN KEY (user_fk) REFERENCES public.user (id),
  CONSTRAINT fk_phishing_simulation_user_vulnerability_content_category FOREIGN KEY (content_category_fk) REFERENCES public.phishing_simulation_content_category (id),
  CONSTRAINT fk_phishing_simulation_user_vulnerability_recognition_feature FOREIGN KEY (recognition_feature_fk) REFERENCES public.phishing_simulation_recognition_feature (id)
);