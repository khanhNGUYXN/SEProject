-- Add a new coupon
INSERT INTO coupons (code, description, discount_percentage, valid_from, expires_at)
VALUES ($1, $2, $3, $4, $5);

-- Retrieve all coupons
SELECT id, code, description, discount_percentage, valid_from, expires_at FROM coupons;

-- Update a coupon
UPDATE coupons SET description = $2, discount_percentage = $3, valid_from = $4, expires_at = $5
WHERE code = $1;

-- Delete a coupon
DELETE FROM coupons WHERE code = $1;
