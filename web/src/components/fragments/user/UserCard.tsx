import { useState } from "react";
import { Link } from "react-router-dom";
import { Card, CardContent } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useFollow } from "@/hooks";
import { type User } from "@/lib/api";
import { UserPlus, UserMinus } from "lucide-react";

interface UserCardProps {
  user: User;
  isFollowing?: boolean;
  onFollowChange?: (isFollowing: boolean) => void;
  showFollowButton?: boolean;
}

export const UserCard = ({ 
  user, 
  isFollowing: initialIsFollowing = false, 
  onFollowChange,
  showFollowButton = true 
}: UserCardProps) => {
  const [isFollowing, setIsFollowing] = useState(initialIsFollowing);
  const { followUser, unfollowUser, loading } = useFollow();

  const handleFollowToggle = async (e: React.MouseEvent) => {
    e.preventDefault();
    e.stopPropagation();
    
    try {
      if (isFollowing) {
        await unfollowUser(user.id);
        setIsFollowing(false);
        onFollowChange?.(false);
      } else {
        await followUser(user.id);
        setIsFollowing(true);
        onFollowChange?.(true);
      }
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

  return (
    <Card className="hover:bg-accent/50 transition-colors">
      <CardContent className="p-4">
        <div className="flex items-start justify-between">
          <Link to={`/profile/${user.username || user.id}`} className="flex items-start gap-3 flex-1">
            <Avatar className="h-12 w-12">
              <AvatarImage src={user.avatar} alt={user.name} />
              <AvatarFallback>{getInitials(user.name)}</AvatarFallback>
            </Avatar>
            
            <div className="flex-1 min-w-0">
              <div className="flex items-center gap-2">
                <h3 className="font-semibold text-sm truncate hover:underline">
                  {user.name}
                </h3>
                {user.verified && (
                  <svg
                    className="w-4 h-4 text-blue-500 flex-shrink-0"
                    fill="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                )}
              </div>
              <p className="text-sm text-muted-foreground">@{user.username}</p>
              {user.bio && (
                <p className="text-sm mt-2 text-foreground/80 line-clamp-2">
                  {user.bio}
                </p>
              )}
            </div>
          </Link>

          {showFollowButton && (
            <Button
              variant={isFollowing ? "outline" : "default"}
              size="sm"
              onClick={handleFollowToggle}
              disabled={loading}
              className="ml-2 flex-shrink-0"
            >
              {isFollowing ? (
                <>
                  <UserMinus className="w-4 h-4 mr-1" />
                  Unfollow
                </>
              ) : (
                <>
                  <UserPlus className="w-4 h-4 mr-1" />
                  Follow
                </>
              )}
            </Button>
          )}
        </div>
      </CardContent>
    </Card>
  );
};
