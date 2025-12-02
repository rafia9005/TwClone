import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Textarea } from "@/components/ui/textarea";
import { useTweets, useUser, useLike } from "@/hooks";
import { useAuth } from "@/context/auth";
import { type Tweet } from "@/lib/api";
import { ArrowLeft, Heart, MessageCircle, Repeat2, Share, Calendar } from "lucide-react";
import { format } from "date-fns";

export default function TweetDetailPage() {
  const { id } = useParams<{ id: string }>();
  const { user: currentUser } = useAuth();
  const { getTweet, createTweet } = useTweets();
  const { fetchUser } = useUser();
  const { likeTweet, unlikeTweet } = useLike();
  
  const [tweet, setTweet] = useState<Tweet | null>(null);
  const [tweetUser, setTweetUser] = useState<any>(null);
  const [loading, setLoading] = useState(true);
  const [isLiked, setIsLiked] = useState(false);
  const [likesCount, setLikesCount] = useState(0);
  const [replyContent, setReplyContent] = useState("");
  const [isReplying, setIsReplying] = useState(false);
  const [replies, setReplies] = useState<Tweet[]>([]);

  useEffect(() => {
    if (id) {
      loadTweet(parseInt(id));
    }
  }, [id]);

  const loadTweet = async (tweetId: number) => {
    try {
      setLoading(true);
      const tweetData = await getTweet(tweetId);
      setTweet(tweetData);
      setLikesCount(tweetData.likes_count || 0);
      
      // Load user data
      if (tweetData.user_id) {
        const userData = await fetchUser(tweetData.user_id);
        setTweetUser(userData);
      }
      
      // TODO: Load replies when backend supports parent_id
      // For now, replies will be empty
      setReplies([]);
    } catch (error) {
      console.error("Failed to load tweet:", error);
    } finally {
      setLoading(false);
    }
  };

  const handleLike = async () => {
    if (!tweet) return;
    
    try {
      const previousLiked = isLiked;
      const previousCount = likesCount;
      
      // Optimistic update
      setIsLiked(!isLiked);
      setLikesCount(isLiked ? likesCount - 1 : likesCount + 1);
      
      if (isLiked) {
        await unlikeTweet(tweet.id);
      } else {
        await likeTweet(tweet.id);
      }
    } catch (error) {
      // Revert on error
      setIsLiked(!isLiked);
      setLikesCount(likesCount);
      console.error("Failed to toggle like:", error);
    }
  };

  const handleReply = async () => {
    if (!replyContent.trim() || !currentUser) return;
    
    try {
      setIsReplying(true);
      // Create a new tweet as reply
      // TODO: Add parent_id support when backend is ready
      await createTweet(replyContent);
      setReplyContent("");
      // Reload tweet to get updated replies count
      if (id) {
        await loadTweet(parseInt(id));
      }
    } catch (error) {
      console.error("Failed to post reply:", error);
    } finally {
      setIsReplying(false);
    }
  };

  const getInitials = (name: string) => {
    return name
      .split(" ")
      .map((n) => n[0])
      .join("")
      .toUpperCase()
      .slice(0, 2);
  };

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
      </div>
    );
  }

  if (!tweet || !tweetUser) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen gap-4">
        <h2 className="text-2xl font-bold">Tweet not found</h2>
        <Link to="/">
          <Button>Go Home</Button>
        </Link>
      </div>
    );
  }

  return (
    <div className="min-h-screen">
      {/* Header */}
      <div className="sticky top-0 z-10 bg-background/80 backdrop-blur-sm border-b">
        <div className="flex items-center gap-4 p-4">
          <Link to="/">
            <Button variant="ghost" size="icon">
              <ArrowLeft className="w-5 h-5" />
            </Button>
          </Link>
          <h1 className="font-bold text-xl">Tweet</h1>
        </div>
      </div>

      {/* Tweet Content */}
      <div className="border-b p-4">
        {/* User Info */}
        <Link to={`/profile/${tweetUser.username || tweetUser.id}`} className="flex items-center gap-3 mb-3">
          <Avatar className="h-12 w-12">
            <AvatarImage src={tweetUser.avatar} alt={tweetUser.name} />
            <AvatarFallback>{getInitials(tweetUser.name)}</AvatarFallback>
          </Avatar>
          <div>
            <div className="flex items-center gap-2">
              <h2 className="font-semibold hover:underline">{tweetUser.name}</h2>
              {tweetUser.verified && (
                <svg
                  className="w-5 h-5 text-blue-500"
                  fill="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              )}
            </div>
            <p className="text-sm text-muted-foreground">@{tweetUser.username}</p>
          </div>
        </Link>

        {/* Tweet Text */}
        <div className="text-xl mb-4">{tweet.content}</div>

        {/* Media */}
        {tweet.media_url && (
          <img
            src={tweet.media_url}
            alt="Tweet media"
            className="w-full rounded-2xl mb-4"
          />
        )}

        {/* Timestamp */}
        <div className="flex items-center gap-2 text-sm text-muted-foreground border-b pb-3 mb-3">
          <Calendar className="w-4 h-4" />
          <time>{format(new Date(tweet.created_at), "h:mm a · MMM d, yyyy")}</time>
        </div>

        {/* Stats */}
        <div className="flex gap-6 text-sm border-b pb-3 mb-3">
          <div>
            <span className="font-bold">{tweet.retweets_count || 0}</span>{" "}
            <span className="text-muted-foreground">Retweets</span>
          </div>
          <div>
            <span className="font-bold">{likesCount}</span>{" "}
            <span className="text-muted-foreground">Likes</span>
          </div>
          <div>
            <span className="font-bold">{tweet.replies_count || 0}</span>{" "}
            <span className="text-muted-foreground">Replies</span>
          </div>
        </div>

        {/* Actions */}
        <div className="flex items-center justify-around py-2">
          <Button variant="ghost" size="sm" className="gap-2">
            <MessageCircle className="w-5 h-5" />
            <span className="text-sm">{tweet.replies_count || 0}</span>
          </Button>
          
          <Button variant="ghost" size="sm" className="gap-2">
            <Repeat2 className="w-5 h-5" />
            <span className="text-sm">{tweet.retweets_count || 0}</span>
          </Button>
          
          <Button
            variant="ghost"
            size="sm"
            className={`gap-2 ${isLiked ? "text-red-500" : ""}`}
            onClick={handleLike}
          >
            <Heart className={`w-5 h-5 ${isLiked ? "fill-current" : ""}`} />
            <span className="text-sm">{likesCount}</span>
          </Button>
          
          <Button variant="ghost" size="sm">
            <Share className="w-5 h-5" />
          </Button>
        </div>
      </div>

      {/* Reply Form */}
      {currentUser && (
        <div className="border-b p-4">
          <div className="flex gap-3">
            <Avatar className="h-10 w-10">
              <AvatarImage src={currentUser.avatar} alt={currentUser.name} />
              <AvatarFallback>{getInitials(currentUser.name)}</AvatarFallback>
            </Avatar>
            <div className="flex-1">
              <Textarea
                placeholder="Tweet your reply"
                value={replyContent}
                onChange={(e) => setReplyContent(e.target.value)}
                className="min-h-[80px] resize-none"
              />
              <div className="flex justify-end mt-2">
                <Button
                  onClick={handleReply}
                  disabled={!replyContent.trim() || isReplying}
                  size="sm"
                >
                  {isReplying ? "Replying..." : "Reply"}
                </Button>
              </div>
            </div>
          </div>
        </div>
      )}

      {/* Replies */}
      <div>
        {replies.length === 0 ? (
          <Card className="m-4">
            <CardContent className="p-8 text-center">
              <p className="text-muted-foreground">No replies yet</p>
            </CardContent>
          </Card>
        ) : (
          replies.map((reply) => (
            <div key={reply.id} className="border-b p-4">
              {/* Reply content - similar to TweetCard */}
              <p>{reply.content}</p>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
