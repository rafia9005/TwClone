"use client"

import { Text } from "@/components/elements/Text"
import { TrendingUp } from "lucide-react"
import { cn } from "@/lib/utils"

interface TrendingTopicProps {
  category: string
  topic: string
  posts: string
  onClick?: () => void
  className?: string
}

export function TrendingTopic({ category, topic, posts, onClick, className }: TrendingTopicProps) {
  return (
    <button
      onClick={onClick}
      className={cn(
        "w-full p-3 rounded-xl hover:bg-muted/60 transition-colors text-left group",
        className
      )}
    >
      <div className="flex items-start justify-between gap-2">
        <div className="flex-1 min-w-0">
          <Text variant="muted" className="text-xs mb-0.5">
            {category}
          </Text>
          <Text className="font-bold leading-tight truncate group-hover:underline">
            {topic}
          </Text>
          <Text variant="muted" className="text-xs mt-0.5">
            {posts}
          </Text>
        </div>
        <TrendingUp size={16} className="text-muted-foreground mt-1 flex-shrink-0" />
      </div>
    </button>
  )
}
