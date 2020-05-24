//   Onix Config Manager - Dbman
//   Copyright (c) 2018-2020 by www.gatblau.org
//   Licensed under the Apache License, Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0
//   Contributors to this project, hereby assign copyright in this code to the project,
//   to be licensed under the same terms as the rest of the code.
package util

const VersionTable = `
DO
  $$
    BEGIN
      ---------------------------------------------------------------------------
      -- VERSION - version of releases (not only database)
      ---------------------------------------------------------------------------
      IF NOT EXISTS(SELECT relname FROM pg_class WHERE relname = 'version')
      THEN
        CREATE TABLE version
        (
          application_version CHARACTER VARYING(25) NOT NULL COLLATE pg_catalog."default",
          database_version    CHARACTER VARYING(25) NOT NULL COLLATE pg_catalog."default",
          description         TEXT COLLATE pg_catalog."default",
          time                timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP(6),
          scripts_source      character varying(250),
          CONSTRAINT version_app_version_db_release_pk PRIMARY KEY (application_version, database_version)
        )
          WITH (
            OIDS = FALSE
          )
          TABLESPACE pg_default;

        ALTER TABLE version
          OWNER to onix;
      END IF;
    END;
    $$
`