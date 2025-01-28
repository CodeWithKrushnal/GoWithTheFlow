# REST API Report

# How HTTP  1, 1.1, 2 and 3 works

HTTP (Hypertext Transfer Protocol) is the foundation of data communication on the web. Over the years, several versions of HTTP have been developed to improve performance, reliability, and security.

HTTP/1.0 opens a new connection per request, while HTTP/1.1 supports persistent connections. HTTP/2 introduces multiplexing and binary protocols, improving performance. HTTP/3, built on QUIC (UDP-based), reduces latency with faster handshakes, better congestion control, and mandatory encryption. HTTP/3 is the most efficient, especially on mobile or unreliable networks.

Reference:

| Feature                    | HTTP/1.0                          | HTTP/1.1                          | HTTP/2                                             | HTTP/3                        |
| -------------------------- | --------------------------------- | --------------------------------- | -------------------------------------------------- | ----------------------------- |
| Protocol Type              | Text-based                        | Text-based                        | Binary                                             | Binary (based on QUIC)        |
| Connection Management      | One connection per request        | Persistent (keep-alive)           | Multiplexed (one connection for multiple requests) | Multiplexed (via QUIC)        |
| Header Compression         | No                                | No                                | HPACK                                              | HPACK (QUIC)                  |
| Multiplexing               | No                                | No                                | Yes                                                | Yes                           |
| Server Push                | No                                | No                                | Yes                                                | Yes                           |
| Latency (Connection Setup) | High (new connection per request) | Moderate (persistent connections) | Lower (multiple requests in one connection)        | Very Low (QUIC's 0-RTT)       |
| Protocol Security          | No encryption (optional)          | Optional TLS/SSL                  | Mandatory encryption (TLS/1.2 or higher)           | Built-in encryption (TLS 1.3) |
| Transport Protocol         | TCP                               | TCP                               | TCP                                                | QUIC (UDP-based)              |
| Head-of-Line Blocking      | Yes                               | Yes                               | No                                                 | No                            |
| Adoption/Support           | Older browsers only               | Widely supported                  | Increasingly supported                             | Emerging, not universal yet   |

# OPTIONS, HEAD http methods

The OPTIONS method is used to describe the communication options available for a resource or server. It allows clients to query the server for supported HTTP methods and capabilities without requesting any specific data. It’s useful for CORS (Cross-Origin Resource Sharing) and pre-flight checks.

The HEAD method is similar to GET, but it only retrieves the headers of a resource, not the body. It’s often used to check if a resource exists, to verify metadata (like modification date), or to check the size of a resource without downloading it. Both methods are typically used for information retrieval or testing.

# X-Rate-Limit-Reset, why there is X in header names

The "X-" prefix in HTTP header names was originally used to denote non-standard or custom headers, which are not part of the official HTTP specification. The X-Rate-Limit-Reset header indicates the time when the rate limit will reset, helping clients manage request limits for APIs.

While the "X-" prefix was widely used for custom headers, its use has been deprecated by the IETF (Internet Engineering Task Force) to avoid confusion. Today, custom headers should follow more specific conventions without the "X-" prefix, but many legacy systems and APIs still use it for custom or non-standard headers like X-Rate-Limit-Reset.

# Difference - Handle and Handlefunc, Handler and Handlerfunc

In Go, Handle and HandleFunc are methods for registering HTTP handlers, but they differ in how they are used:

1. Handle: Takes an http.Handler interface as an argument. It’s used when you have a custom handler that implements the ServeHTTP method (e.g., a struct or a method).

2. HandleFunc: Takes a function with the signature func(http.ResponseWriter, *http.Request) as an argument. It’s a more concise way to register a simple handler function.

Both register handlers with the http.ServeMux router, but HandleFunc is often more convenient for basic routing, while Handle is used for more complex custom handlers.
