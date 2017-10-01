### Communication Between Cloud Services
```
Key-points Covered:
- Communication between microservices and how that can be a bottleneck.
- Reusing Connection to improve performance by avoiding repeated TCP slow-start, congestion-control ramp ups and connection negotiations.
- Faster JSON Marshalling and Unmarshalling that avoids extra time spent reflecting.
- Using protocol buffering instead of JSON for messaging.
- Communicating over RPC using gRPC.
```
