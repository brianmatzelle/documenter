import { forwardRef } from 'react'

type Position = 't' | 'b' | 'l' | 'r'

interface FocusedMsgProps {
  hidden: boolean
  children: React.ReactNode
  position?: Position
  arrow?: boolean
  offset?: number
}

const arrowLabels: Record<Position, string> = {
  // these point towards the element
  't': '↓',
  'b': '↑',
  'l': '→',
  'r': '←',
}

const FocusedMsg = forwardRef<HTMLParagraphElement, FocusedMsgProps>(
  ({ hidden , children, position = 'l', arrow = true }, ref) => {
    if (hidden === true) return null

    // const positionClasses: Record<Position, string> = {
    //   't': `left-1/2 top-0 -translate-x-1/2 -translate-y-[calc(100%+${offset}px)]`,
    //   'b': `left-1/2 bottom-0 -translate-x-1/2 translate-y-[calc(100%+${offset}px)]`,
    //   'l': `left-0 top-1/2 -translate-y-1/2 -translate-x-[calc(100%+${offset}px)]`,
    //   'r': `right-0 top-1/2 -translate-y-1/2 translate-x-[calc(100%+${offset}px)]`,
    // }
    const positionClasses: Record<Position, string> = {
      't': `left-1/2 top-0 -translate-x-1/2 -translate-y-[calc(100%+24px)]`,
      'b': `left-1/2 bottom-0 -translate-x-1/2 translate-y-[calc(100%+24px)]`,
      'l': `left-0 top-1/2 -translate-y-1/2 -translate-x-[calc(100%+24px)]`,
      'r': `right-0 top-1/2 -translate-y-1/2 translate-x-[calc(100%+24px)]`,
    }

    return (
      <p
        ref={ref}
        className={`text-sm text-gray-500 absolute ${positionClasses[position]}`}
      >
        {children} {arrow && arrowLabels[position]}
      </p>
    )
  }
)

FocusedMsg.displayName = 'FocusedMsg'

export default FocusedMsg
