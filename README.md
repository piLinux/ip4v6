## IP4v6 | Find your IPv4/IPv6 address

[![Build Status](https://travis-ci.com/piLinux/ip4v6.svg?branch=master)][01]


#### For users

- Find your IPv4/v6
  - From any web browser, visit: https://ip.pilinux.me
  - From Linux/Mac Terminal or Windows command prompt: `curl ip.pilinux.me`
    - To find your IPv4 address: `curl ip.pilinux.me -4`
    - To find your IPv6 address: `curl ip.pilinux.me -6`

- Find your geolocation
  - From any web browser, visit: https://gl.pilinux.me
  - From Linux/Mac Terminal or Windows command prompt: `curl gl.pilinux.me`

- Find your user-agent type and version:
  - From any web browser, visit: https://ua.pilinux.me
  - From Linux/Mac Terminal or Windows command prompt: `curl ua.pilinux.me`

**Note 1:** These three URLs can be used in your daily work and for private/
commercial usage for free. But kindly do not abuse the services. If you need
to send over 250 HTTP(s) requests/day for your project, please create an issue
for a dedicated service.


#### For developers

- Test and build locally
  
  ```
  go get github.com/piLinux/ip4v6
  cd $(go env GOPATH)/src/github.com/piLinux/ip4v6/
  make test build
  ```

- Run locally
  - Create a file: `echo 'PORT=8081' > .env`
  - `./ip4v6`

- Test locally
  - `curl localhost:8081`


**Note 2:** This repository is for the learning purpose. The binary code
returns only the IPv4/IPv6 address. To find your user-agent and geolocation,
please check the previous section `For users`.


[01]: https://travis-ci.com/piLinux/ip4v6
