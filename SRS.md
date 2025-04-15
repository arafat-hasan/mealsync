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
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_event_addresses_deleted_at ON event_addresses(deleted_at);

CREATE TABLE menu_items (
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

CREATE TABLE menu_sets (
  id SERIAL PRIMARY KEY,
  menu_set_name TEXT NOT NULL,
  menu_set_description TEXT,
  menu_set_is_active BOOLEAN DEFAULT TRUE,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_menu_sets_deleted_at ON menu_sets(deleted_at);

CREATE TABLE menu_set_items (
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

CREATE TABLE meal_events (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  event_date DATE NOT NULL,
  event_duration INT NOT NULL, -- in minutes
  cutoff_time TIMESTAMP NOT NULL,
  is_active BOOLEAN DEFAULT TRUE,
  is_confirmed BOOLEAN DEFAULT FALSE,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id)
);

CREATE INDEX idx_meal_events_deleted_at ON meal_events(deleted_at);

CREATE TABLE meal_event_menu_sets (
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

CREATE TABLE meal_event_addresses (
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

CREATE TABLE meal_requests (
  id SERIAL PRIMARY KEY,
  employee_id INT REFERENCES users(id) ON DELETE CASCADE,
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  event_menu_set_id INT,
  address_id INT,
  deleted_at TIMESTAMP DEFAULT NULL,
  is_confirmed BOOLEAN DEFAULT FALSE,
  confirmed_at TIMESTAMP,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id),
  -- Composite foreign key for event menu set
  CONSTRAINT fk_event_menu_set FOREIGN KEY (meal_event_id, event_menu_set_id) 
    REFERENCES meal_event_menu_sets(meal_event_id, id),
  -- Composite foreign key for event address
  CONSTRAINT fk_event_address FOREIGN KEY (meal_event_id, address_id) 
    REFERENCES meal_event_addresses(meal_event_id, address_id),
  CONSTRAINT uq_employee_event UNIQUE(employee_id, meal_event_id)
);

CREATE INDEX idx_meal_requests_deleted_at ON meal_requests(deleted_at);
CREATE INDEX idx_meal_requests_meal_event_id ON meal_requests(meal_event_id);
CREATE INDEX idx_meal_requests_employee_id ON meal_requests(employee_id);

-- Junction table for meal request items (many-to-many)
CREATE TABLE meal_request_items (
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

CREATE TABLE meal_comments (
  id SERIAL PRIMARY KEY,
  employee_id INT REFERENCES users(id) ON DELETE CASCADE, 
  meal_event_id INT REFERENCES meal_events(id) ON DELETE CASCADE,
  menu_set_id INT,
  menu_item_id INT,
  comment TEXT NOT NULL,
  deleted_at TIMESTAMP DEFAULT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW(),
  created_by INT REFERENCES users(id),
  updated_by INT REFERENCES users(id),
  CONSTRAINT fk_menu_set_event FOREIGN KEY (meal_event_id, menu_set_id) 
    REFERENCES meal_event_menu_sets(meal_event_id, id),
  CONSTRAINT fk_menu_item_set FOREIGN KEY (menu_set_id, menu_item_id) 
    REFERENCES menu_set_items(menu_set_id, menu_item_id)
);

CREATE INDEX idx_meal_comments_deleted_at ON meal_comments(deleted_at);
CREATE INDEX idx_meal_comments_meal_event_id ON meal_comments(meal_event_id);
CREATE INDEX idx_meal_comments_menu_set_id ON meal_comments(menu_set_id);
CREATE INDEX idx_meal_comments_menu_item_id ON meal_comments(menu_item_id);

CREATE TABLE notifications (
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

-- Analytics view for admin reporting: items requested by address
CREATE VIEW admin_address_item_report AS
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
  AND mr.is_confirmed = TRUE
GROUP BY 
  ea.id, ea.address, me.id, me.name, me.event_date, 
  mi.id, mi.name, ms.id, ms.menu_set_name;
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
