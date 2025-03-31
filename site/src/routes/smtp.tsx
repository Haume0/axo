import { createFileRoute, Link } from "@tanstack/react-router";
import { useState } from "react";

export const Route = createFileRoute("/smtp")({
  component: RouteComponent,
});

function RouteComponent() {
  const [mloading, _mloading] = useState(false);
  return (
    <>
      {" "}
      <div>
        <a href="https://haume.me/axo" target="_blank">
          <img
            src="/api/static/axo.webp"
            alt="Vite logo"
            className="h-32 transition-all hover:drop-shadow-[0_0_2em_#646cffaa]"
          />
        </a>
      </div>
      <h1 className="text-3xl font-bold">🪸 Axo Mail System 🌊</h1>
      <p className="">
        To send mails, you need to fill .env or adding required env variables in
        your server. You can see them in ./env.go file.
        <br />
        How to use mail system, you can use the utility function in the
        /mail/seng.go to send mails. More detail included into function
        document.
      </p>
      <Link to="/">
        <button>← Back to Home</button>
      </Link>
      <p className="dark:text-white/40 text-black/20 text-sm">
        License: MIT <br />
        Copyright (c) 2025 Haume
      </p>
      <h4 className="mt-12">Test Mail!</h4>
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
        className={`flex gap-2 w-full ease-in-out duration-700 ${
          mloading && "scale-95 blur-[2px] opacity-60"
        }`}>
        <input
          type="email"
          name="email"
          placeholder="Enter a mail adress!"
          id="email"
        />
        <button>Send</button>
      </form>
    </>
  );
}
