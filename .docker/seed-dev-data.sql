DO $$
DECLARE
  exam_id UUID := uuid_generate_v4();
  q1_id UUID := uuid_generate_v4();
  q2_id UUID := uuid_generate_v4();
  q3_id UUID := uuid_generate_v4();
BEGIN

  -- Insert exam
  INSERT INTO exams (id, title, description)
  VALUES (exam_id, 'Grundlagen', 'Prüfe, dass du verstanden hast, was man unter dem Begriff Phishing versteht und weshalb dies so gefährlich ist.');

  -- Insert questions
  INSERT INTO exam_questions (id, exam_fk, question)
  VALUES
    (q1_id, exam_id, 'Welche der folgenden Aussagen beschreiben mögliche Auswirkungen eines erfolgreichen Phishing-Angriffs?'),
    (q2_id, exam_id, 'Was versteht man unter Phishing?'),
    (q3_id, exam_id, 'Über welche Kanäle kann Phishing verbreitet werden?');

  -- Insert answers for question 1
  INSERT INTO exam_question_answers (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), q1_id, 'Finanzielle Verluste für Einzelpersonen oder Unternehmen', true),
    (uuid_generate_v4(), q1_id, 'Verbesserung der IT-Sicherheit durch erhöhte Aufmerksamkeit', false),
    (uuid_generate_v4(), q1_id, 'Identitätsdiebstahl', true),
    (uuid_generate_v4(), q1_id, 'Rechtliche Konsequenzen für betroffene Unternehmen', true),
    (uuid_generate_v4(), q1_id, 'Automatische Löschung der Phishing-Mail durch das System', false);

  -- Insert answers for question 2
  INSERT INTO exam_question_answers (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), q2_id, 'Der Versuch, durch gefälschte Nachrichten an vertrauliche Informationen zu gelangen', true),
    (uuid_generate_v4(), q2_id, 'Eine Methode zur sicheren Datenübertragung', false),
    (uuid_generate_v4(), q2_id, 'Ein IT-Sicherheitsverfahren zur Mitarbeiterschulung', false),
    (uuid_generate_v4(), q2_id, 'Eine Form von Cyberangriff, bei dem sich Angreifende als vertrauenswürdige Quelle ausgeben', true);

  -- Insert answers for question 3
  INSERT INTO exam_question_answers (id, question_fk, answer, is_correct)
  VALUES
    (uuid_generate_v4(), q3_id, 'E-Mail', true),
    (uuid_generate_v4(), q3_id, 'Soziale Netzwerke', true),
    (uuid_generate_v4(), q3_id, 'SMS / Messenger-Dienste', true),
    (uuid_generate_v4(), q3_id, 'Klassische Briefe', false),
    (uuid_generate_v4(), q3_id, 'Telefonanrufe', true);

END
$$;