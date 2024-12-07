import { forwardRef } from 'react'

const Button = forwardRef<
	HTMLButtonElement,
	React.ButtonHTMLAttributes<HTMLButtonElement> & { children: React.ReactNode, className?: string }
>(({ children, className, ...props }, ref) => {
	return (
		<button
			ref={ref}
			className={`rounded-md border bg-white/15 border-gray-300 hover:bg-white/25 transition-colors ${className || ''}`}
			{...props}
		>
			{children}
		</button>
	)
})

Button.displayName = 'Button'

export default Button
