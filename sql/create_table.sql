-- create blog table
CREATE TABLE blog (
    `blog_id` VARCHAR(50),
    `title` VARCHAR(50) NOT NULL,
    `content` VARCHAR(20000) NOT NULL,
    `image_url` VARCHAR(1000) NOT NULL,
    `published` BOOLEAN NOT NULL,
    `type_id` VARCHAR(50),
    `create_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `update_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`blog_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

alter table blog add constraint `blog_type` foreign key (`type_id`) references `type`(`type_id`);

-- create user table
CREATE TABLE user (
    `user_id` VARCHAR(50),
    `username` VARCHAR(50) NOT NULL,
    `password` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `avatar_url` VARCHAR(1000) NOT NULL,
    `type` TINYINT NOT NULL,
    `create_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    `update_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`user_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- create type table
CREATE TABLE type (
    `type_id` VARCHAR(50),
    `name` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`type_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- create tag table
CREATE TABLE `tag` (
    `tag_id` VARCHAR(50),
    `name` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`tag_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- create relationship table between blog and tag
CREATE TABLE `blog_tag` (
    `id` VARCHAR(50),
    `blog_id` VARCHAR(50) NOT NULL,
    `tag_id` VARCHAR(50) NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`blog_id`) REFERENCES `blog`(`blog_id`),
    FOREIGN KEY (`tag_id`) REFERENCES `tag`(`tag_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- create comment table
CREATE TABLE `comment` (
    `comment_id` VARCHAR(50),
    `nickname` VARCHAR(50) NOT NULL,
    `email` VARCHAR(50) NOT NULL,
    `blog_id` VARCHAR(50) NOT NULL,
    `content` VARCHAR(2000) NOT NULL,
    `avatar_url` VARCHAR(1000),
    `create_at` TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (`comment_id`),
    FOREIGN KEY (`blog_id`) REFERENCES `blog`(`blog_id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;