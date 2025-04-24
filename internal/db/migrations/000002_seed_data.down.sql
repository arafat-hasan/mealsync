-- Delete test menu item comments
DELETE FROM menu_item_comments WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND meal_event_id IN (
    SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date::date = CURRENT_DATE
) AND menu_item_id IN (
    SELECT id FROM menu_items WHERE name = 'Classic Burger'
);

-- Delete test user requested items
DELETE FROM user_requested_items WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND meal_event_id IN (
    SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date::date = CURRENT_DATE
);

-- Delete test meal requests
DELETE FROM meal_requests WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND meal_event_id IN (
    SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date::date = CURRENT_DATE
);

-- Delete test meal event addresses
DELETE FROM meal_event_addresses WHERE meal_event_id IN (
    SELECT id FROM meal_events WHERE name IN ('Breakfast Event', 'Lunch Event') AND event_date::date = CURRENT_DATE
);

-- Delete test meal event sets
DELETE FROM meal_event_sets WHERE meal_event_id IN (
    SELECT id FROM meal_events WHERE name IN ('Breakfast Event', 'Lunch Event') AND event_date::date = CURRENT_DATE
);

-- Delete test meal events
DELETE FROM meal_events WHERE name IN ('Breakfast Event', 'Lunch Event') AND event_date::date = CURRENT_DATE;

-- Delete test menu set items
DELETE FROM menu_set_items WHERE menu_set_id IN (
    SELECT id FROM menu_sets WHERE menu_set_name IN ('Breakfast Set A', 'Lunch Set A', 'Lunch Set B')
);

-- Delete test menu sets
DELETE FROM menu_sets WHERE menu_set_name IN ('Breakfast Set A', 'Lunch Set A', 'Lunch Set B');

-- Delete test menu items
DELETE FROM menu_items WHERE name IN (
    'Classic Burger',
    'Caesar Salad',
    'Margherita Pizza',
    'Chicken Wrap',
    'Chocolate Brownie',
    'Scrambled Eggs'
);

-- Delete test event addresses
DELETE FROM event_addresses WHERE address IN (
    'Main Campus Cafeteria',
    'Downtown Office Cafeteria'
);


-- Delete test notifications for different types and users
DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND type = 'confirmation' AND payload::text LIKE '%Lunch Event%';

DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND type = 'reminder' AND payload::text LIKE '%Breakfast Event%';

DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee1@mealsync.com'
) AND type = 'admin-message' AND payload::text LIKE '%renovations%';

DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee2@mealsync.com'
) AND type = 'confirmation' AND read = true;

DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'admin@mealsync.com'
) AND type = 'admin-message' AND payload::text LIKE '%maintenance%';

DELETE FROM notifications WHERE user_id IN (
    SELECT id FROM users WHERE email = 'employee2@mealsync.com'
) AND type = 'reminder' AND payload::text LIKE '%Deadline approaching%';

-- Delete test users
DELETE FROM users WHERE email IN (
    'admin@mealsync.com',
    'employee1@mealsync.com',
    'employee2@mealsync.com'
);
