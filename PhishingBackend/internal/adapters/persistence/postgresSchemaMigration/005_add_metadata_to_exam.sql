ALTER TABLE exam
  ADD title text,
  ADD description text;

UPDATE exam SET title = '', description = '';

ALTER TABLE exam
  ALTER COLUMN title SET NOT NULL,
  ALTER COLUMN description SET NOT NULL;