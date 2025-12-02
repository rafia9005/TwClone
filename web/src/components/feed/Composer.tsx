import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useState } from "react"

export default function Composer({ user, onTweet }: { user?: any; onTweet: (text: string) => void }) {
  const [text, setText] = useState("")
  const disabled = !user

  return (
    <div className="mb-4 p-4 bg-white rounded-lg shadow-sm">
      <div className="flex gap-3">
        <Avatar>
          {user?.avatar ? <AvatarImage src={user.avatar} /> : <AvatarFallback>{user?.name?.[0] ?? "G"}</AvatarFallback>}
        </Avatar>
        <div className="flex-1">
          <Input
            placeholder={disabled ? "Sign in to post or continue as guest" : "What's happening?"}
            value={text}
            onChange={(e) => setText(e.target.value)}
            disabled={disabled}
          />
          <div className="mt-3 flex justify-end">
            <Button disabled={text.trim().length === 0 || disabled} onClick={() => { onTweet(text.trim()); setText("") }}>
              Tweet
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}
