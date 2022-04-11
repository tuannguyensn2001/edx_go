-- +goose Up
-- +goose StatementBegin
CREATE TABLE `comments`  (
                             `id` int(11) NOT NULL AUTO_INCREMENT,
                             `content` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                             `user_id` int(11) NULL DEFAULT NULL,
                             `parent_id` int(11) NULL DEFAULT NULL,
                             `commentable_model` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
                             `commentable_id` int(11) NULL DEFAULT NULL,
                             `created_at` timestamp NULL DEFAULT NULL,
                             `updated_at` timestamp NULL DEFAULT NULL,
                             `deleted_at` timestamp NULL DEFAULT NULL,
                             PRIMARY KEY (`id`) USING BTREE
)



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `comments`;
-- +goose StatementEnd
