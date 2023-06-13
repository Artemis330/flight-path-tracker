# Flight Path Tracker Microservice API

This microservice API is designed to help track the flight paths of individuals by sorting through their flight records. The API accepts a request that includes a list of flights defined by a source and destination airport code and returns the total flight path starting and ending airports.

## API Endpoint

Endpoint: /calculate

This endpoint accepts a JSON payload containing a list of flights, each defined as a pair of source and destination airport codes. The flights may not be listed in order and will need to be sorted to find the total flight paths starting and ending airports. The endpoint returns a JSON object containing the sorted flight path from the starting airport to the ending airport.

### Input

The input to the API should be a JSON array containing one or more flight segments. Each segment should be an array containing two elements: the source airport code and the destination airport code.

Example Input: [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]

### Output

The output of the API is a JSON array containing the sorted list of airport codes representing the complete flight path from the first airport in the first segment to the last airport in the last segment.

Example Output: ["IND", "EWR"]

## Usage

To use this microservice API, send a POST request to the `/calculate` endpoint with a JSON payload containing the list of flights to be sorted. The response will contain the sorted list of airport codes representing the complete flight path.
