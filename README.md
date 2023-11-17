# AlfaBet Home Test
Ron Shakutai

## API Overview

This project provides a RESTFul API to manage events, detailed as follows:

### Endpoints

#### Schedule a New Event
- **Endpoint**: `POST /events`
- **Description**: Schedules a new event.
- **Body**: JSON payload containing event details.

#### Retrieve a List of All Scheduled Events
- **Endpoint**: `GET /events`
- **Description**: Retrieves a list of events, with optional filtering and sorting.
- **Query Parameters**:
   - `location`: (Optional) Filter events by location. For example, `location=New York` retrieves all events in New York.
   - `sortBy`: (Optional) Sorts events based on certain criteria. Available options:
      - `date`: Sorts events by start time.
      - `popularity`: Sorts events by the number of participants.
      - `creation`: Sorts events by the creation time of the event.
   - **Example**: `http://localhost:8080/events?location=New York&sortBy=popularity` retrieves events in New York, sorted by their popularity.

#### Retrieve Details of a Specific Event
- **Endpoint**: `GET /events/:id`
- **Description**: Retrieves detailed information about a specific event.
- **Parameters**:
   - `id`: The unique identifier of the event.

#### Update Details of a Specific Event
- **Endpoint**: `PUT /events/:id`
- **Description**: Updates the details of a specific event.
- **Parameters**:
   - `id`: The unique identifier of the event.
- **Body**: JSON payload containing updated event details.

#### Delete a Specific Event
- **Endpoint**: `DELETE /events/:id`
- **Description**: Deletes a specific event.
- **Parameters**:
   - `id`: The unique identifier of the event.

### Using the API with Postman
A folder named 'postman' has been included in the project, containing the Postman collection for the API.
A Postman collection is available for easy interaction with the API. Follow these steps to use it:
1. Import the collection into Postman (from the folder).
2. Ensure the API server is running.
3. Send requests to the API via Postman.

## Architecture

The project is structured using the Ports & Adapters (also known as Hexagonal Architecture) methodology. This design pattern promotes the separation of concerns by dividing the application into distinct layers:
- **Domain Layer**: Contains the business logic.
- **Application Layer**: Coordinates high-level application behavior.
- **Adapters Layer**: Connects the application with external concerns like databases and web interfaces.

## Running Unit Tests

To run the unit tests, navigate to the project's root directory and execute the following command:
i have made some unit & integration tests.

```bash
go test ./... -v
```

This command runs all the tests in the project and provides verbose output. The tests are designed to validate the functionality of individual units of code independently of external systems.

## Installation & Setup

Follow these steps to set up the project and install its dependencies:

1. **Clone the Repository**:
   ```bash
   git clone [repository-url]
   ```

2. **Navigate to the Project Directory**:
   ```bash
   cd [project-name]
   ```

3. **Install Dependencies**:
   The project uses Go Modules for dependency management. Run the following command to install all required dependencies:
   ```bash
   go mod tidy
   ```

4. **Start the Server**:
   ```bash
   go run main.go
   ```
   This starts the server on the default port (e.g., `:8080`). Ensure this port is available or adjust the port in the configuration.

