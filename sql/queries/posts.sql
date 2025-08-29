-- name: CreatePost :one
INSERT INTO posts (
    id,
    created_at,
    updated_at,
    title,
    url,
    description,
    published_at,
    feed_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
)
RETURNING *;

-- name: GetPostsForUser :many
SELECT posts.title, posts.url, posts.description, feeds.name FROM posts 
INNER join feed_follows
ON posts.feed_id = feed_follows.feed_id 
INNER join users
ON feed_follows.user_id = users.id 
INNER JOIN feeds
ON feeds.id = posts.feed_id
WHERE users.id = $1 
ORDER BY posts.published_at DESC
LIMIT $2;