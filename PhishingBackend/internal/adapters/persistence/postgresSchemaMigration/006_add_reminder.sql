CREATE TABLE reminder_email_template
(
    id       integer NOT NULL,
    subject  text  NOT NULL,
    template text  NOT NULL,
    CONSTRAINT reminder_email_template_pkey PRIMARY KEY (id)
);

CREATE TABLE reminder
(
    id          uuid    NOT NULL,
    user_fk     uuid    NOT NULL,
    sent_time   date    NOT NULL,
    template_fk integer NOT NULL,
    count       integer NOT NULL,
    CONSTRAINT reminder_pkey PRIMARY KEY (id),
    CONSTRAINT fk_reminder_user FOREIGN KEY (user_fk) REFERENCES public.user (id),
    CONSTRAINT fk_reminder_reminder_email_template FOREIGN KEY (template_fk) REFERENCES public.reminder_email_template (id)
);

INSERT INTO reminder_email_template
VALUES (0, 'Bist du für den nächsten Phishing-Versuch gewappnet?',
E'Hallo {{ .Firstname }}\r\n\r\nCyberkriminelle schlafen nicht – und Phishing wird immer raffinierter.\r\nBesuche jetzt securaware.ch, frische dein Wissen auf und bleib einen Schritt voraus.\r\n\r\nStarte deine nächste Phishing-Simulation oder mach einen kurzen Selbsttest!\r\n\r\nBleib sicher.\r\nDein Securaware-Team');


INSERT INTO reminder_email_template
VALUES (1, 'Phishing wird immer cleverer. Und du?', 'Hallo {{ .Firstname }} \r\n \r\n
Phishing-Angriffe sind heute kaum noch zu erkennen – ausser du bist vorbereitet. \r\n
Komm zurück und hol dir die aktuellsten Tipps auf securaware.ch. \r\n \r\n

Schütze dich mit Wissen. \r\n \r\n

Dein Securaware-Team');


INSERT INTO reminder_email_template
VALUES (2, 'Securaware vermisst dich', 'Hallo {{ .Firstname }} \r\n \r\n
Es dauert nur wenige Minuten, um dich wieder fit gegen Phishing zu machen. \r\n
Komm zurück zu securaware.ch. \r\n \r\n

Lerne, wie du gefährliche Mails erkennst – bevor es zu spät ist. \r\n \r\n

Wir freuen uns auf dich! \r\n
Dein Securaware-Team');


INSERT INTO reminder_email_template
VALUES (3, 'Ein Klick kann den Unterschied machen', 'Hallo {{ .Firstname }} \r\n \r\n
Phishing kann jeden treffen. Aber nicht jeden muss es erwischen. \r\n
- Hol dir das nötige Wissen auf securaware.ch \r\n \r\n
- Erkenne gefährliche E-Mails
- Vermeide teure Fehler \r\n \r\n
Mach den Unterschied mit einem Klick.! \r\n \r\n

Sicher bleiben. \r\n
Dein Securaware-Team');


INSERT INTO reminder_email_template
VALUES (4, 'Phishing? Hoffen reicht nicht', 'Hallo {{ .Firstname }} \r\n \r\n
Phishing erkennen ist kein Zufall, sondern eine Frage des Wissens. \r\n \r\n
Komm jetzt zurück auf securaware.ch und lerne, wie du dich und deine Daten zuverlässig schützt. \r\n \r\n

Wir zeigen dir, wie’s geht – einfach und kostenlos! \r\n \r\n

Dein Securaware-Team');

INSERT INTO reminder_email_template
VALUES (4, '', $$fsdfsd$$);
