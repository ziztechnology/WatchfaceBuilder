package builder

import "fmt"

// generateSimpleTemplate generates the simple template
func (b *Builder) generateSimpleTemplate(options BuildOptions) map[string]string {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="container">
        <div class="time" id="time">00:00:00</div>
        <div class="date" id="date">2025-01-21</div>
    </div>
    <script src="script.js"></script>
</body>
</html>`, options.Name)

	css := `* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    width: 100%%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
    overflow: hidden;
}

.container {
    text-align: center;
    color: white;
}

.time {
    font-size: 4rem;
    font-weight: 300;
    letter-spacing: 0.1em;
    text-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
}

.date {
    font-size: 1.5rem;
    margin-top: 1rem;
    opacity: 0.9;
    font-weight: 300;
}

@media (max-width: 480px) {
    .time {
        font-size: 3rem;
    }
    .date {
        font-size: 1.2rem;
    }
}`

	js := `function updateTime() {
    const now = new Date();

    // Format time as HH:MM:SS
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');
    const timeString = hours + ':' + minutes + ':' + seconds;

    // Format date
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const dateString = year + '-' + month + '-' + day;

    // Update DOM
    document.getElementById('time').textContent = timeString;
    document.getElementById('date').textContent = dateString;
}

// Update every second
updateTime();
setInterval(updateTime, 1000);`

	return map[string]string{
		"index.html": html,
		"style.css":  css,
		"script.js":  js,
	}
}

// generateAnalogTemplate generates the analog clock template
func (b *Builder) generateAnalogTemplate(options BuildOptions) map[string]string {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <canvas id="clock"></canvas>
    <script src="script.js"></script>
</body>
</html>`, options.Name)

	css := `* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    width: 100%%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(135deg, #f5f5f5 0%%, #e0e0e0 100%%);
    font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
    overflow: hidden;
}

#clock {
    border-radius: 50%%;
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.2);
}`

	js := `const canvas = document.getElementById('clock');
const ctx = canvas.getContext('2d');

// Set canvas size
const size = Math.min(window.innerWidth, window.innerHeight) * 0.9;
canvas.width = size;
canvas.height = size;

const centerX = size / 2;
const centerY = size / 2;
const radius = size / 2 - 20;

function drawClock() {
    const now = new Date();
    const hours = now.getHours() %% 12;
    const minutes = now.getMinutes();
    const seconds = now.getSeconds();

    // Clear canvas
    ctx.clearRect(0, 0, size, size);

    // Draw clock face
    ctx.beginPath();
    ctx.arc(centerX, centerY, radius, 0, 2 * Math.PI);
    ctx.fillStyle = '#ffffff';
    ctx.fill();
    ctx.strokeStyle = '#333333';
    ctx.lineWidth = 2;
    ctx.stroke();

    // Draw hour markers
    for (let i = 0; i < 12; i++) {
        const angle = (i * 30 - 90) * Math.PI / 180;
        const x1 = centerX + Math.cos(angle) * (radius - 15);
        const y1 = centerY + Math.sin(angle) * (radius - 15);
        const x2 = centerX + Math.cos(angle) * (radius - 5);
        const y2 = centerY + Math.sin(angle) * (radius - 5);

        ctx.beginPath();
        ctx.moveTo(x1, y1);
        ctx.lineTo(x2, y2);
        ctx.strokeStyle = '#333333';
        ctx.lineWidth = 3;
        ctx.stroke();
    }

    // Draw hour hand
    const hourAngle = ((hours + minutes / 60) * 30 - 90) * Math.PI / 180;
    drawHand(hourAngle, radius * 0.5, 6, '#333333');

    // Draw minute hand
    const minuteAngle = ((minutes + seconds / 60) * 6 - 90) * Math.PI / 180;
    drawHand(minuteAngle, radius * 0.7, 4, '#666666');

    // Draw second hand
    const secondAngle = (seconds * 6 - 90) * Math.PI / 180;
    drawHand(secondAngle, radius * 0.8, 2, '#e74c3c');

    // Draw center dot
    ctx.beginPath();
    ctx.arc(centerX, centerY, 8, 0, 2 * Math.PI);
    ctx.fillStyle = '#e74c3c';
    ctx.fill();
}

function drawHand(angle, length, width, color) {
    ctx.beginPath();
    ctx.moveTo(centerX, centerY);
    ctx.lineTo(
        centerX + Math.cos(angle) * length,
        centerY + Math.sin(angle) * length
    );
    ctx.strokeStyle = color;
    ctx.lineWidth = width;
    ctx.lineCap = 'round';
    ctx.stroke();
}

// Update every second
drawClock();
setInterval(drawClock, 1000);`

	return map[string]string{
		"index.html": html,
		"style.css":  css,
		"script.js":  js,
	}
}

