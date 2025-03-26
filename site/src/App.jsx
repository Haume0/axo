import { useState } from "react";

function App() {
  const [exampleUrl, setExampleUrl] = useState("/api/hello");
  return (
    <main className="w-full relative flex flex-col items-center gap-4 justify-center p-8 text-center">
      <div>
        <a href="https://haume.me/axo" target="_blank">
          <img
            src="/api/static/axo.webp"
            className="h-32 transition-all hover:drop-shadow-[0_0_2em_#646cffaa]"
            alt="Vite logo"
          />
        </a>
      </div>
      <h1 className="text-3xl font-bold">🪸 Welcome to Axo 🌊</h1>
      <p className="">
        AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and
        gorm. <br /> It is designed to be simple, fast, and easy to use. For
        more information, please click the logo above.
      </p>
      <p className="text-white/40 text-sm">
        License: MIT <br />
        Copyright (c) 2025 Haume
      </p>
      <div className="flex justify-center items-center  gap-3">
        <button
          onClick={() => {
            setExampleUrl("/api/hello");
          }}>
          Hello Example
        </button>
        <button
          onClick={() => {
            setExampleUrl("/api/error");
          }}>
          Error Example
        </button>
        <button
          onClick={() => {
            setExampleUrl(
              "/api/image?src=axo.webp&format=jpeg&quality=80&width=200&height=200"
            );
          }}>
          Image Optimization
        </button>
      </div>
      <div className="inline-flex justify-center items-center gap-3 w-[27rem]">
        <input
          type="text"
          name="url"
          value={exampleUrl}
          readOnly
          className="w-full h-11 rounded-lg px-4 text-sm font-mono bg-black/40 outline-none border-none"
        />
        <a href={exampleUrl} target="_blank">
          <button>Open</button>
        </a>
      </div>
      <iframe
        src={exampleUrl}
        className="rounded-2xl w-[27rem] h-[27rem] p-4 bg-black/40"
        frameBorder="0"></iframe>
    </main>
  );
}

export default App;
