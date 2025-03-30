import { useState } from "react";

function App() {
  const [exampleUrl, setExampleUrl] = useState("/api/hello");
  const [mloading, _mloading] = useState(false);
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
      <p className="dark:text-white/40 text-black/20 text-sm">
        License: MIT <br />
        Copyright (c) 2025 Haume
      </p>
      <div className="flex justify-center items-center  gap-3">
        <button
          onClick={() => {
            setExampleUrl("/api/hello");
          }}
        >
          Hello Example
        </button>
        <button
          onClick={() => {
            setExampleUrl("/api/error");
          }}
        >
          Error Example
        </button>
        <button
          onClick={() => {
            setExampleUrl(
              "/api/image?src=axo.webp&format=jpeg&quality=80&width=200&height=200",
            );
          }}
        >
          Image Optimization
        </button>
      </div>
      <div className="inline-flex justify-center items-center gap-3 w-[27rem]">
        <input
          type="text"
          name="url"
          value={exampleUrl}
          readOnly
          className="w-full h-12 rounded-lg px-4 text-sm font-mono bg-black/10 dark:bg-black/20 outline-none border-none"
        />
        <a href={exampleUrl} target="_blank">
          <button>Open</button>
        </a>
      </div>
      <iframe
        src={exampleUrl}
        className="rounded-2xl w-[27rem] h-[27rem] p-4 bg-black/10 dark:bg-black/20"
        frameBorder="0"
      ></iframe>
      <h4 className="text-2xl font-bold">Test Mail!</h4>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          _mloading(true);
          let formdata = new FormData(e.currentTarget);
          let mail = formdata.get("email");
          const res = await fetch(`/api/testmail?mail=${mail}`);
          if (res.ok) {
            const a = await res.json();
            alert(a.message);
          } else {
            const a = await res.json();
            alert(a.error);
          }
          _mloading(false);
        }}
        className={`flex gap-2 w-[27rem] ease-in-out duration-700 ${
          mloading && "scale-95 blur-[2px] opacity-60"
        }`}
      >
        <input
          className="w-full h-12 rounded-lg px-4 text-sm bg-black/10 dark:bg-black/20 outline-none border-none"
          type="email"
          name="email"
          placeholder="Enter a mail adress!"
          id="email"
        />
        <button>Send</button>
      </form>
    </main>
  );
}

export default App;
