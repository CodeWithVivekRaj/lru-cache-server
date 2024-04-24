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

