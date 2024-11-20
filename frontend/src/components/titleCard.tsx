import { colors } from "./ui/colors";

export default function TitleCard() {
	return (
		<div className="flex flex-col items-center justify-center h-full gap-2 sm:gap-4">
			<h1 className="text-2xl sm:text-3xl md:text-4xl font-bold">
				<span className={`text-[${colors.primary}]`}>SageSure</span> Documenter
			</h1>
      <div className="flex flex-col items-center justify-center gap-1">
        <p className="text-base sm:text-lg text-center px-4">
          AI documentation generator for SageSure merge requests,
        </p>
        <p className="text-xs sm:text-sm">
          by brian matzelle
        </p>
      </div>
		</div>
	)
}
