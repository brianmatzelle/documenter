import type { Metadata } from "next";
import "./globals.css";
import StarfieldWrapper from "@/components/starfieldWrapper";
import Footer from "@/app/_footer/footer";
import { Inconsolata } from 'next/font/google'
import { DocProvider } from "@/context/doc-context";

const font = Inconsolata({
  subsets: ['latin'],
  weight: ['400', '500', '700'],
})

export const metadata: Metadata = {
  title: "GitLab Documenter",
  description: "Documenter for GitLab merge requests, by brian matzelle",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${font.className} flex flex-col min-h-screen w-full`}>
        <main className="flex-grow flex flex-col">
          <DocProvider>
            <StarfieldWrapper />
            {children}
            <Footer />
          </DocProvider>
        </main>
      </body>
    </html>
  );
}
