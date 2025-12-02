import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import TweetCard from "@/components/fragments/feed/TweetCard";
import { UserCard } from "@/components/fragments/user/UserCard";
import { useUser, useFollow, useTweets } from "@/hooks";
import { useAuth } from "@/context/auth";
import { type Tweet } from "@/lib/api";
import { ArrowLeft, Calendar, UserPlus, UserMinus } from "lucide-react";
import { format } from "date-fns";

export default function ProfilePage() {
  const { id } = useParams<{ id: string }>();
  const { user: currentUser } = useAuth();
  const { user, loading: userLoading, fetchUser, fetchUsers } = useUser();
  const { followUser, unfollowUser, followers, following, getFollowers, getFollowing } = useFollow();
  const { fetchTweets } = useTweets();
  const [isFollowing, setIsFollowing] = useState(false);
  const [userTweets, setUserTweets] = useState<Tweet[]>([]);
  const [activeTab, setActiveTab] = useState("tweets");
  const [followerUsers, setFollowerUsers] = useState<any[]>([]);
  const [followingUsers, setFollowingUsers] = useState<any[]>([]);

  useEffect(() => {
    if (id) {
      loadUserProfile(id);
    }
  }, [id]);

  const loadUserProfile = async (identifier: string) => {
    try {
      // Check if identifier is a number (user ID) or string (username)
      const userId = parseInt(identifier);
      
      if (isNaN(userId)) {
        // It's a username, need to fetch all users and find by username
        const allUsers = await fetchUsers();
        const foundUser = allUsers.find(u => u.username === identifier);
        
        if (foundUser) {
          await fetchUser(foundUser.id);
          loadFollowData(foundUser.id);
          loadUserTweets(foundUser.id);
        }
      } else {
        // It's a user ID
        await fetchUser(userId);
        loadFollowData(userId);
        loadUserTweets(userId);
      }
    } catch (error) {
      console.error("Failed to load user profile:", error);
    }
  };

  const loadFollowData = async (userId: number) => {
    try {
      const [followersData, followingData, allUsers] = await Promise.all([
        getFollowers(userId),
        getFollowing(userId),
        fetchUsers()
      ]);

      // Map follower_id to user objects
      const followerUsersList = followersData
        .map(f => allUsers.find(u => u.id === f.follower_id))
        .filter(Boolean);
      setFollowerUsers(followerUsersList);

      // Map following_id to user objects
      const followingUsersList = followingData
        .map(f => allUsers.find(u => u.id === f.following_id))
        .filter(Boolean);
      setFollowingUsers(followingUsersList);
    } catch (error) {
      console.error("Failed to load follow data:", error);
    }
  };

  const loadUserTweets = async (userId: number) => {
    try {
      const allTweets = await fetchTweets();
      // Filter tweets by user_id
      const filtered = allTweets.filter((t) => t.user_id === userId);
      setUserTweets(filtered);
    } catch (error) {
      console.error("Failed to load user tweets:", error);
    }
  };

  const handleFollowToggle = async () => {
    if (!user) return;
    
    try {
      if (isFollowing) {
        await unfollowUser(user.id);
        setIsFollowing(false);
      } else {
        await followUser(user.id);
        setIsFollowing(true);
      }
      // Refresh followers/following
      loadFollowData(user.id);
    } catch (error) {
      console.error("Failed to toggle follow:", error);
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

  const isOwnProfile = currentUser?.id === user?.id;

  if (userLoading) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
      </div>
    );
  }

  if (!user) {
    return (
      <div className="flex flex-col items-center justify-center min-h-screen gap-4">
        <h2 className="text-2xl font-bold">User not found</h2>
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
          <div>
            <h1 className="font-bold text-xl">{user.name}</h1>
            <p className="text-sm text-muted-foreground">{userTweets.length} posts</p>
          </div>
        </div>
      </div>

      {/* Banner */}
      <div className="h-48 bg-gradient-to-r from-blue-400 to-purple-500">
        {user.banner && (
          <img src={user.banner} alt="Banner" className="w-full h-full object-cover" />
        )}
      </div>

      {/* Profile Info */}
      <div className="px-4">
        <div className="flex justify-between items-start -mt-16 mb-4">
          <Avatar className="h-32 w-32 border-4 border-background">
            <AvatarImage src={user.avatar} alt={user.name} />
            <AvatarFallback className="text-3xl">{getInitials(user.name)}</AvatarFallback>
          </Avatar>
          
          {!isOwnProfile && (
            <Button
              variant={isFollowing ? "outline" : "default"}
              onClick={handleFollowToggle}
              className="mt-4"
            >
              {isFollowing ? (
                <>
                  <UserMinus className="w-4 h-4 mr-2" />
                  Unfollow
                </>
              ) : (
                <>
                  <UserPlus className="w-4 h-4 mr-2" />
                  Follow
                </>
              )}
            </Button>
          )}
          
          {isOwnProfile && (
            <Button variant="outline" className="mt-4">
              Edit Profile
            </Button>
          )}
        </div>

        <div className="space-y-3">
          <div>
            <div className="flex items-center gap-2">
              <h2 className="text-2xl font-bold">{user.name}</h2>
              {user.verified && (
                <svg
                  className="w-6 h-6 text-blue-500"
                  fill="currentColor"
                  viewBox="0 0 24 24"
                >
                  <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              )}
            </div>
            <p className="text-muted-foreground">@{user.username}</p>
          </div>

          {user.bio && (
            <p className="text-foreground">{user.bio}</p>
          )}

          <div className="flex flex-wrap gap-4 text-sm text-muted-foreground">
            <div className="flex items-center gap-1">
              <Calendar className="w-4 h-4" />
              <span>Joined {format(new Date(user.created_at), "MMMM yyyy")}</span>
            </div>
          </div>

          <div className="flex gap-4 text-sm">
            <button
              onClick={() => setActiveTab("following")}
              className="hover:underline"
            >
              <span className="font-bold">{following.length}</span>{" "}
              <span className="text-muted-foreground">Following</span>
            </button>
            <button
              onClick={() => setActiveTab("followers")}
              className="hover:underline"
            >
              <span className="font-bold">{followers.length}</span>{" "}
              <span className="text-muted-foreground">Followers</span>
            </button>
          </div>
        </div>
      </div>

      {/* Tabs */}
      <Tabs value={activeTab} onValueChange={setActiveTab} className="mt-6">
        <TabsList className="w-full justify-start border-b rounded-none bg-transparent h-auto p-0">
          <TabsTrigger 
            value="tweets" 
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Posts
          </TabsTrigger>
          <TabsTrigger 
            value="replies" 
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Replies
          </TabsTrigger>
          <TabsTrigger 
            value="likes" 
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Likes
          </TabsTrigger>
          <TabsTrigger 
            value="followers" 
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Followers
          </TabsTrigger>
          <TabsTrigger 
            value="following" 
            className="flex-1 rounded-none border-b-2 border-transparent data-[state=active]:border-primary"
          >
            Following
          </TabsTrigger>
        </TabsList>

        <TabsContent value="tweets" className="space-y-0">
          {userTweets.length === 0 ? (
            <Card className="m-4">
              <CardContent className="p-8 text-center">
                <p className="text-muted-foreground">No tweets yet</p>
              </CardContent>
            </Card>
          ) : (
            userTweets.map((tweet) => (
              <TweetCard 
                key={tweet.id} 
                tweet={{
                  ...tweet,
                  user: user ? {
                    id: user.id,
                    name: user.name,
                    username: user.username,
                    avatar: user.avatar
                  } : {
                    id: tweet.user_id,
                    name: 'Unknown',
                    username: 'unknown',
                    avatar: undefined
                  }
                }} 
              />
            ))
          )}
        </TabsContent>

        <TabsContent value="replies" className="space-y-0">
          <Card className="m-4">
            <CardContent className="p-8 text-center">
              <p className="text-muted-foreground">No replies yet</p>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="likes" className="space-y-0">
          <Card className="m-4">
            <CardContent className="p-8 text-center">
              <p className="text-muted-foreground">No liked tweets yet</p>
            </CardContent>
          </Card>
        </TabsContent>

        <TabsContent value="followers" className="space-y-2 p-4">
          {followerUsers.length === 0 ? (
            <Card>
              <CardContent className="p-8 text-center">
                <p className="text-muted-foreground">No followers yet</p>
              </CardContent>
            </Card>
          ) : (
            followerUsers.map((followerUser: any) => (
              <UserCard 
                key={followerUser.id} 
                user={followerUser}
                showFollowButton={followerUser.id !== currentUser?.id}
              />
            ))
          )}
        </TabsContent>

        <TabsContent value="following" className="space-y-2 p-4">
          {followingUsers.length === 0 ? (
            <Card>
              <CardContent className="p-8 text-center">
                <p className="text-muted-foreground">Not following anyone yet</p>
              </CardContent>
            </Card>
          ) : (
            followingUsers.map((followingUser: any) => (
              <UserCard 
                key={followingUser.id} 
                user={followingUser}
                showFollowButton={followingUser.id !== currentUser?.id}
              />
            ))
          )}
        </TabsContent>
      </Tabs>
    </div>
  );
}
