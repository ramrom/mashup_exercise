# mashup_exercise

A service that notifies you of the best restaurants in you location.

Uses google's geoloc API to obtain the location of the client (longitude/latitude) and queries the Yelp API for the top 5 restaurants in the area.

The service can return JSON or HTML or notify you via cell using the Twilio API.

## Installation
- Install the standard go tools
- clone this down and run "go build"
- run the server on a port of your choice
