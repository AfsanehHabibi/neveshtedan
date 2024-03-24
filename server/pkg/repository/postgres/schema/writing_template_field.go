package schema

const WritingTemplateFieldTable = `
CREATE TABLE IF NOT EXISTS template_fields (
    template_id INT,
    name VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    description VARCHAR(100) NOT NULL,
    PRIMARY KEY (template_id, name)
);
`
