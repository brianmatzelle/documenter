import type { Metadata } from "next";
import "./globals.css";
import StarfieldWrapper from "@/components/starfieldWrapper";
import Footer from "@/app/_footer/footer";
import { Inconsolata } from 'next/font/google'

const font = Inconsolata({
  subsets: ['latin'],
  weight: ['400', '500', '700'],
})

export const metadata: Metadata = {
  title: "SageSure Documenter",
  description: "Documenter for SageSure merge requests, by brian matzelle",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${font.className} flex flex-col h-screen w-full`}>
        <StarfieldWrapper />
        {children}
        <Footer />
      </body>
    </html>
  );
}
