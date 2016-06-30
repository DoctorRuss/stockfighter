# stockfighter
Simple Go client for stockfighter.io

Used as a learning experience in Go and Cloud 9.

First step was to make a heartbeat call.

Second step is to make an extra call to venue heartbeat and to unmarshal into JSON.

Next steps were done in a big lump. 
* Added API calls to the Game Master to start and end levels.
* Created JSON types for all API calls.
* Added websockets to get quote or tickertape: bit provisional at the mo.