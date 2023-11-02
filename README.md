Welcome to the TechSystems Coding Test Project - completed by Don Jackson

Instructions:
1) Create a folder and unzip the file contents into that folder

2) Open a terminal and navigate to the project directory.

3) Make sure Docker is running on your machine

4) Start the services by running: 
docker-compose up

*** Optional ***
If there are problems, here are some hints:
- In another terminal, you may make sure both services (tecksystems-app AND mysql:latest) are running: 
docker ps
- Make sure Docker is running. You may need to stop and restart it
- Sometimes you may want to delete all containers in Docker and start from scratch. If any containiners won't delete, run these first:
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
You can build the application separatlely by running:
docker build -t myapp .
- You may see the program image by running
docker images
*****************

5) Test the application using Postman or curl. The application is reachable on localhost:8080. Here are the available endpoints:
a) Retrieve a list of alerts: (replace {my_test_service_id} with the actual service id - same with start and end ts)
   - GET
   - URL:localhost:8080/alerts/?service_id={my_test_service_id}&start_ts={1695643160}&end_ts={1695644360}
   or CURL:
   curl --location 'localhost:8080/alerts/?service_id=my_test_service_id&start_ts=1695643160&end_ts=1695644360'

c) Add an alert:
   - POST
   - URL: http://localhost:8080/alerts
   - BODY
      {
         "alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
         "service_id": "my_test_service_id",
         "service_name": "my_test_service",
         "model": "my_test_model",
         "alert_type": "anomaly",
         "alert_ts": "1695644160",
         "severity": "warning",
         "team_slack": "slack_ch"
      }
   Or CURL:
      curl --location 'localhost/alerts' \
      --header 'Content-Type: application/json' \
      --data '{
         "alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
         "service_id": "my_test_service_id",
         "service_name": "my_test_service",
         "model": "my_test_model",
         "alert_type": "anomaly",
         "alert_ts": "1695644160",
         "severity": "warning",
         "team_slack": "slack_ch"
      }
      '

Please let me know if you have any questions or encounter any issues during setup or testing.
Don Jackson
(408) 858-6490
donj91711@yahoo.com
