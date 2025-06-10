DO $$
DECLARE
  cc_id UUID := uuid_generate_v4();
  feat1_id UUID := uuid_generate_v4();
  feat2_id UUID := uuid_generate_v4();
  val1_id UUID := uuid_generate_v4();
  val2_id UUID := uuid_generate_v4();
  template1_id UUID := uuid_generate_v4();
BEGIN

  -- Insert Content Category
  INSERT INTO phishing_simulation_content_category (id, "name")
  VALUES (cc_id, 'Suspicous-E-Banking-Access');

  -- Insert Recognition Features
  INSERT INTO phishing_simulation_recognition_feature (id, "name", is_always_applicable, title, user_instruction)
  VALUES (feat1_id, 'Domain', true, 'URL', 'Beim Umgang mit E-Mails sollte man Links stets mit Vorsicht behandeln, da Phishing-Versuche oft über täuschend echte URLs erfolgen. Auch wenn eine E-Mail seriös wirkt, kann sich hinter einem Link eine manipulierte Adresse verbergen. Schon kleine Abweichungen oder ungewöhnliche Domains sind Warnzeichen. Das Schloss-Symbol oder „https“ bieten keine Garantie für Sicherheit, da auch betrügerische Seiten verschlüsselt sein können. Geben Sie persönliche Daten niemals über einen Link in einer E-Mail ein. Im Zweifel ist es sicherer, die Webadresse selbst im Browser einzugeben.');

  -- Insert Recognition Feature Values
  INSERT INTO phishing_simulation_recognition_feature_value (id, "value", "level", recognition_feature_fk, content_category_fk)
  VALUES
    (val1_id, 'bank-of-switzerland.ch', 0, feat1_id, cc_id),
    (val2_id, 'bamk-of-switzerland.ch', 1, feat1_id, cc_id);


  -- Insert Content Templates
  INSERT INTO phishing_simulation_content_template (id, "subject", content, content_category_fk)
  VALUES (template1_id, 'Verdächtige Aktivitäten in deinem E-Banking-Zugang erkannt!', '<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8" /><meta name="viewport" content="width=device-width, initial-scale=1.0" /></head><body>Hallo,<br /><br />Wir haben verdächtige Aktivitäten in Ihrem E-Banking-Konto festgestellt. Klicken Sie {{EducationLink hier}} um Ihr Konto zu schützen!</body></html>', cc_id);

END
$$;