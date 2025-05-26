ALTER TABLE users ADD COLUMN participates_in_phishing_simulation boolean NULL;

UPDATE users SET participates_in_phishing_simulation = false;

ALTER TABLE users ALTER COLUMN participates_in_phishing_simulation SET NOT NULL;