import { forwardRef } from 'react'

const Select = forwardRef<
  HTMLSelectElement,
  React.SelectHTMLAttributes<HTMLSelectElement>
>(({ className, children, ...props }, ref) => {
  return (
    <select
      ref={ref}
      className={`w-full p-2 rounded-md border bg-white/10 border-gray-300 ${className || ''}`}
      {...props}
    >
      {children}
    </select>
  )
})

Select.displayName = 'Select'

export default Select