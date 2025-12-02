import { Card, CardContent } from "@/components/ui/card"
import { Avatar, AvatarImage, AvatarFallback } from "@/components/ui/avatar"
import { Button } from "@/components/ui/button"
import { formatDistanceToNow } from "date-fns"

type Tweet = {
  id: number
  user: { id: number; name: string; username: string; avatar?: string }
  content: string
  created_at: string
}

export default function TweetCard({ tweet }: { tweet: Tweet }) {
  return (
    <Card className="mb-3">
      <CardContent>
        <div className="flex gap-3">
          <Avatar>
            {tweet.user.avatar ? (
              <AvatarImage src={tweet.user.avatar} />
            ) : (
              <AvatarFallback>{tweet.user.name?.[0] ?? "U"}</AvatarFallback>
            )}
          </Avatar>
          <div className="flex-1">
            <div className="flex items-center gap-2">
              <div className="font-semibold">{tweet.user.name}</div>
              <div className="text-sm text-muted-foreground">@{tweet.user.username}</div>
              <div className="text-sm text-muted-foreground">•</div>
              <div className="text-sm text-muted-foreground">{formatDistanceToNow(new Date(tweet.created_at), { addSuffix: true })}</div>
            </div>
            <div className="mt-2">{tweet.content}</div>
            <div className="mt-3 flex gap-2">
              <Button variant="ghost" size="sm">Like</Button>
              <Button variant="ghost" size="sm">Reply</Button>
            </div>
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
