import { forwardRef, useState } from 'react'
import Button from './button'
import FocusedMsg from './focused-msg'

interface ListInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  mrLinks: string[]
  onMrLinksChange: (links: string[]) => void
  label?: string
}

const ListInput = forwardRef<HTMLInputElement, ListInputProps>(
  ({ className, mrLinks, onMrLinksChange, label, ...props }, ref) => {
    const [focusedIndex, setFocusedIndex] = useState<number | null>(null)

    const handleAddInput = () => {
      onMrLinksChange([...mrLinks, ''])
    }

    const handleInputChange = (index: number, value: string) => {
      const newLinks = [...mrLinks]
      newLinks[index] = value
      onMrLinksChange(newLinks)
    }

    const handleDeleteInput = (indexToDelete: number) => {
      const newLinks = mrLinks.filter((_, index) => index !== indexToDelete)
      onMrLinksChange(newLinks)
    }

    return (
      <div className="flex flex-col gap-2">
        {mrLinks.map((link, index) => (
          <div key={index} className="flex flex-col relative">
            <FocusedMsg hidden={focusedIndex !== index} position="l">{label}</FocusedMsg>
            <div className="flex">
              <input
                ref={index === mrLinks.length - 1 ? ref : undefined}
                type="text"
                value={link}
                onChange={(e) => handleInputChange(index, e.target.value)}
                onFocus={() => setFocusedIndex(index)}
                onBlur={() => setFocusedIndex(null)}
                className={`w-full p-2 rounded-l-md bg-white/10 border border-r-0 border-gray-300 focus:outline-none ${
                  className || ''
                }`}
                {...props}
              />
              {index === 0 ? (
                <Button
                  onClick={handleAddInput}
                  className="px-4 py-2 rounded-none rounded-r-md border border-l-white/20"
                >
                  +
                </Button>
              ) : (
                <Button
                  onClick={() => handleDeleteInput(index)}
                  className="px-4 py-2 rounded-none rounded-r-md border border-l-white/20"
                >
                  -
                </Button>
              )}
            </div>
          </div>
        ))}
      </div>
    )
  }
)

ListInput.displayName = 'ListInput'

export default ListInput
