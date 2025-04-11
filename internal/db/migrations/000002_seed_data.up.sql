-- Insert test users
-- Password for all users is 'password123' (hashed)
INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active, notification_enabled)
SELECT 'admin@mealsync.com', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'Admin', 'User', 'admin', 'Administration', 'ADMIN001', true, true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@mealsync.com');

INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active, notification_enabled)
SELECT 'employee1@mealsync.com', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'John', 'Doe', 'employee', 'Engineering', 'EMP001', true, true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee1@mealsync.com');

INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active, notification_enabled)
SELECT 'employee2@mealsync.com', '$2a$10$TdFO5yl0tUgVA7UuIlKNmubSnXPey22Qw/BXOgFJWXCcJetrdul3W', 'Jane', 'Smith', 'employee', 'Marketing', 'EMP002', true, true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee2@mealsync.com');

-- Insert test restaurants
INSERT INTO restaurants (name, description, address, phone, email, image)
SELECT 'Campus Cafeteria', 'Main cafeteria on campus', '123 Campus Drive', '555-123-4567', 'cafeteria@mealsync.com', 'https://example.com/cafeteria.jpg'
WHERE NOT EXISTS (SELECT 1 FROM restaurants WHERE name = 'Campus Cafeteria');

INSERT INTO restaurants (name, description, address, phone, email, image)
SELECT 'Quick Bites', 'Fast food restaurant', '456 University Ave', '555-987-6543', 'quickbites@mealsync.com', 'https://example.com/quickbites.jpg'
WHERE NOT EXISTS (SELECT 1 FROM restaurants WHERE name = 'Quick Bites');

-- Insert test menu items
INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'),
    'Classic Burger', 
    'Juicy beef patty with lettuce, tomato, and cheese', 
    12.99, 
    'Main Course', 
    'lunch', 
    'Set A', 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Classic Burger' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'));

INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'),
    'Caesar Salad', 
    'Fresh romaine lettuce with caesar dressing and croutons', 
    8.99, 
    'Salads', 
    'lunch', 
    'Set B', 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Caesar Salad' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'));

INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Quick Bites'),
    'Margherita Pizza', 
    'Traditional pizza with tomato sauce, mozzarella, and basil', 
    14.99, 
    'Main Course', 
    'lunch', 
    'Set A', 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Margherita Pizza' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Quick Bites'));

INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Quick Bites'),
    'Chicken Wrap', 
    'Grilled chicken with vegetables in a tortilla wrap', 
    10.99, 
    'Sandwiches', 
    'lunch', 
    'Set B', 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chicken Wrap' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Quick Bites'));

INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'),
    'Chocolate Brownie', 
    'Rich chocolate brownie with vanilla ice cream', 
    6.99, 
    'Desserts', 
    'snacks', 
    NULL, 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chocolate Brownie' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'));

INSERT INTO menu_items (restaurant_id, name, description, price, category, meal_type, set_name, is_available)
SELECT 
    (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'),
    'Scrambled Eggs', 
    'Fluffy scrambled eggs with toast', 
    7.99, 
    'Breakfast', 
    'breakfast', 
    NULL, 
    true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Scrambled Eggs' AND restaurant_id = (SELECT id FROM restaurants WHERE name = 'Campus Cafeteria'));

-- Insert test meal menus
INSERT INTO meal_menus (date, meal_type, cutoff_time, is_active, created_by)
SELECT 
    CURRENT_DATE, 
    'breakfast', 
    CURRENT_DATE + INTERVAL '1 day' - INTERVAL '2 hours', 
    true, 
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menus 
    WHERE date = CURRENT_DATE AND meal_type = 'breakfast'
);

INSERT INTO meal_menus (date, meal_type, cutoff_time, is_active, created_by)
SELECT 
    CURRENT_DATE, 
    'lunch', 
    CURRENT_DATE + INTERVAL '1 day' - INTERVAL '2 hours', 
    true, 
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menus 
    WHERE date = CURRENT_DATE AND meal_type = 'lunch'
);

INSERT INTO meal_menus (date, meal_type, cutoff_time, is_active, created_by)
SELECT 
    CURRENT_DATE, 
    'snacks', 
    CURRENT_DATE + INTERVAL '1 day' - INTERVAL '2 hours', 
    true, 
    (SELECT id FROM users WHERE email = 'admin@mealsync.com')
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menus 
    WHERE date = CURRENT_DATE AND meal_type = 'snacks'
);

-- Insert test meal menu items
INSERT INTO meal_menu_items (meal_menu_id, menu_item_id, set_name)
SELECT 
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'breakfast'),
    (SELECT id FROM menu_items WHERE name = 'Scrambled Eggs'),
    NULL
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menu_items 
    WHERE meal_menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'breakfast')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Scrambled Eggs')
);

INSERT INTO meal_menu_items (meal_menu_id, menu_item_id, set_name)
SELECT 
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch'),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    'Set A'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menu_items 
    WHERE meal_menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
);

INSERT INTO meal_menu_items (meal_menu_id, menu_item_id, set_name)
SELECT 
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch'),
    (SELECT id FROM menu_items WHERE name = 'Caesar Salad'),
    'Set B'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menu_items 
    WHERE meal_menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Caesar Salad')
);

INSERT INTO meal_menu_items (meal_menu_id, menu_item_id, set_name)
SELECT 
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'snacks'),
    (SELECT id FROM menu_items WHERE name = 'Chocolate Brownie'),
    NULL
WHERE NOT EXISTS (
    SELECT 1 FROM meal_menu_items 
    WHERE meal_menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'snacks')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Chocolate Brownie')
);

-- Insert test meal requests
INSERT INTO meal_requests (user_id, menu_item_id, menu_id, quantity, status, requested_for, notes)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch'),
    1, 
    'pending', 
    CURRENT_DATE, 
    'No onions please'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
    AND menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch')
);

INSERT INTO meal_requests (user_id, menu_item_id, menu_id, quantity, status, requested_for, notes)
SELECT 
    (SELECT id FROM users WHERE email = 'employee2@mealsync.com'),
    (SELECT id FROM menu_items WHERE name = 'Caesar Salad'),
    (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch'),
    2, 
    'approved', 
    CURRENT_DATE, 
    NULL
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee2@mealsync.com')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Caesar Salad')
    AND menu_id = (SELECT id FROM meal_menus WHERE date = CURRENT_DATE AND meal_type = 'lunch')
);

-- Insert test meal comments
INSERT INTO meal_comments (user_id, date, comment)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    CURRENT_DATE, 
    'The food was delicious today!'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_comments 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND date = CURRENT_DATE
);

-- Insert test notifications
INSERT INTO notifications (user_id, title, message, type, read)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    'Meal Request Approved', 
    'Your meal request for Classic Burger has been approved.', 
    'meal_confirmation', 
    false
WHERE NOT EXISTS (
    SELECT 1 FROM notifications 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND title = 'Meal Request Approved'
); 