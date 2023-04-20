# Echo server (with debug abilities)

A simple echo server that listens for different HTTP Methods and sends back a response that can contain information about:

- the request's parameters,
- caller's details,
- running environment details,
- and others

But, it can also be used to load the running environment's CPU and/or memory by providing:

- some CLI flags,
- or some Headers with the HTTP Request,
- or querry params,
- or body payload,
- etc.

It is written in two programming languages:
- Go - ./go_echo_server
   - [README.go.md][go_readme]
- Rust - ./rust_echo_server
   - [README.rust.md][rust_readme]

[go_readme]: ./go_echo_server/README.go.md
[rust_readme]: ./rust_echo_server/README.rust.md

