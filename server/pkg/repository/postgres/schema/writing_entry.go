package schema

const WritingEntryTable = `
CREATE TABLE IF NOT EXISTS writings (
	id SERIAL PRIMARY KEY,
	template_id INT,
	user_id INT
  );  
`
