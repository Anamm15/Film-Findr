-- Materialzed view for reviews
CREATE MATERIALIZED VIEW weekly_review AS
SELECT
  date_trunc('week', created_at) AS weekly,
  COUNT(*) AS total_review
FROM reviews
GROUP BY 1
ORDER BY 1;

-- indexing for better performance
CREATE INDEX idx_review_created_at ON reviews(created_at);
CREATE UNIQUE INDEX idx_weekly_review ON weekly_review(weekly);

-- refreshing manually
REFRESH MATERIALIZED VIEW CONCURRENTLY weekly_review;
-- or using pg_cron
--- 0 0 * * 1 psql -U postgres -d database_name -c "REFRESH MATERIALIZED VIEW weekly_review"


-- materialized view for rating film
CREATE MATERIALIZED VIEW rating_film AS
SELECT
  film_id,
  AVG(rating) AS rating
FROM reviews
GROUP BY 1;

-- indexing for better performance
CREATE INDEX idx_review_film_id ON reviews(film_id);
CREATE UNIQUE INDEX idx_rating_film ON rating_film(film_id);

-- refreshing
REFRESH MATERIALIZED VIEW CONCURRENTLY rating_film;
-- or 
--- 0 0 * * 1 psql -U postgres -d database_name -c "REFRESH MATERIALIZED VIEW rating_film"


-- materialized view for user
CREATE MATERIALIZED VIEW weekly_user AS
SELECT
  date_trunc('week', created_at) AS weekly,
  COUNT(*) AS total_user
FROM users
GROUP BY 1
ORDER BY 1;

-- indexing for better performance
CREATE INDEX idx_users_created_at ON users(created_at);
CREATE UNIQUE INDEX idx_weekly_user ON weekly_user(weekly);

-- refreshing
REFRESH MATERIALIZED VIEW CONCURRENTLY weekly_user;
-- or 
--- 0 0 * * 1 psql -U postgres -d database_name -c "REFRESH MATERIALIZED VIEW weekly_user"
