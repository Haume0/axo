import "./App.css";

function App() {
  return (
    <>
      <div>
        <a href="https://haume.me/axo" target="_blank">
          <img src="/api/static/axo.webp" className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>🪸 Welcome to Axo 🌊</h1>
      <p>
        AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and
        gorm. <br /> It is designed to be simple, fast, and easy to use. For
        more information, please click the logo above.
      </p>
      <span
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          marginTop: "1.2rem",
          gap: "0.725rem",
        }}>
        <a href="/api/hello">
          <button>Hello Example</button>
        </a>
        <a href="/api/error">
          <button>Error Example</button>
        </a>
        <a href="/api/image?src=/api/static/axo.webp&format=webp&quality=75&width=480&height=480">
          <button>Image Optimization</button>
        </a>
      </span>
      <p className="read-the-docs">
        License: MIT <br />
        Copyright (c) 2025 Haume
      </p>
    </>
  );
}

export default App;
