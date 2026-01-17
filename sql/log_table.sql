--  log watchlist for top film and trending film all time
CREATE TABLE log_watchlist (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  film_id UUID NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (film_id) REFERENCES films(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- trigger if watchlist added
CREATE OR REPLACE FUNCTION log_watchlist_trigger()
RETURNS trigger AS $$
BEGIN
  INSERT INTO log_watchlist (user_id, film_id)
  VALUES (NEW.user_id, NEW.film_id);
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE TRIGGER trg_log_watchlist
AFTER INSERT ON user_films
FOR EACH ROW
EXECUTE FUNCTION log_watchlist_trigger();


-- materialized view for top film
CREATE MATERIALIZED VIEW top_film_watchlist AS
SELECT
  film_id,
  COUNT(*) AS total_add
FROM log_watchlist
GROUP BY film_id
ORDER BY total_add DESC;

CREATE UNIQUE INDEX IF NOT EXISTS top_film
ON top_film_watchlist(film_id);

REFRESH MATERIALIZED VIEW CONCURRENTLY top_film_watchlist;

-- DROP INDEX top_film;
-- DROP MATERIALIZED VIEW top_film_watchlist;

-- how to get top film
SELECT * FROM top_film_watchlist ORDER BY total_add DESC LIMIT 10;


-- trending film
CREATE MATERIALIZED VIEW trending_film_weekly AS
SELECT
  film_id,
  COUNT(*) AS total_added,
  MAX(created_at) AS last_added_at
FROM log_watchlist
WHERE created_at >= date_trunc('week', NOW())
GROUP BY film_id
ORDER BY total_added DESC;

CREATE UNIQUE INDEX IF NOT EXISTS idx_trending_film_film_id
ON trending_film_weekly(film_id);

-- DROP INDEX idx_trending_film_film_id;

REFRESH MATERIALIZED VIEW CONCURRENTLY trending_film_weekly;

SELECT * FROM trending_film_weekly
 JOIN film ON film.id = trending_film_weekly.film_id
 ORDER BY total_added DESC LIMIT 10;