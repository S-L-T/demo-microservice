# Demo Microservice

## Assumptions during development & further improvements

### Description
The app is fully dockerized and easily managed by Makefile commands.
On the assignment's points:
* **Provide an HTTP/gRPC API**
* * The app exposes both HTTP (`/user:8080`) and gRPC (`:8081`) endpoints.
* **Use a sensible storage mechanism for the Users**
* * Even though the original schema was in JSON, I decided to use a MySQL database instead of a NoSQL one because it made more sense to me for this particular assignment. The structure of User data is pretty stable, without any differentiations or diverse fields. The database's connection string is `demo:demo@tcp(db:3306)/users`.
* **Have the ability to notify other interested services of changes to User entities**
* * The app is able to raise an event at domain and presentation level. Either at domain level (`domain/use_case/user.go`), where an event would always be raised, or at presentation level (`presentation/http/server.go` and `presentation/grpc/user/user_server.go`), where an event would be raised only when the API was consumed by HTTP or gRPC respectively.
* **Have meaningful logs**
* * The app uses [logrus](https://github.com/sirupsen/logrus), which logs the date, time, error message, file and exact line of the error.
* **Be well documented**
* * That's fairly subjective, but I believe self-documenting code is the best form of documentation. I have added very few comments to the code, but I'm more than willing to discuss this.
* **Have a health check**
* * I have only implemented rudimentary healthcheck capabilities. If the server is running, a healthcheck response will be `200` for HTTP and `SERVING` for gRPC. A more sophisticated response could include data source response time & used capacity, number of active connections, API version, uptime etc.


### Improvements
This app was developed under the assumption that it was a test assignment; not an app to be deployed to production. The following points are some examples of what might have been different if this app was to be deployed.
* Use of environmental variables to hide sensitive data and improve configurability (passwords, ports etc.)
* `docker-compose.yaml` shouldn't be part of the app, but part of the entire platform's development environment.
* Safeguards against attacks; API gateways such as Kong, authentication or API keys could prevent unauthorized users from consuming the API.
* Event streaming platforms (such as Kafka) could be employed so that 3rd parties could be notified when something important happens on the app.
* The app logs events on the container's console instead of a clour logging provider.
* Multiple app instances (using a shared data source with appropriate locking/syncing configuration) with a load balancer to handle traffic would allow for seamless scaling of the service.
* When deploying, a cloud hosting platform such as AWS could be used, along with CI/CD tools to assist in faster and safer deployments.

## Makefile commands
When running the app for the first time, run `make start`. When the application starts, run `make reset-db`.

| Command          | Description                                |
|------------------|--------------------------------------------|
| `start`          | Starts the application.                    |
| `start-no-cache` | Rebuilds and starts the application.       |
| `stop`           | Stops the application.                     |
| `generate-proto` | Generates gRPC code from .proto files.     |
| `reset-db`       | Resets the database with some initial data |
| `run-tests`      | Runs all the application's tests.          |

## Tests
I have only included a unit test for the HTTP server. I considered integration tests to be out of scope for this assignment, but sadly did not have time for more thorough unit testing.

## Example HTTP requests
| Method    | Payload                                                                                                                                                                                                                             |
|-----------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `GET`     | `{ "filters": {}, "page_num": 1, "page_size": 3 }`                                                                                                                                                                                  |
| `PUT`     | `{ "user": { "first_name": "FirstName", "last_name": "LastName", "nickname": "Nickname", "password": "Password", "email": "email@email.com", "country": "Country" } }`                                                              |
| `PATCH`   | `{ "id": "3625b6fef5a511ecbcc10242ac170002", "user": { "first_name": "NewFirstName", "last_name": "NewLastName", "nickname": "NewNickname", "password": "NewPassword", "email": "new_email@email.com", "country": "NewCountry" } }` |
| `DELETE`  | `{ "id": "3624561af5a511ecbcc10242ac170002" }`                                                                                                                                                                                      |
| `OPTIONS` |                                                                                                                                                                                                                                     |
