package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func getCommonStyles() string {
	return `
		@import url('https://fonts.googleapis.com/css2?family=Press+Start+2P&family=Orbitron:wght@400;700;900&display=swap');

		body {
			background-color: #0a0a0f;
			color: #00ff41;
			font-family: 'Orbitron', monospace;
			margin: 0;
			padding: 0;
			min-height: 100vh;
			animation: fadeIn 1s ease-in-out;
			background-image: 
				radial-gradient(circle at 25% 25%, #1a0033 0%, transparent 50%),
				radial-gradient(circle at 75% 75%, #003300 0%, transparent 50%);
		}

		@keyframes fadeIn {
			from { opacity: 0; transform: translateY(-20px); }
			to { opacity: 1; transform: translateY(0); }
		}

		.navbar {
			background-color: #111116;
			padding: 1rem 0;
			border-bottom: 2px solid #00ff41;
			box-shadow: 0 2px 20px rgba(0, 255, 65, 0.3);
			backdrop-filter: blur(10px);
		}

		.nav-container {
			max-width: 1200px;
			margin: 0 auto;
			display: flex;
			justify-content: center;
			align-items: center;
			gap: 2rem;
		}

		.nav-link {
			color: #00ff41;
			text-decoration: none;
			padding: 0.8rem 1.5rem;
			border-radius: 5px;
			transition: all 0.3s ease;
			font-weight: 500;
			border: 1px solid transparent;
			font-family: 'Press Start 2P', monospace;
			font-size: 0.8rem;
			text-transform: uppercase;
			letter-spacing: 1px;
		}

		.nav-link:hover {
			background-color: rgba(0, 255, 65, 0.1);
			color: #ffffff;
			transform: translateY(-2px);
			border-color: #00ff41;
			box-shadow: 0 4px 15px rgba(0, 255, 65, 0.4);
			text-shadow: 0 0 10px #00ff41;
		}

		.nav-link.active {
			background-color: #00ff41;
			color: #0a0a0f;
			font-weight: bold;
			box-shadow: 0 0 20px rgba(0, 255, 65, 0.6);
		}

		.container {
			padding: 3rem;
			display: flex;
			justify-content: center;
			align-items: center;
			min-height: calc(100vh - 100px);
			text-align: center;
		}

		.content {
			padding: 2rem 3rem;
			background-color: rgba(17, 17, 22, 0.9);
			border-radius: 15px;
			box-shadow: 
				0 4px 30px rgba(0, 255, 65, 0.2),
				inset 0 1px 0 rgba(255, 255, 255, 0.1);
			border: 2px solid rgba(0, 255, 65, 0.3);
			max-width: 600px;
			width: 100%;
			backdrop-filter: blur(10px);
		}

		h1 {
			color: #00ff41;
			margin-bottom: 0.5rem;
			font-size: 2.5rem;
			text-shadow: 
				0 0 10px rgba(0, 255, 65, 0.8),
				0 0 20px rgba(0, 255, 65, 0.4),
				0 0 30px rgba(0, 255, 65, 0.2);
			transition: all 0.3s ease;
			font-family: 'Press Start 2P', monospace;
			text-transform: uppercase;
			letter-spacing: 2px;
		}

		p {
			font-size: 1.2rem;
			color: #cccccc;
			line-height: 1.8;
			text-shadow: 0 0 5px rgba(0, 255, 65, 0.3);
		}

		.thing {
			font-size: 1rem;
			background-color: rgba(0, 255, 65, 0.1);
			padding: 0.4rem 0.8rem;
			border-radius: 8px;
			display: inline-block;
			border: 1px solid rgba(0, 255, 65, 0.3);
			font-family: 'Press Start 2P', monospace;
			font-size: 0.7rem;
			text-transform: uppercase;
			letter-spacing: 1px;
		}

		button {
			background: linear-gradient(45deg, #00ff41, #00cc33);
			color: #0a0a0f;
			border: none;
			padding: 15px 30px;
			font-size: 1rem;
			font-family: 'Press Start 2P', monospace;
			font-weight: bold;
			border-radius: 8px;
			cursor: pointer;
			margin-top: 2rem;
			transition: all 0.3s ease;
			text-transform: uppercase;
			letter-spacing: 1px;
			font-size: 0.8rem;
			box-shadow: 
				0 4px 15px rgba(0, 255, 65, 0.4),
				inset 0 1px 0 rgba(255, 255, 255, 0.2);
		}

		button:hover {
			background: linear-gradient(45deg, #00cc33, #00ff41);
			transform: translateY(-3px) scale(1.05);
			box-shadow: 
				0 8px 25px rgba(0, 255, 65, 0.6),
				0 0 30px rgba(0, 255, 65, 0.4);
			text-shadow: 0 0 10px rgba(0, 0, 0, 0.5);
		}

		.feature-grid {
			display: grid;
			grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
			gap: 1.5rem;
			margin-top: 2rem;
		}

		.feature-card {
			background: linear-gradient(135deg, rgba(0, 255, 65, 0.1), rgba(0, 204, 51, 0.1));
			padding: 1.5rem;
			border-radius: 12px;
			transition: all 0.3s ease;
			border: 1px solid rgba(0, 255, 65, 0.2);
			backdrop-filter: blur(5px);
		}

		.feature-card:hover {
			transform: translateY(-8px) scale(1.02);
			box-shadow: 
				0 12px 35px rgba(0, 255, 65, 0.3),
				0 0 25px rgba(0, 255, 65, 0.2);
			border-color: rgba(0, 255, 65, 0.5);
		}

		.feature-icon {
			font-size: 2.5rem;
			margin-bottom: 1rem;
			filter: drop-shadow(0 0 10px rgba(0, 255, 65, 0.6));
		}
	`
}

