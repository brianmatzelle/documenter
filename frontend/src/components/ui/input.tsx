import { forwardRef, useState } from 'react'
import FocusedMsg from './focused-msg'

const Input = forwardRef<
  HTMLInputElement,
  React.InputHTMLAttributes<HTMLInputElement> & { label: string }
>(({ className, label, ...props }, ref) => {
  const [isFocused, setIsFocused] = useState(false)

  return (
    <div className="flex flex-col relative w-full">
      <FocusedMsg hidden={!isFocused} position="l">{label}</FocusedMsg>
      <input
        ref={ref}
        type="text"
        onFocus={() => setIsFocused(true)}
        onBlur={() => setIsFocused(false)}
        className={`w-full p-2 rounded-md border bg-white/10 border-gray-300 focus:outline-none ${className || ''}`}
        {...props}
      />
    </div>
  )
})

Input.displayName = 'Input'

export default Input
