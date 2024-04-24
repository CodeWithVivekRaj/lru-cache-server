<!-- Title -->
<h1 align="center">LRU Cache Server</h1>

<!-- Description -->
<p align="center">Lightweight HTTP server built with GoLang for simple caching mechanism.</p>

<!-- Table of Contents -->
## Table of Contents

- [About](#about)
- [Installation](#installation)
- [Usage](#usage)
- [Endpoints](#endpoints)

<!-- About -->
## About

LRU Cache Server is a lightweight HTTP server built with GoLang that provides a simple caching mechanism. It allows users to store key-value pairs with expiration times and retrieve them later via HTTP API endpoints.

<!-- Installation -->
## Installation

To install and run the Cache Server locally, follow these steps:

1. **Clone the repository:**

   git clone <repository-url>

2. **Navigate to the project directory::**

   cd lru-cache-server

3. **Run the server::**

   go run main.go

<!-- Usage -->
<h2 id="usage">## Usage</h2>

<p>Once the server is running, you can interact with the cache using the following API endpoints:</p>

<ul>
  <li><strong>GET /get:</strong> Retrieve a value from the cache.</li>
  <li><strong>POST /set:</strong> Set a value in the cache.</li>
  <li><strong>DELETE /delete:</strong> Delete a value from the cache.</li>
</ul>

<h3>Example</h3>

<p>Set a key-value pair in the cache:</p>

<pre><code class="language-bash">curl -X POST "http://localhost:8080/set?key=mykey&amp;value=myvalue&amp;expiration=10s"
</code></pre>

<p>Retrieve the value associated with a key:</p>

<pre><code class="language-bash">curl -X GET "http://localhost:8080/get?key=mykey"
</code></pre>

<p>Delete a key-value pair from the cache:</p>

<pre><code class="language-bash">curl -X DELETE "http://localhost:8080/delete?key=mykey"
</code></pre>


<h2 id="endpoints">## Endpoints</h2>

<ul>
  <li><strong>GET /get:</strong> Retrieve a value from the cache.</li>
  <li><strong>POST /set:</strong> Set a value in the cache.</li>
  <li><strong>DELETE /delete:</strong> Delete a value from the cache.</li>
</ul>
