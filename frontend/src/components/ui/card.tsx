import { forwardRef } from 'react'

interface CardProps extends React.HTMLAttributes<HTMLDivElement> {
	children: React.ReactNode
}

const Card = forwardRef<HTMLDivElement, CardProps>(
	({ children, className, ...props }, ref) => {
		return (
			<div 
				ref={ref}
				className={`w-full backdrop-blur-sm p-4 rounded-md border bg-white/10 border-gray-300 flex flex-col items-center justify-center max-w-xl gap-10 ${className ?? ''}`}
				{...props}
			>
				{children}
			</div>
		)
	}
)

Card.displayName = 'Card'

export default Card
