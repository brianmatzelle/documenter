export default function TitleCard() {
	return (
		<div className="flex flex-col items-center justify-center h-full gap-4">
			<h1 className="text-4xl font-bold">
				<span className="text-[hsl(169,100%,20%)]">SageSure</span> Documenter
				{/* SageSure Documenter */}
			</h1>
			<p className="text-lg">
				AI documenter for SageSure merge requests, by brian matzelle
			</p>
		</div>
	)
}
