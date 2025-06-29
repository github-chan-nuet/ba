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
  INSERT INTO phishing_simulation_recognition_feature_value (id, "value", "level", user_instruction, recognition_feature_fk, content_category_fk)
  VALUES
    (val1_id, 'bank-of-switzerland.ch', 0, '', feat1_id, cc_id),
    (val2_id, 'bamk-of-switzerland.ch', 1, 'Phishing-Seiten nutzen oft URLs, die der echten Webseite sehr ähnlich sehen - zum Beispiel „bamk-of-switzerland.ch“ statt „bank-of-switzerland.ch“. Achte genau auf kleine Tippfehler oder ungewöhnliche Schreibweisen, um sicherzugehen, dass du wirklich auf der richtigen Seite bist. Im Zweifelsfall solltest du unter keinen Umständen dem Link folgen. Navigiere stattdessen manuell zur Webseite der Organisation, indem du die Webseite selbst im Suchfeld eintippst.', feat1_id, cc_id);


  -- Insert Content Templates
  INSERT INTO phishing_simulation_content_template (id, "subject", content, content_category_fk)
  VALUES (
    template1_id,
    'Verdächtige Aktivitäten in deinem E-Banking-Zugang erkannt!',
    $html$
<!DOCTYPE html>
<html lang="de">
<head>
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <title>Bank of Switzerland - Wichtiger Sicherheitshinweis</title>
  <style>
    body {
      margin:0;
      padding:0;
      font-family: Arial, sans-serif;
      background-color: #f4f4f4;
      color: #333333;
    }
    .container {
      width:100%;
      padding:20px 0;
    }
    .content {
      max-width:600px;
      margin:0 auto;
      background-color:#ffffff;
      border:1px solid #dddddd;
      border-radius:4px;
      padding:20px;
    }
    .header {
      text-align:center;
      padding-bottom:20px;
    }
    .logo {
      max-width:150px;
      margin-bottom:10px;
    }
    .button {
      display:inline-block;
      padding:12px 24px;
      margin-top:20px;
      background-color:#0056b3;
      color:#ffffff;
      text-decoration:none;
      border-radius:4px;
      font-weight:bold;
    }
    .footer {
      font-size:12px;
      color:#888888;
      text-align:center;
      margin-top:20px;
    }
  </style>
</head>
<body>
  <div class="container">
    <div class="content">
      <div class="header">
        <img src="https://securaware.ch/logos/bank-of-switzerland.png" alt="Bank of Switzerland" class="logo">
      </div>
      <p>Sehr geehrte Kundin, sehr geehrter Kunde,</p>
      <p>wir haben ungewöhnliche Aktivitäten in Ihrem Online-Banking-Konto festgestellt. Aus Sicherheitsgründen wurde der Zugriff vorübergehend eingeschränkt.</p>
      <p>Um Ihr Konto zu überprüfen und den vollen Zugriff wiederherzustellen, bestätigen Sie bitte Ihre Identität über den untenstehenden Link:</p>
      <p style="text-align:center;">
        {{EducationLink Verifizieren}}
      </p>
      <p>Falls Sie diese Aktivität nicht autorisiert haben, kontaktieren Sie uns bitte umgehend über den Kundenservice.</p>
      <p>Wir danken Ihnen für Ihre Aufmerksamkeit und Ihr Verständnis.</p>
      <p>Freundliche Grüße,<br>
      Ihr Sicherheitsteam der Bank of Switzerland</p>
      <div class="footer">
        Dies ist eine automatisch generierte Nachricht. Bitte antworten Sie nicht auf diese E-Mail.
      </div>
    </div>
  </div>
</body>
</html>
    $html$,
    cc_id
  );

END
$$;