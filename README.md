# Go Web Application

A modern, interactive web application built with Go's native HTTP server and enhanced with CSS animations and JavaScript interactions.

## 🚀 Features

- **Multi-page Navigation**: Seamless navigation between Home and About pages
- **Interactive Animations**: Dynamic color-changing buttons and spinning card animations
- **Retro-Futuristic Theme**: Cyberpunk-inspired design with neon green matrix aesthetics
- **Responsive Design**: Mobile-friendly layout that works on all devices
- **Zero Dependencies**: Built entirely with Go's standard library
- **Single Binary**: Portable and easy to deploy

## 🛠️ Tech Stack

- **Backend**: Go (Golang) with `net/http` and `html/template`
- **Frontend**: Vanilla JavaScript with CSS3 animations
- **Styling**: Custom retro CSS fontw (Press Start 2P & Orbitron)
- **Architecture**: Server-side rendered templates

## 📦 Installation

### Prerequisites

- Go 1.16+ installed on your system

### Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-q
   ```

2. **Run the application**
   ```bash
   go run main.go
   ```

3. **Open your browser**
   Navigate to `http://localhost:8080`

### Build for Production

```bash
go build -o webapp main.go
./webapp
```

## 🎯 Usage

The application serves two main pages:

- **Home (`/`)**: Welcome page with interactive color-changing animations
- **About (`/about`)**: Project information with animated feature cards

### Interactive Features

- **Color Changer**: Click the button on the home page to cycle through retro neon color schemes
- **Card Animations**: Click "Animate Features!" on the about page to see spinning card animations
- **Smooth Navigation**: Hover effects and glowing transition animations throughout the site

## 📁 Project Structure

```
go-q/
├── main.go          # Main application file
├── go.mod           # Go module definition
└── README.md        # This file
```

## 🎨 Design Philosophy

This project demonstrates:

- **Minimalist Architecture**: Everything in a single Go file for simplicity
- **Retro-Futuristic UI/UX**: Dark theme with neon green glowing effects
- **Performance First**: Lightweight with fast load times
- **Developer Friendly**: Clean, readable code with proper error handling

## 🔧 Configuration

The server runs on port `:3000` by default. You can modify the port in `main.go`:

```go
http.ListenAndServe(":3000", nil)
```

## 🎭 Animation Details

### Home Page
- **Retro Color Cycling**: 6 vibrant neon color schemes with smooth transitions
- **Scale Animation**: Heading scales up briefly when colors change
- **Glow Effects**: Dynamic text-shadow effects

### About Page
- **360° Card Spin**: Feature cards perform full rotations with staggered timing
- **3D Heading Flip**: Y-axis rotation animation for the main heading
- **Hover Effects**: Cards lift and glow with neon green effects on mouse hover

## 🚦 API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/` | GET | Home page with welcome message and interactive button |
| `/about` | GET | About page with project features and animations |

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is open source and available under the [MIT License](LICENSE).

## 🙋‍♂️ Support

If you have any questions or need help with the project:

- Open an issue on GitHub
- Check the code comments for implementation details
- Review the inline documentation

## 🔮 Future Enhancements

- [ ] Add more page templates
- [ ] Implement theme switching
- [ ] Add form handling capabilities
- [ ] Include API endpoints for data
- [ ] Add session management
- [ ] Implement logging middleware

---

**Built with ❤️ using Go and modern web technologies**