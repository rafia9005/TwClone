"use client"

import { useState } from "react"
import { Heart, MessageCircle, Repeat2, ImageIcon } from "lucide-react"

interface PostComposerProps {
  onPost: (content: string) => void
}

export default function PostComposer({ onPost }: PostComposerProps) {
  const [content, setContent] = useState("")

  const handlePost = () => {
    if (content.trim()) {
      onPost(content)
      setContent("")
    }
  }

  return (
    <div className="border-b border-border p-4 flex gap-4">
      <div className="w-12 h-12 bg-muted rounded-full flex-shrink-0" />

      <div className="flex-1">
        <textarea
          placeholder="What's happening!?"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="w-full bg-transparent text-2xl placeholder-muted-foreground outline-none resize-none"
          rows={3}
        />

        <div className="flex items-center justify-between mt-4 pt-4 border-t border-border">
          <div className="flex gap-4 text-primary">
            <button className="hover:bg-primary/10 p-2 rounded-full transition-colors">
              <ImageIcon size={20} />
            </button>
            <button className="hover:bg-primary/10 p-2 rounded-full transition-colors">
              <Heart size={20} />
            </button>
            <button className="hover:bg-primary/10 p-2 rounded-full transition-colors">
              <Repeat2 size={20} />
            </button>
            <button className="hover:bg-primary/10 p-2 rounded-full transition-colors">
              <MessageCircle size={20} />
            </button>
          </div>

          <button
            onClick={handlePost}
            disabled={!content.trim()}
            className="bg-primary hover:bg-primary/90 disabled:opacity-50 disabled:cursor-not-allowed text-primary-foreground font-bold py-2 px-6 rounded-full transition-all"
          >
            Post
          </button>
        </div>
      </div>
    </div>
  )
}
