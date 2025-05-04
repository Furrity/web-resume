-- name: AddLanguage :exec
INSERT INTO languages (code, name)
VALUES (
    ?1,
    ?2
);
