# MealSync - Software Requirements Specification (SRS)

## 1. Overview

**Project Name:** MealSync  
**Purpose:** MealSync is a RESTful API service designed to streamline meal planning, employee meal preferences, and final meal estimations for organizations. It aims to reduce food wastage and eliminate manual errors in meal planning by providing a structured and automated system for meal management.

**Business Objective:**  
MealSync enables employees to confirm their participation in scheduled meals (e.g., breakfast, lunch, snacks) and allows administrators to efficiently plan, track, and report meal requests. The system integrates with attendance systems to ensure accurate meal estimations and reduce food wastage. By automating meal planning and providing real-time insights, MealSync helps organizations save costs and improve employee satisfaction.

**Key Features:**

- Employees can view upcoming meal events and submit meal requests.
- Administrators can create meal events, add menu sets, and track meal preferences.
- Real-time integration with attendance systems to exclude employees on leave.
- Detailed reporting and analytics for meal planning and wastage reduction.

**Technology Stack:**

- Language: Go
- Web Framework: Gin
- ORM: GORM
- Authentication: JWT (Access/Refresh tokens)
- Database: PostgreSQL
- Documentation: Swagger/OpenAPI

## 2. User Roles

### 2.1 Admin

Admins are responsible for managing meal events and overseeing the entire meal planning process. Their key responsibilities include:

- Setting up meal events up to 30 days in advance.
- Creating and managing menu sets and menu items.
- Assigning multiple locations to meal events.
- Viewing and managing employee meal requests (private to admins).
- Canceling or assigning meals for employees.
- Viewing dashboards and generating reports (e.g., stats, trends, item popularity, wastage reduction).
- Exporting final meal estimation data to Excel.
- Viewing breakdowns of meal requests by department, item, and date.
- Integrating with external attendance/leave systems via API to exclude employees on leave.

### 2.2 Employee (User)

Employees are the end-users who interact with the system to manage their meal preferences. Their key responsibilities include:

- Logging in and registering using JWT-based authentication.
- Viewing upcoming meal events created by the admin.
- Submitting meal requests for available meal events.
- Selecting one menu set from the options provided for a meal event.
- Selecting or deselecting individual items within the chosen menu set.
- Choosing one location from the available options for a meal event.
- Withdrawing meal requests before the respective cutoff times.
- Receiving meal confirmation notifications after cutoff times.
- Enabling or disabling notification preferences.
- Adding public comments on meal items for specific meal events (visible to all users).

## 3. Functional Requirements

### 3.1 Meal Planning (Admin)

- Create/edit/delete meal events for future dates (max 30 days).
- Add multiple menu sets sets per event.
- Set cutoff times

### 3.2 Meal Requesting (Employee)

- Submit meal requests for available meal events.
- Choose preferred set
- Withdraw request before cutoff time.
- Get confirmation once cutoff is reached.
- Set notification preferences (opt-in/opt-out).
- Receive reminders to submit meal requests.

### 3.3 Meal Comments (Employee)

- Post public comments on meal item of a meal event.
- View all employee comments.
- Comments are visible to all; meal requests are private.

### 3.4 Estimation (Admin)

- View requested meal sets by day.
- Breakdown of meal choices by department, item, date.
- Only confirmed meals (post-cutoff) count toward final estimation.
- Cancel or assign meals to any user.
- Export estimation to Excel.
- Connect to external attendance system to exclude employees on leave.

## 4. Non-Functional Requirements

### 4.1 Architecture

- Follow clean MVC architecture with isolated concerns:
  - DTO layer
  - Controller/API layer
  - Service layer
  - Repository layer
- Use interfaces for dependency injection.
- Implement structured logging and error handling.
- Modular project structure for scalability.
- Ensure the application is stateless to support horizontal scaling.
- Use containerization (e.g., Docker) and orchestration (e.g., Kubernetes) for deployment.
- Add observability tools like Prometheus and Grafana for monitoring and alerting.
- Implement distributed tracing (e.g., OpenTelemetry) for debugging and performance analysis.
- Use caching (e.g., Redis) for frequently accessed data to reduce database load.
- Add circuit breakers (e.g., Resilience4j) and retry mechanisms with exponential backoff for resilience.

### 4.2 API Design

- Fully RESTful JSON APIs.
- No frontend logic in this project.
- Swagger documentation for all endpoints.
- Public endpoints:
  - Login
  - Register
  - Token refresh
- Private endpoints (JWT-protected): All others
- Pagination and filtering support for admin data views.
- Request validation via separate validation layer/helpers.
- Add API versioning (e.g., `/v1/`) to ensure backward compatibility.
- Implement rate limiting to prevent abuse and ensure fair usage.
- Consider adding GraphQL for flexible querying, especially for admin dashboards and reporting.
- Define a consistent error response format (e.g., HTTP status codes, error codes, and messages).

### 4.3 Security

- JWT-based auth with access and refresh tokens.
- Role-based middleware access control.
- Secure password storage (bcrypt).
- Input validation and sanitization.
- Encrypt sensitive data at rest (e.g., using AES-256) and in transit (e.g., using HTTPS/TLS).
- Add audit logs for critical actions (e.g., admin creating meal events, users submitting requests).
- Ensure compliance with OWASP Top 10 security practices (e.g., preventing SQL injection, XSS).
- Add optional Multi-Factor Authentication (MFA) for admin accounts to enhance security.

