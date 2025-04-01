import { createFileRoute, Link } from "@tanstack/react-router";
import { useState, useEffect } from "react";

export const Route = createFileRoute("/")({
  component: App,
});

function App() {
  const [exampleUrl, setExampleUrl] = useState("/api/hello");
  return (
    <>
      <h1>🪸 Welcome to Axo 🌊</h1>
      <p className="">
        AxoScaffold is a Restful API scaffold for Go, built on top of stdlib and
        gorm. <br /> It is designed to be simple, fast, and easy to use. For
        more information, please click the logo above.
      </p>
      <h4 className="mt-6">Examples</h4>
      <div className="flex justify-center w-full items-center gap-3">
        <Link to="/image" className="w-full">
          <button className="w-full">Image Optimization</button>
        </Link>
        <Link to="/smtp" className="w-full">
          <button className="w-full">Mail</button>
        </Link>
        <Link to="/database" className="w-full">
          <button className="w-full">Database</button>
        </Link>
      </div>
      <h4 className="mt-6">Hello World Examples</h4>
      <span className="w-full flex flex-col gap-2 p-2 rounded-xl">
        <div className="flex justify-center w-full items-center  gap-3">
          <button
            className="w-full"
            onClick={() => {
              setExampleUrl("/api/hello");
            }}>
            Hello Example
          </button>
          <button
            className="w-full"
            onClick={() => {
              setExampleUrl("/api/error");
            }}>
            Error Example
          </button>
        </div>
        <div className="inline-flex justify-center items-center gap-3 w-full">
          <input
            type="text"
            name="url"
            value={exampleUrl}
            readOnly
            className="font-mono"
          />
          <a href={exampleUrl} target="_blank">
            <button className="size-12 !p-3 items-center justify-center flex">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 20 20">
                <path
                  fill="currentColor"
                  d="M6.25 4.5A1.75 1.75 0 0 0 4.5 6.25v7.5c0 .966.784 1.75 1.75 1.75h7.5a1.75 1.75 0 0 0 1.75-1.75v-2a.75.75 0 0 1 1.5 0v2A3.25 3.25 0 0 1 13.75 17h-7.5A3.25 3.25 0 0 1 3 13.75v-7.5A3.25 3.25 0 0 1 6.25 3h2a.75.75 0 0 1 0 1.5zm4.25-.75a.75.75 0 0 1 .75-.75h5a.75.75 0 0 1 .75.75v5a.75.75 0 0 1-1.5 0V5.56l-3.72 3.72a.75.75 0 1 1-1.06-1.06l3.72-3.72h-3.19a.75.75 0 0 1-.75-.75"
                />
              </svg>
            </button>
          </a>
        </div>
        <iframe
          src={exampleUrl}
          className="rounded-2xl w-full h-[27rem] p-4 bg-comet-500"
        />
      </span>
    </>
  );
}

export default App;
