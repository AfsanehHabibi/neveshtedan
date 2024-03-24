package schema

const WritingEntryFieldTable = `
CREATE TABLE IF NOT EXISTS writing_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    value VARCHAR(50) NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;

const WritingEntryTextFieldTable = `
CREATE TABLE IF NOT EXISTS writing_text_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    text VARCHAR(50) NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;

const WritingEntryNumberFieldTable = `
CREATE TABLE IF NOT EXISTS writing_number_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    value REAL NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;

const WritingEntryImageFieldTable = `
CREATE TABLE IF NOT EXISTS writing_image_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(50) NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;

const WritingEntryVideoFieldTable = `
CREATE TABLE IF NOT EXISTS writing_video_fields (
    entry_id INT,
    name VARCHAR(50) NOT NULL,
    url VARCHAR(50) NOT NULL,
    PRIMARY KEY (entry_id, name)
);
`;
