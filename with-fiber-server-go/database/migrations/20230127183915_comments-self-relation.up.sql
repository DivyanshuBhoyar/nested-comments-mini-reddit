BEGIN ;

-- // add a new column parent_id to the table with self referencing foreign key
ALTER TABLE comments ADD parent_id INT NULL DEFAULT NULL ;
ALTER TABLE comments ADD CONSTRAINT comments_parent_id_fk FOREIGN KEY (parent_id) REFERENCES comments(id) ON DELETE CASCADE ON UPDATE CASCADE;

COMMIT ;