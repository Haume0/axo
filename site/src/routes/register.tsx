import { createFileRoute, useNavigate } from "@tanstack/react-router";

export const Route = createFileRoute("/register")({
  component: RouteComponent,
});

function RouteComponent() {
  const navigate = useNavigate();
  return (
    <>
      <h1>Welcome to Axo</h1>
      <h4>Create an account!</h4>
      <form
        onSubmit={async (e) => {
          e.preventDefault();
          const formData = new FormData(e.currentTarget);

          const res = await fetch("/api/auth/register", {
            method: "POST",
            body: formData,
          });
          if (res.ok) {
            navigate({
              to: "/",
            });
          } else {
            const data = await res.json();
            alert(data.error);
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
    </>
  );
}
