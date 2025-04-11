-- Delete test notifications
DELETE FROM notifications WHERE title = 'Meal Request Approved';

-- Delete test meal comments
DELETE FROM meal_comments WHERE date = CURRENT_DATE;

-- Delete test meal requests
DELETE FROM meal_requests WHERE requested_for = CURRENT_DATE;

-- Delete test meal menu items
DELETE FROM meal_menu_items WHERE meal_menu_id IN (
    SELECT id FROM meal_menus WHERE date = CURRENT_DATE
);

-- Delete test meal menus
DELETE FROM meal_menus WHERE date = CURRENT_DATE;

-- Delete test menu items
DELETE FROM menu_items WHERE name IN (
    'Classic Burger',
    'Caesar Salad',
    'Margherita Pizza',
    'Chicken Wrap',
    'Chocolate Brownie',
    'Scrambled Eggs'
);

-- Delete test restaurants
DELETE FROM restaurants WHERE name IN (
    'Campus Cafeteria',
    'Quick Bites'
);

-- Delete test users
DELETE FROM users WHERE email IN (
    'admin@mealsync.com',
    'employee1@mealsync.com',
    'employee2@mealsync.com'
); 