import Image from "next/image";
import Link from "next/link";

const Logo = () => {
	return <Link href="https://sagesure.com/" className="hover:scale-110 transition-all duration-300 backdrop-blur-sm">
		<Image src="/logo-removed-bg.png" alt="SageSure" width={100} height={100} />
	</Link>
}

export default Logo;