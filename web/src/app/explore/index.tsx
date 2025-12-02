import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import TweetCard from "@/components/fragments/feed/TweetCard";
import { UserCard } from "@/components/fragments/user/UserCard";
import { useTweets, useUser } from "@/hooks";
import { ArrowLeft, Search, TrendingUp } from "lucide-react";
import { type Tweet, type User } from "@/lib/api";

export default function ExplorePage() {
  const { tweets, loading: tweetsLoading, fetchTweets } = useTweets();
  const { users, loading: usersLoading, fetchUsers } = useUser();
  const [searchQuery, setSearchQuery] = useState("");
  const [filteredTweets, setFilteredTweets] = useState<Tweet[]>([]);
  const [filteredUsers, setFilteredUsers] = useState<User[]>([]);
  const [activeTab, setActiveTab] = useState("tweets");

  useEffect(() => {
    fetchTweets();
    fetchUsers();
  }, []);

  useEffect(() => {
    if (searchQuery.trim()) {
      // Filter tweets by content
      const tweetResults = tweets.filter((tweet) =>
        tweet.content.toLowerCase().includes(searchQuery.toLowerCase())
      );
      setFilteredTweets(tweetResults);

      // Filter users by name or username
      const userResults = users.filter(
        (user) =>
          user.name.toLowerCase().includes(searchQuery.toLowerCase()) ||
          user.username.toLowerCase().includes(searchQuery.toLowerCase())
      );
      setFilteredUsers(userResults);
    } else {
      setFilteredTweets(tweets);
      setFilteredUsers(users);
    }
  }, [searchQuery, tweets, users]);

  const trendingTopics = [
    { id: 1, topic: "#golang", tweets: "12.5K tweets" },
    { id: 2, topic: "#typescript", tweets: "8.3K tweets" },
    { id: 3, topic: "#react", tweets: "15.2K tweets" },
    { id: 4, topic: "#webdev", tweets: "21.8K tweets" },
    { id: 5, topic: "#shadcn", tweets: "5.1K tweets" },
  ];

  return (
    <div className="min-h-screen">
      {/* Header */}
      <div className="sticky top-0 z-10 bg-background/80 backdrop-blur-sm border-b">
        <div className="p-4 space-y-4">
          <div className="flex items-center gap-4">
            <Link to="/">
              <Button variant="ghost" size="icon">
                <ArrowLeft className="w-5 h-5" />
              </Button>
            </Link>
            <h1 className="font-bold text-xl">Explore</h1>
          </div>

          {/* Search Bar */}
          <div className="relative">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 w-5 h-5 text-muted-foreground" />
            <Input
              type="text"
              placeholder="Search tweets and users..."
              value={searchQuery}
              onChange={(e) => setSearchQuery(e.target.value)}
              className="pl-10"
            />
          </div>
        </div>
      </div>

      {/* Trending Topics */}
      {!searchQuery && (
        <Card className="m-4">
          <CardContent className="p-4">
            <div className="flex items-center gap-2 mb-4">
              <TrendingUp className="w-5 h-5 text-primary" />
              <h2 className="font-bold text-lg">Trending Topics</h2>
            </div>
            <div className="space-y-3">
              {trendingTopics.map((trend) => (
                <button
                  key={trend.id}
                  className="w-full text-left p-3 rounded-lg hover:bg-accent/50 transition-colors"
                  onClick={() => setSearchQuery(trend.topic)}
                >
                  <p className="font-semibold text-primary">{trend.topic}</p>
                  <p className="text-sm text-muted-foreground">{trend.tweets}</p>
                </button>
              ))}
            </div>
          </CardContent>
        </Card>
      )}

      {/* Search Results */}
      {searchQuery && (
        <div className="p-4">
          <p className="text-sm text-muted-foreground mb-4">
            {filteredTweets.length + filteredUsers.length} results for "{searchQuery}"
          </p>
        </div>
      )}

      {/* Tabs */}
      <Tabs value={activeTab} onValueChange={setActiveTab} className="mt-2">
        <TabsList className="w-full justify-start border-b rounded-none bg-transparent h-auto p-0">
          <TabsTrigger
            value="tweets"
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Tweets
            {searchQuery && (
              <span className="ml-2 text-xs bg-accent px-2 py-0.5 rounded-full">
                {filteredTweets.length}
              </span>
            )}
          </TabsTrigger>
          <TabsTrigger
            value="users"
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Users
            {searchQuery && (
              <span className="ml-2 text-xs bg-accent px-2 py-0.5 rounded-full">
                {filteredUsers.length}
              </span>
            )}
          </TabsTrigger>
        </TabsList>

        <TabsContent value="tweets" className="space-y-0">
          {tweetsLoading ? (
            <div className="flex items-center justify-center p-8">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
            </div>
          ) : filteredTweets.length === 0 ? (
            <Card className="m-4">
              <CardContent className="p-8 text-center">
                <p className="text-muted-foreground">
                  {searchQuery ? "No tweets found" : "No tweets available"}
                </p>
              </CardContent>
            </Card>
          ) : (
            filteredTweets.map((tweet) => (
              <TweetCard
                key={tweet.id}
                tweet={{
                  ...tweet,
                  user: tweet.user || {
                    id: tweet.user_id,
                    name: "Unknown User",
                    username: "unknown",
                    avatar: undefined,
                  },
                }}
              />
            ))
          )}
        </TabsContent>

        <TabsContent value="users" className="space-y-2 p-4">
          {usersLoading ? (
            <div className="flex items-center justify-center p-8">
              <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
            </div>
          ) : filteredUsers.length === 0 ? (
            <Card>
              <CardContent className="p-8 text-center">
                <p className="text-muted-foreground">
                  {searchQuery ? "No users found" : "No users available"}
                </p>
              </CardContent>
            </Card>
          ) : (
            filteredUsers.map((user) => (
              <UserCard key={user.id} user={user} showFollowButton={true} />
            ))
          )}
        </TabsContent>
      </Tabs>
    </div>
  );
}
