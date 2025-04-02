import { Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <>
      <main className="min-w-[27rem] max-w-[38rem] shrink-0 mx-auto relative flex flex-col items-center gap-4 justify-center p-8 text-center">
        <div>
          <a href="https://haume.me/axo" target="_blank">
            <img
              src="/api/axo.webp"
              alt="Vite logo"
              className="h-32 transition-all hover:drop-shadow-[0_0_2em_#646cffaa]"
            />
          </a>
        </div>
        <Outlet />
        <p className="text-white/30 text-sm">
          License: MIT <br />
          Copyright (c) 2025 Haume
        </p>
      </main>
    </>
  ),
});
