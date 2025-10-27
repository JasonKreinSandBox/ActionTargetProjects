This Program uses Server Sent Events to update statistics on basic ping request to a host.
I chose this method as it was a quick way to determine the online status of the host, there would
be no client -> server data needed. Its a simple program which will ping the host and return response
time. Data returned in this case will be simply number of ping attempts, response time for each ping
and average of response times for all attempts.

The build file was made on a cross platform (windows) and will make a linux executable file called main-linux
Once the main-linux file is run, open a browser to http://localhost:8080/events to view the returned results.