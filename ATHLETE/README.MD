# ATHLETE SERVICE

- CreateAthlete -> For creating an athlete, uses Message broker
- UpdateAthlete -> For updating the data of an athlete, uses Message broker
- GetAthleteById -> For getting athlete data from the server by their id, message broker is not used
- GetAthleteByName -> For getting athlete data from the server by their name, message broker is not used
- GetAthleteBySurname -> For getting athlete data from the server by their surname, message broker is not used
- GetAthletesBySportId -> For getting athletes data from the server by the id of the sport, message broker is not used
- GetAthletesByCountryId -> For getting athlete data from the server by their country id, message broker is not used
- DeleteAthleteById -> For deleting an athlete data from the server by their id, message broker is used


## Overview

The Athlete Service provides various endpoints to manage athlete data. The service uses a message broker for creating and updating athlete data, while direct queries are used for retrieving athlete information.

## Services

### CreateAthlete
- **Description**: Stores athlete data on the server. Receives data from the message broker as a consumer. Does not return any data to the client. Any errors encountered are logged.
- **Usage**: Message broker

### UpdateAthlete
- **Description**: Updates the data of an athlete. Receives data from the message broker as a consumer. Does not return any data to the client. Any errors encountered are logged.
- **Usage**: Message broker

### GetAthleteById
- **Description**: Retrieves athlete data from the server by their ID.
- **Usage**: Direct query (no message broker)

### GetAthleteByName
- **Description**: Retrieves athlete data from the server by their name.
- **Usage**: Direct query (no message broker)

### GetAthleteBySurname
- **Description**: Retrieves athlete data from the server by their surname.
- **Usage**: Direct query (no message broker)

### GetAthletesBySportId
- **Description**: Retrieves athlete data from the server by the ID of the sport.
- **Usage**: Direct query (no message broker)

### GetAthletesByCountryId
- **Description**: Retrieves athlete data from the server by their country ID.
- **Usage**: Direct query (no message broker)

## Service Details

### Create Athlete
- **Purpose**: The `CreateAthlete` service is used to store athlete data on the server. The server receives data from the message broker as a consumer. No data is returned to the client. Any errors encountered are logged.

### Update Athlete
- **Purpose**: The `UpdateAthlete` service is used to update existing athlete data on the server. Similar to the `CreateAthlete` service, data is received from the message broker as a consumer and no data is returned to the client. Any errors encountered are logged.

### Get Athlete By ID
- **Purpose**: The `GetAthleteById` service retrieves athlete data from the server using their unique ID. This service does not use the message broker.

### Get Athlete By Name
- **Purpose**: The `GetAthleteByName` service retrieves athlete data from the server using their name. This service does not use the message broker.

### Get Athlete By Surname
- **Purpose**: The `GetAthleteBySurname` service retrieves athlete data from the server using their surname. This service does not use the message broker.

### Get Athletes By Sport ID
- **Purpose**: The `GetAthletesBySportId` service retrieves athlete data from the server using the ID of the sport. This service does not use the message broker.

### Get Athletes By Country ID
- **Purpose**: The `GetAthletesByCountryId` service retrieves athlete data from the server using the country ID. This service does not use the message broker.

### Delete Athlete By ID
- **Purpose**: The `DeleteAthleteById` service retreives an athlet's id, and deletes it from the server softly. This service uses message broker.

## Error Handling
Any errors encountered during the `CreateAthlete` or `UpdateAthlete` operations are logged by the server. The service does not return any error information to the client for these operations.

## Usage Instructions
- To create or update athlete data, ensure the data is sent to the message broker. The service will consume the data and process it accordingly.
- For retrieving athlete data, use the appropriate endpoint with the required parameters (ID, name, surname, sport ID, or country ID).
