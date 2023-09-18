CREATE TABLE IF NOT EXISTS `users` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `username` varchar(64) NOT NULL,
  `password` varchar(256),
  `email` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  CONSTRAINT `users_unique_with_username` UNIQUE (`username`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `user_3rdparty` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_id` bigint NOT NULL,
  `provider` text CHECK (`provider` IN ('github', 'gitlab', 'gitea')) NOT NULL DEFAULT 'github',
  `account_id` varchar(256),
  `token` varchar(256),
  `refresh_token` varchar(256),
  `cr_last_update_timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `cr_last_update_status` text CHECK (`cr_last_update_status` IN ('Success', 'Failed', 'Doing')) NOT NULL DEFAULT 'Doing',
  `cr_last_update_message` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `user_3rdparty_unique_with_account_id` UNIQUE (`provider`, `account_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `code_repository_clone_credentials` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_3rdparty_id` bigint NOT NULL,
  `type` text CHECK (`type` IN ('none', 'ssh', 'username', 'token')) NOT NULL DEFAULT 'none',
  `ssh_key` BLOB,
  `username` varchar(256),
  `password` varchar(256),
  `token` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_3rdparty_id`) REFERENCES `user_3rdparty` (`id`)
);

CREATE TABLE IF NOT EXISTS `code_repository_owners` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_3rdparty_id` bigint NOT NULL,
  `is_org` integer NOT NULL DEFAULT 0,
  `owner_id` varchar(256) NOT NULL,
  `owner` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_3rdparty_id`) REFERENCES `user_3rdparty` (`id`),
  CONSTRAINT `code_repository_owners_unique_with_name` UNIQUE (`user_3rdparty_id`, `owner_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `code_repositories` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_3rdparty_id` integer NOT NULL,
  `repository_id` varchar(256) NOT NULL,
  `is_org` integer NOT NULL DEFAULT 0,
  `owner_id` varchar(256) NOT NULL,
  `owner` varchar(256) NOT NULL,
  `name` varchar(256) NOT NULL,
  `ssh_url` varchar(256) NOT NULL,
  `clone_url` varchar(256) NOT NULL,
  `oci_repo_count` integer NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_3rdparty_id`) REFERENCES `user_3rdparty` (`id`),
  CONSTRAINT `code_repositories_unique_with_name` UNIQUE (`user_3rdparty_id`, `owner_id`, `repository_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `code_repository_branches` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `code_repository_id` integer NOT NULL,
  `name` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0,
  FOREIGN KEY (`code_repository_id`) REFERENCES `code_repositories` (`id`),
  CONSTRAINT `code_repository_branches_unique_with_name` UNIQUE (`code_repository_id`, `name`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `user_recover_codes` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_id` integer NOT NULL,
  `code` varchar(256) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `user_recover_codes_unique_with_use_id` UNIQUE (`user_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `namespaces` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `name` varchar(64) NOT NULL,
  `description` varchar(256),
  `visibility` text CHECK (`visibility` IN ('public', 'private')) NOT NULL DEFAULT 'private',
  `size_limit` integer NOT NULL DEFAULT 0,
  `size` integer NOT NULL DEFAULT 0,
  `repository_limit` integer NOT NULL DEFAULT 0,
  `repository_count` integer NOT NULL DEFAULT 0,
  `tag_limit` integer NOT NULL DEFAULT 0,
  `tag_count` integer NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  CONSTRAINT `namespaces_unique_with_name` UNIQUE (`name`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `audits` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `user_id` bigint NOT NULL,
  `namespace_id` bigint,
  `action` text CHECK (`action` IN ('create', 'update', 'delete', 'pull', 'push')) NOT NULL,
  `resource_type` text CHECK (`resource_type` IN ('namespace', 'repository', 'tag', 'builder')) NOT NULL,
  `resource` varchar(256) NOT NULL,
  `before_raw` BLOB,
  `req_raw` BLOB,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`namespace_id`) REFERENCES `namespaces` (`id`)
);

CREATE TABLE IF NOT EXISTS `repositories` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `name` varchar(64) NOT NULL,
  `description` varchar(255),
  `overview` BLOB,
  `visibility` text CHECK (`visibility` IN ('public', 'private')) NOT NULL DEFAULT 'private',
  `size_limit` integer NOT NULL DEFAULT 0,
  `size` integer NOT NULL DEFAULT 0,
  `tag_limit` integer NOT NULL DEFAULT 0,
  `tag_count` integer NOT NULL DEFAULT 0,
  `namespace_id` integer NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0,
  FOREIGN KEY (`namespace_id`) REFERENCES `namespaces` (`id`),
  CONSTRAINT `repositories_unique_with_namespace` UNIQUE (`namespace_id`, `name`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `artifacts` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `repository_id` integer NOT NULL,
  `digest` varchar(256) NOT NULL,
  `size` integer NOT NULL DEFAULT 0,
  `blobs_size` integer NOT NULL DEFAULT 0,
  `content_type` varchar(256) NOT NULL,
  `raw` BLOB NOT NULL,
  `config_raw` BLOB,
  `config_media_type` varchar(256),
  `type` text CHECK (`type` IN ('image', 'imageIndex', 'chart', 'cnab', 'wasm', 'provenance', 'unknown')) NOT NULL DEFAULT 'unknown',
  `pushed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_pull` timestamp,
  `pull_times` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`repository_id`) REFERENCES `repositories` (`id`),
  CONSTRAINT `artifacts_unique_with_repo` UNIQUE (`repository_id`, `digest`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `artifact_sboms` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `artifact_id` integer NOT NULL,
  `raw` BLOB,
  `result` BLOB,
  `status` varchar(64) NOT NULL,
  `stdout` BLOB,
  `stderr` BLOB,
  `message` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`artifact_id`) REFERENCES `artifacts` (`id`),
  CONSTRAINT `artifact_sbom_unique_with_artifact` UNIQUE (`artifact_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `artifact_vulnerabilities` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `artifact_id` integer NOT NULL,
  `metadata` BLOB,
  `raw` BLOB,
  `result` BLOB,
  `status` varchar(64) NOT NULL,
  `stdout` BLOB,
  `stderr` BLOB,
  `message` varchar(256),
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`artifact_id`) REFERENCES `artifacts` (`id`),
  CONSTRAINT `artifact_vulnerability_unique_with_artifact` UNIQUE (`artifact_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `tags` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `repository_id` integer NOT NULL,
  `artifact_id` integer NOT NULL,
  `name` varchar(64) NOT NULL,
  `pushed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_pull` timestamp,
  `pull_times` integer NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`repository_id`) REFERENCES `repositories` (`id`),
  FOREIGN KEY (`artifact_id`) REFERENCES `artifacts` (`id`),
  CONSTRAINT `tags_unique_with_repo` UNIQUE (`repository_id`, `name`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `blobs` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `digest` varchar(256) NOT NULL UNIQUE,
  `size` integer NOT NULL,
  `content_type` varchar(256) NOT NULL,
  `pushed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `last_pull` timestamp,
  `pull_times` integer NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  CONSTRAINT `blobs_unique_with_digest` UNIQUE (`digest`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `blob_uploads` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `part_number` int NOT NULL,
  `upload_id` varchar(256) NOT NULL,
  `etag` varchar(256) NOT NULL,
  `repository` varchar(256) NOT NULL,
  `file_id` varchar(256) NOT NULL,
  `size` integer NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  CONSTRAINT `blob_uploads_unique_with_upload_id_etag` UNIQUE (`upload_id`, `etag`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `artifact_artifacts` (
  `artifact_id` integer NOT NULL,
  `artifact_index_id` integer NOT NULL,
  PRIMARY KEY (`artifact_id`, `artifact_index_id`),
  CONSTRAINT `fk_artifact_artifacts_artifact` FOREIGN KEY (`artifact_id`) REFERENCES `artifacts` (`id`),
  CONSTRAINT `fk_artifact_artifacts_artifact_index` FOREIGN KEY (`artifact_index_id`) REFERENCES `artifacts` (`id`)
);

CREATE TABLE IF NOT EXISTS `artifact_blobs` (
  `artifact_id` integer NOT NULL,
  `blob_id` integer NOT NULL,
  PRIMARY KEY (`artifact_id`, `blob_id`),
  CONSTRAINT `fk_artifact_blobs_artifact` FOREIGN KEY (`artifact_id`) REFERENCES `artifacts` (`id`),
  CONSTRAINT `fk_artifact_blobs_blob` FOREIGN KEY (`blob_id`) REFERENCES blobs (`id`)
);

CREATE TABLE IF NOT EXISTS `daemon_logs` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `namespace_id` integer,
  `type` text CHECK (`type` IN ('Gc', 'Vulnerability', 'Sbom')) NOT NULL,
  `action` text CHECK (`action` IN ('create', 'update', 'delete', 'pull', 'push')) NOT NULL,
  `resource` varchar(256) NOT NULL,
  `status` text CHECK (`status` IN ('Success', 'Failed')) NOT NULL,
  `message` BLOB,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0
);

CREATE TABLE `casbin_rules` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `ptype` text,
  `v0` text,
  `v1` text,
  `v2` text,
  `v3` text,
  `v4` text,
  `v5` text,
  CONSTRAINT `idx_casbin_rules` UNIQUE (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
);

-- ptype type
-- v0 sub
-- v1 dom
-- v2 url
-- v3 attr
-- v4 method
-- v5 allow or deny
INSERT INTO `casbin_rules` (`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
  VALUES ('p', 'admin', '*', '*', '*', '*', 'allow'),
  ('p', 'anonymous', '/*', '/v2/', 'public|private', 'GET', 'allow'),
  ('p', 'anonymous', '/*', 'DS$*/**$blobs$*', 'public', 'GET|HEAD', 'allow'),
  ('p', 'anonymous', '/*', 'DS$*/**$manifests$*', 'public', 'GET|HEAD', 'allow'),
  ('p', 'namespace_reader', '/*', 'DS$*/**$blobs$*', 'public|private', 'GET|HEAD', 'allow'), -- get blob
  ('p', 'namespace_reader', '/*', 'DS$*/**$manifests$*', 'public|private', 'GET|HEAD', 'allow'), -- get manifest
  ('p', 'namespace_reader', '/*', 'API$*/**$namespaces/*', 'public|private', 'GET', 'allow'), -- get namespace
  ('p', 'namespace_reader', '/*', 'API$*/**$namespaces/*/artifacts/*', 'public|private', 'GET', 'allow'), -- get artifact
  ('p', 'namespace_reader', '/*', 'API$*/**$namespaces/*/artifacts/', 'public|private', 'GET', 'allow'), -- list artifacts
  ('p', 'namespace_reader', '/*', 'API$*/**$namespaces/*/repositories/', 'public|private', 'GET', 'allow'), -- list repositories
  ('p', 'namespace_reader', '/*', 'API$*/**$namespaces/*/repositories/*', 'public|private', 'GET', 'allow'), -- get repository
  ('p', 'namespace_admin', '/*', '*', 'public', 'GET|HEAD', 'allow'),
  ('p', 'namespace_owner', '/*', '*', 'public', 'GET|HEAD', 'allow');

INSERT INTO `namespaces` (`name`, `visibility`)
  VALUES ('library', 'public');

CREATE TABLE IF NOT EXISTS `webhooks` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `namespace_id` integer,
  `url` varchar(128) NOT NULL,
  `secret` varchar(63),
  `ssl_verify` integer NOT NULL DEFAULT 1,
  `retry_times` integer NOT NULL DEFAULT 3,
  `retry_duration` integer NOT NULL DEFAULT 5,
  `event_namespace` integer,
  `event_repository` integer NOT NULL DEFAULT 1,
  `event_tag` integer NOT NULL DEFAULT 1,
  `event_pull_push` integer NOT NULL DEFAULT 1,
  `event_member` integer NOT NULL DEFAULT 1,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS `webhook_logs` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `webhook_id` integer NOT NULL,
  `event` varchar(128) NOT NULL,
  `status_code` integer NOT NULL,
  `req_header` BLOB NOT NULL,
  `req_body` BLOB NOT NULL,
  `resp_header` BLOB NOT NULL,
  `resp_body` BLOB,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` integer NOT NULL DEFAULT 0,
  FOREIGN KEY (`webhook_id`) REFERENCES `webhooks` (`id`)
);

CREATE TABLE IF NOT EXISTS `builders` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `repository_id` integer NOT NULL,
  `source` text CHECK (`source` IN ('SelfCodeRepository', 'CodeRepository', 'Dockerfile')) NOT NULL,
  -- source SelfCodeRepository
  `scm_credential_type` varchar(16) NOT NULL,
  `scm_repository` varchar(256) NOT NULL,
  `scm_ssh_key` BLOB,
  `scm_token` varchar(256),
  `scm_username` varchar(30),
  `scm_password` varchar(30),
  -- source CodeRepository
  `code_repository_id` integer,
  -- source Dockerfile
  `dockerfile` BLOB,
  -- common settings
  `scm_branch` varchar(256),
  `scm_depth` MEDIUMINT,
  `scm_submodule` integer,
  -- cron settings
  `cron_rule` varchar(30),
  `cron_branch` varchar(256),
  `cron_tag_template` varchar(256),
  `cron_next_trigger` timestamp,
  -- webhook settings
  `webhook_branch_name` varchar(256),
  `webhook_branch_tag_template` varchar(256),
  `webhook_tag_tag_template` varchar(256),
  -- buildkit settings
  `buildkit_insecure_registries` varchar(256),
  `buildkit_context` varchar(30),
  `buildkit_dockerfile` varchar(256),
  `buildkit_platforms` varchar(256) NOT NULL DEFAULT 'linux/amd64',
  `buildkit_build_args` varchar(256),
  -- other fields
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`repository_id`) REFERENCES `repositories` (`id`),
  CONSTRAINT `builders_unique_with_repository` UNIQUE (`repository_id`, `deleted_at`)
);

CREATE TABLE IF NOT EXISTS `builder_runners` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `builder_id` bigint NOT NULL,
  `log` BLOB,
  `status` text CHECK (`status` IN ('Success', 'Failed', 'Pending', 'Scheduling', 'Building')) NOT NULL,
  -- common settings
  `tag` varchar(30) NOT NULL, -- image tag
  `scm_branch` varchar(30) NOT NULL DEFAULT 'main',
  `buildkit_platforms` varchar(256) NOT NULL DEFAULT 'linux/amd64',
  -- other fields
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0,
  FOREIGN KEY (`builder_id`) REFERENCES `builders` (`id`)
);

CREATE TABLE IF NOT EXISTS `work_queue` (
  `id` integer PRIMARY KEY AUTOINCREMENT,
  `topic` varchar(30) NOT NULL,
  `payload` BLOB NOT NULL,
  `times` integer NOT NULL DEFAULT 0,
  `version` varchar(30) NOT NULL,
  `status` text CHECK (`status` IN ('Success', 'Failed', 'Pending', 'Scheduling', 'Building')) NOT NULL DEFAULT 'Pending',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint NOT NULL DEFAULT 0
);

