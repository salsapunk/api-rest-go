INSERT INTO tasks(title, description, created_at, created_by) VALUES($1, $2, $3, $4) RETURN id;
