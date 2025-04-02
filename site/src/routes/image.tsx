import { createFileRoute, Link } from "@tanstack/react-router";
import Img from "../../AxoComponents/Img";
import { useEffect, useState } from "react";

export const Route = createFileRoute("/image")({
  component: RouteComponent,
});

function RouteComponent() {
  const [imgSpecs, setImgSpecs] = useState({
    src: "/api/axo.webp",
    quality: 80,
    width: 200,
    height: 200,
    format: "jpeg",
  });
  const [imageUrl, setImageUrl] = useState(
    "/api/image?src=/api/axo.webp&format=jpeg&quality=80&width=200&height=200"
  );

  useEffect(() => {
    const { quality, src, width, height, format } = imgSpecs;
    setImageUrl(
      `/api/image?src=${src}&format=${format}&quality=${quality}&width=${width}&height=${height}`
    );
  }, [imgSpecs]);

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
      <h4 className="mt-12">Parameters</h4>
      <p className="text-xs text-white/20">
        Warning, due to the active breakpoint system, the image may not change
        until a certain amount of height and width is adjusted.
      </p>
      <div className="grid w-full grid-cols-3 grid-flow-row gap-2">
        <label htmlFor="quality" className="flex flex-col items-start">
          <span className="text-sm text-white/30">Quality</span>
          <input
            type="number"
            name="quality"
            defaultValue={80}
            min={1}
            max={100}
            onChange={(e) =>
              setImgSpecs((prev) => ({
                ...prev,
                quality: Number(e.target.value),
              }))
            }
          />
        </label>
        <label htmlFor="width" className="flex flex-col items-start">
          <span className="text-sm text-white/30">Width</span>
          <input
            type="number"
            name="width"
            defaultValue={200}
            min={1}
            onChange={(e) =>
              setImgSpecs((prev) => ({
                ...prev,
                width: Number(e.target.value),
              }))
            }
          />
        </label>
        <label htmlFor="height" className="flex flex-col items-start">
          <span className="text-sm text-white/30">Height</span>
          <input
            type="number"
            name="height"
            defaultValue={200}
            min={1}
            onChange={(e) =>
              setImgSpecs((prev) => ({
                ...prev,
                height: Number(e.target.value),
              }))
            }
          />
        </label>
        <select
          name="format"
          id="format"
          onChange={(e) => {
            setImgSpecs((prev) => ({
              ...prev,
              format: e.target.value,
            }));
          }}
          className="block">
          <option value="jpeg">jpeg</option>
          <option value="jpg">jpg</option>
          <option value="webp">webp</option>
          <option value="png">png</option>
        </select>
        <label htmlFor="img" className="flex flex-col col-span-2 items-start">
          <input
            type="text"
            name="img"
            defaultValue={imgSpecs.src}
            onChange={(e) =>
              setImgSpecs((prev) => ({
                ...prev,
                src: e.target.value,
              }))
            }
          />
        </label>
      </div>
      <h4 className="mt-12">API Example</h4>
      <div className="inline-flex justify-center items-center gap-3 w-full">
        <input
          type="text"
          name="url"
          value={imageUrl}
          readOnly
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
        src="/api/axo.webp"
        alt=""
        className="rounded-2xl object-contain w-full h-[27rem] p-4 bg-comet-500"
        // width={200} //? Component can calculate the required width
        // height={200} //? Component can calculate the required height
        quality={100}
        loading="lazy"
        format={"webp"} //default format is jpeg but you can set it to webp, png, jpeg, jpg
      />
    </>
  );
}
