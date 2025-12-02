import { useEffect } from "react";
import { Link } from "react-router-dom";
import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { useNotifications } from "@/hooks";
import { ArrowLeft, Heart, MessageCircle, UserPlus, CheckCheck } from "lucide-react";
import { formatDistanceToNow } from "date-fns";

export default function NotificationsPage() {
  const { notifications, unreadCount, loading, fetchNotifications, markAsRead, markAllAsRead } = useNotifications();

  useEffect(() => {
    fetchNotifications();
  }, []);

  const getNotificationIcon = (type: string) => {
    switch (type) {
      case "like":
        return <Heart className="w-5 h-5 text-red-500 fill-red-500" />;
      case "reply":
      case "mention":
        return <MessageCircle className="w-5 h-5 text-blue-500" />;
      case "follow":
        return <UserPlus className="w-5 h-5 text-green-500" />;
      default:
        return <CheckCheck className="w-5 h-5 text-primary" />;
    }
  };

  const getNotificationMessage = (notification: any) => {
    switch (notification.type) {
      case "like":
        return "liked your tweet";
      case "reply":
        return "replied to your tweet";
      case "mention":
        return "mentioned you in a tweet";
      case "follow":
        return "started following you";
      case "retweet":
        return "retweeted your tweet";
      default:
        return notification.content || "new notification";
    }
  };

  if (loading && notifications.length === 0) {
    return (
      <div className="flex items-center justify-center min-h-screen">
        <div className="animate-spin rounded-full h-8 w-8 border-b-2 border-primary"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen">
      {/* Header */}
      <div className="sticky top-0 z-10 bg-background/80 backdrop-blur-sm border-b">
        <div className="flex items-center justify-between p-4">
          <div className="flex items-center gap-4">
            <Link to="/">
              <Button variant="ghost" size="icon">
                <ArrowLeft className="w-5 h-5" />
              </Button>
            </Link>
            <div>
              <h1 className="font-bold text-xl">Notifications</h1>
              {unreadCount > 0 && (
                <p className="text-sm text-muted-foreground">
                  {unreadCount} unread notification{unreadCount > 1 ? "s" : ""}
                </p>
              )}
            </div>
          </div>
          
          {unreadCount > 0 && (
            <Button 
              variant="ghost" 
              size="sm"
              onClick={() => markAllAsRead()}
            >
              Mark all as read
            </Button>
          )}
        </div>
      </div>

      {/* Notifications List */}
      <div className="divide-y">
        {notifications.length === 0 ? (
          <Card className="m-4">
            <CardContent className="p-8 text-center">
              <div className="flex flex-col items-center gap-2">
                <CheckCheck className="w-12 h-12 text-muted-foreground" />
                <h3 className="font-semibold text-lg">No notifications yet</h3>
                <p className="text-sm text-muted-foreground">
                  When you get notifications, they'll show up here
                </p>
              </div>
            </CardContent>
          </Card>
        ) : (
          notifications.map((notification) => (
            <div
              key={notification.id}
              className={`p-4 hover:bg-accent/50 transition-colors cursor-pointer ${
                !notification.is_read ? "bg-accent/20" : ""
              }`}
              onClick={() => !notification.is_read && markAsRead(notification.id)}
            >
              <div className="flex gap-3">
                {/* Icon */}
                <div className="flex-shrink-0">
                  {getNotificationIcon(notification.type)}
                </div>

                {/* Content */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-start gap-3">
                    {/* Avatar - for now using placeholder */}
                    <Avatar className="h-10 w-10">
                      <AvatarFallback>
                        {notification.type[0].toUpperCase()}
                      </AvatarFallback>
                    </Avatar>

                    <div className="flex-1">
                      <p className="text-sm">
                        <span className="text-muted-foreground">
                          {getNotificationMessage(notification)}
                        </span>
                      </p>
                      
                      {notification.content && notification.type !== "follow" && (
                        <p className="text-sm text-muted-foreground mt-1 line-clamp-2">
                          {notification.content}
                        </p>
                      )}
                      
                      <p className="text-xs text-muted-foreground mt-1">
                        {formatDistanceToNow(new Date(notification.created_at), {
                          addSuffix: true,
                        })}
                      </p>
                    </div>

                    {/* Unread indicator */}
                    {!notification.is_read && (
                      <div className="flex-shrink-0">
                        <div className="w-2 h-2 bg-blue-500 rounded-full"></div>
                      </div>
                    )}
                  </div>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}
