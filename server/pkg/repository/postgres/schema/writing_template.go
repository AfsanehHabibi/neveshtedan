package schema

const WritingTemplateTable = `
CREATE TABLE IF NOT EXISTS templates (
	id SERIAL PRIMARY KEY,
	title VARCHAR(100) NOT NULL,
    description VARCHAR(100) NOT NULL,
	user_id INT
  );
`
