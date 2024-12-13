import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: false,
  async rewrites() {
    return [
      {
        source: "/:path*",
        destination: process.env.NODE_ENV === "production" 
          ? "http://api:8080/:path*"  // Docker internal network
          : "http://localhost:8080/:path*",  // Local development
      },
    ];
  },
};

export default nextConfig;