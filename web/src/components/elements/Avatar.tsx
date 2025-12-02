import { cn } from "@/lib/utils"

interface AvatarProps {
  src?: string
  alt?: string
  size?: "sm" | "md" | "lg" | "xl"
  className?: string
}

const sizeClasses = {
  sm: "w-8 h-8",
  md: "w-10 h-10",
  lg: "w-12 h-12",
  xl: "w-16 h-16",
}

export function Avatar({ src, alt = "User avatar", size = "md", className }: AvatarProps) {
  return (
    <div className={cn("rounded-full bg-muted overflow-hidden flex-shrink-0", sizeClasses[size], className)}>
      {src ? (
        <img src={src} alt={alt} className="w-full h-full object-cover" />
      ) : (
        <div className="w-full h-full flex items-center justify-center bg-gradient-to-br from-primary/20 to-primary/40">
          <span className="text-primary font-semibold">{alt.charAt(0).toUpperCase()}</span>
        </div>
      )}
    </div>
  )
}
