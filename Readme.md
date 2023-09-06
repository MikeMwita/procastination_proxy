# Procrastination-Blocking Web Proxy

Self-discipline is a valuable trait, but sometimes we all need a little assistance in staying focused. It's easy to get sidetracked by websites and waste precious time scrolling endlessly. This project aims to help you regain control of your online activities by creating a web proxy server in Go that can block requests to specified websites or domains.

## Table of Contents

1. [Blocking Everything](#blocking-everything)
2. [Allowing Everything](#allowing-everything)
3. [Static Blocklist](#static-blocklist)
4. [Configurable Blocklist](#configurable-blocklist)
5. [Dynamic Blocklist](#dynamic-blocklist)
6. [Office Hours](#office-hours)
7. [Productization](#productization)
8. [Extensions](#extensions)

## 1. Blocking Everything

The simplest functionality to implement is blocking everything. Write a program that listens on a specified port for HTTP requests and responds to all of them with a status code of 403 (Forbidden).

```go
// Example HTTP client configuration to use the proxy:
client := &http.Client{
    Transport: &http.Transport{
        Proxy: http.ProxyURL(&url.URL{
            Scheme: "http",
            Host:   "localhost:XXXX",
        }),
    },
}
resp, err := client.Get("some URL")



**Allowing Everything**
Instead of responding with a 403 status code, modify the proxy to make the request and return the response to the user.

Static Blocklist
Enhance the proxy to block requests only if they match a global, static blocklist, allowing all other requests.

Configurable Blocklist
Allow the proxy server administrator to configure the blocklist when starting the server. This can be achieved by passing the blocklist as a parameter during server startup.

Dynamic Blocklist
Implement a special 'admin' endpoint in the proxy server to enable administrators to connect via HTTP and block or unblock specific domains. Admins can make requests like:

/admin/block/reddit.com
/admin/unblock/facebook.com
Ensure the blocklist is concurrency-safe to avoid data races during simultaneous updates.
Office Hours
Add an 'office hours' feature that enforces the blocklist only during specific hours of the day. The exact times should be configurable by the administrator when starting the server.

Productization
To make this project a shipping product, you need to add several components:

TLS and WebSocket Support: Secure the proxy with TLS for encrypted communication and support WebSocket connections.
Error Handling: Handle and log errors gracefully to enhance stability.
Documentation: Create comprehensive documentation, including installation, configuration, and usage instructions.
Release Automation: Implement a release process, such as versioning and automated deployment.
Extensions
Feel free to add additional features to enhance the proxy's functionality:

Content-Aware Blocking: Block web pages containing specific keywords or content.
Per-User Blocklists: Allow individual users to have custom blocklists.
Time-Limited Blocking: Implement time quotas for accessing blocked sites.
Warning-Only Mode: Show a warning message on blocked sites, informing users of potential distractions.
Content Caching: Cache content to improve performance and reduce bandwidth usage.