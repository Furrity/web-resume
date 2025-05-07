-- name: GetProfileInfo :one
SELECT 
    t_name.text as name,
    t_title.text as title,
    t_about.text as about,
    p.image_url as image_url
FROM profile p
JOIN languages l ON l.code = ?
join translations t_name ON t_name.key = p.name_key AND t_name.language_id = l.id
join translations t_title ON t_title.key = p.title_key and t_title.language_id = l.id
join translations t_about ON t_about.key = p.title_key and t_about.language_id = l.id
where id = 1;
