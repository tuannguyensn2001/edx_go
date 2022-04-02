-- +goose Up
-- +goose StatementBegin
CREATE TABLE `lessons`(
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NULL DEFAULT NULL,
    `video_url` varchar(255)  NULL DEFAULT NULL,
    `video_type` varchar(255)  NULL DEFAULT NULL,
    `order` INT(11)  NULL DEFAULT NULL,
    `chapter_id` INT(11)  NULL DEFAULT NULL,
    `status` varchar(255)  NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE  IF EXISTS `lessons`
-- +goose StatementEnd
