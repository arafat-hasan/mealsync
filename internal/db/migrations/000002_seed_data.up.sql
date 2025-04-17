-- Insert test users
-- Password for all users is 'password123' (hashed)
INSERT INTO users (employee_id, username, password_hash, name, email, department, role, is_active)
SELECT 1001, 'admin', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'Admin User', 'admin@mealsync.com', 'Administration', 'admin', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@mealsync.com');

INSERT INTO users (employee_id, username, password_hash, name, email, department, role, is_active)
SELECT 1002, 'john.doe', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'John Doe', 'employee1@mealsync.com', 'Engineering', 'user', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee1@mealsync.com');

INSERT INTO users (employee_id, username, password_hash, name, email, department, role, is_active)
SELECT 1003, 'jane.smith', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'Jane Smith', 'employee2@mealsync.com', 'Marketing', 'user', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee2@mealsync.com');

-- Insert test event addresses
INSERT INTO event_addresses (address, is_active, created_by, updated_by)
SELECT 'Main Campus Cafeteria', true, (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM event_addresses WHERE address = 'Main Campus Cafeteria');

INSERT INTO event_addresses (address, is_active, created_by, updated_by)
SELECT 'Downtown Office Cafeteria', true, (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM event_addresses WHERE address = 'Downtown Office Cafeteria');

-- Insert test menu items
INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Classic Burger', 'Juicy beef patty with lettuce, tomato, and cheese', 'https://example.com/burger.jpg', true, 
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Classic Burger');

INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Caesar Salad', 'Fresh romaine lettuce with caesar dressing and croutons', 'https://example.com/salad.jpg', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Caesar Salad');

INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Margherita Pizza', 'Traditional pizza with tomato sauce, mozzarella, and basil', 'https://example.com/pizza.jpg', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Margherita Pizza');

INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Chicken Wrap', 'Grilled chicken with vegetables in a tortilla wrap', 'https://example.com/wrap.jpg', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chicken Wrap');

INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Chocolate Brownie', 'Rich chocolate brownie with vanilla ice cream', 'https://example.com/brownie.jpg', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chocolate Brownie');

INSERT INTO menu_items (name, description, image_url, is_active, created_by, updated_by)
SELECT 'Scrambled Eggs', 'Fluffy scrambled eggs with toast', 'https://example.com/eggs.jpg', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Scrambled Eggs');

-- Insert test menu sets
INSERT INTO menu_sets (menu_set_name, menu_set_description, is_active, created_by, updated_by)
SELECT 'Breakfast Set A', 'Standard breakfast options', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_sets WHERE menu_set_name = 'Breakfast Set A');

INSERT INTO menu_sets (menu_set_name, menu_set_description, is_active, created_by, updated_by)
SELECT 'Lunch Set A', 'Main course options', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_sets WHERE menu_set_name = 'Lunch Set A');

INSERT INTO menu_sets (menu_set_name, menu_set_description, is_active, created_by, updated_by)
SELECT 'Lunch Set B', 'Light meal options', true,
       (SELECT id FROM users WHERE email = 'admin@mealsync.com'), (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (SELECT 1 FROM menu_sets WHERE menu_set_name = 'Lunch Set B');

-- Insert test menu set items
INSERT INTO menu_set_items (menu_set_id, menu_item_id, created_by, updated_by)
SELECT 
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Breakfast Set A'),
    (SELECT id FROM menu_items WHERE name = 'Scrambled Eggs'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM menu_set_items 
    WHERE menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Breakfast Set A')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Scrambled Eggs')
);

INSERT INTO menu_set_items (menu_set_id, menu_item_id, created_by, updated_by)
SELECT 
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A'),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM menu_set_items 
    WHERE menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
);

INSERT INTO menu_set_items (menu_set_id, menu_item_id, created_by, updated_by)
SELECT 
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A'),
    (SELECT id FROM menu_items WHERE name = 'Margherita Pizza'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM menu_set_items 
    WHERE menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Margherita Pizza')
);

INSERT INTO menu_set_items (menu_set_id, menu_item_id, created_by, updated_by)
SELECT 
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B'),
    (SELECT id FROM menu_items WHERE name = 'Caesar Salad'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM menu_set_items 
    WHERE menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Caesar Salad')
);

INSERT INTO menu_set_items (menu_set_id, menu_item_id, created_by, updated_by)
SELECT 
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B'),
    (SELECT id FROM menu_items WHERE name = 'Chicken Wrap'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM menu_set_items 
    WHERE menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Chicken Wrap')
);

