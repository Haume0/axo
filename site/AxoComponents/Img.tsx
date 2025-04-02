import { useState, useEffect, useRef } from "react";

interface ImgProps {
  src: string;
  alt: string;
  width?: number;
  height?: number;
  quality?: number;
  loading?: "lazy" | "eager";
  format?: "webp" | "png" | "jpeg" | "jpg" | string;
  className?: string;
}

export default function Img({
  src,
  alt,
  width,
  height,
  quality,
  loading,
  format,
  className,
}: ImgProps) {
  const [myImg, setMyImg] = useState<string>("");
  const imgRef = useRef<HTMLImageElement | null>(null);

  useEffect(() => {
    if (!imgRef.current) return;

    const imgWidth = width || imgRef.current.clientWidth;

    const queryParams = new URLSearchParams({
      src,
      ...(imgWidth && { width: imgWidth.toString() }),
      ...(height && { height: height.toString() }),
      ...(quality && { quality: quality.toString() }),
      ...(format && { format }),
    }).toString();

    setMyImg(`/api/image?${queryParams}`);
  }, [src, width, height, quality, format]);

  return (
    <img
      ref={imgRef}
      src={myImg}
      alt={alt}
      className={className}
      loading={loading}
    />
  );
}
