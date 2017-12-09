-- +goose Up
INSERT INTO config (ckey, cset) VALUES
  ('Developer.mr_robot:Database.processing', '{"host": "localhost", "port": "5432", "database": "devdb", "user": "mr_robot", "password": "secret", "schema": "public"}'),
  ('Developer.mr_robot:Deploy.dev', '{"hosts": ["hdth.dev.loc", "hdph.dev.loc", "zhru.dev.loc"], "repo": "http://github.com/zigi/st", "user": "deployer", "migrations": "http://st.zigi.st.migrations.dev"}'),
  ('Developer.mr_robot:Deploy.prod', '{"hosts": ["hdth.prod", "hdph.prod", "zhru.prod"], "repo": "http://github.com/zigi/st", "user": "deployer", "migrations": "http://st.zigi.st.migrations.prod"}'),
  ('Developer.mr_robot:Deploy.regress', '{"hosts": ["hdth.regress", "hdph.regress", "zhru.regress"], "repo": "http://github.com/zigi/st", "user": "deployer", "migrations": "http://zigi.st.migrations.regress"}'),
  ('StressTester.stress_engine', '{"script": "/etc/stress/run.sh", "user": "deployer", "result": {"host": "http://stress.result.gh", "api_key": "03630b562df15c64c77a8e12cb1c216a"}, "migrations": "http://zigi.st.migrations.regress"}');

-- +goose Down
TRUNCATE config;