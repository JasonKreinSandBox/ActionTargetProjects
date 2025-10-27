This Program uses Server Sent Events to update statistics on basic ping request to a host.
I chose this method as it was a quick way to determine the online status of the host, there would
be no client -> server data needed. Its a simple program which will ping the host and return response
time. Data returned in this case will be simply number of ping attempts, response time for each ping
and average of response times for all attempts.

The build file was made on a cross platform (windows) and will make a linux executable file called main-linux
Once the main-linux file is run, open a browser to http://localhost:8080/events to view the returned results.

Project Notes: I chose to try the project using Golang, as I knew this was the language used at Action Targets. I was unfamiliar
with this language and I have little experience with front end programming. However I enjoy learning and am not afraid to tread
in unfamiliar territory to learn new skills, and techniques. Due to this and my limited time available (2 days), the project 
is still half baked.

TODOs:
    Impliment robust ICMP function to include use of port flag as basic ping doesn't include port handling
    Create unit tests to validate functions and enhance error handling
    Fix client side html and beautify index.html to make returned data more readable
    Setup and configure linux environment to test build file
