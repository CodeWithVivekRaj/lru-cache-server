<!-- Title -->
<h1 align="center">LRU Cache Server</h1>

<!-- Description -->
<p align="center">Lightweight HTTP server built with GoLang for simple caching mechanism with real-time updates using WebSocket.</p>

<!-- Table of Contents -->
## Table of Contents

- [About](#about)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)

<!-- About -->
## About

LRU Cache Server is a lightweight HTTP server built with GoLang that provides a simple caching mechanism. It allows users to store key-value pairs with expiration times and retrieve them later via HTTP API endpoints and also serves a socket connection to get realtime data.

### Key Features

- **Simple Caching Mechanism:** Allows users to store key-value pairs with expiration times.
- **HTTP API Endpoints:** Provides endpoints for setting, retrieving, and deleting values from the cache.
- **Real-time Updates:** Supports WebSocket for real-time updates about cache changes.
- **Efficient LRU Algorithm:** Uses a Least Recently Used (LRU) eviction policy to maintain cache size and performance.

### Purpose

The purpose of this project is to offer a lightweight and easy-to-use caching solution for applications that require temporary storage of key-value data. Whether it's caching API responses, session data, or frequently accessed resources, the LRU Cache Server provides a reliable and efficient way to manage and access cached data.

### Goals

- Provide a simple and intuitive API for interacting with the cache.
- Ensure high performance and efficiency in storing and accessing cached data.
- Support real-time updates to connected clients through WebSocket.
- Maintain code simplicity and readability for easy maintenance and future enhancements.
  

<!-- Installation -->
## Installation

To install and run the Cache Server locally, follow these steps:

1. **Clone the repository:**

   <pre><code class="language-json">git clone <repository-url></code></pre>

2. **Navigate to the project directory::**

   <pre><code class="language-json">cd lru-cache-server</code></pre>

3. **Run the server::**

   <pre><code class="language-json">go run main.go</code></pre>


<!-- Usage -->
<h2 id="usage">Usage</h2>

<p>Once the server is running, you can interact with the cache using the following API endpoints:</p>

<ul>
  <li><strong>GET /get:</strong> Retrieve a value from the cache.</li>
  <li><strong>POST /set:</strong> Set a value in the cache.</li>
  <li><strong>DELETE /delete:</strong> Delete a value from the cache.</li>
</ul>

<h3>Example</h3>

<p>Set a key-value pair in the cache:</p>

<pre><code class="language-bash">curl -X POST "http://localhost:8080/set?key=mykey1&amp;value=value1&amp;expiration=10s"
</code></pre>

<p>Response:</p>

<pre><code class="language-json">{
  "message": "Set mykey1=value1 with expiration 10s"
}
</code></pre>

<p>Retrieve the value associated with a key:</p>

<pre><code class="language-bash">curl -X GET "http://localhost:8080/get?key=mykey1"
</code></pre>

<p>Response:</p>

<pre><code class="language-json">{
  "key": "mykey1",
  "value": "value1"
   "message": "key value fethed",
}
</code></pre>

<p>Delete a key-value pair from the cache:</p>

<pre><code class="language-bash">curl -X DELETE "http://localhost:8080/delete?key=mykey1"
</code></pre>

<p>Response:</p>

<pre><code class="language-json">{
  "message": "Deleted mykey1"
}
</code></pre>




<!-- Endpoints -->
<h2 id="endpoints">API Endpoints</h2>

<p>The Cache Server provides the following API endpoints:</p>

<ul>
  <li><strong>GET /get:</strong> Retrieve a value from the cache. Specify the key as a query parameter to retrieve the corresponding value.</li>
  <li><strong>POST /set:</strong> Set a value in the cache. Provide the key, value, and expiration time (optional) as query parameters.</li>
  <li><strong>DELETE /delete:</strong> Delete a value from the cache. Specify the key as a query parameter to remove the corresponding key-value pair.</li>
</ul>

### GET /get

- **Description:** Retrieve a value from the cache.
- **Request:** GET /get?key=<key>
- **Response:** JSON object containing the key-value pair.

### POST /set

- **Description:** Set a value in the cache.
- **Request:** POST /set with JSON body: {"key": "example_key", "value": "example_value", "expiration": "10s"}
- **Response:** JSON object confirming the key-value pair was set.

### DELETE /delete

- **Description:** Delete a value from the cache.
- **Request:** DELETE /delete?key=<key>
- **Response:** JSON object confirming the key was deleted.

### WebSocket Endpoint

- **Description:** The WebSocket endpoint provides real-time updates about cache changes. Clients can connect to this endpoint to receive JSON messages containing cache updates.
  
- **Endpoint:** `ws://localhost:8080/ws`
  
- **Response Format:** JSON object representing cache updates. Each object includes the following fields:
  - `"key"`: The key of the cache item.
  - `"value"`: The value associated with the key.
  - `"duration"`: The remaining time until the expiration of the cache item.

**Example Response:**
```json
{
  "key": "example_key",
  "value": "example_value",
  "duration": "5m"
}


