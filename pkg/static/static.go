package static

var (
	InsertTweetQuery = `
		INSERT INTO 
			tweets (
				id,
				content,
				user_id
			) VALUES ($1, $2, $3)
	`

	GetTweetByIDQuery = `
		SELECT 
			id,
			content,
			user_id,
			created_at, 
			updated_at
		FROM 
			tweets
		WHERE 
			id = $1
		`

	GetListTweetBaseQuery = `
		SELECT 
			t.id,
			t.content,
			t.user_id,
			t.created_at,
			t.updated_at,
			JSON_AGG(
				JSON_BUILD_OBJECT(
					'id', tm.id,
					'tweet_id', tm.tweet_id,
					'url', tm.url,
					'media_type', tm.media_type
				)
			) AS medias
		FROM 
			tweets AS t
		LEFT JOIN 
			tweet_media AS tm ON t.id = tm.tweet_id
		`

	GetListTweetGroupQuery = `
		GROUP BY
			t.id, t.content, t.user_id, t.created_at, t.updated_at
		`

	GetListTweetQueryCount = `
		SELECT 	
			COUNT(*)
		FROM 
			tweets
		`

	UpdateTweetQuery = `
		UPDATE 
			tweets
		SET 
			content = $1, 
			user_id = $2,
			updated_at = NOW()
		WHERE
			id = $3
		`

	DeleteTweetQuery = `
		DELETE 
			FROM 
				tweets
		WHERE 
			id = $1
		`
)