// generateDigitalTemplate generates the digital clock template
func (b *Builder) generateDigitalTemplate(options BuildOptions) map[string]string {
	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>%s</title>
    <link rel="stylesheet" href="style.css">
</head>
<body>
    <div class="container">
        <div class="time" id="time">
            <span id="hours">00</span>
            <span class="separator">:</span>
            <span id="minutes">00</span>
            <span class="separator">:</span>
            <span id="seconds">00</span>
        </div>
        <div class="date" id="date">2025-01-21 星期二</div>
    </div>
    <script src="script.js"></script>
</body>
</html>`, options.Name)

	css := `* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
}

body {
    width: 100%%;
    height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(135deg, #0a0a1e 0%%, #1e1e3c 100%%);
    font-family: 'Courier New', monospace;
    overflow: hidden;
}

.container {
    text-align: center;
}

.time {
    font-size: 5rem;
    font-weight: bold;
    color: #00ffff;
    text-shadow:
        0 0 10px #00ffff,
        0 0 20px #00ffff,
        0 0 30px #00ffff;
    letter-spacing: 0.1em;
}

.separator {
    animation: blink 1s infinite;
}

@keyframes blink {
    0%%, 49%% { opacity: 1; }
    50%%, 100%% { opacity: 0; }
}

.date {
    font-size: 1.5rem;
    margin-top: 2rem;
    color: #00cccc;
    opacity: 0.8;
}

@media (max-width: 480px) {
    .time {
        font-size: 3rem;
    }
    .date {
        font-size: 1.2rem;
    }
}`

	js := `function updateTime() {
    const now = new Date();

    // Format time
    const hours = String(now.getHours()).padStart(2, '0');
    const minutes = String(now.getMinutes()).padStart(2, '0');
    const seconds = String(now.getSeconds()).padStart(2, '0');

    // Format date with day of week
    const year = now.getFullYear();
    const month = String(now.getMonth() + 1).padStart(2, '0');
    const day = String(now.getDate()).padStart(2, '0');
    const weekdays = ['星期日', '星期一', '星期二', '星期三', '星期四', '星期五', '星期六'];
    const weekday = weekdays[now.getDay()];
    const dateString = year + '-' + month + '-' + day + ' ' + weekday;

    // Update DOM
    document.getElementById('hours').textContent = hours;
    document.getElementById('minutes').textContent = minutes;
    document.getElementById('seconds').textContent = seconds;
    document.getElementById('date').textContent = dateString;
}

// Update every second
updateTime();
setInterval(updateTime, 1000);`

	return map[string]string{
		"index.html": html,
		"style.css":  css,
		"script.js":  js,
	}
}

// generateCustomTemplate generates a custom template
func (b *Builder) generateCustomTemplate(options BuildOptions) map[string]string {
	files := map[string]string{}

	if options.CustomHTML != "" {
		files["index.html"] = options.CustomHTML
	}
	if options.CustomCSS != "" {
		files["style.css"] = options.CustomCSS
	}
	if options.CustomJS != "" {
		files["script.js"] = options.CustomJS
	}

	return files
}
