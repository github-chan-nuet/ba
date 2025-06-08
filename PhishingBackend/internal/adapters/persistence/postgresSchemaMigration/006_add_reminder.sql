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

-- INSERT INTO reminder_email_template
-- VALUES (0, 'Bist du für den nächsten Phishing-Versuch gewappnet?',
-- E'Hallo {{ .Firstname }}\r\n\r\nCyberkriminelle schlafen nicht – und Phishing wird immer raffinierter.\r\nBesuche jetzt securaware.ch, frische dein Wissen auf und bleib einen Schritt voraus.\r\n\r\nStarte deine nächste Phishing-Simulation oder mach einen kurzen Selbsttest!\r\n\r\nBleib sicher.\r\nDein Securaware-Team');
--
--
-- INSERT INTO reminder_email_template
-- VALUES (1, 'Phishing wird immer cleverer. Und du?',
--         E'Hallo {{ .Firstname }}\r\n\r\nPhishing-Angriffe sind heute kaum noch zu erkennen – ausser du bist vorbereitet.\r\nKomm zurück und hol dir die aktuellsten Tipps auf securaware.ch.\r\n\r\nSchütze dich mit Wissen.\r\n\r\nDein Securaware-Team');
--
--
-- INSERT INTO reminder_email_template
-- VALUES (2, 'Securaware vermisst dich',
--         E'Hallo {{ .Firstname }}\r\n\r\nEs dauert nur wenige Minuten, um dich wieder fit gegen Phishing zu machen. \r\nKomm zurück zu securaware.ch.\r\n\r\nLerne, wie du gefährliche Mails erkennst – bevor es zu spät ist.\r\n\r\nWir freuen uns auf dich!\r\nDein Securaware-Team');
--
--
-- INSERT INTO reminder_email_template
-- VALUES (3, 'Ein Klick kann den Unterschied machen',
--         E'Hallo {{ .Firstname }}\r\n\r\nPhishing kann jeden treffen. Aber nicht jeden muss es erwischen.\r\n- Hol dir das nötige Wissen auf securaware.ch\r\n- Erkenne gefährliche E-Mails\r\n- Vermeide teure Fehler\r\n\r\nMach den Unterschied mit einem Klick.!\r\n\r\nSicher bleiben.\r\nDein Securaware-Team');
--
--
-- INSERT INTO reminder_email_template
-- VALUES (4, 'Phishing? Hoffen reicht nicht',
--         E'Hallo {{ .Firstname }}\r\n\r\nPhishing erkennen ist kein Zufall, sondern eine Frage des Wissens.\r\n\r\nKomm jetzt zurück auf securaware.ch und lerne, wie du dich und deine Daten zuverlässig schützt.\r\n\r\nWir zeigen dir, wie’s geht – einfach und kostenlos!\r\n\r\nDein Securaware-Team');
