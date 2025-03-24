import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <div>
        <a href="https://haume.me/axo" target="_blank">
          <img src="/static/axo.webp" className="logo" alt="Vite logo" />
        </a>
      </div>
      <h1>🪐 Welcome to Axo ✨</h1>
      <p>
        AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and
        gorm. <br /> It is designed to be simple, fast, and easy to use. For
        more information, please click the logo above.
      </p>
      <p className="read-the-docs">
        Licanse: MIT <br />
        Copyright (c) 2025 Haume
      </p>
    </>
  );
}

export default App;
