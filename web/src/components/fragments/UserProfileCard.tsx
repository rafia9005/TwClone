"use client"

import { Avatar } from "@/components/elements/Avatar"
import { Text } from "@/components/elements/Text"
import { MoreHorizontal } from "lucide-react"
import { cn } from "@/lib/utils"

interface UserProfileCardProps {
  user: {
    name: string
    username: string
    avatar?: string
  }
  onClick?: () => void
  showMore?: boolean
  className?: string
}

export function UserProfileCard({ user, onClick, showMore = true, className }: UserProfileCardProps) {
  return (
    <button
      onClick={onClick}
      className={cn(
        "flex items-center gap-3 p-3 rounded-xl w-full hover:bg-muted/60 transition-colors",
        className
      )}
    >
      <Avatar src={user.avatar} alt={user.name} size="lg" />
      
      <div className="flex-1 min-w-0 text-left">
        <Text className="font-semibold leading-tight truncate">{user.name}</Text>
        <Text variant="muted" className="truncate">@{user.username}</Text>
      </div>

      {showMore && (
        <MoreHorizontal size={18} className="text-muted-foreground flex-shrink-0" />
      )}
    </button>
  )
}
