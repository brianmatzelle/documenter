import { forwardRef } from 'react'

const Button = forwardRef<
	HTMLButtonElement,
	React.ButtonHTMLAttributes<HTMLButtonElement> & { children: React.ReactNode, className?: string }
>(({ children, className, ...props }, ref) => {
	return (
		<button
			ref={ref}
			className={`w-full p-2 rounded-md border bg-white/10 border-gray-300 hover:bg-white/20 transition-colors ${className || ''}`}
			{...props}
		>
			{children}
		</button>
	)
})

Button.displayName = 'Button'

export default Button
