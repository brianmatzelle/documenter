import { forwardRef } from 'react'

const SelectItem = forwardRef<
  HTMLOptionElement,
  React.OptionHTMLAttributes<HTMLOptionElement>
>(({ className, ...props }, ref) => {
  return (
    <option
      ref={ref}
      className={`bg-white/10 ${className || ''}`}
      {...props}
    />
  )
})

SelectItem.displayName = 'SelectItem'

export default SelectItem