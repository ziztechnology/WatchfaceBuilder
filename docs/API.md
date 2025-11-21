# API Documentation

## Overview

Watchface Builder can be run as a web service, exposing REST APIs for building watchfaces programmatically.

## Starting the Server

```bash
watchface-builder serve --port 8080
```

## Endpoints

### Build Watchface

**Endpoint**: `POST /api/build`

**Description**: Build a watchface package with specified options.

**Request Body**:
```json
{
  "name": "My Watchface",
  "version": "1.0.0",
  "author": "Your Name",
  "description": "A beautiful watchface",
  "tags": ["minimal", "digital"],
  "template": "simple",
  "generatePreview": true
}
```

**Parameters**:
- `name` (string, required): Watchface name
- `version` (string): Version number (default: "1.0.0")
- `author` (string): Author name (default: "Anonymous")
- `description` (string): Watchface description
- `tags` ([]string): Array of tags
- `template` (string, required): Template type (`simple`, `analog`, `digital`, `custom`)
- `customHTML` (string): Custom HTML content (for custom template)
- `customCSS` (string): Custom CSS content (for custom template)
- `customJS` (string): Custom JS content (for custom template)
- `generatePreview` (boolean): Whether to generate preview image (default: true)

**Response**:
```json
{
  "success": true,
  "zipPath": "My_Watchface_v1.0.0_20250121_100000.zip",
  "downloadURL": "/api/download/My_Watchface_v1.0.0_20250121_100000.zip",
  "fileHash": "a3f5c8...",
  "size": 15360,
  "fileCount": 5,
  "files": [
    "manifest.json",
    "index.html",
    "style.css",
    "script.js",
    "preview.png"
  ],
  "manifest": {
    "name": "My Watchface",
    "version": "1.0.0",
    "author": "Your Name",
    "entrypoint": "index.html"
  }
}
```

### List Templates

**Endpoint**: `GET /api/templates`

**Description**: Get a list of available templates.

**Response**:
```json
{
  "templates": [
    {
      "id": "simple",
      "name": "Simple Digital Clock",
      "description": "Minimalist digital clock with gradient background",
      "category": "digital",
      "difficulty": "easy",
      "features": ["time", "date", "gradient", "responsive"]
    },
    {
      "id": "analog",
      "name": "Analog Clock",
      "description": "Classic clock with Canvas rendering",
      "category": "analog",
      "difficulty": "medium",
      "features": ["hour hand", "minute hand", "second hand", "canvas"]
    },
    {
      "id": "digital",
      "name": "Digital Clock",
      "description": "Tech-style digital clock with neon effects",
      "category": "digital",
      "difficulty": "easy",
      "features": ["time", "date", "day of week", "neon effects"]
    }
  ]
}
```

### Download Watchface

**Endpoint**: `GET /api/download/:filename`

**Description**: Download a built watchface package.

**Response**: ZIP file download

## Examples

### Using cURL

**Build a watchface**:
```bash
curl -X POST http://localhost:8080/api/build \
  -H "Content-Type: application/json" \
  -d '{
    "name": "My Watchface",
    "template": "analog",
    "version": "1.0.0",
    "author": "Your Name"
  }'
```

**List templates**:
```bash
curl http://localhost:8080/api/templates
```

**Download watchface**:
```bash
curl -O http://localhost:8080/api/download/My_Watchface_v1.0.0_20250121_100000.zip
```

### Using JavaScript/Fetch

```javascript
// Build watchface
async function buildWatchface() {
  const response = await fetch('http://localhost:8080/api/build', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      name: 'My Watchface',
      template: 'simple',
      version: '1.0.0',
      author: 'Your Name',
      tags: ['minimal', 'digital']
    })
  });

  const result = await response.json();
  console.log('Build result:', result);

  // Download the file
  window.location.href = `http://localhost:8080${result.downloadURL}`;
}

// List templates
async function listTemplates() {
  const response = await fetch('http://localhost:8080/api/templates');
  const templates = await response.json();
  console.log('Available templates:', templates);
}
```

### Using Python

```python
import requests

# Build watchface
def build_watchface():
    url = 'http://localhost:8080/api/build'
    data = {
        'name': 'My Watchface',
        'template': 'analog',
        'version': '1.0.0',
        'author': 'Your Name',
        'tags': ['classic', 'analog']
    }

    response = requests.post(url, json=data)
    result = response.json()
    print('Build result:', result)

    # Download file
    if result['success']:
        download_url = f"http://localhost:8080{result['downloadURL']}"
        file_response = requests.get(download_url)
        with open(result['zipPath'], 'wb') as f:
            f.write(file_response.content)
        print(f"Downloaded: {result['zipPath']}")

# List templates
def list_templates():
    url = 'http://localhost:8080/api/templates'
    response = requests.get(url)
    templates = response.json()
    print('Available templates:', templates)

if __name__ == '__main__':
    build_watchface()
    list_templates()
```

## Error Handling

All endpoints return errors in the following format:

```json
{
  "success": false,
  "error": "Error message description"
}
```

**Common Errors**:
- `400 Bad Request`: Invalid parameters
- `500 Internal Server Error`: Build failed

## Rate Limiting

By default, no rate limiting is applied. For production use, consider adding rate limiting middleware.

## Authentication

By default, the API is open. For production use, consider adding authentication:

```bash
watchface-builder serve --port 8080 --auth-token YOUR_SECRET_TOKEN
```

Then include the token in requests:

```bash
curl -X POST http://localhost:8080/api/build \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_SECRET_TOKEN" \
  -d '{"name": "My Watchface", "template": "simple"}'
```

## CORS Configuration

CORS is enabled by default for all origins. To restrict origins:

```bash
watchface-builder serve --port 8080 --cors-origins "https://example.com,https://app.example.com"
```
