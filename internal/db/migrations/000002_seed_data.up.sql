-- Insert test users
-- Password for all users is 'password123' (hashed)
INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active)
SELECT 'admin@mealsync.com', '$2a$10$ZjXzGXkNTL5u.rcKGYC9N.cQoVYfV.kX4.nFEj0EONvK3O9F5YgYi', 'Admin', 'User', 'admin', 'Administration', 'ADMIN001', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'admin@mealsync.com');

INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active)
SELECT 'employee1@mealsync.com', '$2a$10$ZjXzGXkNTL5u.rcKGYC9N.cQoVYfV.kX4.nFEj0EONvK3O9F5YgYi', 'John', 'Doe', 'employee', 'Engineering', 'EMP001', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee1@mealsync.com');

INSERT INTO users (email, password, first_name, last_name, role, department, employee_id, is_active)
SELECT 'employee2@mealsync.com', '$2a$10$ZjXzGXkNTL5u.rcKGYC9N.cQoVYfV.kX4.nFEj0EONvK3O9F5YgYi', 'Jane', 'Smith', 'employee', 'Marketing', 'EMP002', true
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'employee2@mealsync.com');

-- Insert test menu items
INSERT INTO menu_items (name, description, price, category, is_available)
SELECT 'Classic Burger', 'Juicy beef patty with lettuce, tomato, and cheese', 12.99, 'Main Course', true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Classic Burger');

INSERT INTO menu_items (name, description, price, category, is_available)
SELECT 'Caesar Salad', 'Fresh romaine lettuce with caesar dressing and croutons', 8.99, 'Salads', true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Caesar Salad');

INSERT INTO menu_items (name, description, price, category, is_available)
SELECT 'Margherita Pizza', 'Traditional pizza with tomato sauce, mozzarella, and basil', 14.99, 'Main Course', true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Margherita Pizza');

INSERT INTO menu_items (name, description, price, category, is_available)
SELECT 'Chicken Wrap', 'Grilled chicken with vegetables in a tortilla wrap', 10.99, 'Sandwiches', true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chicken Wrap');

INSERT INTO menu_items (name, description, price, category, is_available)
SELECT 'Chocolate Brownie', 'Rich chocolate brownie with vanilla ice cream', 6.99, 'Desserts', true
WHERE NOT EXISTS (SELECT 1 FROM menu_items WHERE name = 'Chocolate Brownie');

-- Insert test meal requests (only if they don't exist)
INSERT INTO meal_requests (user_id, menu_item_id, quantity, status)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM menu_items WHERE name = 'Classic Burger'),
    1, 'pending'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Classic Burger')
);

INSERT INTO meal_requests (user_id, menu_item_id, quantity, status)
SELECT 
    (SELECT id FROM users WHERE email = 'employee2@mealsync.com'),
    (SELECT id FROM menu_items WHERE name = 'Caesar Salad'),
    2, 'approved'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee2@mealsync.com')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Caesar Salad')
);

INSERT INTO meal_requests (user_id, menu_item_id, quantity, status)
SELECT 
    (SELECT id FROM users WHERE email = 'employee1@mealsync.com'),
    (SELECT id FROM menu_items WHERE name = 'Chocolate Brownie'),
    1, 'completed'
WHERE NOT EXISTS (
    SELECT 1 FROM meal_requests 
    WHERE user_id = (SELECT id FROM users WHERE email = 'employee1@mealsync.com')
    AND menu_item_id = (SELECT id FROM menu_items WHERE name = 'Chocolate Brownie')
); 