DO $$
DECLARE
  exam1_id UUID := uuid_generate_v4();
  exam2_id UUID := uuid_generate_v4();
  exam1_q1_id UUID := uuid_generate_v4();
  exam1_q2_id UUID := uuid_generate_v4();
  exam1_q3_id UUID := uuid_generate_v4();
  exam2_q1_id UUID := uuid_generate_v4();
  exam2_q2_id UUID := uuid_generate_v4();
  exam2_q3_id UUID := uuid_generate_v4();
  exam2_q4_id UUID := uuid_generate_v4();
BEGIN

  -- Insert exams
  INSERT INTO exam (id, title, description)
  VALUES
    (exam1_id, 'Grundlagen', 'Prüfe, dass du verstanden hast, was man unter dem Begriff Phishing versteht und weshalb dies so gefährlich ist.'),
    (exam2_id, 'Angriffsvektoren', 'Verstehe, welche Angriffsmethoden für Phishing oft verwendet werden und welche Merkmale diese aufweisen.');

  -- Insert questions into exam 1
  INSERT INTO exam_question (id, exam_fk, question)
  VALUES
    (exam1_q1_id, exam1_id, 'Welche der folgenden Aussagen beschreiben mögliche Auswirkungen eines erfolgreichen Phishing-Angriffs?'),
    (exam1_q2_id, exam1_id, 'Was versteht man unter Phishing?'),
    (exam1_q3_id, exam1_id, 'Über welche Kanäle kann Phishing verbreitet werden?');

  -- Insert answers for exam 1, question 1
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam1_q1_id, 'Finanzielle Verluste für Einzelpersonen oder Unternehmen', true),
    (uuid_generate_v4(), exam1_q1_id, 'Verbesserung der IT-Sicherheit durch erhöhte Aufmerksamkeit', false),
    (uuid_generate_v4(), exam1_q1_id, 'Identitätsdiebstahl', true),
    (uuid_generate_v4(), exam1_q1_id, 'Rechtliche Konsequenzen für betroffene Unternehmen', true),
    (uuid_generate_v4(), exam1_q1_id, 'Automatische Löschung der Phishing-Mail durch das System', false);

  -- Insert answers for exam 1, question 2
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam1_q2_id, 'Der Versuch, durch gefälschte Nachrichten an vertrauliche Informationen zu gelangen', true),
    (uuid_generate_v4(), exam1_q2_id, 'Ein Verfahren zur Verschlüsselung von E-Mails', false),
    (uuid_generate_v4(), exam1_q2_id, 'Ein IT-Sicherheitsverfahren zur Mitarbeiterschulung', false),
    (uuid_generate_v4(), exam1_q2_id, 'Eine Form von Cyberangriff, bei dem sich Angreifende als vertrauenswürdige Quelle ausgeben', true);

  -- Insert answers for exam 1, question 3
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam1_q3_id, 'E-Mail', true),
    (uuid_generate_v4(), exam1_q3_id, 'Soziale Netzwerke', true),
    (uuid_generate_v4(), exam1_q3_id, 'SMS / Messenger-Dienste', true),
    (uuid_generate_v4(), exam1_q3_id, 'Klassische Briefe', true),
    (uuid_generate_v4(), exam1_q3_id, 'Telefonanrufe', true);

  -- Insert questions into exam 2
  INSERT INTO exam_question (id, exam_fk, question)
  VALUES
    (exam2_q1_id, exam2_id, 'In welcher der Methoden werden Links verwendet um an vertrauliche Daten zu gelangen?'),
    (exam2_q2_id, exam2_id, 'Welcher der Angriffsvektoren wird auch als Vishing bezeichnet?'),
    (exam2_q3_id, exam2_id, 'Weshalb wird Smishing (Phishing über SMS) als besonders gefährlich eingestuft?'),
    (exam2_q4_id, exam2_id, 'Was hilft dir, um eine verdächtige E-Mail korrekt zu erkennen?');

  -- Insert answers for exam 2, question 1
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam2_q1_id, 'E-Mail', true),
    (uuid_generate_v4(), exam2_q1_id, 'SMS', true),
    (uuid_generate_v4(), exam2_q1_id, 'Telefonanruf', false);

  -- Insert answers for exam 2, question 2
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam2_q2_id, 'E-Mail', false),
    (uuid_generate_v4(), exam2_q2_id, 'SMS', false),
    (uuid_generate_v4(), exam2_q2_id, 'Telefonanruf', true);

  -- Insert answers for exam 2, question 3
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam2_q3_id, 'Weil über SMS keine Dringlichkeit ausgedrückt werden kann', false),
    (uuid_generate_v4(), exam2_q3_id, 'Weil mobile Geräte oft schlechter geschützt sind als klassische Computer', true),
    (uuid_generate_v4(), exam2_q3_id, 'Weil SMS automatisch auf dem Gerät gespeichert werden und dadurch Schadsoftware installiert werden kann', false),
    (uuid_generate_v4(), exam2_q3_id, 'Weil SMS ausschließlich von vertrauenswürdigen Absendern stammen dürfen und Angreifer das gezielt ausnutzen', false);

  -- Insert answers for exam 2, question 4
  INSERT INTO exam_question_answer (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), exam2_q4_id, 'Technisches Grundwissen', true),
    (uuid_generate_v4(), exam2_q4_id, 'Aufmerksamkeit', true),
    (uuid_generate_v4(), exam2_q4_id, 'Die Webseite hinter dem Link auf verdächtige Elemente prüfen', false),
    (uuid_generate_v4(), exam2_q4_id, 'Skepsis', true);

END
$$;