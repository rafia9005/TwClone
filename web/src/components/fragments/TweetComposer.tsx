"use client"

import { useState } from "react"
import { Avatar } from "@/components/elements/Avatar"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { ImageIcon, Smile, MapPin, BarChart3 } from "lucide-react"
import { cn } from "@/lib/utils"

interface TweetComposerProps {
  user?: {
    name: string
    username: string
    avatar?: string
  }
  onPost?: (content: string) => void
  placeholder?: string
  autoFocus?: boolean
}

export function TweetComposer({ user, onPost, placeholder = "What's happening?", autoFocus }: TweetComposerProps) {
  const [content, setContent] = useState("")
  const maxLength = 280

  const handlePost = () => {
    if (content.trim() && content.length <= maxLength) {
      onPost?.(content)
      setContent("")
    }
  }

  const progress = (content.length / maxLength) * 100
  const isOverLimit = content.length > maxLength

  return (
    <div className="flex gap-3 p-4 border-b border-border">
      <Avatar src={user?.avatar} alt={user?.name || "User"} size="lg" />
      
      <div className="flex-1 min-w-0">
        <Textarea
          value={content}
          onChange={(e: React.ChangeEvent<HTMLTextAreaElement>) => setContent(e.target.value)}
          placeholder={placeholder}
          autoFocus={autoFocus}
          className="min-h-[120px] text-xl border-0 resize-none focus-visible:ring-0 focus-visible:ring-offset-0 p-0 placeholder:text-muted-foreground"
        />

        <div className="flex items-center justify-between mt-4 pt-4 border-t border-border">
          <div className="flex items-center gap-1">
            <Button variant="ghost" size="icon" className="h-9 w-9 text-primary hover:bg-primary/10">
              <ImageIcon className="h-5 w-5" />
            </Button>
            <Button variant="ghost" size="icon" className="h-9 w-9 text-primary hover:bg-primary/10">
              <Smile className="h-5 w-5" />
            </Button>
            <Button variant="ghost" size="icon" className="h-9 w-9 text-primary hover:bg-primary/10">
              <MapPin className="h-5 w-5" />
            </Button>
            <Button variant="ghost" size="icon" className="h-9 w-9 text-primary hover:bg-primary/10">
              <BarChart3 className="h-5 w-5" />
            </Button>
          </div>

          <div className="flex items-center gap-3">
            {content.length > 0 && (
              <div className="flex items-center gap-2">
                <svg className="w-8 h-8 -rotate-90" viewBox="0 0 32 32">
                  <circle
                    cx="16"
                    cy="16"
                    r="14"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="3"
                    className="text-muted/20"
                  />
                  <circle
                    cx="16"
                    cy="16"
                    r="14"
                    fill="none"
                    stroke="currentColor"
                    strokeWidth="3"
                    strokeDasharray={`${2 * Math.PI * 14}`}
                    strokeDashoffset={`${2 * Math.PI * 14 * (1 - progress / 100)}`}
                    className={cn(
                      "transition-all",
                      isOverLimit ? "text-destructive" : progress > 90 ? "text-yellow-500" : "text-primary"
                    )}
                  />
                </svg>
                <span className={cn("text-sm", isOverLimit && "text-destructive")}>
                  {maxLength - content.length}
                </span>
              </div>
            )}
            <Button
              onClick={handlePost}
              disabled={!content.trim() || isOverLimit}
              className="rounded-full font-bold px-6"
            >
              Post
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}
