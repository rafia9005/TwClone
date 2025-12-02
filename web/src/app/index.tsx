"use client"

import { useState, useEffect } from "react"
import MainLayout from "@/components/layouts/MainLayout"
import { TweetComposer } from "@/components/fragments/TweetComposer"
import { TweetCard } from "@/components/fragments/TweetCard"
import { useAuth } from "@/context/auth"
import { tweetAPI, likeAPI, type Tweet, type User } from "@/lib/api"
import Cookies from "js-cookie"

interface TweetWithUser extends Tweet {
  user: User
  likes_count: number
  replies_count: number
  retweets_count: number
  bookmarks_count?: number
  is_liked?: boolean
}

export default function Index() {
  const { user } = useAuth()
  const [tweets, setTweets] = useState<TweetWithUser[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchTweets()
  }, [])

  const fetchTweets = async () => {
    try {
      setLoading(true)
      const data = await tweetAPI.getAll()
      // Transform to match TweetCard interface
      const tweetsWithUser = data.map(tweet => ({
        ...tweet,
        user: tweet.user || {
          id: tweet.user_id,
          name: "Unknown User",
          username: "unknown",
          email: "",
          created_at: "",
          updated_at: "",
        },
        likes_count: tweet.likes_count || 0,
        replies_count: tweet.replies_count || 0,
        retweets_count: tweet.retweets_count || 0,
        bookmarks_count: tweet.bookmarks_count || 0,
        is_liked: false,
      }))
      setTweets(tweetsWithUser)
    } catch (err: any) {
      console.error("Failed to fetch tweets:", err)
      setError(err.message || "Failed to load tweets")
      // Show sample tweets if API fails
      setTweets(getSampleTweets())
    } finally {
      setLoading(false)
    }
  }

  const getSampleTweets = (): TweetWithUser[] => [
    {
      id: 1,
      user_id: 1,
      user: {
        id: 1,
        name: "TwClone Team",
        username: "twclone",
        email: "team@twclone.com",
        avatar: undefined,
        verified: true,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
      },
      content: "Welcome to TwClone! 🎉\n\nThis is a modern Twitter clone built with React, TypeScript, Tailwind CSS, and shadcn/ui components following atomic design principles.\n\nFeel free to explore and post your first tweet!",
      created_at: new Date(Date.now() - 1000 * 60 * 30).toISOString(),
      updated_at: new Date(Date.now() - 1000 * 60 * 30).toISOString(),
      likes_count: 42,
      replies_count: 8,
      retweets_count: 15,
      bookmarks_count: 5,
      is_liked: false,
    },
  ]

  const handlePost = async (content: string, mediaUrl?: string) => {
    try {
      // Check if user is authenticated
      const token = Cookies.get("accessToken")
      if (!token) {
        alert("Please login to post a tweet")
        return
      }

      const newTweet = await tweetAPI.create({ content, media_url: mediaUrl })
      
      // Add the new tweet to the top of the list with user info
      const tweetWithUser: TweetWithUser = {
        ...newTweet,
        user: user || {
          id: newTweet.user_id,
          name: "Unknown User",
          username: "unknown",
          email: "",
          created_at: "",
          updated_at: "",
        },
        likes_count: 0,
        replies_count: 0,
        retweets_count: 0,
        bookmarks_count: 0,
        is_liked: false,
      }
      
      setTweets([tweetWithUser, ...tweets])
    } catch (err: any) {
      console.error("Failed to create tweet:", err)
      alert(err.response?.data?.message || "Failed to post tweet. Please try again.")
    }
  }

  const handleLike = async (id: number) => {
    try {
      const tweet = tweets.find(t => t.id === id)
      if (!tweet) return

      if (tweet.is_liked) {
        await likeAPI.unlike(id)
      } else {
        await likeAPI.like(id)
      }

      // Optimistic update
      setTweets(
        tweets.map((t) =>
          t.id === id
            ? {
                ...t,
                is_liked: !t.is_liked,
                likes_count: t.is_liked ? t.likes_count - 1 : t.likes_count + 1,
              }
            : t
        )
      )
    } catch (err: any) {
      console.error("Failed to like tweet:", err)
      // Revert on error by refetching
      fetchTweets()
    }
  }

  return (
    <MainLayout>
      {/* Header */}
      <div className="sticky top-0 z-40 bg-background/80 backdrop-blur-md border-b border-border">
        <div className="flex items-center justify-between px-4 h-[53px]">
          <h1 className="text-xl font-bold">Home</h1>
        </div>
      </div>

      {/* Tweet Composer */}
      {user && (
        <TweetComposer
          user={{
            name: user.name,
            username: user.username,
            avatar: user.avatar,
          }}
          onPost={handlePost}
        />
      )}

      {/* Loading State */}
      {loading && (
        <div className="p-8 text-center text-muted-foreground">
          <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary mx-auto mb-2"></div>
          Loading tweets...
        </div>
      )}

      {/* Error State */}
      {error && !loading && (
        <div className="p-8 text-center">
          <p className="text-muted-foreground mb-2">Could not connect to the backend</p>
          <p className="text-sm text-muted-foreground">Showing sample data instead</p>
        </div>
      )}

      {/* Feed */}
      {!loading && (
        <div>
          {tweets.length === 0 ? (
            <div className="p-8 text-center text-muted-foreground">
              <p className="text-lg font-semibold mb-2">No tweets yet</p>
              <p className="text-sm">Be the first to post!</p>
            </div>
          ) : (
            tweets.map((tweet) => (
              <TweetCard
                key={tweet.id}
                tweet={{
                  id: tweet.id,
                  user: {
                    name: tweet.user.name,
                    username: tweet.user.username,
                    avatar: tweet.user.avatar,
                    verified: (tweet.user as any).verified,
                  },
                  content: tweet.content,
                  created_at: tweet.created_at,
                  image: tweet.media_url,
                  likes: tweet.likes_count,
                  replies: tweet.replies_count,
                  retweets: tweet.retweets_count,
                  bookmarks: tweet.bookmarks_count,
                }}
                liked={tweet.is_liked}
                onLike={handleLike}
              />
            ))
          )}
        </div>
      )}
    </MainLayout>
  )
}