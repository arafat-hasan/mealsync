-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    employee_id INT UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    department VARCHAR(100),
    role VARCHAR(50) CHECK (role IN ('admin', 'user')) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_deleted_at ON users(deleted_at);

-- Create event_addresses table
CREATE TABLE IF NOT EXISTS event_addresses (
    id SERIAL PRIMARY KEY,
    address TEXT NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_event_addresses_deleted_at ON event_addresses(deleted_at);

-- Create menu_items table
CREATE TABLE IF NOT EXISTS menu_items (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    image_url TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_items_deleted_at ON menu_items(deleted_at);

-- Create menu_sets table
CREATE TABLE IF NOT EXISTS menu_sets (
    id SERIAL PRIMARY KEY,
    menu_set_name TEXT NOT NULL,
    menu_set_description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_sets_deleted_at ON menu_sets(deleted_at);

-- Create menu_set_items table
CREATE TABLE IF NOT EXISTS menu_set_items (
    id SERIAL PRIMARY KEY,
    menu_set_id INT REFERENCES menu_sets(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    UNIQUE(menu_set_id, menu_item_id),
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_set_items_deleted_at ON menu_set_items(deleted_at);

-- Create meal_events table
CREATE TABLE IF NOT EXISTS meal_events (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    event_date DATE NOT NULL,
    event_duration INT NOT NULL, -- in minutes
    cutoff_time TIMESTAMP NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    confirmed_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_meal_events_deleted_at ON meal_events(deleted_at);

-- Create meal_event_menu_sets table
CREATE TABLE IF NOT EXISTS meal_event_menu_sets (
    id SERIAL PRIMARY KEY,
    meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
    menu_set_id INT REFERENCES menu_sets(id) ON DELETE CASCADE,
    label TEXT,
    note TEXT,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    UNIQUE(meal_event_id, menu_set_id)
);

CREATE INDEX idx_meal_event_menu_sets_deleted_at ON meal_event_menu_sets(deleted_at);
CREATE INDEX idx_meal_event_menu_sets_meal_event_id ON meal_event_menu_sets(meal_event_id);
CREATE INDEX idx_meal_event_menu_sets_menu_set_id ON meal_event_menu_sets(menu_set_id);

-- Create meal_event_addresses table
CREATE TABLE IF NOT EXISTS meal_event_addresses (
    id SERIAL PRIMARY KEY,
    meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
    address_id INT REFERENCES event_addresses(id) ON DELETE CASCADE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    UNIQUE(meal_event_id, address_id)
);

CREATE INDEX idx_meal_event_addresses_deleted_at ON meal_event_addresses(deleted_at);
CREATE INDEX idx_meal_event_addresses_meal_event_id ON meal_event_addresses(meal_event_id);
CREATE INDEX idx_meal_event_addresses_address_id ON meal_event_addresses(address_id);

-- Create meal_requests table
CREATE TABLE IF NOT EXISTS meal_requests (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES users(id),
    meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
    event_menu_set_id INT,
    meal_event_address_id INT,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    CONSTRAINT fk_event_menu_set FOREIGN KEY (event_menu_set_id) 
        REFERENCES meal_event_menu_sets(id),
    CONSTRAINT fk_event_address FOREIGN KEY (meal_event_address_id) 
        REFERENCES meal_event_addresses(id),
    CONSTRAINT uq_employee_event UNIQUE(employee_id, meal_event_id)
);

CREATE INDEX idx_meal_requests_deleted_at ON meal_requests(deleted_at);
CREATE INDEX idx_meal_requests_meal_event_id ON meal_requests(meal_event_id);
CREATE INDEX idx_meal_requests_employee_id ON meal_requests(employee_id);

-- Create meal_request_items table
CREATE TABLE IF NOT EXISTS meal_request_items (
    id SERIAL PRIMARY KEY,
    meal_request_id INT REFERENCES meal_requests(id) ON DELETE CASCADE,
    menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
    menu_set_id INT REFERENCES menu_sets(id) ON DELETE CASCADE,
    is_selected BOOLEAN DEFAULT TRUE NOT NULL,
    quantity INT DEFAULT 1 NOT NULL,
    notes TEXT,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    UNIQUE(meal_request_id, menu_item_id)
);

CREATE INDEX idx_meal_request_items_deleted_at ON meal_request_items(deleted_at);
CREATE INDEX idx_meal_request_items_selected ON meal_request_items(is_selected);
CREATE INDEX idx_meal_request_items_meal_request_id ON meal_request_items(meal_request_id);
CREATE INDEX idx_meal_request_items_menu_item_id ON meal_request_items(menu_item_id);

-- Create meal_comments table
CREATE TABLE IF NOT EXISTS meal_comments (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES users(id),
    meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
    event_menu_set_id INT,
    menu_item_id INT,
    comment TEXT NOT NULL,
    rating INT CHECK (rating >= 1 AND rating <= 5),
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id),
    CONSTRAINT fk_menu_set_event FOREIGN KEY (event_menu_set_id) 
        REFERENCES meal_event_menu_sets(id),
    CONSTRAINT fk_menu_item_event_set FOREIGN KEY (menu_item_id) 
        REFERENCES menu_items(id) ON DELETE CASCADE
);

CREATE INDEX idx_meal_comments_deleted_at ON meal_comments(deleted_at);
CREATE INDEX idx_meal_comments_meal_event_id ON meal_comments(meal_event_id);
CREATE INDEX idx_meal_comments_event_menu_set_id ON meal_comments(event_menu_set_id);
CREATE INDEX idx_meal_comments_menu_item_id ON meal_comments(menu_item_id);

-- Create notifications table
CREATE TABLE IF NOT EXISTS notifications (
    id SERIAL PRIMARY KEY,
    employee_id INT REFERENCES users(id) ON DELETE CASCADE,
    type TEXT CHECK(type IN ('reminder', 'confirmation', 'admin-message')),
    payload JSONB,
    read BOOLEAN DEFAULT FALSE,
    deleted_at TIMESTAMP DEFAULT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    created_by INT REFERENCES users(id),
    updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_notifications_deleted_at ON notifications(deleted_at);

-- Create admin_address_item_report view
CREATE OR REPLACE VIEW admin_address_item_report AS
SELECT 
    ea.id AS address_id,
    ea.address AS address_name,
    me.id AS meal_event_id,
    me.name AS meal_event_name,
    me.event_date,
    mi.id AS menu_item_id,
    mi.name AS menu_item_name,
    ms.id AS menu_set_id,
    ms.menu_set_name,
    COUNT(mri.id) AS item_request_count,
    SUM(CASE WHEN mri.is_selected THEN mri.quantity ELSE 0 END) AS total_quantity
FROM meal_requests mr
JOIN meal_event_addresses mel ON mr.meal_event_id = mel.meal_event_id AND mr.meal_event_address_id = mel.address_id
JOIN event_addresses ea ON mel.address_id = ea.id
JOIN meal_events me ON mr.meal_event_id = me.id
JOIN meal_event_menu_sets ems ON mr.event_menu_set_id = ems.id
JOIN menu_sets ms ON ems.menu_set_id = ms.id
JOIN meal_request_items mri ON mr.id = mri.meal_request_id
JOIN menu_items mi ON mri.menu_item_id = mi.id
WHERE mr.deleted_at IS NULL
    AND mri.deleted_at IS NULL
    AND ea.deleted_at IS NULL
    AND me.deleted_at IS NULL
    AND ms.deleted_at IS NULL
    AND mi.deleted_at IS NULL
    AND me.confirmed_at IS NOT NULL
GROUP BY 
    ea.id, ea.address, me.id, me.name, me.event_date, 
    mi.id, mi.name, ms.id, ms.menu_set_name; 