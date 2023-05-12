-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id       BIGSERIAL PRIMARY KEY NOT NULL unique,
    login    varchar(255)          NOT NULL DEFAULT '' unique,
    password varchar(255)          NOT NULL DEFAULT '',
    status   boolean               NOT NULL default true,
    role     varchar(50)           NOT NULL default 'worker'
);

CREATE TABLE IF NOT EXISTS session
(
    id         BIGSERIAL PRIMARY KEY                          NOT NULL unique,
    user_id    int references users (id) on delete cascade    not null,
    project_id int references projects (id) on delete cascade not null,
    keylog     text                                           not null default '',
    screens    bytea[]                                        not null,
    start      timestamp                                      not null default now(),
    finish     timestamp                                      not null default now()
);

CREATE TABLE IF NOT EXISTS tasks
(
    id              BIGSERIAL PRIMARY KEY NOT NULL unique,
    title           text                  NOT NULL DEFAULT '',
    description     text                  NOT NULL DEFAULT '',
    progress_status boolean               NOT NULL default true
);

CREATE TABLE IF NOT EXISTS projects
(
    id          BIGSERIAL PRIMARY KEY NOT NULL unique,
    title       text                  NOT NULL DEFAULT '',
    description text                  NOT NULL default '',
    status      boolean               NOT NULL default true
);

CREATE TABLE IF NOT EXISTS task_reports
(
    id          BIGSERIAL PRIMARY KEY NOT NULL unique,
    title       text                  NOT NULL DEFAULT '',
    description text                  NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS user_task_list
(
    id         BIGSERIAL PRIMARY KEY                          NOT NULL unique,
    user_id    int references users (id) on delete cascade    not null,
    task_id    int references tasks (id) on delete cascade    not null,
    project_id int references projects (id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS user_project_list
(
    id         BIGSERIAL PRIMARY KEY                          NOT NULL unique,
    user_id    int references users (id) on delete cascade    not null,
    project_id int references projects (id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS task_report_list
(
    id        BIGSERIAL PRIMARY KEY                              NOT NULL unique,
    task_id   int references tasks (id) on delete cascade        not null,
    report_id int references task_reports (id) on delete cascade not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_task_list, user_project_list, task_report_list, session, users, tasks, projects, task_reports;
-- +goose StatementEnd
