-- First, drop any triggers
DROP TRIGGER IF EXISTS trg_validate_meal_request ON meal_requests;
DROP TRIGGER IF EXISTS trg_validate_requested_item ON user_requested_items;
DROP TRIGGER IF EXISTS trg_validate_menu_item_comment ON menu_item_comments;
DROP TRIGGER IF EXISTS trg_rating_insert ON menu_item_comments;
DROP TRIGGER IF EXISTS trg_rating_update ON menu_item_comments;
DROP TRIGGER IF EXISTS trg_rating_delete ON menu_item_comments;

-- Then drop tables in correct dependency order
DROP TABLE IF EXISTS notifications;
DROP TABLE IF EXISTS menu_item_comments;
DROP TABLE IF EXISTS user_requested_items;
DROP TABLE IF EXISTS meal_requests;
DROP TABLE IF EXISTS meal_event_addresses;
DROP TABLE IF EXISTS meal_event_sets;
DROP TABLE IF EXISTS meal_events;
DROP TABLE IF EXISTS menu_set_items;
DROP TABLE IF EXISTS menu_sets;
DROP TABLE IF EXISTS menu_items;
DROP TABLE IF EXISTS event_addresses;
DROP TABLE IF EXISTS users;

-- Finally drop functions
DROP FUNCTION IF EXISTS validate_meal_request();
DROP FUNCTION IF EXISTS validate_requested_item();
DROP FUNCTION IF EXISTS validate_menu_item_comment();
DROP FUNCTION IF EXISTS update_menu_item_rating();
