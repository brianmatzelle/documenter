export default function Card({ children }: { children: React.ReactNode }) {
	return <div className="w-full p-4 rounded-md border bg-white/10 border-gray-300 flex flex-col items-center justify-center max-w-xl gap-10">
		{children}
	</div>
}
