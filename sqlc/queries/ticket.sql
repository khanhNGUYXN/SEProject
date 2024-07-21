-- File: queries/ticket_details.sql

-- Get detailed information about tickets
SELECT
    t.id AS ticket_id,
    u.username,
    u.email,
    m.title AS movie_title,
    m.overview,
    c.name AS cinema_name,
    s.schedule_movie_date,
    s.schedule_movie_start,
    ss.status AS seat_status,
    ss.price
FROM ticket t
JOIN users u ON t.user_id = u.user_id
JOIN movie m ON t.movie_id = m.movie_id
JOIN cinema c ON t.cinema_id = c.cinema_id
JOIN Schedule s ON t.schedule_id = s.schedule_id
JOIN seatSchedule ss ON t.seat_id = ss.seat_id AND t.schedule_id = ss.schedule_id
WHERE t.id = $1;  -- $1 is the ticket ID parameter