-- Insert test meal events
INSERT INTO meal_events (name, description, event_date, event_duration, cutoff_time, is_active, created_by, updated_by)
SELECT 
    'Breakfast Event', 
    'Daily breakfast service', 
    CURRENT_DATE, 
    60, 
    CURRENT_DATE + INTERVAL '1 day' - INTERVAL '2 hours', 
    true, 
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_events 
    WHERE name = 'Breakfast Event' AND event_date = CURRENT_DATE
);

INSERT INTO meal_events (name, description, event_date, event_duration, cutoff_time, is_active, created_by, updated_by)
SELECT 
    'Lunch Event', 
    'Daily lunch service', 
    CURRENT_DATE, 
    90, 
    CURRENT_DATE + INTERVAL '1 day' - INTERVAL '2 hours', 
    true, 
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_events 
    WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE
);

-- Insert test meal event menu sets
INSERT INTO meal_event_menu_sets (meal_event_id, menu_set_id, label, note, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Breakfast Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Breakfast Set A'),
    'Standard Breakfast',
    'Includes eggs and toast',
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_menu_sets 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Breakfast Event' AND event_date = CURRENT_DATE)
    AND menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Breakfast Set A')
);

INSERT INTO meal_event_menu_sets (meal_event_id, menu_set_id, label, note, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A'),
    'Main Course Options',
    'Includes burger or pizza',
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_menu_sets 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
    AND menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A')
);

INSERT INTO meal_event_menu_sets (meal_event_id, menu_set_id, label, note, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B'),
    'Light Meal Options',
    'Includes salad or wrap',
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_menu_sets 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
    AND menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set B')
);

-- Insert test meal event addresses
INSERT INTO meal_event_addresses (meal_event_id, address_id, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Breakfast Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM event_addresses WHERE address = 'Main Campus Cafeteria'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_addresses 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Breakfast Event' AND event_date = CURRENT_DATE)
    AND address_id = (SELECT id FROM event_addresses WHERE address = 'Main Campus Cafeteria')
);

INSERT INTO meal_event_addresses (meal_event_id, address_id, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM event_addresses WHERE address = 'Main Campus Cafeteria'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_addresses 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
    AND address_id = (SELECT id FROM event_addresses WHERE address = 'Main Campus Cafeteria')
);

INSERT INTO meal_event_addresses (meal_event_id, address_id, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM event_addresses WHERE address = 'Downtown Office Cafeteria'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_event_addresses 
    WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
    AND address_id = (SELECT id FROM event_addresses WHERE address = 'Downtown Office Cafeteria')
);

-- Insert test meal requests
INSERT INTO meal_requests (employee_id, meal_event_id, event_menu_set_id, meal_event_address_id, created_by, updated_by)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM meal_event_menu_sets WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE) 
     AND menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A')),
    (SELECT id FROM meal_event_addresses WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE) 
     AND address_id = (SELECT id FROM event_addresses WHERE address = 'Main Campus Cafeteria')),
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE employee_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
);

-- Insert test meal request items
INSERT INTO meal_request_items (meal_request_id, menu_item_id, menu_set_id, is_selected, quantity, notes, created_by, updated_by)
SELECT 
    (SELECT id FROM meal_requests WHERE employee_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
     AND meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A'),
    true,
    1,
    'No onions please',
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_request_items 
    WHERE meal_request_id = (SELECT id FROM meal_requests WHERE employee_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
                            AND meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE))
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
);

-- Insert test meal comments
INSERT INTO meal_comments (employee_id, meal_event_id, event_menu_set_id, menu_item_id, comment, rating, created_by, updated_by)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE),
    (SELECT id FROM meal_event_menu_sets WHERE meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE) 
     AND menu_set_id = (SELECT id FROM menu_sets WHERE menu_set_name = 'Lunch Set A')),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    'The burger was delicious today!',
    5,
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_comments 
    WHERE employee_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND meal_event_id = (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE)
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
);

-- Insert test notifications
INSERT INTO notifications (employee_id, type, payload, read, created_by, updated_by)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    'confirmation',
    '{"meal_event_id": ' || (SELECT id FROM meal_events WHERE name = 'Lunch Event' AND event_date = CURRENT_DATE) || 
    ', "message": "Your meal request for Classic Burger has been confirmed."}',
    false,
    (SELECT id FROM users WHERE email = 'admin@mealsync.com'),
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM notifications 
    WHERE employee_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND type = 'confirmation'
); 