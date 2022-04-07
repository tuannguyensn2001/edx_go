-- +goose Up
-- +goose StatementBegin
CREATE TABLE `carts`
(
    `user_id`   int(11) NOT NULL,
    `course_id` int(11) NULL DEFAULT NULL
)


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `carts`;
-- +goose StatementEnd