### 4.4 Performance

- Define specific performance metrics (e.g., API response time < 200ms for 95% of requests).
- Conduct load testing to ensure the system can handle peak traffic.

### 4.5 Disaster Recovery

- Implement automated backups and a disaster recovery plan with defined RPO (Recovery Point Objective) and RTO (Recovery Time Objective).

### 4.6 Internationalization (i18n)

- Add support for multiple languages and time zones to make the system globally usable.

### 4.7 Compliance

- Ensure compliance with relevant regulations (e.g., GDPR for data privacy).

## 5. Integration

### 5.1 Attendance API Integration

- MealSync will call an external service to determine employee presence or leave status.
- On final meal confirmation (post-cutoff), any user on leave will have their request auto-cancelled.

## 6. Database Schema Design

### Core Tables

```sql
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
            FROM menu_item_comments
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

CREATE INDEX idx_meal_event_sets_deleted_at ON meal_event_sets(deleted_at);
CREATE INDEX idx_meal_event_sets_meal_event_id ON meal_event_sets(meal_event_id);
CREATE INDEX idx_meal_event_sets_menu_set_id ON meal_event_sets(menu_set_id);

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

CREATE TABLE menu_item_comments (
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
BEFORE INSERT OR UPDATE ON menu_item_comments
FOR EACH ROW
EXECUTE FUNCTION validate_menu_item_comment();

CREATE TRIGGER trg_rating_insert
AFTER INSERT ON menu_item_comments
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();

CREATE TRIGGER trg_rating_update
AFTER UPDATE OF rating ON menu_item_comments
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();

CREATE TRIGGER trg_rating_delete
AFTER DELETE ON menu_item_comments
FOR EACH ROW
EXECUTE FUNCTION update_menu_item_rating();


CREATE INDEX idx_menu_item_comment_deleted_at ON menu_item_comments(deleted_at);
CREATE INDEX idx_menu_item_comment_meal_event_id ON menu_item_comments(meal_event_id);
CREATE INDEX idx_menu_item_comment_menu_item_id ON menu_item_comments(menu_item_id);
CREATE INDEX idx_menu_item_comment_rating ON menu_item_comments(rating);
```

---

## 7. Success Metrics

Define measurable success criteria to evaluate the system's effectiveness:

- Reduction in food wastage by 20% within the first 6 months.
- Average API response time under 200ms for 95% of requests.
- User adoption rate: 80% of employees actively using the system within 3 months of deployment.
- Admin satisfaction score of 90% based on internal surveys.

## 8. Assumptions and Constraints

### 8.1 Assumptions

- Employees will have access to the internet to use the system.
- Admins will provide accurate meal event details and cutoff times.
- The system will primarily be used during business hours.
- Employees will use modern browsers that support the latest web standards.

### 8.2 Constraints

- The system must support up to 10,000 users concurrently.
- Deployment is limited to PostgreSQL as the database.
- The system must comply with GDPR for data privacy.
- Initial deployment will be on a cloud provider (e.g., AWS, GCP, or Azure).

## 9. Testing Requirements

### 9.1 Unit Testing

- Ensure coverage for all critical business logic, including meal requests, menu set selection, and reporting.

### 9.2 Integration Testing

- Validate seamless interaction between modules (e.g., meal requests and reporting).

### 9.3 Load Testing

- Simulate peak traffic to ensure the system can handle up to 10,000 concurrent users.

### 9.4 Security Testing

- Validate against OWASP Top 10 vulnerabilities, including SQL injection, XSS, and CSRF.

### 9.5 User Acceptance Testing (UAT)

- Conduct UAT with a sample group of employees and admins to ensure the system meets business requirements.

## 10. Deployment and Maintenance

### 10.1 Deployment

- Use a CI/CD pipeline for automated builds, testing, and deployment.
- Use containerization (e.g., Docker) and orchestration (e.g., Kubernetes) for scalability.

### 10.2 Maintenance

- Schedule regular maintenance windows for updates and patches.
- Implement automated backups with a retention period of 30 days.
- Monitor system health using tools like Prometheus and Grafana.

## 11. User Stories

### 11.1 Employee User Stories

- **As an employee**, I want to select a menu set and location for a meal event so that I can confirm my participation.
- **As an employee**, I want to deselect items within a menu set so that I can customize my meal.
- **As an employee**, I want to withdraw my meal request before the cutoff time so that I can update my plans.

### 11.2 Admin User Stories

- **As an admin**, I want to create meal events with multiple menu sets and locations so that employees can make their selections.
- **As an admin**, I want to view reports on meal requests by location so that I can plan food quantities accurately.
- **As an admin**, I want to export meal request data to Excel so that I can share it with catering services.

## 12. Future Enhancements

- Connect with google chat for employees to manage meal requests on the go.
- Add AI-based meal quantity prediction to further reduce wastage.
- Integrate with payment systems for paid meal events.
- Enable push notifications for meal request reminders and confirmations.
