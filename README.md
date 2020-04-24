# arbitrary
A http mock server for a arbitrary problem test

###Config example

    routes:
    - path: /foo/bar
      method: get
      throttling:
        max_simultaneous_requests: 100
        max_throttled_requests: 100
        delay_on_response: 200ms
        delay_on_throttled_response: 10s
      successful_response:
        status_code: 200
        body: |
          {"id": "xpto"}
      throttling_response:
        status_code: 429
        body: |
          {"error": "cannot handle the request"}