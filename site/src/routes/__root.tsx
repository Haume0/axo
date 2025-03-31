import { Outlet, createRootRoute } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <>
      <main className="min-w-[27rem] max-w-[38rem] shrink-0 mx-auto relative flex flex-col items-center gap-4 justify-center p-8 text-center">
        <Outlet />
      </main>
    </>
  ),
});
