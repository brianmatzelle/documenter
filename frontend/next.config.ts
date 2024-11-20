import { BACKEND_URL } from "@/lib/config";
import type { NextConfig } from "next";

const nextConfig: NextConfig = {
  reactStrictMode: false,
  async rewrites() {
    return [
      {
        source: "/:path*",
        destination: `${BACKEND_URL}/:path*`,
      },
    ];
  },
};

export default nextConfig;
