CREATE TABLE student (
    id          VARCHAR(126) PRIMARY KEY NOT NULL,
    name        VARCHAR(126) NOT NULL,
    bio         VARCHAR(510),
    since       TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    email       VARCHAR(126) NOT NULL UNIQUE,
    password    CHAR(60) NOT NULL,
    img_path    UUID
);

CREATE TABLE img (
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL UNIQUE,
    name        VARCHAR(126) NOT NULL UNIQUE
);

CREATE TABLE markdown (
    id              UUID DEFAULT gen_random_uuid() PRIMARY KEY NOT NULL UNIQUE,
    student_id      VARCHAR(126) NOT NULL,
    title           VARCHAR(255) NOT NULL,
    content_path    VARCHAR(255) NOT NULL UNIQUE,
    since           TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated         TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    FOREIGN KEY     (student_id) REFERENCES student(id)
);

CREATE TABLE markdown_img_rel (
    markdown_id UUID NOT NULL,
    img_id      UUID NOT NULL,
    FOREIGN KEY (markdown_id) REFERENCES markdown(id),
    FOREIGN KEY (img_id) REFERENCES img(id),
    PRIMARY KEY (markdown_id, img_id)
);