"use client"

import { useEffect, useState } from "react"
import { Home, Hash, Bell, Mail, Bookmark, User, Sun, Moon, Feather, Search } from "lucide-react"
import { NavItem } from "@/components/fragments/NavItem"
import { UserProfileCard } from "@/components/fragments/UserProfileCard"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { useAuth } from "@/context/auth"

export default function MainLayout({ children }: { children: React.ReactNode }) {
  const { user } = useAuth()
  const [theme, setTheme] = useState<string>(() => {
    if (typeof window === "undefined") return "light"
    return localStorage.getItem("theme") || (document.documentElement.classList.contains("dark") ? "dark" : "light")
  })

  useEffect(() => {
    if (theme === "dark") {
      document.documentElement.classList.add("dark")
      localStorage.setItem("theme", "dark")
    } else {
      document.documentElement.classList.remove("dark")
      localStorage.setItem("theme", "light")
    }
  }, [theme])

  return (
    <div className="app min-h-screen">
      <div className="max-w-7xl mx-auto flex">
        {/* Left Sidebar - Navigation */}
        <aside className="hidden lg:flex flex-col justify-between py-4 px-2 xl:px-4 sticky top-0 h-screen w-[88px] xl:w-[275px]">
          <div className="flex flex-col gap-1">
            {/* Logo */}
            <div className="flex items-center justify-center xl:justify-start gap-3 mb-4 px-3">
              <div className="w-12 h-12 flex items-center justify-center rounded-full bg-primary text-primary-foreground">
                <Feather size={28} />
              </div>
              <span className="hidden xl:block text-2xl font-extrabold">TwClone</span>
            </div>

            {/* Navigation */}
            <nav className="flex flex-col items-center xl:items-start">
              <NavItem icon={Home} label="Home" href="/" active />
              <NavItem icon={Search} label="Explore" href="/explore" />
              <NavItem icon={Bell} label="Notifications" href="/notifications" />
              <NavItem icon={Mail} label="Messages" href="/messages" />
              <NavItem icon={Bookmark} label="Bookmarks" href="/bookmarks" />
              <NavItem icon={User} label="Profile" href={user ? `/profile/${user.username}` : "/login"} />
            </nav>

            {/* Post Button */}
            <div className="mt-4 px-3">
              <Button className="w-full rounded-full font-bold text-base h-[52px] hidden xl:flex items-center justify-center gap-2">
                <Feather size={20} /> Post
              </Button>
              <Button size="icon" className="xl:hidden w-[52px] h-[52px] rounded-full">
                <Feather size={24} />
              </Button>
            </div>
          </div>

          {/* User Profile */}
          <div className="mt-auto">
            <UserProfileCard
              user={{
                name: user?.name || "Guest",
                username: user?.username || "guest",
              }}
              className="relative"
            />
            <button
              onClick={() => setTheme(theme === "dark" ? "light" : "dark")}
              className="mt-2 w-full p-3 rounded-full hover:bg-muted/60 transition flex items-center justify-center xl:justify-start gap-3"
              aria-label="Toggle theme"
            >
              {theme === "dark" ? <Sun size={20} /> : <Moon size={20} />}
              <span className="hidden xl:inline font-semibold">
                {theme === "dark" ? "Light mode" : "Dark mode"}
              </span>
            </button>
          </div>
        </aside>

        {/* Main Content */}
        <main className="flex-1 min-w-0 border-x border-border">
          {children}
        </main>

        {/* Right Sidebar - Search & Trends */}
        <aside className="hidden xl:block w-[350px] py-4 px-6">
          <div className="sticky top-4 space-y-4">
            {/* Search */}
            <div className="relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-muted-foreground" size={20} />
              <Input
                placeholder="Search"
                className="pl-12 rounded-full bg-muted/50 border-0 focus-visible:ring-1"
              />
            </div>

            {/* Subscribe Card */}
            <div className="bg-card rounded-2xl p-4 border border-border">
              <h3 className="text-xl font-bold mb-2">Subscribe to Premium</h3>
              <p className="text-sm text-muted-foreground mb-3">
                Subscribe to unlock new features and if eligible, receive a share of revenue.
              </p>
              <Button className="rounded-full font-bold">Subscribe</Button>
            </div>
          </div>
        </aside>
      </div>

      {/* Mobile Bottom Navigation */}
      <nav className="fixed bottom-0 left-0 right-0 bg-card/80 backdrop-blur-lg border-t border-border px-4 py-2 flex justify-around lg:hidden z-50">
        <a href="/" className="p-3 rounded-full hover:bg-primary/10 transition">
          <Home size={24} />
        </a>
        <a href="/explore" className="p-3 rounded-full hover:bg-primary/10 transition">
          <Hash size={24} />
        </a>
        <a href="/notifications" className="p-3 rounded-full hover:bg-primary/10 transition relative">
          <Bell size={24} />
          <span className="absolute top-2 right-2 w-2 h-2 bg-primary rounded-full" />
        </a>
        <a href="/messages" className="p-3 rounded-full hover:bg-primary/10 transition">
          <Mail size={24} />
        </a>
      </nav>
    </div>
  )
}
