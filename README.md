<div align="center">

#### $`\textcolor{#342ca8}{\text{SCORE PUBLISHER SERVICE}}`$



[![Golang][golang-shield]][golang-url]
[![Gin][gin-shield]][gin-url]
[![GRPC][grpc-shield]][grpc-url]
[![Pgx][pgx-shield]][pgx-url]

</div>


##

## Overview


## Architecture


The architecture can be separated into 3 layers, including  `Service`, `Transport`, and `Repository`.


- `Transport` handles input request things, such as HTTP, websockets, gRPC request routing, authentication, access control, and parameter validation.
- `Repository` handle output requests, such as accessing DB, communicate with other services.
- `Service` handles use cases 


<details><summary>Test usage</summary>
  
- [testify](https://github.com/stretchr/testify)
- [mockgen](https://github.com/golang/mock)

</details>