import { createFileRoute, Link } from "@tanstack/react-router";
import Img from "/AxoComponents/Img";
import { useState } from "react";

export const Route = createFileRoute("/image")({
  component: RouteComponent,
});

function RouteComponent() {
  const [imageUrl, setImageUrl] = useState(
    "/api/image?src=axo.webp&format=jpeg&quality=80&width=200&height=200"
  );
  return (
    <>
      <h1 className="text-3xl font-bold">🪸 Axo Image Optimization 🌊</h1>
      <p className="">
        This optional feature helps you optimize images at runtime before
        sending them to users. It has two sides: <br />
        <code>api</code>: a side that processes and optimizes images and <br />
        <code>component</code>: a side that calculates the required dimensions
        of images and sends them to the api.
      </p>
      <Link to="/">
        <button>← Back to Home</button>
      </Link>
      <h4 className="mt-12">API Example</h4>
      <div className="inline-flex justify-center items-center gap-3 w-full">
        <input
          type="text"
          name="url"
          value={imageUrl}
          onChange={(e) => setImageUrl(e.target.value)}
          className="font-mono"
        />
        <a href={imageUrl} target="_blank">
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
      <img
        src={imageUrl}
        className="rounded-2xl object-contain w-full h-[27rem] p-4 bg-comet-500"
        alt=""
      />
      <h4 className="mt-12">Component Example</h4>
      <p>
        This component takes the dimensions of the image client-side if not
        provided and communicates them to the server, allowing the image to be
        created in the required dimensions.
      </p>
      <p className="text-white/30 text-sm">
        To see the URL of the image created by the component, right-click and
        select "open image in new tab".
      </p>
      <Img
        src="/api/static/axo.webp"
        alt="Vite logo"
        className="rounded-2xl object-contain w-full h-[27rem] p-4 bg-comet-500"
        // width={200} * Component can calculate the required width
        // height={200} * Component can calculate the required height
        quality={80}
        loading="lazy"
        format="jpeg" //default format is jpeg but you can set it to webp, png, jpeg, jpg
      />
    </>
  );
}
