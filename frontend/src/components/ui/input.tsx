import { forwardRef } from 'react'

const Input = forwardRef<
  HTMLInputElement,
  React.InputHTMLAttributes<HTMLInputElement>
>(({ className, ...props }, ref) => {
  return (
    <input
      ref={ref}
      type="text"
      className={`w-full p-2 rounded-md border bg-white/10 border-gray-300 ${className || ''}`}
      {...props}
    />
  )
})

Input.displayName = 'Input'

export default Input
