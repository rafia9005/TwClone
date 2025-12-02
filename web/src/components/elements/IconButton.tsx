import type { LucideIcon } from "lucide-react"
import { cn } from "@/lib/utils"

interface IconButtonProps {
  icon: LucideIcon
  label?: string
  active?: boolean
  count?: number
  onClick?: () => void
  className?: string
  size?: "sm" | "md" | "lg"
}

const sizeClasses = {
  sm: "p-1.5",
  md: "p-2",
  lg: "p-3",
}

const iconSizes = {
  sm: 16,
  md: 18,
  lg: 20,
}

export function IconButton({
  icon: Icon,
  label,
  active = false,
  count,
  onClick,
  className,
  size = "md",
}: IconButtonProps) {
  return (
    <button
      onClick={onClick}
      aria-label={label}
      className={cn(
        "group flex items-center gap-2 rounded-full transition-all hover:bg-primary/10",
        sizeClasses[size],
        active && "text-primary",
        className
      )}
    >
      <Icon
        size={iconSizes[size]}
        className={cn("transition-colors", active && "fill-current")}
      />
      {count !== undefined && (
        <span className={cn("text-sm transition-colors", active && "text-primary")}>
          {count > 999 ? `${(count / 1000).toFixed(1)}K` : count}
        </span>
      )}
    </button>
  )
}
