ALTER TABLE blueprints ADD incremental_id INT

ALTER TABLE blueprints DROP PRIMARY KEY;
ALTER TABLE blueprints ADD CONSTRAINT UNIQUE (id);

ALTER TABLE blueprints MODIFY COLUMN incremental_id  INT AUTO_INCREMENT;

ALTER TABLE blueprints ADD PRIMARY KEY (incremental_id);
