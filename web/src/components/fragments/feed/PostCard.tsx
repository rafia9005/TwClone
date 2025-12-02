"use client"

import { Heart, MessageCircle, Repeat2, Share, MoreHorizontal } from "lucide-react"

interface Post {
  id: number
  author: string
  handle: string
  avatar: string
  content: string
  timestamp: string
  likes: number
  replies: number
  retweets: number
  liked: boolean
  image?: string
}

interface PostCardProps {
  post: Post
  onLike: () => void
}

export default function PostCard({ post, onLike }: PostCardProps) {
  return (
    <div className="border-b border-border p-4 hover:bg-secondary/5 transition-colors cursor-pointer group">
      <div className="flex gap-4">
        <div className="w-12 h-12 bg-muted rounded-full flex-shrink-0" />

        <div className="flex-1 min-w-0">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2 min-w-0">
              <span className="font-bold hover:underline">{post.author}</span>
              <span className="text-muted-foreground">{post.handle}</span>
              <span className="text-muted-foreground">·</span>
              <span className="text-muted-foreground">{post.timestamp}</span>
            </div>
            <button className="opacity-0 group-hover:opacity-100 text-muted-foreground hover:text-primary transition-all p-2 rounded-full hover:bg-primary/10">
              <MoreHorizontal size={18} />
            </button>
          </div>

          <p className="text-base mt-2 leading-normal text-foreground break-words">{post.content}</p>

          {post.image && (
            <div className="mt-3 rounded-2xl overflow-hidden border border-border">
              <img src={post.image || "/placeholder.svg"} alt="Post image" className="w-full h-auto" />
            </div>
          )}

          <div className="flex justify-between mt-4 text-muted-foreground text-sm max-w-xs">
            <div className="flex items-center gap-2 hover:text-primary group/item cursor-pointer">
              <div className="group-hover/item:bg-primary/10 p-2 rounded-full transition-colors">
                <MessageCircle size={18} />
              </div>
              <span className="group-hover/item:text-primary transition-colors">{post.replies}</span>
            </div>

            <div className="flex items-center gap-2 hover:text-primary group/item cursor-pointer">
              <div className="group-hover/item:bg-primary/10 p-2 rounded-full transition-colors">
                <Repeat2 size={18} />
              </div>
              <span className="group-hover/item:text-primary transition-colors">{post.retweets}</span>
            </div>

            <div className="flex items-center gap-2 hover:text-primary group/item cursor-pointer" onClick={onLike}>
              <div className="group-hover/item:bg-primary/10 p-2 rounded-full transition-colors">
                <Heart
                  size={18}
                  fill={post.liked ? "currentColor" : "none"}
                  color={post.liked ? "#ff5a7e" : "currentColor"}
                />
              </div>
              <span className="group-hover/item:text-primary transition-colors">{post.likes}</span>
            </div>

            <div className="flex items-center gap-2 hover:text-primary group/item cursor-pointer">
              <div className="group-hover/item:bg-primary/10 p-2 rounded-full transition-colors">
                <Share size={18} />
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
