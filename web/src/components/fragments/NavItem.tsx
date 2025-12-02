"use client"

import type { LucideIcon } from "lucide-react"
import { cn } from "@/lib/utils"

interface NavItemProps {
  icon: LucideIcon
  label: string
  href: string
  active?: boolean
  badge?: number
}

export function NavItem({ icon: Icon, label, href, active = false, badge }: NavItemProps) {
  return (
    <a
      href={href}
      className={cn(
        "flex items-center gap-4 px-4 py-3 rounded-full font-semibold text-lg transition-all hover:bg-primary/10 w-fit",
        active && "font-bold"
      )}
    >
      <div className="relative">
        <Icon size={26} className={cn("transition-colors", active && "text-primary")} />
        {badge !== undefined && badge > 0 && (
          <span className="absolute -top-1 -right-1 bg-primary text-primary-foreground text-xs font-bold rounded-full min-w-[18px] h-[18px] flex items-center justify-center px-1">
            {badge > 9 ? "9+" : badge}
          </span>
        )}
      </div>
      <span>{label}</span>
    </a>
  )
}
