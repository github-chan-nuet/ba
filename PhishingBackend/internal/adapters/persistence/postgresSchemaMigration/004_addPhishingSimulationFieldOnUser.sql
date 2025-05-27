ALTER TABLE "user" ADD COLUMN participates_in_phishing_simulation boolean NULL;

UPDATE "user" SET participates_in_phishing_simulation = false;

ALTER TABLE "user" ALTER COLUMN participates_in_phishing_simulation SET NOT NULL;