import { useState } from "react"
import { Link } from "react-router-dom"
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { useLike } from "@/hooks"
import { formatDistanceToNow } from "date-fns"
import { Heart, MessageCircle, Repeat2, Share, Bookmark } from "lucide-react"

type Tweet = {
  id: number
  user: { id: number; name: string; username: string; avatar?: string; verified?: boolean }
  content: string
  created_at: string
  likes_count?: number
  replies_count?: number
  retweets_count?: number
  is_liked?: boolean
}

export default function TweetCard({ tweet }: { tweet: Tweet }) {
  const { likeTweet, unlikeTweet } = useLike()
  const [isLiked, setIsLiked] = useState(tweet.is_liked || false)
  const [likesCount, setLikesCount] = useState(tweet.likes_count || 0)

  const handleLike = async (e: React.MouseEvent) => {
    e.preventDefault()
    e.stopPropagation()
    
    const previousLiked = isLiked
    const previousCount = likesCount
    
    // Optimistic update
    setIsLiked(!isLiked)
    setLikesCount(isLiked ? likesCount - 1 : likesCount + 1)
    
    try {
      if (isLiked) {
        await unlikeTweet(tweet.id)
      } else {
        await likeTweet(tweet.id)
      }
    } catch (error) {
      // Revert on error
      setIsLiked(previousLiked)
      setLikesCount(previousCount)
      console.error("Failed to toggle like:", error)
    }
  }

  const handleShare = async (e: React.MouseEvent) => {
    e.preventDefault()
    e.stopPropagation()
    
    if (navigator.share) {
      try {
        await navigator.share({
          title: `${tweet.user.name} on TwClone`,
          text: tweet.content,
          url: `${window.location.origin}/tweet/${tweet.id}`,
        })
      } catch (error) {
        console.log("Share cancelled")
      }
    } else {
      // Fallback: copy to clipboard
      navigator.clipboard.writeText(`${window.location.origin}/tweet/${tweet.id}`)
      alert("Link copied to clipboard!")
    }
  }

  return (
    <Link to={`/tweet/${tweet.id}`}>
      <article className="p-4 border-b hover:bg-accent/50 transition-colors cursor-pointer">
        <div className="flex gap-3">
          <Link 
            to={`/profile/${tweet.user.username || tweet.user.id}`}
            onClick={(e) => e.stopPropagation()}
            className="flex-shrink-0"
          >
            <Avatar className="h-12 w-12">
              {tweet.user.avatar ? (
                <AvatarImage src={tweet.user.avatar} alt={tweet.user.name} />
              ) : (
                <AvatarFallback>{tweet.user.name?.[0] ?? "U"}</AvatarFallback>
              )}
            </Avatar>
          </Link>
          
          <div className="flex-1 min-w-0">
            <header className="flex items-center gap-2 flex-wrap">
              <Link 
                to={`/profile/${tweet.user.username || tweet.user.id}`}
                onClick={(e) => e.stopPropagation()}
                className="font-semibold hover:underline"
              >
                {tweet.user.name}
              </Link>
              {tweet.user.verified && (
                <svg className="w-4 h-4 text-blue-500 flex-shrink-0" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              )}
              <span className="text-sm text-muted-foreground">@{tweet.user.username}</span>
              <span className="text-sm text-muted-foreground">·</span>
              <time className="text-sm text-muted-foreground">
                {formatDistanceToNow(new Date(tweet.created_at), { addSuffix: true })}
              </time>
            </header>

            <div className="mt-2 text-sm leading-relaxed whitespace-pre-wrap break-words">
              {tweet.content}
            </div>

            <footer className="mt-3 flex items-center justify-between max-w-md">
              <Button
                variant="ghost"
                size="sm"
                className="gap-2 text-muted-foreground hover:text-blue-500 hover:bg-blue-500/10"
                onClick={(e) => {
                  e.preventDefault()
                  e.stopPropagation()
                }}
              >
                <MessageCircle className="w-4 h-4" />
                <span className="text-xs">{tweet.replies_count || 0}</span>
              </Button>

              <Button
                variant="ghost"
                size="sm"
                className="gap-2 text-muted-foreground hover:text-green-500 hover:bg-green-500/10"
                onClick={(e) => {
                  e.preventDefault()
                  e.stopPropagation()
                }}
              >
                <Repeat2 className="w-4 h-4" />
                <span className="text-xs">{tweet.retweets_count || 0}</span>
              </Button>

              <Button
                variant="ghost"
                size="sm"
                className={`gap-2 hover:bg-red-500/10 ${
                  isLiked ? "text-red-500" : "text-muted-foreground hover:text-red-500"
                }`}
                onClick={handleLike}
              >
                <Heart className={`w-4 h-4 ${isLiked ? "fill-current" : ""}`} />
                <span className="text-xs">{likesCount}</span>
              </Button>

              <Button
                variant="ghost"
                size="sm"
                className="gap-2 text-muted-foreground hover:text-blue-500 hover:bg-blue-500/10"
                onClick={(e) => {
                  e.preventDefault()
                  e.stopPropagation()
                }}
              >
                <Bookmark className="w-4 h-4" />
              </Button>

              <Button
                variant="ghost"
                size="sm"
                className="text-muted-foreground hover:text-blue-500 hover:bg-blue-500/10"
                onClick={handleShare}
              >
                <Share className="w-4 h-4" />
              </Button>
            </footer>
          </div>
        </div>
      </article>
    </Link>
  )
}
