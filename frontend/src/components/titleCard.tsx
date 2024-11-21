import { splitTextIntoColorChunks } from "@/lib/colors";

export default function TitleCard({ companyName, companyColors }: { 
  companyName: string, 
  companyColors: string[]
}) {
	return (
		<div className="flex flex-col items-center justify-center h-full gap-2 sm:gap-4">
			<h1 className="text-2xl sm:text-3xl md:text-4xl font-bold">
				{splitTextIntoColorChunks(companyName, companyColors).map(({ text, color }, index) => (
					<span key={index} style={{ color }}>{text}</span>
				))} Documenter
			</h1>
			<div className="flex flex-col items-center justify-center gap-1">
				<p className="text-base sm:text-lg text-center px-4">
					AI documentation generator for GitLab merge requests,
				</p>
				<p className="text-xs sm:text-sm text-zinc-400">
					by brian matzelle
				</p>
			</div>
		</div>
	)
}
