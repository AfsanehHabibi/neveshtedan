package schema

const WritingEntryFieldTable = `
CREATE TABLE IF NOT EXISTS writing_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    value VARCHAR(50) NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;

