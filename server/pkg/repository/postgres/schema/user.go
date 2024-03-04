package schema

const UserTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	password VARCHAR(100) NOT NULL,
    username VARCHAR(100) UNIQUE NOT NULL
  );  
`
