package schema

const WritingEntryTable = `
CREATE TABLE IF NOT EXISTS writings (
	id INT PRIMARY KEY AUTO_INCREMENT,
	template_id INT,
	user_id INT
  );  
`
