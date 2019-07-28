CREATE TABLE tasks
(
    `id`         bigint(20)                       NOT NULL AUTO_INCREMENT,
    `word`       varchar(255)                     NOT NULL,
    `difficulty` tinyint unsigned                 NOT NULL,
    `type`       enum ('DRAW', 'SHOW', 'EXPLAIN') NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COLLATE = utf8_unicode_ci;