func main() {
	// Home page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Home - Go Web App</title>
	<style>` + getCommonStyles() + `</style>
</head>
<body>
	<nav class="navbar">
		<div class="nav-container">
			<a href="/" class="nav-link active" id="home-link">🏠 Home</a>
			<a href="/about" class="nav-link" id="about-link">📋 About</a>
		</div>
	</nav>

	<div class="container">
		<div class="content">
			<h1 id="main-heading">Home</h1>
			<p>Welcome!</p>
			<p><span class="thing">{{.Message}}</span></p>
			<button id="actionButton">{{.ButtonText}}</button>
		</div>
	</div>

	<script>
		document.addEventListener('DOMContentLoaded', () => {
			const button = document.getElementById('actionButton');
			const heading = document.getElementById('main-heading');
			const colors = ['#00ff41', '#ff0080', '#00ffff', '#ffff00', '#ff8000', '#8000ff'];
			let colorIndex = 0;

			button.addEventListener('click', () => {
				colorIndex = (colorIndex + 1) % colors.length;
				heading.style.color = colors[colorIndex];
				heading.style.textShadow = '0 0 10px ' + colors[colorIndex] + '80';
				
				// Add some fun animations
				heading.style.transform = 'scale(1.1)';
				setTimeout(() => {
					heading.style.transform = 'scale(1)';
				}, 200);
			});

			// Add smooth scroll effect when clicking nav links
			document.querySelectorAll('.nav-link').forEach(link => {
				link.addEventListener('click', (e) => {
					// Add loading animation
					document.body.style.opacity = '0.8';
					setTimeout(() => {
						document.body.style.opacity = '1';
					}, 100);
				});
			});
		});
	</script>
</body>
</html>`

		t, err := template.New("home").Parse(tmpl)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Message    string
			ButtonText string
		}{
			Message: `
			This is a simple Go web server that serves dynamic HTML page templates. Feel free to explore and modify the code to suit your needs.
			`,
			ButtonText: `
			🎨 Change Colors!
			`,
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// About page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl := `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>About - Go Web App</title>
	<style>` + getCommonStyles() + `</style>
</head>
<body>
	<nav class="navbar">
		<div class="nav-container">
			<a href="/" class="nav-link" id="home-link">🏠 Home</a>
			<a href="/about" class="nav-link active" id="about-link">📋 About</a>
		</div>
	</nav>

	<div class="container">
		<div class="content">
			<h1 id="about-heading">About This Project</h1>
			<p>{{.Description}}</p>
			
			<div class="feature-grid">
				{{range .Features}}
				<div class="feature-card">
					<div class="feature-icon">{{.Icon}}</div>
					<h3 style="color: #00ff41; margin: 0 0 0.5rem 0; font-family: 'Press Start 2P', monospace; font-size: 0.9rem; text-transform: uppercase; letter-spacing: 1px; text-shadow: 0 0 10px rgba(0, 255, 65, 0.6);">{{.Title}}</h3>
					<p style="margin: 0; font-size: 1rem;">{{.Description}}</p>
				</div>
				{{end}}
			</div>

			<button id="interactiveButton">{{.ButtonText}}</button>
		</div>
	</div>

	<script>
		document.addEventListener('DOMContentLoaded', () => {
			const button = document.getElementById('interactiveButton');
			const heading = document.getElementById('about-heading');
			const cards = document.querySelectorAll('.feature-card');
			
			let isAnimating = false;

			button.addEventListener('click', () => {
				if (isAnimating) return;
				isAnimating = true;

				// Animate heading
				heading.style.transform = 'rotateY(360deg)';
				heading.style.transition = 'transform 1s ease-in-out';

				// Animate cards with staggered effect - full 360 degree spin
				cards.forEach((card, index) => {
					setTimeout(() => {
						card.style.transform = 'scale(1.1) rotateZ(360deg)';
						card.style.transition = 'transform 0.8s ease-in-out';
						
						setTimeout(() => {
							card.style.transform = 'scale(1) rotateZ(0deg)';
							card.style.transition = 'transform 0.3s ease';
						}, 800);
					}, index * 150);
				});

				// Reset button state
				setTimeout(() => {
					heading.style.transform = 'rotateY(0deg)';
					isAnimating = false;
				}, 1000);
			});

			// Add hover effects to feature cards
			cards.forEach(card => {
				card.addEventListener('mouseenter', () => {
					if (!isAnimating) {
						card.style.transform = 'translateY(-8px) scale(1.02)';
					}
				});

				card.addEventListener('mouseleave', () => {
					if (!isAnimating) {
						card.style.transform = 'translateY(0) scale(1)';
					}
				});
			});

			// Smooth navigation
			document.querySelectorAll('.nav-link').forEach(link => {
				link.addEventListener('click', (e) => {
					document.body.style.opacity = '0.8';
					setTimeout(() => {
						document.body.style.opacity = '1';
					}, 100);
				});
			});
		});
	</script>
</body>
</html>`

		t, err := template.New("about").Parse(tmpl)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Description string
			ButtonText  string
			Features    []struct {
				Icon        string
				Title       string
				Description string
			}
		}{
			Description: "This is a modern Go web application built with native Go templates and enhanced with CSS animations and JavaScript features.",
			ButtonText:  "🎭 Animate Features!",
			Features: []struct {
				Icon        string
				Title       string
				Description string
			}{
				{"🚀", "Fast Performance", "Built with Go's efficient HTTP server for lightning-fast response times."},
				{"🎨", "Beautiful UI", "Modern dark theme with smooth animations and responsive design."},
				{"📱", "Mobile Friendly", "Fully responsive layout that works great on all devices."},
				{"⚡", "Interactive", "Dynamic JavaScript interactions with smooth animations."},
				{"🔒", "Single file", "Portable and easy to deploy."},
			},
		}

		if err := t.Execute(w, data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Starting server on :3000")
	http.ListenAndServe(":3000", nil)
}

// q
