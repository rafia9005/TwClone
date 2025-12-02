"use client"

import { useState } from "react"
import { Avatar } from "@/components/elements/Avatar"
import { IconButton } from "@/components/elements/IconButton"
import { Text } from "@/components/elements/Text"
import { Heart, MessageCircle, Repeat2, Share, MoreHorizontal, Bookmark } from "lucide-react"
import { cn } from "@/lib/utils"
import { formatDistanceToNow } from "date-fns"

interface TweetCardProps {
  tweet: {
    id: number
    user: {
      name: string
      username: string
      avatar?: string
      verified?: boolean
    }
    content: string
    created_at: string
    image?: string
    likes?: number
    replies?: number
    retweets?: number
    bookmarks?: number
  }
  onLike?: (id: number) => void
  onReply?: (id: number) => void
  onRetweet?: (id: number) => void
  onBookmark?: (id: number) => void
  liked?: boolean
  bookmarked?: boolean
}

export function TweetCard({
  tweet,
  onLike,
  onReply,
  onRetweet,
  onBookmark,
  liked = false,
  bookmarked = false,
}: TweetCardProps) {
  const [isLiked, setIsLiked] = useState(liked)
  const [isBookmarked, setIsBookmarked] = useState(bookmarked)

  const handleLike = () => {
    setIsLiked(!isLiked)
    onLike?.(tweet.id)
  }

  const handleBookmark = () => {
    setIsBookmarked(!isBookmarked)
    onBookmark?.(tweet.id)
  }

  return (
    <article className="flex gap-3 p-4 border-b border-border hover:bg-muted/30 transition-colors cursor-pointer group">
      <Avatar src={tweet.user.avatar} alt={tweet.user.name} size="lg" />

      <div className="flex-1 min-w-0">
        <div className="flex items-start justify-between gap-2">
          <div className="flex items-center gap-2 min-w-0 flex-wrap">
            <span className="font-bold hover:underline truncate">{tweet.user.name}</span>
            {tweet.user.verified && (
              <svg className="w-4 h-4 text-primary flex-shrink-0" viewBox="0 0 24 24" fill="currentColor">
                <path d="M8.52 3.59a1.5 1.5 0 012.96 0l.9 3.17a1.5 1.5 0 001.42 1.03h3.34a1.5 1.5 0 01.88 2.71l-2.7 1.96a1.5 1.5 0 00-.55 1.68l.9 3.17a1.5 1.5 0 01-2.31 1.68l-2.7-1.96a1.5 1.5 0 00-1.76 0l-2.7 1.96a1.5 1.5 0 01-2.31-1.68l.9-3.17a1.5 1.5 0 00-.55-1.68l-2.7-1.96a1.5 1.5 0 01.88-2.71h3.34a1.5 1.5 0 001.42-1.03l.9-3.17z" />
              </svg>
            )}
            <Text variant="muted" as="span" className="truncate">
              @{tweet.user.username}
            </Text>
            <Text variant="muted" as="span" className="flex-shrink-0">
              ·
            </Text>
            <Text variant="muted" as="span" className="flex-shrink-0">
              {formatDistanceToNow(new Date(tweet.created_at), { addSuffix: true })}
            </Text>
          </div>

          <button
            className="opacity-0 group-hover:opacity-100 p-2 rounded-full hover:bg-primary/10 transition-all flex-shrink-0"
            aria-label="More options"
          >
            <MoreHorizontal size={18} className="text-muted-foreground" />
          </button>
        </div>

        <div className="mt-2">
          <Text className="whitespace-pre-wrap break-words">{tweet.content}</Text>
        </div>

        {tweet.image && (
          <div className="mt-3 rounded-2xl overflow-hidden border border-border">
            <img src={tweet.image} alt="Tweet image" className="w-full h-auto" />
          </div>
        )}

        <div className="flex items-center justify-between mt-3 max-w-md">
          <IconButton
            icon={MessageCircle}
            count={tweet.replies}
            onClick={() => onReply?.(tweet.id)}
            label="Reply"
          />

          <IconButton
            icon={Repeat2}
            count={tweet.retweets}
            onClick={() => onRetweet?.(tweet.id)}
            label="Retweet"
          />

          <IconButton
            icon={Heart}
            count={tweet.likes}
            active={isLiked}
            onClick={handleLike}
            label="Like"
            className={cn(isLiked && "text-pink-600 hover:bg-pink-600/10")}
          />

          <IconButton
            icon={Bookmark}
            active={isBookmarked}
            onClick={handleBookmark}
            label="Bookmark"
            className={cn(isBookmarked && "text-primary")}
          />

          <IconButton icon={Share} label="Share" />
        </div>
      </div>
    </article>
  )
}
