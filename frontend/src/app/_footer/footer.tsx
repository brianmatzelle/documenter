import Logo from "@/components/ui/logo";
import Link from "next/link";
import Image from "next/image";
export default function Footer() {
	return (
		<div className="relative flex items-center w-full px-4 pb-4">
			<div className="absolute right-4">
				<Link href="https://github.com/brianmatzelle/documenter" className="flex items-center gap-2">
					<p className={`hover:text-[hsl(169,100%,20%)] transition-all duration-150 hidden md:block`}>SageSure Documenter</p>
					<Image src="/github-mark-white.svg" alt="GitHub" width={20} height={20} />
				</Link>
			</div>
			<div className="w-full flex justify-center">
				<Logo />
			</div>
		</div>
	)
}
