import Link from "next/link";
import Image from "next/image";
import { splitTextIntoColorChunks, gitlabColors } from "@/lib/colors";

export default function Footer() {
	return (
		<div className="relative flex items-center w-full px-4 pb-4">
			<div className="absolute right-4">
				<Link href="https://github.com/brianmatzelle/documenter" className="flex items-center gap-2">
					<p className={`hover:underline transition-all duration-150 hidden md:block`}>
						{splitTextIntoColorChunks("GitLab", Object.values(gitlabColors)).map(({ text, color }, index) => (
							<span key={index} style={{ color }}>{text}</span>
						))}
						{" "}
						<span className="">Documenter</span>
					</p>
					<Image src="/github-mark-white.svg" alt="GitHub" width={20} height={20} />
				</Link>
			</div>
		</div>
	)
}
