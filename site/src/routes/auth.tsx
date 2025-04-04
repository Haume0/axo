import { createFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/auth")({
  component: RouteComponent,
});

function RouteComponent() {
  const [screen, setScreen] = useState("");
  const navigate = useNavigate();
  const [user, setUser] = useState<any>(undefined);
  return (
    <>
      <h1 className="text-3xl font-bold">🪸 Axo Auth & Role System 🌊</h1>
      <p className="">
        A sample page for you to test the Auth system.
        <br />
        On this page, you can perform user registration and login operations.
      </p>
      <Link to="/">
        <button>← Back to Home</button>
      </Link>
      <span className="flex flex-col gap-1 w-full">
        <p className="text-xs">Server response message:</p>
        <p className="bg-comet-700 p-2 min-h-9 rounded-lg w-full font-mono">
          {screen}
        </p>
      </span>
      {user != undefined && (
        <div className="rounded-xl bg-comet-700 p-2 w-full flex flex-col gap-1">
          <h3 className="font-bold text-xl text-comet-200 bg-comet-600 rounded-md">
            Current User
          </h3>
          <div className="grid grid-cols-2 gap-1">
            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Email:
            </span>
            <span className="text-comet-100 bg-comet-600 p-2 rounded-md w-full">
              {user.email}
            </span>

            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              User ID:
            </span>
            <span className="text-comet-100 bg-comet-600 p-2 rounded-md w-full">
              {user.id}
            </span>

            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Role:
            </span>
            <span className="text-comet-100 bg-comet-600 p-2 rounded-md w-full">
              {user.role.name || "Not assigned"}
            </span>

            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Status:
            </span>
            <div className="flex items-center justify-center gap-2 bg-comet-600 p-2 rounded-md w-full">
              <span
                className={`h-2 w-2 rounded-full ${user.active ? "bg-green-500" : "bg-red-500"}`}></span>
              <span className="text-comet-100">
                {user.active ? "Active" : "Inactive"}
              </span>
            </div>

            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Verified:
            </span>
            <div className="flex items-center justify-center gap-2 bg-comet-600 p-2 rounded-md w-full">
              <span
                className={`h-2 w-2 rounded-full ${user.verified ? "bg-green-500" : "bg-yellow-500"}`}></span>
              <span className="text-comet-100">
                {user.verified ? "Verified" : "Not Verified"}
              </span>
            </div>

            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Joined:
            </span>
            <span className="text-comet-100 bg-comet-600 p-2 rounded-md w-full">
              {new Date(user.created_at).toLocaleDateString()}
            </span>
            <span className="font-semibold text-comet-300 bg-comet-600 p-2 rounded-md w-full">
              Edited:
            </span>
            <span className="text-comet-100 bg-comet-600 p-2 rounded-md w-full">
              {new Date(user.updated_at).toLocaleDateString()}
            </span>
          </div>
          <button
            className="bg-comet-600 hover:bg-comet-500 p-2 w-full"
            onClick={() => setUser(undefined)}>
            Logout
          </button>
        </div>
      )}
      <section className="p-2 w-full rounded-xl flex flex-col gap-2 bg-comet-700">
        <h4 className="w-full bg-comet-500 rounded-lg py-2">
          Log in to an account!
        </h4>
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const formData = new FormData(e.currentTarget);
            const res = await fetch("/api/auth/login", {
              method: "POST",
              body: formData,
            });
            if (res.ok) {
              const usr = await res.json();
              setScreen("Successfully logged in!");
              setUser(usr);
            } else {
              const data = await res.json();
              setScreen(data.error);
            }
          }}
          className="grid w-full grid-cols-2 grid-flow-row gap-2">
          <input
            type="email"
            name="email"
            className="col-span-2"
            placeholder="Enter your email"
          />
          <input
            type="password"
            name="password"
            placeholder="Choose a password"
            className="col-span-2"
          />
          <button className="col-span-2">Log in</button>
        </form>
      </section>
      <section className="p-2 w-full rounded-xl flex flex-col gap-2 bg-comet-700">
        <h4 className="w-full bg-comet-500 rounded-lg py-2">
          Create an account!
        </h4>
        <form
          onSubmit={async (e) => {
            e.preventDefault();
            const formData = new FormData(e.currentTarget);
            if (formData.get("password") !== formData.get("repassword")) {
              setScreen("Passwords do not match!");
              return;
            }
            const res = await fetch("/api/auth/register", {
              method: "POST",
              body: formData,
            });
            if (res.ok) {
              const usr = await res.json();
              setScreen("User created successfully!");
              setUser(usr);
            } else {
              const data = await res.json();
              setScreen(data.error);
            }
          }}
          className="grid w-full grid-cols-2 grid-flow-row gap-2">
          <input
            type="email"
            name="email"
            className="col-span-2"
            placeholder="Enter your email"
          />
          <input
            type="password"
            name="password"
            placeholder="Choose a password"
          />
          <input
            type="password"
            name="repassword"
            placeholder="Enter password again"
          />
          <button className="col-span-2">Create</button>
        </form>
      </section>
    </>
  );
}
