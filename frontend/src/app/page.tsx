import TitleCard from "@/components/titleCard";
import Input from "@/components/ui/input";
import Card from "@/components/ui/card";
import Button from "@/components/ui/button";
export default function Home() {
	return (
		<div className="flex flex-col items-center justify-around h-full min-h-screen">
				<Card>
				<TitleCard />
				<div className="w-full flex gap-2">
					<Input className="flex-[4]" placeholder="Enter your GitLab MR link here" />
					<Button className="flex-1">Generate</Button>
				</div>
			</Card>

			<div></div>

		</div>
	)
}

