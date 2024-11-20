import TitleCard from "@/components/titleCard";
import Input from "@/components/ui/input";
import Card from "@/components/ui/card";
import Button from "@/components/ui/button";
export default function Home() {
	return (
		<div className="flex flex-col items-center justify-around h-full p-4">
			<Card className="w-full max-w-xs sm:max-w-md md:max-w-lg lg:max-w-xl">
				<TitleCard />
				<div className="w-full flex flex-col sm:flex-row gap-2">
					<Input 
						className="flex-[4]" 
						placeholder="Enter your GitLab MR link here" 
					/>
					<Button className="w-full sm:w-auto sm:flex-1">Generate</Button>
				</div>
			</Card>

			<div></div>

		</div>
	)
}

