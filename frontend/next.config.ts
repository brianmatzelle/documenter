import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: false,
  async rewrites() {
    return [
      {
        source: "/:path*",
        destination: "http://api:8080/:path*",
      },
    ];
  },
};

export default nextConfig;