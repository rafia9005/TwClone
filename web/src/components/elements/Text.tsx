import { cn } from "@/lib/utils"

interface TextProps {
  children: React.ReactNode
  variant?: "body" | "small" | "caption" | "muted"
  className?: string
  as?: "p" | "span" | "div"
}

const variantClasses = {
  body: "text-base leading-normal",
  small: "text-sm",
  caption: "text-xs",
  muted: "text-sm text-muted-foreground",
}

export function Text({ children, variant = "body", className, as: Component = "p" }: TextProps) {
  return <Component className={cn(variantClasses[variant], className)}>{children}</Component>
}
