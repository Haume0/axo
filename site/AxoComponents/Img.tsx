import { useState, useEffect, useRef } from "react";

interface ImgProps {
  src: string;
  alt: string;
  width?: number;
  height?: number;
  quality?: number;
  loading?: "lazy" | "eager";
  format?: "webp" | "png" | "jpeg" | "jpg";
  className?: string;
}

export default function Img(props: ImgProps) {
  const {
    src,
    alt,
    width,
    height,
    quality,
    loading = "lazy",
    format,
    className,
  } = props;

  const [myImg, setMyImg] = useState<string>("");
  const unique = useRef(Math.random().toString(36).substring(7)).current;
  const imgRef = useRef<HTMLImageElement>(null);

  useEffect(() => {
    const imgElement = document.getElementById(`${unique}_image`);
    if (!imgElement) return;

    let imgWidth = width || imgElement.clientWidth;

    const queryParams = new URLSearchParams({
      src,
      ...(imgWidth && { width: imgWidth.toString() }),
      ...(height && { height: height.toString() }),
      ...(quality && { quality: quality.toString() }),
      ...(format && { format }),
    }).toString();

    const imgUrl = `/api/image?${queryParams}`;
    setMyImg(imgUrl);
  }, [src, width, height, quality, format, unique]);

  return (
    <img
      id={`${unique}_image`}
      src={myImg}
      alt={alt}
      className={className}
      loading={loading}
      ref={imgRef}
    />
  );
}
