CREATE OR REPLACE FUNCTION validate_meal_request()
RETURNS TRIGGER AS $$
BEGIN

    -- Check if the selected menu_set is available for the meal_event
    IF NOT EXISTS (
        SELECT 1
        FROM meal_event_sets
        WHERE meal_event_id = NEW.meal_event_id
          AND menu_set_id = NEW.menu_set_id
    ) THEN
        RAISE EXCEPTION 'Invalid selection: Menu set % is not available for event %',
            NEW.menu_set_id, NEW.meal_event_id;
    END IF;

    -- Check if the selected address is available for the meal_event
    IF NOT EXISTS (
        SELECT 1
        FROM meal_event_addresses
        WHERE meal_event_id = NEW.meal_event_id
          AND address_id = NEW.event_address_id
    ) THEN
        RAISE EXCEPTION 'Invalid selection: Event address % is not available for event %',
            NEW.event_address_id, NEW.meal_event_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION validate_requested_item()
RETURNS TRIGGER AS $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM meal_requests mr
        JOIN menu_set_items msi ON mr.menu_set_id = msi.menu_set_id
        WHERE mr.user_id = NEW.user_id
          AND mr.meal_event_id = NEW.meal_event_id
          AND msi.menu_item_id = NEW.menu_item_id
    ) THEN
        RAISE EXCEPTION 'Invalid request: Menu item % not in selected menu set for user % in event %',
            NEW.menu_item_id, NEW.user_id, NEW.meal_event_id;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION validate_menu_item_comment()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.meal_event_id IS NOT NULL THEN
        -- Check that the menu_item appears in any set in that meal_event
        IF NOT EXISTS (
            SELECT 1
            FROM meal_event_sets mes
            JOIN menu_set_items msi ON mes.menu_set_id = msi.menu_set_id
            WHERE mes.meal_event_id = NEW.meal_event_id
              AND msi.menu_item_id = NEW.menu_item_id
        ) THEN
            RAISE EXCEPTION 'Invalid comment: Item % not served at event %',
                NEW.menu_item_id, NEW.meal_event_id;
        END IF;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_menu_item_rating()
RETURNS TRIGGER AS $$
BEGIN
    -- Only update rating if rating is not null
    IF NEW.rating IS NOT NULL THEN
        UPDATE menu_items
        SET average_rating = (
            SELECT ROUND(AVG(rating)::numeric, 2)
            FROM menu_item_comment
            WHERE menu_item_id = NEW.menu_item_id
              AND rating IS NOT NULL
        )
        WHERE id = NEW.menu_item_id;
    END IF;

    RETURN NULL;
END;
$$ LANGUAGE plpgsql;





CREATE TABLE users (
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

CREATE TABLE event_addresses (
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

CREATE TABLE menu_items (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE,
  description TEXT,
  average_rating NUMERIC(3,2),
  image_url TEXT,
  is_active BOOLEAN DEFAULT TRUE,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_items_deleted_at ON menu_items(deleted_at);

CREATE TABLE menu_sets (
  id SERIAL PRIMARY KEY,
  menu_set_name VARCHAR(100) NOT NULL UNIQUE,
  menu_set_description TEXT,
  is_active BOOLEAN DEFAULT TRUE,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_sets_deleted_at ON menu_sets(deleted_at);

CREATE TABLE menu_set_items (
  menu_set_id INT REFERENCES menu_sets(id) ON DELETE CASCADE,
  menu_item_id INT REFERENCES menu_items(id) ON DELETE CASCADE,
  PRIMARY KEY (menu_set_id, menu_item_id),
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_set_items_deleted_at ON menu_set_items(deleted_at);


CREATE TABLE meal_events (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  event_date TIMESTAMP NOT NULL,
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

CREATE TABLE meal_event_sets (
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  menu_set_id INT REFERENCES menu_sets(id) ON DELETE CASCADE,
  PRIMARY KEY (meal_event_id, menu_set_id),
  label TEXT,
  note TEXT,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_meal_event_menu_sets_deleted_at ON meal_event_sets(deleted_at);
CREATE INDEX idx_meal_event_menu_sets_meal_event_id ON meal_event_sets(meal_event_id);
CREATE INDEX idx_meal_event_menu_sets_menu_set_id ON meal_event_sets(menu_set_id);

CREATE TABLE meal_event_addresses (
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  address_id INT REFERENCES event_addresses(id) ON DELETE CASCADE,
  PRIMARY KEY (meal_event_id, address_id),
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_meal_event_addresses_deleted_at ON meal_event_addresses(deleted_at);
CREATE INDEX idx_meal_event_addresses_meal_event_id ON meal_event_addresses(meal_event_id);
CREATE INDEX idx_meal_event_addresses_address_id ON meal_event_addresses(address_id);

CREATE TABLE meal_requests (
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  PRIMARY KEY (user_id, meal_event_id),
  menu_set_id INT REFERENCES menu_sets(id),
  event_address_id INT REFERENCES event_addresses(id),
  confirmed_at TIMESTAMP DEFAULT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE TRIGGER trg_validate_meal_request
BEFORE INSERT OR UPDATE ON meal_requests
FOR EACH ROW
EXECUTE FUNCTION validate_meal_request();


CREATE INDEX idx_meal_requests_deleted_at ON meal_requests(deleted_at);
CREATE INDEX idx_meal_requests_meal_event_id ON meal_requests(meal_event_id);
CREATE INDEX idx_meal_requests_user_id ON meal_requests(user_id);
CREATE UNIQUE INDEX idx_unique_meal_request ON meal_requests(user_id, meal_event_id) WHERE deleted_at IS NULL;


-- Junction table for meal request items (many-to-many)
CREATE TABLE user_requested_items (
  user_id INT,
  meal_event_id INT,
  menu_item_id INT REFERENCES menu_items(id),
  PRIMARY KEY (user_id, meal_event_id, menu_item_id),
  quantity INT DEFAULT 1 NOT NULL CHECK (quantity > 0),
  notes TEXT,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE TRIGGER trg_validate_requested_item
BEFORE INSERT OR UPDATE ON user_requested_items
FOR EACH ROW
EXECUTE FUNCTION validate_requested_item();


CREATE INDEX idx_user_requested_items_deleted_at ON user_requested_items(deleted_at);
CREATE INDEX idx_user_requested_items_menu_item_id ON user_requested_items(menu_item_id);

CREATE TABLE menu_item_comment (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id) ON DELETE CASCADE,
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  menu_item_id INT NOT NULL REFERENCES menu_items(id),
  comment TEXT NOT NULL,
  rating SMALLINT CHECK (rating BETWEEN 1 AND 5) NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);


CREATE TRIGGER trg_validate_menu_item_comment
BEFORE INSERT OR UPDATE ON menu_item_comment
FOR EACH ROW
EXECUTE FUNCTION validate_menu_item_comment();

CREATE TRIGGER trg_rating_insert
AFTER INSERT ON menu_item_comment
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();

CREATE TRIGGER trg_rating_update
AFTER UPDATE OF rating ON menu_item_comment
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();

CREATE TRIGGER trg_rating_delete
AFTER DELETE ON menu_item_comment
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();


CREATE INDEX idx_menu_item_comment_deleted_at ON menu_item_comment(deleted_at);
CREATE INDEX idx_menu_item_comment_meal_event_id ON menu_item_comment(meal_event_id);
CREATE INDEX idx_menu_item_comment_menu_item_id ON menu_item_comment(menu_item_id);
CREATE INDEX idx_menu_item_comment_rating ON menu_item_comment(rating);
