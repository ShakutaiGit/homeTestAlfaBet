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


## Database Setup and CGO Configuration

### Enabling CGO for Database Support ( can be passed if you use windows)

This project uses a database that requires cgo for successful integration. 
CGO enables the Go compiler to call C code directly, which is essential for some database drivers, particularly those that rely on C libraries.

To ensure smooth operation with the database, the following steps were taken:

1. **CGO Enabled Flag**:
    - The project is compiled with CGO enabled. This is achieved by setting the environment variable `CGO_ENABLED=1`.
    - Enabling CGO allows the use of the database driver that depends on C libraries.

2. **Installing CGO Support**:
    - A specific installer was used to support CGO on the development environment. For example, if you're on Windows, you might have installed MinGW or Cygwin to provide a GCC compiler environment, which is a requirement for CGO to compile C code.
    - Ensure that the appropriate tools are installed and correctly configured in your environment to compile projects with CGO dependencies.


## Windows Specific Setup for CGO

### Enabling CGO on Windows with TDM-GCC

CGO allows Go programs to call C code, which is essential for some database drivers in Go projects. To enable CGO on Windows:

#### Install TDM-GCC (GCC Compiler)

1. **Download and Install TDM-GCC**:
    - TDM-GCC is a compiler suite for Windows. It's a more user-friendly way to install GCC on Windows.
    - Download TDM-GCC from [TDM-GCC's Official Site](https://jmeubank.github.io/tdm-gcc/).
    - Run the installer and follow the instructions. It's recommended to use the default settings unless you have specific requirements.

#### Set Environment Variables

2. **Configure CGO_ENABLED**:
    - CGO must be explicitly enabled in Windows.
    - Set the environment variable `CGO_ENABLED` to `1`.
    - To do this, follow these steps:
        - Right-click on 'This PC' or 'My Computer' and select 'Properties'.
        - Click on 'Advanced system settings'.
        - In the System Properties window, go to the 'Advanced' tab and click on 'Environment Variables'.
        - Under 'System variables', click 'New' to add a new variable.
            - Variable name: `CGO_ENABLED`
            - Variable value: `1`
        - Click 'OK' to save the new environment variable.

3. **Add TDM-GCC to the System Path (it might work without it)**: 
    - Ensure that the TDM-GCC bin directory is added to your system's PATH environment variable.
    - The typical path to add is `C:\TDM-GCC-64\bin` (adjust the path if you installed TDM-GCC in a different location).
    - To add TDM-GCC to the PATH:
        - In the Environment Variables window, find and select the `Path` variable in the 'System variables' section, then click 'Edit'.
        - Click 'New' and add the TDM-GCC bin directory path.
        - Click 'OK' to close all dialogs.